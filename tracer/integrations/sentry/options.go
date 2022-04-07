package leafSentryTracer

import (
	"github.com/getsentry/sentry-go"
	leafTime "github.com/paulusrobin/leaf-utilities/time"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/ext"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/messageDestinationType"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/spanType"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
	"net/http"
	"time"
)

// =============
//  StartOption
// =============

type (
	StartOption interface {
		Apply(o *startOption)
	}
	startOption struct {
		sentryOptions      sentry.ClientOptions
		deferFlushDuration time.Duration
	}
)

func defaultStartOption() startOption {
	return startOption{
		sentryOptions:      sentry.ClientOptions{},
		deferFlushDuration: 2 * time.Second,
	}
}

type withSentryOptions struct{ sentry.ClientOptions }

func (w withSentryOptions) Apply(o *startOption) {
	o.sentryOptions = w.ClientOptions
}

func WithSentryOptions(options sentry.ClientOptions) StartOption {
	return withSentryOptions{options}
}

type withDeferFlushDuration struct{ time.Duration }

func (w withDeferFlushDuration) Apply(o *startOption) {
	o.deferFlushDuration = w.Duration
}

func WithDeferFlushDuration(duration time.Duration) StartOption {
	return withDeferFlushDuration{duration}
}

// =================
//  StartSpanOption
// =================

type (
	StartSpanOption = leafTracer.StartSpanOption

	DataStoreOption struct {
		Collection         string
		Operation          string
		ParameterizedQuery string
		QueryParameters    []interface{}
		DatabaseName       string
		DatastoreProduct   string
	}
	MessageProducerOption struct {
		Library              string
		DestinationType      messageDestinationType.Enum
		DestinationName      string
		DestinationTemporary bool
	}
)

func WithSpanType(t spanType.Enum) StartSpanOption {
	return tracer.Tag(ext.SpanType, t.Code())
}

func WithDataStore(option DataStoreOption) StartSpanOption {
	return tracer.Tag(ext.SpanTypeDataStore, option)
}

func WithMessageProducer(option MessageProducerOption) StartSpanOption {
	return tracer.Tag(ext.SpanTypeMessageProducer, option)
}

func WithExternal(option *http.Request) StartSpanOption {
	return tracer.Tag(ext.SpanTypeExternal, option)
}

func defaultStartSpanConfig() leafTracer.StartSpanConfig {
	return leafTracer.StartSpanConfig{
		Parent:    nil,
		StartTime: leafTime.Now(),
		Tags:      make(map[string]interface{}),
		SpanID:    0,
	}
}
