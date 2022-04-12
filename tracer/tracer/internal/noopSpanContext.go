package internal

import "context"

// NoopSpanContext is an implementation of leafTracer.SpanContext that is a no-op.
type NoopSpanContext struct{}

// SpanID implements leafTracer.SpanContext.
func (NoopSpanContext) SpanID() string { return "" }

// TraceID implements leafTracer.SpanContext.
func (NoopSpanContext) TraceID() string { return "" }

// ForeachBaggageItem implements leafTracer.SpanContext.
func (NoopSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}

// Context implements leafTracer.SpanContext.
func (NoopSpanContext) Context() context.Context { return context.Background() }
