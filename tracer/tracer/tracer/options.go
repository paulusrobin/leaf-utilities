package tracer

import (
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"time"
)

// ==========================
//     Start Span Option
// ==========================

// WithSpanID sets the SpanID on the started span, instead of using a random number.
// If there is no parent Span (eg from ChildOf), then the TraceID will also be set to the
// value given here.
func WithSpanID(id uint64) leafTracer.StartSpanOption {
	return func(cfg *leafTracer.StartSpanConfig) {
		cfg.SpanID = id
	}
}

// ChildOf tells StartSpan to use the given span context as a parent for the
// created span.s
func ChildOf(ctx leafTracer.SpanContext) leafTracer.StartSpanOption {
	return func(cfg *leafTracer.StartSpanConfig) {
		cfg.Parent = ctx
	}
}

// Tag sets the given key/value pair as a tag on the started Span.
func Tag(k string, v interface{}) leafTracer.StartSpanOption {
	return func(cfg *leafTracer.StartSpanConfig) {
		if cfg.Tags == nil {
			cfg.Tags = map[string]interface{}{}
		}
		cfg.Tags[k] = v
	}
}

// StartTime sets a custom time as the start time for the created span. By
// default a span is started using the creation time.
func StartTime(t time.Time) leafTracer.StartSpanOption {
	return func(cfg *leafTracer.StartSpanConfig) {
		cfg.StartTime = t
	}
}

// =======================
//     Finish Option
// =======================

// FinishTime sets the given time as the finishing time for the span. By default,
// the current time is used.
func FinishTime(t time.Time) leafTracer.FinishOption {
	return func(cfg *leafTracer.FinishConfig) {
		cfg.FinishTime = t
	}
}

// WithError marks the span as having had an error. It uses the information from
// err to set tags such as the error message, error type and stack trace. It has
// no effect if the error is nil.
func WithError(err error) leafTracer.FinishOption {
	return func(cfg *leafTracer.FinishConfig) {
		cfg.Error = err
	}
}

// NoDebugStack prevents any error presented using the WithError finishing option
// from generating a stack trace. This is useful in situations where errors are frequent
// and performance is critical.
func NoDebugStack() leafTracer.FinishOption {
	return func(cfg *leafTracer.FinishConfig) {
		cfg.NoDebugStack = true
	}
}

// StackFrames limits the number of stack frames included into erroneous spans to n, starting from skip.
func StackFrames(n, skip uint) leafTracer.FinishOption {
	if n == 0 {
		return NoDebugStack()
	}
	return func(cfg *leafTracer.FinishConfig) {
		cfg.StackFrames = n
		cfg.SkipStackFrames = skip
	}
}
