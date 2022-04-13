package leafWorkerMiddleware

import (
	"context"
	"fmt"
	leafWorker "github.com/enricodg/leaf-utilities/appRunner/worker"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func Tracer() leafWorker.MiddlewareFunc {
	return func(next leafWorker.MiddlewareHandlerFunc, runner leafWorker.IRunner) leafWorker.MiddlewareHandlerFunc {
		return func(ctx context.Context) error {
			operationName := fmt.Sprintf("WORKER - %s", runner.OperationName())
			if ctx == nil {
				ctx = context.Background()
			}

			span, ctx := tracer.StartSpanFromContext(ctx, operationName)

			mandatory := leafMandatory.FromContext(ctx)
			span.SetTag("requestId", mandatory.TraceID())

			err := next(ctx)
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
