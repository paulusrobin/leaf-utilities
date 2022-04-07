package leafSentryTracer

import (
	"bytes"
	"github.com/getsentry/sentry-go"
	"github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/ext"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"io/ioutil"
	"net/http"
)

func getExternalSegment(operationName string, cfg leafTracer.StartSpanConfig) *sentry.Span {
	var (
		parent = cfg.Parent.(*SentrySpanContext)
		span   = sentry.StartSpan(parent.Context(), operationName)
	)

	request := &http.Request{}
	_, found := cfg.Tags[ext.SpanTypeExternal]
	if found {
		request = cfg.Tags[ext.SpanTypeExternal].(*http.Request)
	} else {
		return span
	}

	if nil == span.Data {
		span.Data = make(map[string]interface{})
	}

	span.Data["Request Method"] = request.Method
	span.Data["Request Path"] = request.URL.String()
	span.Data["Request Header"] = request.Header
	if body, err := ioutil.ReadAll(request.Body); err == nil {
		span.Data["Request Body"] = string(body)
		request.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
	return span
}

func newExternalSegment(tracer *SentryTracer, operationName string, cfg leafTracer.StartSpanConfig) *SentrySpan {
	span := getExternalSegment(operationName, cfg)
	return &SentrySpan{
		tracer:        tracer,
		span:          span,
		context:       newSpanContext(span),
		operationName: operationName,
		spanID:        span.SpanID.String(),
		traceID:       span.TraceID.String(),
		parentID:      cfg.Parent.SpanID(),
	}
}
