package leafGooglePubsub

import (
	"context"
	"errors"
	"fmt"
	leafHeader "github.com/paulusrobin/leaf-utilities/common/header"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	"strings"
	"time"

	google "cloud.google.com/go/pubsub"
)

type (
	publisher struct {
		option          option
		client          *google.Client
		publisherTopics map[string]*google.Topic
	}
)

func (p *publisher) Publish(ctx context.Context, topic string, msg leafMQ.Message) error {
	//span := startMessagingProducerSpan(ctx, topic)
	//defer span.Finish()

	if _, found := p.publisherTopics[topic]; !found {
		pubTopic := p.client.Topic(topic)
		isExist, err := pubTopic.Exists(ctx)
		if err != nil {
			p.option.logger.Error(leafLogger.BuildMessage(ctx, fmt.Sprintf("ERROR on googlePubsub.Publish.Exist: %+v", err.Error())))
			return err
		}

		if !isExist {
			err = errors.New("Topic " + topic + " is not exists in project " + p.option.googleProject)
			p.option.logger.Error(leafLogger.BuildMessage(ctx, fmt.Sprintf("ERROR on googlePubsub.Topic: %+v", err.Error())))
			return err
		}

		p.publisherTopics[topic] = pubTopic
	}

	if msg.Attributes == nil {
		msg.Attributes = make(map[string]string)
	}

	mandatory := leafMandatory.FromContext(ctx)
	msg.Attributes[leafHeader.MessagingTraceID] = mandatory.TraceID()
	msg.Attributes[leafHeader.MessagingAppVersion] = mandatory.Device().AppVersion()
	msg.Attributes[leafHeader.MessagingAuthorization] = mandatory.Authorization().Authorization()
	msg.Attributes[leafHeader.MessagingServiceID] = mandatory.Authorization().ServiceID()
	msg.Attributes[leafHeader.MessagingServiceSecret] = mandatory.Authorization().ServiceSecret()
	msg.Attributes[leafHeader.MessagingApiKey] = mandatory.Authorization().ApiKey()
	msg.Attributes[leafHeader.MessagingDeviceID] = mandatory.Device().DeviceID()
	msg.Attributes[leafHeader.MessagingUserAgent] = mandatory.UserAgent().Value()
	msg.Attributes[leafHeader.MessagingIpAddress] = strings.Join(mandatory.IpAddresses(), ",")
	msg.Attributes[leafHeader.MessagingPublishTime] = time.Now().Format(time.RFC3339)
	result := p.publisherTopics[topic].Publish(ctx, &google.Message{
		Data:       msg.Data,
		Attributes: msg.Attributes,
	})

	_, err := result.Get(ctx)
	if err != nil {
		p.option.logger.Error(leafLogger.BuildMessage(ctx, fmt.Sprintf("ERROR on googlePubsub.Publish: %+v", err.Error())))
		return err
	}
	p.option.logger.Debug(leafLogger.BuildMessage(ctx, fmt.Sprintf("SUCCESS publish message to topic: %s with message: %+v", topic, msg)))

	return nil
}

func (p *publisher) Close() error {
	return p.client.Close()
}
