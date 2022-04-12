package leafSentryTracer

import (
	"context"
	"github.com/getsentry/sentry-go"
)

// SentrySpanContext is an implementation of leafTracer.SpanContext that is for newrelic.
type (
	SentrySpanContext struct {
		span *sentry.Span
	}
)

// SpanID implements leafTracer.SpanContext.
func (s *SentrySpanContext) SpanID() string {
	if nil == s.span {
		return randomIDString()
	}
	return s.span.SpanID.String()
}

// TraceID implements leafTracer.SpanContext.
func (s *SentrySpanContext) TraceID() string {
	if nil == s.span {
		return randomIDString()
	}
	return s.span.TraceID.String()
}

// ForeachBaggageItem implements leafTracer.SpanContext.
func (s *SentrySpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}

// Context implements leafTracer.SpanContext.
func (s *SentrySpanContext) Context() context.Context {
	if nil == s.span {
		return context.Background()
	}
	return s.span.Context()
}

func newSpanContext(span *sentry.Span) *SentrySpanContext {
	return &SentrySpanContext{
		span: span,
	}
}
