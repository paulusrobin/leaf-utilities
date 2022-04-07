package sentry

import (
	"github.com/getsentry/sentry-go"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
	"net/http"
)

func newTransactionSpan(t *SentryTracer, operationName string, cfg leafTracer.StartSpanConfig) *SentrySpan {
	var parent = &SentrySpanContext{}
	if cfg.Parent != nil {
		parent = cfg.Parent.(*SentrySpanContext)
	}

	spanOptions := []sentry.SpanOption{sentry.TransactionName(operationName)}
	if cfg.Tags != nil {
		request, requestOk := cfg.Tags[tracer.HttpRequestKey].(*http.Request)
		if requestOk {
			spanOptions = append(spanOptions, sentry.ContinueFromRequest(request))
		}
	}

	span := sentry.StartSpan(parent.Context(), operationName, spanOptions...)
	return &SentrySpan{
		span:          span,
		tracer:        t,
		context:       newSpanContext(span),
		operationName: operationName,
		spanID:        span.SpanID.String(),
		traceID:       span.TraceID.String(),
		parentID:      "",
	}
}
