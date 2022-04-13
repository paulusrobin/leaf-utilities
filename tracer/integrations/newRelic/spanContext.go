package leafNewRelicTracer

import "context"

// NrSpanContext is an implementation of taniTracer.SpanContext that is for newrelic.
type (
	NrSpanContext struct {
		ctx      context.Context
		traceID  string
		spanID   string
		parentID string
	}
)

// SpanID implements taniTracer.SpanContext.
func (nr *NrSpanContext) SpanID() string {
	return nr.spanID
}

// TraceID implements taniTracer.SpanContext.
func (nr *NrSpanContext) TraceID() string {
	return nr.traceID
}

// ForeachBaggageItem implements taniTracer.SpanContext.
func (nr *NrSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}

// Context implements taniTracer.SpanContext.
func (nr *NrSpanContext) Context() context.Context { return nr.ctx }

func newSpanContext(spanID string, parent *NrSpanContext) *NrSpanContext {
	if parent == nil {
		parent = &NrSpanContext{ctx: context.Background(), spanID: spanID}
	}
	return &NrSpanContext{
		ctx:      parent.ctx,
		spanID:   spanID,
		traceID:  parent.TraceID(),
		parentID: parent.SpanID(),
	}
}
