package tracer

import (
	"context"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/internal"
)

const (
	HttpRequestKey        = `http-request`
	HttpResponseWriterKey = `http-response-writer`
)

type contextKey struct{}

var activeSpanKey = contextKey{}

// ContextWithSpan returns a copy of the given context which includes the span s.
func ContextWithSpan(ctx context.Context, s Span) context.Context {
	return context.WithValue(ctx, activeSpanKey, s)
}

// SpanFromContext returns the span contained in the given context. A second return
// value indicates if a span was found in the context. If no span is found, a no-op
// span is returned.
func SpanFromContext(ctx context.Context) (Span, bool) {
	if ctx == nil {
		return &internal.NoopSpan{}, false
	}
	v := ctx.Value(activeSpanKey)
	if s, ok := v.(leafTracer.Span); ok {
		return s, true
	}
	return &internal.NoopSpan{}, false
}

// StartSpanFromContext returns a new span with the given operation name and options. If a span
// is found in the context, it will be used as the parent of the resulting span. If the ChildOf
// option is passed, the span from context will take precedence over it as the parent span.
func StartSpanFromContext(ctx context.Context, operationName string, opts ...leafTracer.StartSpanOption) (Span, context.Context) {
	if ctx == nil {
		// default to context.Background() to avoid panics on Go >= 1.15
		ctx = context.Background()
	}
	if s, ok := SpanFromContext(ctx); ok {
		opts = append(opts, ChildOf(s.Context()))
	}
	s := StartSpan(operationName, opts...)
	return s, ContextWithSpan(ctx, s)
}
