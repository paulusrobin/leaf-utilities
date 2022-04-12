package leafGooglePubsub

import (
	"context"
	"encoding/json"
	"fmt"
	leafHeader "github.com/paulusrobin/leaf-utilities/common/header"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	leafSlack "github.com/paulusrobin/leaf-utilities/slack"
	"sync"
	"time"

	google "cloud.google.com/go/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	subscriber struct {
		option            option
		subscription      *google.Subscription
		dispatcher        leafMQ.Dispatcher
		slackNotification leafSlack.Integration
		middlewares       []leafMQ.MiddlewareFunc
	}
	messageShown struct {
		MsgID         string            `json:"msg_id"`
		MsgData       interface{}       `json:"msg_data"`
		MsgAttributes map[string]string `json:"msg_attributes"`
	}
)

var (
	maskedAttributes = map[string]string{
		"authorization": `***token***`,
	}
)

func (s *subscriber) Dispatcher() leafMQ.Dispatcher {
	return s.dispatcher
}

func (s *subscriber) Listen(topic string, listening chan bool) {
	var wg sync.WaitGroup
	ch := make(chan *google.Message, s.option.bufSize)

	wg.Add(1)
	go func(topic string) {
		for {
			select {
			case msg := <-ch:
				if msg != nil {
					s.processMessage(topic, msg)
				}
			}
		}
	}(topic)

	go func() {
		listening <- true
		err := s.subscription.Receive(context.Background(), func(ctx context.Context, msg *google.Message) {
			ch <- msg
		})
		if err != nil && status.Code(err) != codes.Canceled {
			s.option.logger.Error(leafLogger.BuildMessage(context.Background(), fmt.Sprintf("PULL ERROR %v", err)))
			wg.Done()
		}
	}()

	wg.Wait()
	close(ch)
}

func (s *subscriber) processMessage(topic string, msg *google.Message) {
	requestID := msg.Attributes[leafHeader.MessagingID]
	msgType := msg.Attributes["message"]

	mqMsg := leafMQ.Message{
		Data:       msg.Data,
		Attributes: msg.Attributes,
	}
	mqMsg.SetID(msg.ID)

	message := leafMQ.DispatchDTO{
		Type:      leafMQ.Handle,
		Source:    fmt.Sprintf("Google PubSub - %s", topic),
		RequestID: requestID,
		MsgType:   msgType,
		Msg:       mqMsg,
		Log:       s.option.logger,
	}
	if err := s.dispatcher.Dispatch(message, s.middlewares...); err != nil {
		publishTime, _ := time.Parse(time.RFC3339, msg.Attributes[leafHeader.MessagingPublishTime])
		if time.Since(publishTime) < s.option.failedDeadline {
			msg.Nack()
			return
		}
		errMessage := leafMQ.DispatchDTO{
			Type:      leafMQ.Error,
			Source:    fmt.Sprintf("Google PubSub - %s", topic),
			RequestID: requestID,
			MsgType:   msgType,
			Msg:       mqMsg,
			Log:       s.option.logger,
			Err:       err,
		}
		_ = s.dispatcher.Dispatch(errMessage)
		s.pushNotificationToSlack(message)
		msg.Ack()
	} else {
		msg.Ack()
	}
}

func (s *subscriber) pushNotificationToSlack(msg leafMQ.DispatchDTO) {
	if !s.option.slackNotification.Active || s.slackNotification == nil {
		return
	}

	for key := range msg.Msg.Attributes {
		if val, ok := maskedAttributes[key]; ok {
			msg.Msg.Attributes[key] = val
		}
	}

	msgShown := messageShown{
		MsgID:         msg.Msg.GetID(),
		MsgData:       string(msg.Msg.Data),
		MsgAttributes: msg.Msg.Attributes,
	}

	var msgData map[string]interface{}
	if err := json.Unmarshal(msg.Msg.Data, &msgData); err == nil {
		msgShown.MsgData = msgData
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
		s.option.logger.Error(leafLogger.BuildMessage(context.Background(), fmt.Sprintf("Error on create push notification message to slack: %+v", err.Error())))
	}

	if err := s.slackNotification.Push(context.Background(), messageBody); err != nil {
		s.option.logger.Error(leafLogger.BuildMessage(context.Background(), fmt.Sprintf("Error on push notification to slack: %+v", err.Error())))
	}
}
