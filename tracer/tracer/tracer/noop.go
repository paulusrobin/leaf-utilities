package tracer

import (
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/internal"
)

func NoopTracer() Tracer {
	return &internal.NoopTracer{}
}

func NoopSpan() Span {
	return &internal.NoopSpan{}
}

func NoopSpanContext() SpanContext {
	return &internal.NoopSpanContext{}
}
