package leafSentryTracer

import (
	"github.com/getsentry/sentry-go"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/ext"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/spanType"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
)

type SentryTracer struct {
	cfg startOption
}

// StartSpan implements leafTracer.Tracer.
func (s *SentryTracer) StartSpan(operationName string, opts ...leafTracer.StartSpanOption) leafTracer.Span {
	cfg := defaultStartSpanConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.Parent == nil {
		return newTransactionSpan(s, operationName, cfg)
	}

	tagSpanType, found := cfg.Tags[ext.SpanType]
	if found {
		var t spanType.Enum
		if err := t.UnmarshalParam(tagSpanType.(string)); err != nil {
			return newSegment(s, operationName, cfg)
		}

		switch t {
		case spanType.DataStore:
			return newDataStoreSegment(s, operationName, cfg)
		case spanType.External:
			return newExternalSegment(s, operationName, cfg)
		case spanType.MessageProducer:
			return newMessageProducerSegment(s, operationName, cfg)
		}
	}

	return newSegment(s, operationName, cfg)
}

// SetServiceInfo implements leafTracer.Tracer.
func (s *SentryTracer) SetServiceInfo(name, app, appType string) {

}

// Extract implements leafTracer.Tracer.
func (s *SentryTracer) Extract(carrier interface{}) (leafTracer.SpanContext, error) {
	return &SentrySpanContext{}, nil
}

// Inject implements leafTracer.Tracer.
func (s *SentryTracer) Inject(context leafTracer.SpanContext, carrier interface{}) error { return nil }

// Stop implements leafTracer.Tracer.
func (s *SentryTracer) Stop() {
	sentry.Flush(s.cfg.deferFlushDuration)
}

func InitTracing(options ...StartOption) (leafTracer.Tracer, error) {
	cfg := defaultStartOption()
	for _, opt := range options {
		opt.Apply(&cfg)
	}

	err := sentry.Init(cfg.sentryOptions)
	if err != nil {
		return nil, err
	}

	return &SentryTracer{
		cfg: cfg,
	}, nil
}
