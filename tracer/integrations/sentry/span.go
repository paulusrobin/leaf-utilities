package sentry

import (
	"fmt"
	"github.com/getsentry/sentry-go"
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

func spanID(cfg leafTracer.StartSpanConfig) uint64 {
	spanID := randomID()
	if cfg.SpanID != 0 {
		spanID = cfg.SpanID
	}
	return spanID
}

// =============
//  Sentry Span
// =============

type SentrySpan struct {
	span    *sentry.Span
	tracer  *SentryTracer
	context *SentrySpanContext

	operationName string
	traceID       string
	spanID        string
	parentID      string
	baggage       map[string]string
	tag           map[string]interface{}

	rw sync.RWMutex
}

// SetTag implements leafTracer.Span.
func (s *SentrySpan) SetTag(key string, value interface{}) {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.tag == nil {
		s.tag = make(map[string]interface{})
	}
	s.tag[key] = value
	s.span.SetTag(key, fmt.Sprintf("%+v", value))
}

// SetOperationName implements leafTracer.Span.
func (s *SentrySpan) SetOperationName(operationName string) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.operationName = operationName
}

// BaggageItem implements leafTracer.Span.
func (s *SentrySpan) BaggageItem(key string) string {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.baggage == nil {
		return ""
	}

	if _, found := s.baggage[key]; found {
		return s.baggage[key]
	}

	if s.span.Data == nil {
		return ""
	}
	if _, found := s.span.Data[key]; found {
		return s.span.Data[key].(string)
	}
	return ""
}

// SetBaggageItem implements leafTracer.Span.
func (s *SentrySpan) SetBaggageItem(key, val string) {
	s.rw.Lock()
	defer s.rw.Unlock()

	if s.span.Data == nil {
		s.span.Data = make(map[string]interface{})
	}
	s.span.Data[key] = val

	if s.baggage == nil {
		s.baggage = make(map[string]string)
	}
	s.baggage[key] = val
}

// Finish implements leafTracer.Span.
func (s *SentrySpan) Finish(opts ...leafTracer.FinishOption) {
	if s.span == nil {
		return
	}

	cfg := leafTracer.FinishConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.Error != nil {
		sentry.CaptureException(cfg.Error)
	}
	s.span.Finish()
}

// Tracer implements leafTracer.Span.
func (s *SentrySpan) Tracer() leafTracer.Tracer { return s.tracer }

// Context implements leafTracer.Span.
func (s *SentrySpan) Context() leafTracer.SpanContext { return s.context }
