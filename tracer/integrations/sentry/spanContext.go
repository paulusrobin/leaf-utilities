package leafSentryTracer

import (
	"context"
	"github.com/getsentry/sentry-go"
)

// SentrySpanContext is an implementation of taniTracer.SpanContext that is for newrelic.
type (
	SentrySpanContext struct {
		span *sentry.Span
	}
)

// SpanID implements taniTracer.SpanContext.
func (s *SentrySpanContext) SpanID() string {
	if nil == s.span {
		return randomIDString()
	}
	return s.span.SpanID.String()
}

// TraceID implements taniTracer.SpanContext.
func (s *SentrySpanContext) TraceID() string {
	if nil == s.span {
		return randomIDString()
	}
	return s.span.TraceID.String()
}

// ForeachBaggageItem implements taniTracer.SpanContext.
func (s *SentrySpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}

// Context implements taniTracer.SpanContext.
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
