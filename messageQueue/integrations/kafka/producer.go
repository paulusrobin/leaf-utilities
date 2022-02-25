package leafKafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	leafHeader "github.com/paulusrobin/leaf-utilities/common/header"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	"strings"
	"time"
)

const (
	EventPublish = "kafka_publish"
)

type (
	producer struct {
		asyncProducer sarama.AsyncProducer
	}
)

func (p *producer) Publish(ctx context.Context, topic string, msg leafMQ.Message) error {
	//span := startMessagingProducerSpan(ctx, topic)
	//defer span.Finish()

	headers := make([]sarama.RecordHeader, 0)

	if msg.Attributes == nil {
		msg.Attributes = make(map[string]string)
	}

	if msg.GetID() == "" {
		msg.SetID(uuid.New().String())
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

	for key, attr := range msg.Attributes {
		headers = append(headers, sarama.RecordHeader{
			Key:   []byte(key),
			Value: []byte(attr),
		})
	}
	headers = append(headers,
		sarama.RecordHeader{
			Key:   []byte(leafHeader.MessagingPublishTime),
			Value: []byte(time.Now().Format(time.RFC3339)),
		},
		sarama.RecordHeader{
			Key:   []byte(leafHeader.MessagingID),
			Value: []byte(msg.GetID()),
		},
	)

	if "" == msg.Ordering {
		msg.Ordering = uuid.New().String()
	}

	p.asyncProducer.Input() <- &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(msg.Ordering),
		Value:     sarama.StringEncoder(string(msg.Data)),
		Headers:   headers,
		Timestamp: time.Now(),
	}
	return nil
}

func (p *producer) Close() error {
	return p.asyncProducer.Close()
}
