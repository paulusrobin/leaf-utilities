package sentry

import (
	"github.com/getsentry/sentry-go"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
)

func newSegment(tracer *SentryTracer, operationName string, cfg leafTracer.StartSpanConfig) *SentrySpan {
	var parent = cfg.Parent.(*SentrySpanContext)
	span := sentry.StartSpan(parent.Context(), operationName)
	return &SentrySpan{
		tracer:        tracer,
		span:          span,
		context:       newSpanContext(span),
		operationName: operationName,
		spanID:        span.SpanID.String(),
		traceID:       span.TraceID.String(),
		parentID:      parent.span.SpanID.String(),
	}
}
