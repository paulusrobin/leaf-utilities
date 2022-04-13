package leafNewRelicTracer

import leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"

func newSegment(tracer *NrTracer, operationName string, cfg leafTracer.StartSpanConfig) *NrSpan {
	var (
		segment Segment
		spanID  = spanID(cfg)
		txn     = TransactionFromContext(cfg.Parent.Context())
	)

	if txn != nil {
		segment = txn.StartSegment(operationName)
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
