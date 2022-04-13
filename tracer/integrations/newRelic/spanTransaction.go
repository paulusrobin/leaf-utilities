package leafNewRelicTracer

import (
	"context"
	"github.com/newrelic/go-agent/v3/newrelic"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
	"net/http"
)

const transactionKey = `nr-txn`

func transactionContext(parent context.Context, txn *newrelic.Transaction) context.Context {
	return context.WithValue(parent, transactionKey, txn)
}

func TransactionFromContext(parent context.Context) *newrelic.Transaction {
	if nil == parent {
		return nil
	}
	h, _ := parent.Value(transactionKey).(*newrelic.Transaction)
	if nil != h {
		return h
	}
	return nil
}

func newTransactionSpan(t *NrTracer, operationName string, cfg leafTracer.StartSpanConfig) *NrTransactionSpan {
	spanID := spanID(cfg)
	txn := t.app.StartTransaction(operationName)

	parent := &NrSpanContext{ctx: context.Background()}
	if cfg.Parent != nil {
		parent = cfg.Parent.(*NrSpanContext)
	}

	if cfg.Tags != nil {
		request, requestOk := cfg.Tags[tracer.HttpRequestKey].(*http.Request)
		if requestOk {
			txn.SetWebRequestHTTP(request)
		}

		responseWriter, responseWriterOk := cfg.Tags[tracer.HttpResponseWriterKey].(http.ResponseWriter)
		if responseWriterOk {
			txn.SetWebResponse(responseWriter)
		}
	}

	traceID := randomIDString()
	parent.ctx = transactionContext(parent.ctx, txn)
	return &NrTransactionSpan{
		txn:           txn,
		tracer:        t,
		context:       newSpanContext(spanID, parent),
		operationName: operationName,
		spanID:        spanID,
		traceID:       traceID,
		parentID:      "",
	}
}
