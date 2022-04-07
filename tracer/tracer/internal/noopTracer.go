package internal

import leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"

// NoopTracer is an implementation of leafTracer.Tracer that is a no-op.
type NoopTracer struct{}

// StartSpan implements leafTracer.Tracer.
func (NoopTracer) StartSpan(operationName string, opts ...leafTracer.StartSpanOption) leafTracer.Span {
	return &NoopSpan{}
}

// SetServiceInfo implements leafTracer.Tracer.
func (NoopTracer) SetServiceInfo(name, app, appType string) {}

// Extract implements leafTracer.Tracer.
func (NoopTracer) Extract(carrier interface{}) (leafTracer.SpanContext, error) {
	return &NoopSpanContext{}, nil
}

// Inject implements leafTracer.Tracer.
func (NoopTracer) Inject(context leafTracer.SpanContext, carrier interface{}) error { return nil }

// Stop implements leafTracer.Tracer.
func (NoopTracer) Stop() {}
