package leafMessagingMiddleware

import (
	"context"
	"fmt"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func Tracer() leafMQ.MiddlewareFunc {
	return func(next leafMQ.MiddlewareHandlerFunc) leafMQ.MiddlewareHandlerFunc {
		return func(ctx context.Context, msg leafMQ.DispatchDTO) error {
			operationName := fmt.Sprintf("MESSAGING - %s", msg.Source)
			if ctx == nil {
				ctx = context.Background()
			}

			span, ctx := tracer.StartSpanFromContext(ctx, operationName)

			mandatory := leafMandatory.FromContext(ctx)
			span.SetTag("requestId", mandatory.TraceID())
			span.SetTag("requestBody", msg)

			err := next(ctx, msg)
			if nil != err {
				if !leafFunctions.SkipNoticeError(ctx) {
					span.Finish(tracer.WithError(err))
				}
				return err
			}

			span.Finish()
			return nil
		}
	}
}
