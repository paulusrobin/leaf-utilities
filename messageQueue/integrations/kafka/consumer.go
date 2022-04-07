package leafKafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	leafHeader "github.com/paulusrobin/leaf-utilities/common/header"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	leafSlack "github.com/paulusrobin/leaf-utilities/slack"
	"sync"
)

type (
	Strategy string
	consumer struct {
		mu                *sync.Mutex
		topics            map[string]leafMQ.Dispatcher
		ready             chan bool
		listening         bool
		option            option
		middlewares       []leafMQ.MiddlewareFunc
		slackNotification leafSlack.Integration
	}
	messageShown struct {
		ID         string            `json:"id"`
		Ordering   string            `json:"ordering"`
		Data       interface{}       `json:"data"`
		Attributes map[string]string `json:"attributes"`
	}
)

var (
	maskedAttributes = map[string]string{
		"authorization": `***token***`,
	}
)

func (c *consumer) Setup(session sarama.ConsumerGroupSession) error {
	if c.mu == nil {
		c.mu = &sync.Mutex{}
	}

	if c.topics == nil {
		c.topics = make(map[string]leafMQ.Dispatcher)
	}

	c.option.logger.StandardLogger().Info("Start Listening")
	c.listening = true
	return nil
}

func (c *consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		dispatcher := c.topics[message.Topic]

		if message == nil {
			continue
		}

		if dispatcher == nil {
			continue
		}

		c.processMessage(message.Topic, session, message, dispatcher)
	}

	return nil
}

func (c *consumer) getMessage(msg sarama.ConsumerMessage) leafMQ.Message {
	attributes := make(map[string]string)
	for _, attr := range msg.Headers {
		attributes[string(attr.Key)] = string(attr.Value)
	}

	msgID, found := attributes[leafHeader.MessagingID]
	if !found {
		msgID = uuid.New().String()
	}

	mqMessage := leafMQ.Message{
		Ordering:   string(msg.Key),
		Data:       msg.Value,
		Attributes: attributes,
	}
	mqMessage.SetID(msgID)
	return mqMessage
}

func (c *consumer) processMessage(topic string, session sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage, dispatcher leafMQ.Dispatcher) {
	messageData := c.getMessage(*msg)
	traceID := messageData.Attributes[leafHeader.MessagingTraceID]
	msgType := messageData.Attributes["message"]

	var err error
	var message = leafMQ.DispatchDTO{
		Type:      leafMQ.Handle,
		Source:    fmt.Sprintf("Kafka - %s", topic),
		RequestID: traceID,
		MsgType:   msgType,
		Msg:       messageData,
		Log:       c.option.logger,
	}
	for i := 0; i <= c.option.consumerRetryMax; i++ {
		message.Err = nil
		if err = dispatcher.Dispatch(message, c.middlewares...); err == nil {
			session.MarkMessage(msg, "")
			return
		} else {
			c.option.logger.StandardLogger().Error("error on dispatch message from kafka: ", err.Error())
			message.Err = err
		}
	}

	errMessage := leafMQ.DispatchDTO{
		Type:      leafMQ.Error,
		Source:    fmt.Sprintf("Kafka - %s", topic),
		RequestID: traceID,
		MsgType:   msgType,
		Msg:       messageData,
		Log:       c.option.logger,
		Err:       err,
	}
	_ = dispatcher.Dispatch(errMessage, c.middlewares...)
	c.pushNotificationToSlack(message)
	session.MarkMessage(msg, "")
}

func (c *consumer) pushNotificationToSlack(msg leafMQ.DispatchDTO) {
	if !c.option.slackNotification.Active || c.slackNotification == nil {
		return
	}

	for key := range msg.Msg.Attributes {
		if val, ok := maskedAttributes[key]; ok {
			msg.Msg.Attributes[key] = val
		}
	}

	msgShown := messageShown{
		ID:         msg.Msg.GetID(),
		Ordering:   msg.Msg.Ordering,
		Data:       string(msg.Msg.Data),
		Attributes: msg.Msg.Attributes,
	}

	var msgData map[string]interface{}
	if err := json.Unmarshal(msg.Msg.Data, &msgData); err == nil {
		msgShown.Data = msgData
	}

	jsonByte, _ := json.MarshalIndent(msgShown, "", "  ")
	messageBody, err := leafSlack.NewMessage(
		leafSlack.WithBlock(leafSlack.Block{
			Type: "header",
			Text: leafSlack.Text{
				Type: "plain_text",
				Text: fmt.Sprintf("Error on consuming Kafka Message: [%s]", msg.RequestID),
			},
		}),
		leafSlack.WithBlock(leafSlack.Block{
			Type: "section",
			Text: leafSlack.Text{
				Type: "mrkdwn",
				Text: fmt.Sprintf("<!channel> error on processing this message on %s", msg.Source),
			},
		}),
		leafSlack.WithBlock(leafSlack.Block{
			Type: "section",
			Text: leafSlack.Text{
				Type: "mrkdwn",
				Text: fmt.Sprintf("with error: ```%s```", msg.Err.Error()),
			},
		}),
		leafSlack.WithBlock(leafSlack.Block{
			Type: "section",
			Text: leafSlack.Text{
				Type: "mrkdwn",
				Text: fmt.Sprintf("with params: ```%s```", string(jsonByte)),
			},
		}),
	)
	if err != nil {
		c.option.logger.StandardLogger().Errorf("Error on create push notification message to slack: %+v", err.Error())
	}

	if err := c.slackNotification.Push(context.Background(), messageBody); err != nil {
		c.option.logger.StandardLogger().Errorf("Error on push notification to slack: %+v", err.Error())
	}
}
