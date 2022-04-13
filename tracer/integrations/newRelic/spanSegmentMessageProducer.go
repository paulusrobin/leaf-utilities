package leafNewRelicTracer

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/ext"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
)

func getMessageProducerSegment(trx *newrelic.Transaction, cfg leafTracer.StartSpanConfig) Segment {
	messageProducer := MessageProducerOption{}
	_, found := cfg.Tags[ext.SpanTypeMessageProducer]
	if found {
		messageProducer = cfg.Tags[ext.SpanTypeMessageProducer].(MessageProducerOption)
	} else {
		return &noopSegment{}
	}

	return &newrelic.MessageProducerSegment{
		StartTime:            trx.StartSegmentNow(),
		Library:              messageProducer.Library,
		DestinationType:      newrelic.MessageDestinationType(messageProducer.DestinationType.Code()),
		DestinationName:      messageProducer.DestinationName,
		DestinationTemporary: messageProducer.DestinationTemporary,
	}
}

func newMessageProducerSegment(tracer *NrTracer, operationName string, cfg leafTracer.StartSpanConfig) *NrSpan {
	var (
		segment Segment
		spanID  = spanID(cfg)
		txn     = TransactionFromContext(cfg.Parent.Context())
	)

	if txn != nil {
		segment = getMessageProducerSegment(txn, cfg)
	}
	return &NrSpan{
		tracer:        tracer,
		segment:       segment,
		context:       newSpanContext(spanID, cfg.Parent.(*NrSpanContext)),
		operationName: operationName,
		spanID:        spanID,
		traceID:       cfg.Parent.TraceID(),
		parentID:      cfg.Parent.SpanID(),
	}
}
