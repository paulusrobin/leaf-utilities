package internal

import "context"

// NoopSpanContext is an implementation of taniTracer.SpanContext that is a no-op.
type NoopSpanContext struct{}

// SpanID implements taniTracer.SpanContext.
func (NoopSpanContext) SpanID() string { return "" }

// TraceID implements taniTracer.SpanContext.
func (NoopSpanContext) TraceID() string { return "" }

// ForeachBaggageItem implements taniTracer.SpanContext.
func (NoopSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}

// Context implements taniTracer.SpanContext.
func (NoopSpanContext) Context() context.Context { return context.Background() }
