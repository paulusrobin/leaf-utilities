package leafGooglePubsub

import (
	"context"
	"errors"
	"fmt"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	leafSlack "github.com/paulusrobin/leaf-utilities/slack"
	"sync"
	"time"

	google "cloud.google.com/go/pubsub"
	goOption "google.golang.org/api/option"
)

const (
	EventPublish = "google_pubsub_publish"
)

type pubSub struct {
	option          option
	client          *google.Client
	publisherTopics map[string]*google.Topic
	subscriptions   map[string]*subscriber
	listening       bool
}

func (p *pubSub) Publisher() leafMQ.Publisher {
	return p
}

func (p *pubSub) Consumer() leafMQ.Consumer {
	return p
}

func (p *pubSub) Use(middlewareFunc ...leafMQ.MiddlewareFunc) {
	for i := range p.subscriptions {
		p.subscriptions[i].middlewares = append(p.subscriptions[i].middlewares, middlewareFunc...)
	}
}

func (p *pubSub) Ping(ctx context.Context) error {
	if p.listening {
		return nil
	}
	return errors.New("pubsub is not connected")
}

func (p *pubSub) Listen() {
	var wg sync.WaitGroup
	wg.Add(len(p.subscriptions))

	for topic, subscriberData := range p.subscriptions {
		var ch = make(chan bool)

		go func(ch chan bool) {
			select {
			case <-ch:
				wg.Done()
				close(ch)
			}
		}(ch)
		go func(topic string, subscriber *subscriber, ch chan bool) {
			subscriber.Listen(topic, ch)
		}(topic, subscriberData, ch)
	}

	wg.Wait()
	p.listening = true
	p.option.logger.Info(leafLogger.BuildMessage(context.Background(), "start listening google pubsub"))
}

func (p *pubSub) Subscribe(topic string, dispatcher leafMQ.Dispatcher) error {
	var ctx = context.Background()

	subscriptionName, ok := p.option.subscription[topic]
	if !ok {
		return errors.New(fmt.Sprintf("no subscription name found for topic %s", topic))
	}
	subscription := p.client.Subscription(subscriptionName)
	if exists, err := subscription.Exists(ctx); !exists {
		_, err = p.client.CreateSubscription(
			ctx, subscriptionName,
			google.SubscriptionConfig{
				Topic:               p.client.Topic(topic),
				AckDeadline:         p.option.ackDeadline,
				RetainAckedMessages: false,
			},
		)

		if err != nil {
			p.option.logger.Error(leafLogger.BuildMessage(ctx, fmt.Sprintf("ERROR on googlePubsub.Subscribe.Create: %+v", err.Error())))
			return err
		}

		subscription = p.client.Subscription(subscriptionName)
	}

	subscription.ReceiveSettings.Synchronous = !p.option.asynchronous
	subscription.ReceiveSettings.MaxExtension = p.option.maxExtensionDeadline
	subscription.ReceiveSettings.MaxOutstandingMessages = p.option.maxOutstandingMessages

	if p.option.asynchronous {
		numGoroutine := p.option.numGoroutines
		if numGoroutine == 0 {
			numGoroutine = google.DefaultReceiveSettings.NumGoroutines
		}
		subscription.ReceiveSettings.NumGoroutines = numGoroutine
	}

	p.subscriptions[topic] = &subscriber{
		option:       p.option,
		subscription: subscription,
		dispatcher:   dispatcher,
	}

	if p.option.slackNotification.Active {
		slackNotification, err := leafSlack.Notification(
			leafSlack.WithHook(p.option.slackNotification.Hook),
			leafSlack.WithTimeout(p.option.slackNotification.Timeout))
		if err != nil {
			return err
		}
		p.subscriptions[topic].slackNotification = slackNotification
	}
	return nil
}

func (p *pubSub) Publish(ctx context.Context, topic string, msg leafMQ.Message) error {
	publisher := &publisher{option: p.option, client: p.client}
	defer func() {
		if err := publisher.Close(); err != nil {
			p.option.logger.StandardLogger().Error(fmt.Errorf("failed to Close producer: %+v", err))
		}
	}()

	return publisher.Publish(ctx, topic, msg)
}

func (p *pubSub) Close() error {
	return p.client.Close()
}

func New(options ...Option) (leafMQ.MessagingQueue, error) {
	opt := defaultOption()

	for _, option := range options {
		option.Apply(&opt)
	}

	if err := opt.validate(); err != nil {
		return nil, err
	}

	if opt.failedDeadline >= opt.ackDeadline {
		opt.failedDeadline = opt.ackDeadline - (100 * time.Millisecond)
	}

	var err error
	var client *google.Client
	if opt.googleCredentialPath != "" {
		client, err = google.NewClient(context.Background(), opt.googleProject,
			goOption.WithCredentialsFile(opt.googleCredentialPath))
	} else {
		client, err = google.NewClient(context.Background(), opt.googleProject)
	}

	if err != nil {
		opt.logger.Error(leafLogger.BuildMessage(context.Background(), fmt.Sprintf("Init client: %s", err.Error())))
		return nil, err
	}

	return &pubSub{
		option:          opt,
		client:          client,
		publisherTopics: make(map[string]*google.Topic),
		subscriptions:   make(map[string]*subscriber),
	}, nil
}
