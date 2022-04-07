package sentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/ext"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
)

func getMessageProducerSegment(operationName string, cfg leafTracer.StartSpanConfig) *sentry.Span {
	var (
		parent = cfg.Parent.(*SentrySpanContext)
		span   = sentry.StartSpan(parent.Context(), operationName)
	)

	messageProducer := MessageProducerOption{}
	_, found := cfg.Tags[ext.SpanTypeMessageProducer]
	if found {
		messageProducer = cfg.Tags[ext.SpanTypeMessageProducer].(MessageProducerOption)
	} else {
		return span
	}

	if nil == span.Data {
		span.Data = make(map[string]interface{})
	}

	span.Data["Library"] = messageProducer.Library
	span.Data["Destination Name"] = messageProducer.DestinationName
	span.Data["Destination Type"] = messageProducer.DestinationType.Code() + " - " + messageProducer.DestinationType.Name()
	return span
}

func newMessageProducerSegment(tracer *SentryTracer, operationName string, cfg leafTracer.StartSpanConfig) *SentrySpan {
	span := getMessageProducerSegment(operationName, cfg)
	return &SentrySpan{
		tracer:        tracer,
		span:          span,
		context:       newSpanContext(span),
		operationName: operationName,
		spanID:        span.SpanID.String(),
		traceID:       span.TraceID.String(),
		parentID:      cfg.Parent.SpanID(),
	}
}
