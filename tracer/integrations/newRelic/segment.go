package leafNewRelicTracer

import (
	"fmt"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
)

type (
	Segment interface {
		AddAttribute(key string, val interface{})
		End()
	}
	noopSegment struct{}
)

func (no noopSegment) AddAttribute(key string, val interface{}) {}
func (no noopSegment) End()                                     {}

func spanID(cfg leafTracer.StartSpanConfig) string {
	spanID := randomIDString()
	if cfg.SpanID != 0 {
		spanID = fmt.Sprintf("%d", cfg.SpanID)
	}
	return spanID
}
