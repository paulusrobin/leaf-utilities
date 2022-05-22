package leafRedis

import (
	"context"
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	leafNewRelicTracer "github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic"
	leafNewRelicSpanType "github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/spanType"
	leafSentryTracer "github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry"
	leafSentrySpanType "github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/spanType"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

type dataStoreParam struct {
	db                 int
	operationName      string
	collectionName     string
	parameterizedQuery string
	queryParameters    []interface{}
}

func startDataStoreSpan(ctx *context.Context, param dataStoreParam) leafTracer.Span {
	var span leafTracer.Span
	span, *ctx = tracer.StartSpanFromContext(*ctx,
		`Redis `+param.operationName,
		leafNewRelicTracer.WithSpanType(leafNewRelicSpanType.DataStore),
		leafNewRelicTracer.WithDataStore(leafNewRelicTracer.DataStoreOption{
			Collection:         param.collectionName,
			Operation:          `Redis ` + param.operationName,
			ParameterizedQuery: param.parameterizedQuery,
			QueryParameters:    param.queryParameters,
			DatabaseName:       fmt.Sprintf("Redis#%d", param.db),
			DatastoreProduct:   newrelic.DatastoreRedis,
		}),
		leafSentryTracer.WithSpanType(leafSentrySpanType.DataStore),
		leafSentryTracer.WithDataStore(leafSentryTracer.DataStoreOption{
			Collection:         param.collectionName,
			Operation:          `Redis ` + param.operationName,
			ParameterizedQuery: param.parameterizedQuery,
			QueryParameters:    param.queryParameters,
			DatabaseName:       fmt.Sprintf("Redis#%d", param.db),
			DatastoreProduct:   "Redis",
		}),
	)
	return span
}
