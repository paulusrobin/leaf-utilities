package leafNewRelicTracer

import (
	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/ext"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/spanType"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"time"
)

type NrTracer struct {
	app *newrelic.Application
	cfg startOption
}

// StartSpan implements leafTracer.Tracer.
func (nr *NrTracer) StartSpan(operationName string, opts ...leafTracer.StartSpanOption) leafTracer.Span {
	cfg := defaultStartSpanConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.Parent == nil {
		return newTransactionSpan(nr, operationName, cfg)
	}

	tagSpanType, found := cfg.Tags[ext.SpanType]
	if found {
		var t spanType.Enum
		if err := t.UnmarshalParam(tagSpanType.(string)); err != nil {
			return newSegment(nr, operationName, cfg)
		}

		switch t {
		case spanType.DataStore:
			return newDataStoreSegment(nr, operationName, cfg)
		case spanType.External:
			return newExternalSegment(nr, operationName, cfg)
		case spanType.MessageProducer:
			return newMessageProducerSegment(nr, operationName, cfg)
		}
	}

	return newSegment(nr, operationName, cfg)
}

// SetServiceInfo implements leafTracer.Tracer.
func (nr *NrTracer) SetServiceInfo(name, app, appType string) {

}

// Extract implements leafTracer.Tracer.
func (nr *NrTracer) Extract(carrier interface{}) (leafTracer.SpanContext, error) {
	return &NrSpanContext{}, nil
}

// Inject implements leafTracer.Tracer.
func (nr *NrTracer) Inject(context leafTracer.SpanContext, carrier interface{}) error { return nil }

// Stop implements leafTracer.Tracer.
func (nr *NrTracer) Stop() {
	nr.app.Shutdown(12 * time.Second)
}

func InitTracing(options ...StartOption) (leafTracer.Tracer, error) {
	cfg := defaultStartOption()
	for _, opt := range options {
		opt.Apply(&cfg)
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.appName),
		newrelic.ConfigLicense(cfg.license),
		newrelic.ConfigDistributedTracerEnabled(cfg.distributedTracer),
		newrelic.ConfigFromEnvironment(),
		nrlogrus.ConfigStandardLogger(),
	)
	if err != nil {
		return nil, err
	}

	return &NrTracer{
		app: app,
		cfg: cfg,
	}, nil
}
