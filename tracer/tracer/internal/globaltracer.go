package internal

import (
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"sync"
)

var (
	mu           sync.RWMutex      // guards globalTracer
	globalTracer leafTracer.Tracer = &NoopTracer{}
)

// SetGlobalTracer sets the global tracer to t.
func SetGlobalTracer(t leafTracer.Tracer) {
	mu.Lock()
	defer mu.Unlock()
	globalTracer = t
}

// GetGlobalTracer returns the currently active tracer.
func GetGlobalTracer() leafTracer.Tracer {
	mu.RLock()
	defer mu.RUnlock()
	return globalTracer
}

var _ leafTracer.Tracer = (*NoopTracer)(nil)

var _ leafTracer.Span = (*NoopSpan)(nil)

var _ leafTracer.SpanContext = (*NoopSpanContext)(nil)
