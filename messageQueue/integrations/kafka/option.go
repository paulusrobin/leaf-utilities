package leafKafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	leafZap "github.com/paulusrobin/leaf-utilities/logger/integrations/zap"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"time"
)

const (
	DefaultConsumerWorker       = 10
	DefaultConsumerRetryMax     = 3
	DefaultConsumerRetryBackoff = 2 * time.Second
	DefaultFailedDeadline       = 60 * time.Second
	DefaultStrategy             = "BalanceStrategyRoundRobin"
	DefaultHeartbeat            = 3
	DefaultProducerMaxBytes     = 1000000
	DefaultProducerRetryMax     = 3
	DefaultProducerRetryBackoff = 100
	BalanceStrategySticky       = "BalanceStrategySticky"
	BalanceStrategyRoundRobin   = "BalanceStrategyRoundRobin"
	BalanceStrategyRange        = "BalanceStrategyRange"
)

type (
	Option interface {
		Apply(o *option)
	}
	option struct {
		host                 []string
		clientID             string
		consumerWorker       int
		consumerGroup        string
		consumerRetryMax     int
		consumerRetryBackoff time.Duration
		strategy             Strategy
		heartbeat            int
		producerMaxBytes     int
		producerRetryMax     int
		producerRetryBackOff time.Duration
		kafkaVersion         string
		logger               leafLogger.Logger
		withoutProducer      bool
		withoutConsumer      bool
		slackNotification    SlackNotification
	}
	SlackNotification struct {
		Active  bool
		Hook    string
		Timeout time.Duration
	}
)

func defaultOption() option {
	return option{
		host:                 make([]string, 0),
		consumerWorker:       DefaultConsumerWorker,
		consumerGroup:        "",
		consumerRetryMax:     DefaultConsumerRetryMax,
		consumerRetryBackoff: DefaultConsumerRetryBackoff,
		strategy:             DefaultStrategy,
		heartbeat:            DefaultHeartbeat,
		producerMaxBytes:     DefaultProducerMaxBytes,
		producerRetryMax:     DefaultProducerRetryMax,
		producerRetryBackOff: DefaultProducerRetryBackoff,
		kafkaVersion:         "",
		logger:               leafZap.DefaultLog(),
		withoutProducer:      false,
		withoutConsumer:      false,
		slackNotification: SlackNotification{
			Active:  true,
			Hook:    "https://hooks.slack.com/services/",
			Timeout: 5 * time.Second,
		},
	}
}

func validate(option option) error {
	if len(option.host) < 1 {
		return fmt.Errorf("invalid kafka host")
	}
	if option.consumerGroup == "" {
		return fmt.Errorf("invalid kafka consumer group")
	}
	if option.kafkaVersion == "" {
		return fmt.Errorf("invalid kafka version")
	}
	return nil
}

func getStrategy(option option) sarama.BalanceStrategy {
	if option.strategy == BalanceStrategyRange {
		return sarama.BalanceStrategyRange
	}

	if option.strategy == BalanceStrategyRoundRobin {
		return sarama.BalanceStrategyRoundRobin
	}

	return sarama.BalanceStrategySticky
}

type withHost []string

func WithHost(host []string) Option {
	return withHost(host)
}

func (w withHost) Apply(o *option) {
	o.host = w
}

type withClientID string

func WithClientID(clientID string) Option {
	return withClientID(clientID)
}

func (w withClientID) Apply(o *option) {
	o.clientID = string(w)
}

type withConsumerWorker int

func WithConsumerWorker(worker int) Option {
	return withConsumerWorker(worker)
}

func (w withConsumerWorker) Apply(o *option) {
	o.consumerWorker = int(w)
}

type withConsumerGroup string

func WithConsumerGroup(group string) Option {
	return withConsumerGroup(group)
}

func (w withConsumerGroup) Apply(o *option) {
	o.consumerGroup = string(w)
}

type withConsumerRetryMax int

func WithConsumerRetryMax(maxRetry int) Option {
	return withProducerRetryMax(maxRetry)
}

func (w withConsumerRetryMax) Apply(o *option) {
	o.producerRetryMax = int(w)
}

type withStrategy Strategy

func WithStrategy(strategy Strategy) Option {
	return withStrategy(strategy)
}

func (w withStrategy) Apply(o *option) {
	o.strategy = Strategy(w)
}

type withHeartbeat int

func WithHeartbeat(heartbeat int) Option {
	return withHeartbeat(heartbeat)
}

func (w withHeartbeat) Apply(o *option) {
	o.heartbeat = int(w)
}

type withProducerMaxBytes int

func WithProducerMaxBytes(maxBytes int) Option {
	return withProducerMaxBytes(maxBytes)
}

func (w withProducerMaxBytes) Apply(o *option) {
	o.producerMaxBytes = int(w)
}

type withProducerRetryMax int

func WithProducerRetryMax(maxRetry int) Option {
	return withProducerRetryMax(maxRetry)
}

func (w withProducerRetryMax) Apply(o *option) {
	o.producerRetryMax = int(w)
}

type withProducerRetryBackOff time.Duration

func WithProducerRetryBackOff(retryBackoff time.Duration) Option {
	return withProducerRetryBackOff(retryBackoff)
}

func (w withProducerRetryBackOff) Apply(o *option) {
	o.producerRetryBackOff = time.Duration(w)
}

type withKafkaVersion string

func WithKafkaVersion(version string) Option {
	return withKafkaVersion(version)
}

func (w withKafkaVersion) Apply(o *option) {
	o.kafkaVersion = string(w)
}

type withLog struct{ leafLogger.Logger }

func WithLog(logger leafLogger.Logger) Option {
	return withLog{logger}
}

func (w withLog) Apply(o *option) {
	o.logger = w
}

type withoutProducer bool

func WithoutProducer() Option {
	return withoutProducer(true)
}

func (w withoutProducer) Apply(o *option) {
	o.withoutProducer = bool(w)
}

type withoutConsumer bool

func WithoutConsumer() Option {
	return withoutConsumer(true)
}

func (w withoutConsumer) Apply(o *option) {
	o.withoutConsumer = bool(w)
}

type withSlackNotification SlackNotification

func (w withSlackNotification) Apply(o *option) {
	o.slackNotification = SlackNotification(w)
}

func WithSlackNotification(notification SlackNotification) Option {
	return withSlackNotification(notification)
}

type withoutSlackNotification bool

func (w withoutSlackNotification) Apply(o *option) {
	o.slackNotification.Active = false
}

func WithoutSlackNotification() Option {
	return withoutSlackNotification(true)
}
