package leafGooglePubsub

import (
	"errors"
	leafZap "github.com/paulusrobin/leaf-utilities/logger/integrations/zap"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"time"

	google "cloud.google.com/go/pubsub"
)

type (
	Option interface {
		Apply(o *option)
	}
	option struct {
		// Logger is interface used for logging purpose
		// Logger default value is logs.DefaultLogger
		logger leafLogger.Logger

		// googleProject is used for connection project to google pubsub
		// googleProject is required
		googleProject string

		// subscription is used for connection subscription to google pubsub
		// subscription is required
		// Map key is subscription topic
		// Map value is subscription name
		subscription map[string]string

		// GoogleCredentialPath is used for connection project to google pubsub
		googleCredentialPath string

		// MaxRetry is used for retrying handler process when returning error
		// before call error handler
		//
		// MaxRetry configuration can be disabled by specifying a
		// number less than (or equal to) 0.
		maxRetry int

		// AckDeadline is the maximum period for which the subscription should
		// automatically extend the ack deadline for each message.
		//
		// The subscription will automatically extend the ack deadline of all
		// fetched Messages up to the duration specified. Automatic deadline
		// extension beyond the initial receipt may be disabled by specifying a
		// duration less than 0.
		ackDeadline time.Duration

		// MaxExtensionDeadline is the maximum duration by which to extend the ack
		// deadline at a time. The ack deadline will continue to be extended by up
		// to this duration until MaxExtension is reached. Setting MaxExtensionPeriod
		// bounds the maximum amount of time before a message redelivery in the
		// event the subscriber fails to extend the deadline.
		//
		// MaxExtensionDeadline configuration can be disabled by specifying a
		// duration less than (or equal to) 0.
		maxExtensionDeadline time.Duration

		// MaxOutstandingMessages is the maximum number of unprocessed messages
		// (unacknowledged but not yet expired). If MaxOutstandingMessages is 0, it
		// will be treated as if it were DefaultReceiveSettings.MaxOutstandingMessages.
		// If the value is negative, then there will be no limit on the number of
		// unprocessed messages.
		maxOutstandingMessages int

		// MaxOutstandingBytes is the maximum size of unprocessed messages
		// (unacknowledged but not yet expired). If MaxOutstandingBytes is 0, it will
		// be treated as if it were DefaultReceiveSettings.MaxOutstandingBytes. If
		// the value is negative, then there will be no limit on the number of bytes
		// for unprocessed messages.
		maxOutstandingBytes int

		// BufSize is used for buffer channel to listen messages
		// BufSize default value is 100
		bufSize int

		// If Asynchronous is false, then no more than MaxOutstandingMessages will be in
		// memory at one time. (In contrast, when ASynchronous is true, more than
		// MaxOutstandingMessages may have been received from the service and in memory
		// before being processed.) MaxOutstandingBytes still refers to the total bytes
		// processed, rather than in memory.
		// The default is false.
		//
		// NumGoroutines is ignored when the Asynchronous is false (Synchronous)
		asynchronous bool

		// NumGoroutines only used if Asynchronous is true
		//
		// NumGoroutines is the number of goroutines that each data structure along
		// the Receive path will spawn. Adjusting this value adjusts concurrency
		// along the receive path.
		//
		// NumGoroutines defaults to DefaultReceiveSettings.NumGoroutines.
		//
		// NumGoroutines does not limit the number of messages that can be processed
		// concurrently. Even with one goroutine, many messages might be processed at
		// once, because that goroutine may continually receive messages and invoke the
		// function passed to Receive on them. To limit the number of messages being
		// processed concurrently, set MaxOutstandingMessages.
		numGoroutines  int
		failedDeadline time.Duration

		slackNotification SlackNotification
	}

	SlackNotification struct {
		Active  bool
		Hook    string
		Timeout time.Duration
	}
)

func (o option) validate() error {
	if o.googleProject == "" {
		return errors.New("google project is required")
	}

	if o.slackNotification.Active && o.slackNotification.Hook == "" {
		return errors.New("hook is required when slack notification is active")
	}

	if len(o.subscription) == 0 {
		return errors.New("subscription is required")
	}

	for _, s := range o.subscription {
		if s == "" {
			return errors.New("subscription is required")
		}
	}

	return nil
}

func defaultOption() option {
	return option{
		logger:                 leafZap.DefaultLog(),
		maxRetry:               0,
		ackDeadline:            60 * time.Second,
		maxExtensionDeadline:   60 * time.Second,
		failedDeadline:         60 * time.Second,
		maxOutstandingMessages: google.DefaultReceiveSettings.MaxOutstandingMessages,
		maxOutstandingBytes:    google.DefaultReceiveSettings.MaxOutstandingBytes,
		bufSize:                100,
		asynchronous:           false,
		numGoroutines:          google.DefaultReceiveSettings.NumGoroutines,
		slackNotification: SlackNotification{
			Active:  true,
			Hook:    "https://hooks.slack.com/services/",
			Timeout: 5 * time.Second,
		},
	}
}

type withLog struct{ leafLogger.Logger }

func WithLog(logger leafLogger.Logger) Option {
	return withLog{logger}
}

func (w withLog) Apply(o *option) {
	o.logger = w
}

type withGoogleProject string

func (w withGoogleProject) Apply(o *option) {
	o.googleProject = string(w)
}

func WithGoogleProject(googleProject string) Option {
	return withGoogleProject(googleProject)
}

type withSubscription map[string]string

func (w withSubscription) Apply(o *option) {
	o.subscription = w
}

func WithSubscription(subscription map[string]string) Option {
	return withSubscription(subscription)
}

type withGoogleCredentialPath string

func (w withGoogleCredentialPath) Apply(o *option) {
	o.googleCredentialPath = string(w)
}

func WithGoogleCredentialPath(GoogleCredentialPath string) Option {
	return withGoogleCredentialPath(GoogleCredentialPath)
}

type withMaxRetry int

func (w withMaxRetry) Apply(o *option) {
	o.maxRetry = int(w)
}

func WithMaxRetry(MaxRetry int) Option {
	return withMaxRetry(MaxRetry)
}

type withAckDeadline time.Duration

func (w withAckDeadline) Apply(o *option) {
	o.ackDeadline = time.Duration(w)
}

func WithAckDeadline(AckDeadline time.Duration) Option {
	return withAckDeadline(AckDeadline)
}

type withMaxExtensionDeadline time.Duration

func (w withMaxExtensionDeadline) Apply(o *option) {
	o.maxExtensionDeadline = time.Duration(w)
}

func WithMaxExtensionDeadline(MaxExtensionDeadline time.Duration) Option {
	return withMaxExtensionDeadline(MaxExtensionDeadline)
}

type withFailedDeadline time.Duration

func (w withFailedDeadline) Apply(o *option) {
	o.failedDeadline = time.Duration(w)
}

func WithFailedDeadline(failedDeadline time.Duration) Option {
	return withFailedDeadline(failedDeadline)
}

type withMaxOutstandingMessages int

func (w withMaxOutstandingMessages) Apply(o *option) {
	o.maxOutstandingMessages = int(w)
}

func WithMaxOutstandingMessages(MaxOutstandingMessages int) Option {
	return withMaxOutstandingMessages(MaxOutstandingMessages)
}

type withMaxOutstandingBytes int

func (w withMaxOutstandingBytes) Apply(o *option) {
	o.maxOutstandingBytes = int(w)
}

func WithMaxOutstandingBytes(MaxOutstandingBytes int) Option {
	return withMaxOutstandingBytes(MaxOutstandingBytes)
}

type withBufSize int

func (w withBufSize) Apply(o *option) {
	o.bufSize = int(w)
}

func WithBufSize(BufSize int) Option {
	return withBufSize(BufSize)
}

type withAsynchronous bool

func (w withAsynchronous) Apply(o *option) {
	o.asynchronous = bool(w)
}
func WithAsynchronous(Asynchronous bool) Option {
	return withAsynchronous(Asynchronous)
}

type withNumGoroutines int

func (w withNumGoroutines) Apply(o *option) {
	o.numGoroutines = int(w)
}

func WithNumGoroutines(NumGoroutines int) Option {
	return withNumGoroutines(NumGoroutines)
}

type withSlackNotification SlackNotification

func (w withSlackNotification) Apply(o *option) {
	if o.slackNotification.Active {
		o.slackNotification = SlackNotification(w)
	}
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
