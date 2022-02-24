package leafKafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	"sync"
)

type kafka struct {
	option        option
	client        sarama.Client
	consumerGroup sarama.ConsumerGroup
	consumer      *consumer
}

func (k *kafka) Publisher() leafMQ.Publisher {
	return k
}

func (k *kafka) Consumer() leafMQ.Consumer {
	return k
}

func New(options ...Option) (leafMQ.MessagingQueue, error) {
	var err error

	option := defaultOption()
	for _, opt := range options {
		opt.Apply(&option)
	}

	if err := validate(option); err != nil {
		return nil, err
	}

	l := kafka{
		option: option,
	}

	version, err := sarama.ParseKafkaVersion(l.option.kafkaVersion)
	if err != nil {
		return nil, err
	}

	sarama.Logger = l.option.logger.StandardLogger()

	cfg := sarama.NewConfig()
	cfg.Version = version

	if len(l.option.clientID) > 0 {
		cfg.ClientID = l.option.clientID
	}

	if !option.withoutConsumer {
		// - consumer
		cfg.Consumer.Group.Rebalance.Strategy = getStrategy(l.option)
		cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
		cfg.Consumer.Retry.Backoff = l.option.consumerRetryBackoff
		cfg.Consumer.Return.Errors = true
		l.consumerGroup, err = sarama.NewConsumerGroup(l.option.host, l.option.consumerGroup, cfg)
		if err != nil {
			return nil, err
		}
	}

	if !option.withoutProducer {
		// - producer
		cfg.Producer.Return.Errors = true
		cfg.Producer.Return.Successes = true
		cfg.Producer.MaxMessageBytes = l.option.producerMaxBytes
		cfg.Producer.Retry.Max = l.option.producerRetryMax
		cfg.Producer.Retry.Backoff = l.option.producerRetryBackOff
	}

	l.client, err = sarama.NewClient(l.option.host, cfg)
	if err != nil {
		return nil, err
	}

	if !option.withoutConsumer {
		l.consumer = &consumer{
			mu:          &sync.Mutex{},
			topics:      make(map[string]leafMQ.Dispatcher),
			ready:       make(chan bool),
			middlewares: make([]leafMQ.MiddlewareFunc, 0),
			option:      l.option,
		}

		if l.option.slackNotification.Active {
			//slackNotification, err := taniSlack.Notification(
			//	taniSlack.WithHook(l.option.slackNotification.Hook),
			//	taniSlack.WithTimeout(l.option.slackNotification.Timeout))
			//if err != nil {
			//	return nil, err
			//}
			//l.consumer.slackNotification = slackNotification
		}
	}

	return &l, nil
}

func (k *kafka) Subscribe(topic string, dispatcher leafMQ.Dispatcher) error {
	if k.option.withoutConsumer {
		return fmt.Errorf("kafka is initialize without consumer")
	}

	k.consumer.mu.Lock()
	defer k.consumer.mu.Unlock()

	k.consumer.topics[topic] = dispatcher
	return nil
}

func (k *kafka) Use(middlewareFunc ...leafMQ.MiddlewareFunc) {
	k.consumer.mu.Lock()
	defer k.consumer.mu.Unlock()
	k.consumer.middlewares = append(k.consumer.middlewares, middlewareFunc...)
}

func (k *kafka) Listen() {
	if k.option.withoutConsumer {
		k.option.logger.StandardLogger().Error(fmt.Errorf("kafka is initialize without consumer"))
		return
	}

	if k.consumer.listening {
		k.option.logger.StandardLogger().Info("already listening to kafka")
		return
	}

	var (
		topics = make([]string, 0)
	)

	for key := range k.consumer.topics {
		topics = append(topics, key)
	}

	go func() {
		for {
			if err := k.consumerGroup.Consume(context.Background(), topics, k.consumer); err != nil {
				k.option.logger.StandardLogger().Errorf("error from consumer: %s", err.Error())
			}

			k.consumer.ready <- true
		}
	}()

	// Await till the consumer has been set up
	<-k.consumer.ready
}

func (k *kafka) Publish(ctx context.Context, topic string, msg leafMQ.Message) error {
	if k.option.withoutProducer {
		return fmt.Errorf("kafka is initialize without producer")
	}

	asyncProducer, err := sarama.NewAsyncProducerFromClient(k.client)
	if err != nil {
		return fmt.Errorf("error create async client message: %s", err.Error())
	}

	producer := &producer{asyncProducer: asyncProducer}
	defer func() {
		if err := producer.Close(); err != nil {
			k.option.logger.StandardLogger().Error(fmt.Errorf("failed to Close producer: %+v", err))
		}
	}()

	return producer.Publish(ctx, topic, msg)
}

func (k *kafka) Close() error {
	if k.consumerGroup != nil {
		if err := k.consumerGroup.Close(); err != nil {
			return fmt.Errorf("failed to close consumer: %+v", err)
		}
	}

	if k.client != nil {
		if err := k.client.Close(); err != nil {
			return fmt.Errorf("failed to close producer: %+v", err)
		}
	}

	return nil
}

func (k *kafka) Ping(ctx context.Context) error {
	if k.consumer.listening {
		return nil
	}
	return fmt.Errorf("kafka is not rebalanced")
}
