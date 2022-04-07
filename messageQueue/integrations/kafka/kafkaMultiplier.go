package leafKafka

import (
	"context"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	"math/rand"
)

type (
	kafkaMultipleClient struct {
		multiplier int
		publisher  leafMQ.Publisher
		consumer   leafMQ.Consumer
	}
	kafkaMultipleClientPublisher struct {
		clients    []leafMQ.Publisher
		multiplier int
	}
	kafkaMultipleClientConsumer struct {
		clients    []leafMQ.Consumer
		multiplier int
	}
)

// ===========
//  Consumer
// ===========

func (k kafkaMultipleClientConsumer) Ping(ctx context.Context) error {
	for _, client := range k.clients {
		if err := client.Ping(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (k kafkaMultipleClientConsumer) Use(middlewareFunc ...leafMQ.MiddlewareFunc) {
	for _, client := range k.clients {
		client.Use(middlewareFunc...)
	}
}

func (k kafkaMultipleClientConsumer) Listen() {
	for _, client := range k.clients {
		client.Listen()
	}
}

func (k kafkaMultipleClientConsumer) Subscribe(topic string, dispatcher leafMQ.Dispatcher) error {
	for _, client := range k.clients {
		if err := client.Subscribe(topic, dispatcher); err != nil {
			return err
		}
	}
	return nil
}

func (k kafkaMultipleClientConsumer) Close() error {
	for _, client := range k.clients {
		if err := client.Close(); err != nil {
			return err
		}
	}
	return nil
}

// ===========
//  Publisher
// ===========

func (k kafkaMultipleClientPublisher) Ping(ctx context.Context) error {
	for _, client := range k.clients {
		if err := client.Ping(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (k kafkaMultipleClientPublisher) Publish(ctx context.Context, topic string, msg leafMQ.Message) error {
	return k.clients[rand.Intn(k.multiplier-1)].Publish(ctx, topic, msg)
}

func (k kafkaMultipleClientPublisher) Close() error {
	for _, client := range k.clients {
		if err := client.Close(); err != nil {
			return err
		}
	}
	return nil
}

// ========
//  Client
// ========

func (k *kafkaMultipleClient) Ping(ctx context.Context) error {
	if err := k.publisher.Ping(ctx); err != nil {
		return err
	}
	return k.consumer.Ping(ctx)
}

func (k *kafkaMultipleClient) Publish(ctx context.Context, topic string, msg leafMQ.Message) error {
	return k.publisher.Publish(ctx, topic, msg)
}

func (k *kafkaMultipleClient) Use(middlewareFunc ...leafMQ.MiddlewareFunc) {
	k.consumer.Use(middlewareFunc...)
}

func (k *kafkaMultipleClient) Listen() {
	k.consumer.Listen()
}

func (k *kafkaMultipleClient) Subscribe(topic string, dispatcher leafMQ.Dispatcher) error {
	return k.consumer.Subscribe(topic, dispatcher)
}

func (k *kafkaMultipleClient) Close() error {
	if err := k.publisher.Close(); err != nil {
		return err
	}
	return k.consumer.Close()
}

func (k *kafkaMultipleClient) Publisher() leafMQ.Publisher {
	return k.publisher
}

func (k *kafkaMultipleClient) Consumer() leafMQ.Consumer {
	return k.consumer
}

func NewMultiple(multiplier int, options ...Option) (leafMQ.MessagingQueue, error) {
	var (
		consumers  = make([]leafMQ.Consumer, multiplier)
		publishers = make([]leafMQ.Publisher, multiplier)
	)
	for i := 0; i < multiplier; i++ {
		mq, err := New(options...)
		if err != nil {
			return nil, err
		}
		consumers = append(consumers, mq)
		publishers = append(publishers, mq)
	}
	return &kafkaMultipleClient{
		multiplier: multiplier,
		publisher:  &kafkaMultipleClientPublisher{publishers, multiplier},
		consumer:   &kafkaMultipleClientConsumer{consumers, multiplier},
	}, nil
}
