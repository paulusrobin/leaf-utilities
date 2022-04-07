package internal

import leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"

// NoopSpan is an implementation of leafTracer.Span that is a no-op.
type NoopSpan struct{}

// SetTag implements leafTracer.Span.
func (NoopSpan) SetTag(key string, value interface{}) {}

// SetOperationName implements leafTracer.Span.
func (NoopSpan) SetOperationName(operationName string) {}

// BaggageItem implements leafTracer.Span.
func (NoopSpan) BaggageItem(key string) string { return "" }

// SetBaggageItem implements leafTracer.Span.
func (NoopSpan) SetBaggageItem(key, val string) {}

// Finish implements leafTracer.Span.
func (NoopSpan) Finish(opts ...leafTracer.FinishOption) {}

// Tracer implements leafTracer.Span.
func (NoopSpan) Tracer() leafTracer.Tracer { return &NoopTracer{} }

// Context implements leafTracer.Span.
func (NoopSpan) Context() leafTracer.SpanContext { return &NoopSpanContext{} }
