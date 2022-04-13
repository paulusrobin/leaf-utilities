package leafNewRelicTracer

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/ext"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"strconv"
)

func getDataStoreSegment(trx *newrelic.Transaction, cfg leafTracer.StartSpanConfig) Segment {
	dataStore := DataStoreOption{}
	_, found := cfg.Tags[ext.SpanTypeDataStore]
	if found {
		dataStore = cfg.Tags[ext.SpanTypeDataStore].(DataStoreOption)
	} else {
		return &noopSegment{}
	}

	var queryParams map[string]interface{}
	if dataStore.QueryParameters != nil {
		queryParams = map[string]interface{}{}
		for i, value := range dataStore.QueryParameters {
			queryParams[strconv.Itoa(i)] = value
		}
	}

	return &newrelic.DatastoreSegment{
		Product:            dataStore.DatastoreProduct,
		Collection:         dataStore.Collection,
		Operation:          dataStore.Operation,
		ParameterizedQuery: dataStore.ParameterizedQuery,
		QueryParameters:    queryParams,
		DatabaseName:       dataStore.DatabaseName,
		StartTime:          trx.StartSegmentNow(),
	}
}

func newDataStoreSegment(tracer *NrTracer, operationName string, cfg leafTracer.StartSpanConfig) *NrSpan {
	var (
		segment Segment
		spanID  = spanID(cfg)
		txn     = TransactionFromContext(cfg.Parent.Context())
	)

	if txn != nil {
		segment = getDataStoreSegment(txn, cfg)
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
