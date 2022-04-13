package leafNewRelicTracer

import (
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"math/rand"
	"sync"
)

func randomID() uint64 {
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}

func randomIDString() string {
	return fmt.Sprintf("%d", randomID())
}

// ================
//  NewRelic Span
// ================

type NrSpan struct {
	segment Segment
	tracer  *NrTracer
	context *NrSpanContext

	operationName string
	traceID       string
	spanID        string
	parentID      string
	baggage       map[string]string
	tag           map[string]interface{}

	rw sync.RWMutex
}

// SetTag implements leafTracer.Span.
func (nr *NrSpan) SetTag(key string, value interface{}) {
	nr.rw.Lock()
	defer nr.rw.Unlock()

	nr.segment.AddAttribute(key, value)
	if nr.tag == nil {
		nr.tag = make(map[string]interface{})
	}
	nr.tag[key] = value
}

// SetOperationName implements leafTracer.Span.
func (nr *NrSpan) SetOperationName(operationName string) {
	nr.rw.Lock()
	defer nr.rw.Unlock()
	nr.operationName = operationName
}

// BaggageItem implements leafTracer.Span.
func (nr *NrSpan) BaggageItem(key string) string {
	nr.rw.Lock()
	defer nr.rw.Unlock()

	if nr.baggage == nil {
		return ""
	}

	if _, found := nr.baggage[key]; found {
		return nr.baggage[key]
	}
	return ""
}

// SetBaggageItem implements leafTracer.Span.
func (nr *NrSpan) SetBaggageItem(key, val string) {
	nr.rw.Lock()
	defer nr.rw.Unlock()

	nr.segment.AddAttribute(key, val)
	if nr.baggage == nil {
		nr.baggage = make(map[string]string)
	}
	nr.baggage[key] = val
}

// Finish implements leafTracer.Span.
func (nr *NrSpan) Finish(opts ...leafTracer.FinishOption) {
	if nr.segment == nil {
		return
	}

	cfg := leafTracer.FinishConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.Error != nil {
		nr.segment.AddAttribute("error", cfg.Error)
	}
	nr.segment.End()
}

// Tracer implements leafTracer.Span.
func (nr *NrSpan) Tracer() leafTracer.Tracer { return nr.tracer }

// Context implements leafTracer.Span.
func (nr *NrSpan) Context() leafTracer.SpanContext { return nr.context }

// =======================
//  NewRelic Transaction
// =======================

type NrTransactionSpan struct {
	txn     *newrelic.Transaction
	tracer  *NrTracer
	context *NrSpanContext

	operationName string
	traceID       string
	spanID        string
	parentID      string
	baggage       map[string]string
	tag           map[string]interface{}

	rw sync.RWMutex
}

// SetTag implements leafTracer.Span.
func (nr *NrTransactionSpan) SetTag(key string, value interface{}) {
	nr.rw.Lock()
	defer nr.rw.Unlock()

	nr.txn.AddAttribute(key, value)
	if nr.tag == nil {
		nr.tag = make(map[string]interface{})
	}
	nr.tag[key] = value
}

// SetOperationName implements leafTracer.Span.
func (nr *NrTransactionSpan) SetOperationName(operationName string) {
	nr.rw.Lock()
	defer nr.rw.Unlock()
	nr.operationName = operationName
}

// BaggageItem implements leafTracer.Span.
func (nr *NrTransactionSpan) BaggageItem(key string) string {
	nr.rw.Lock()
	defer nr.rw.Unlock()

	if nr.baggage == nil {
		return ""
	}

	if _, found := nr.baggage[key]; found {
		return nr.baggage[key]
	}
	return ""
}

// SetBaggageItem implements leafTracer.Span.
func (nr *NrTransactionSpan) SetBaggageItem(key, val string) {
	nr.rw.Lock()
	defer nr.rw.Unlock()

	nr.txn.AddAttribute(key, val)
	if nr.baggage == nil {
		nr.baggage = make(map[string]string)
	}
	nr.baggage[key] = val
}

// Finish implements leafTracer.Span.
func (nr *NrTransactionSpan) Finish(opts ...leafTracer.FinishOption) {
	if nr.txn == nil {
		return
	}

	cfg := leafTracer.FinishConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.Error != nil {
		nr.txn.SetWebResponse(nil)
		nr.txn.NoticeError(cfg.Error)
	}
	nr.txn.End()
}

// Tracer implements leafTracer.Span.
func (nr *NrTransactionSpan) Tracer() leafTracer.Tracer { return nr.tracer }

// Context implements leafTracer.Span.
func (nr *NrTransactionSpan) Context() leafTracer.SpanContext { return nr.context }
