package leafGoMongo

import (
	"context"
	leafSentryTracer "github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry"
	leafSentrySpanType "github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/spanType"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

type dataStoreParam struct {
	databaseName       string
	operationName      string
	collectionName     string
	parameterizedQuery string
	queryParameters    []interface{}
}

func startDataStoreSpan(ctx *context.Context, param dataStoreParam) leafTracer.Span {
	var span leafTracer.Span

	span, *ctx = tracer.StartSpanFromContext(*ctx, param.operationName,
		//taniNewRelicTracer.WithSpanType(taniNewRelicSpanType.DataStore),
		//taniNewRelicTracer.WithDataStore(taniNewRelicTracer.DataStoreOption{
		//	Collection:         param.collectionName,
		//	Operation:          param.operationName,
		//	ParameterizedQuery: param.parameterizedQuery,
		//	QueryParameters:    param.queryParameters,
		//	DatabaseName:       param.databaseName,
		//	DatastoreProduct:   newrelic.DatastoreMongoDB,
		//}),
		leafSentryTracer.WithSpanType(leafSentrySpanType.DataStore),
		leafSentryTracer.WithDataStore(leafSentryTracer.DataStoreOption{
			Collection:         param.collectionName,
			Operation:          param.operationName,
			ParameterizedQuery: param.parameterizedQuery,
			QueryParameters:    param.queryParameters,
			DatabaseName:       param.databaseName,
			DatastoreProduct:   "MongoDB",
		}),
	)
	return span
}
