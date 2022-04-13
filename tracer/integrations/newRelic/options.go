package leafNewRelicTracer

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	leafTime "github.com/paulusrobin/leaf-utilities/time"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/ext"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/messageDestinationType"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/spanType"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
	"net/http"
)

// =============
//  StartOption
// =============

type (
	StartOption interface {
		Apply(o *startOption)
	}
	startOption struct {
		appName           string
		license           string
		distributedTracer bool
	}
)

func defaultStartOption() startOption {
	return startOption{
		appName:           "",
		license:           "",
		distributedTracer: true,
	}
}

type withAppName string

func (w withAppName) Apply(o *startOption) {
	o.appName = string(w)
}

func WithAppName(appName string) StartOption {
	return withAppName(appName)
}

type withLicense string

func (w withLicense) Apply(o *startOption) {
	o.license = string(w)
}

func WithLicense(license string) StartOption {
	return withLicense(license)
}

type withDistributedTracing bool

func (w withDistributedTracing) Apply(o *startOption) {
	o.distributedTracer = bool(w)
}

func WithDistributedTracing(distributedTracing bool) StartOption {
	return withDistributedTracing(distributedTracing)
}

// ==================
//  StartSpanOption
// ==================

type (
	StartSpanOption = leafTracer.StartSpanOption

	DataStoreOption struct {
		Collection         string
		Operation          string
		ParameterizedQuery string
		QueryParameters    []interface{}
		DatabaseName       string
		DatastoreProduct   newrelic.DatastoreProduct
	}
	MessageProducerOption struct {
		Library              string
		DestinationType      messageDestinationType.Enum
		DestinationName      string
		DestinationTemporary bool
	}
)

func WithSpanType(t spanType.Enum) StartSpanOption {
	return tracer.Tag(ext.SpanType, t.Code())
}

func WithDataStore(option DataStoreOption) StartSpanOption {
	return tracer.Tag(ext.SpanTypeDataStore, option)
}

func WithMessageProducer(option MessageProducerOption) StartSpanOption {
	return tracer.Tag(ext.SpanTypeMessageProducer, option)
}

func WithExternal(option *http.Request) StartSpanOption {
	return tracer.Tag(ext.SpanTypeExternal, option)
}

func defaultStartSpanConfig() leafTracer.StartSpanConfig {
	return leafTracer.StartSpanConfig{
		Parent:    nil,
		StartTime: leafTime.Now(),
		Tags:      make(map[string]interface{}),
		SpanID:    0,
	}
}
