package leafSentryTracer

import (
	"github.com/getsentry/sentry-go"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/ext"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"strconv"
)

func getDataStoreSegment(operationName string, cfg leafTracer.StartSpanConfig) *sentry.Span {
	var (
		parent = cfg.Parent.(*SentrySpanContext)
		span   = sentry.StartSpan(parent.Context(), operationName)
	)

	dataStore := DataStoreOption{}
	_, found := cfg.Tags[ext.SpanTypeDataStore]
	if found {
		dataStore = cfg.Tags[ext.SpanTypeDataStore].(DataStoreOption)
	} else {
		return span
	}

	var queryParams map[string]interface{}
	if dataStore.QueryParameters != nil {
		queryParams = map[string]interface{}{}
		for i, value := range dataStore.QueryParameters {
			queryParams[strconv.Itoa(i)] = value
		}
	}

	if nil == span.Data {
		span.Data = make(map[string]interface{})
	}

	span.Data["Collection"] = dataStore.Collection
	span.Data["Operation"] = dataStore.Operation
	span.Data["Parameterized Query"] = dataStore.ParameterizedQuery
	span.Data["Query Parameters"] = dataStore.QueryParameters
	span.Data["Database Name"] = dataStore.DatabaseName
	span.Data["Datastore Product"] = dataStore.DatastoreProduct
	return span
}

func newDataStoreSegment(tracer *SentryTracer, operationName string, cfg leafTracer.StartSpanConfig) *SentrySpan {
	span := getDataStoreSegment(operationName, cfg)
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
