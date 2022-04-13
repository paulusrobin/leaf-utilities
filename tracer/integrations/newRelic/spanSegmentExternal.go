package leafNewRelicTracer

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/ext"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"net/http"
)

func getExternalSegment(trx *newrelic.Transaction, cfg leafTracer.StartSpanConfig) Segment {
	request := &http.Request{}
	_, found := cfg.Tags[ext.SpanTypeExternal]
	if found {
		request = cfg.Tags[ext.SpanTypeExternal].(*http.Request)
	} else {
		return &noopSegment{}
	}

	return newrelic.StartExternalSegment(trx, request)
}

func newExternalSegment(tracer *NrTracer, operationName string, cfg leafTracer.StartSpanConfig) *NrSpan {
	var (
		segment Segment
		spanID  = spanID(cfg)
		txn     = TransactionFromContext(cfg.Parent.Context())
	)

	if txn != nil {
		segment = getExternalSegment(txn, cfg)
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
