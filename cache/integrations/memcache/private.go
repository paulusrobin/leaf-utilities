package leafMemcache

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
	operationName      string
	collectionName     string
	parameterizedQuery string
	queryParameters    []interface{}
}

func startDataStoreSpan(ctx *context.Context, param dataStoreParam) leafTracer.Span {
	var span leafTracer.Span
	span, *ctx = tracer.StartSpanFromContext(*ctx,
		`Memcache `+param.operationName,
		leafNewRelicTracer.WithSpanType(leafNewRelicSpanType.DataStore),
		leafNewRelicTracer.WithDataStore(leafNewRelicTracer.DataStoreOption{
			Collection:         param.collectionName,
			Operation:          `Memcache ` + param.operationName,
			ParameterizedQuery: param.parameterizedQuery,
			QueryParameters:    param.queryParameters,
			DatabaseName:       fmt.Sprintf("Memcache"),
			DatastoreProduct:   newrelic.DatastoreMemcached,
		}),
		leafSentryTracer.WithSpanType(leafSentrySpanType.DataStore),
		leafSentryTracer.WithDataStore(leafSentryTracer.DataStoreOption{
			Collection:         param.collectionName,
			Operation:          `Memcache ` + param.operationName,
			ParameterizedQuery: param.parameterizedQuery,
			QueryParameters:    param.queryParameters,
			DatabaseName:       fmt.Sprintf("Memcache"),
			DatastoreProduct:   "Memcache",
		}),
	)
	return span
}
