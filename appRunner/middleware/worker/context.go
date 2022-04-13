package leafWorkerMiddleware

import (
	"context"
	"fmt"
	leafMiddleware "github.com/enricodg/leaf-utilities/appRunner/middleware"
	leafWorker "github.com/enricodg/leaf-utilities/appRunner/worker"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	leafTime "github.com/paulusrobin/leaf-utilities/time"
	"time"
)

func buildMandatory(builder leafMandatory.MandatoryBuilder) leafMandatory.Mandatory {
	return builder.
		WithUserAgent("system").
		WithTraceID(leafMiddleware.GenerateUUID()).
		WithUserPhone(0, "system", "").Build()
}

func AppContext() leafWorker.MiddlewareFunc {
	return func(next leafWorker.MiddlewareHandlerFunc, runner leafWorker.IRunner) leafWorker.MiddlewareHandlerFunc {
		return func(ctx context.Context) error {
			if ctx == nil {
				ctx = context.Background()
			}

			builder, err := leafMandatory.NewMandatoryBuilder()
			if err != nil {
				return err
			}

			ctx = leafMandatory.Context(ctx, buildMandatory(builder))
			return next(ctx)
		}
	}
}

func AppContextWithLogger(logger leafLogger.Logger) leafWorker.MiddlewareFunc {
	return func(next leafWorker.MiddlewareHandlerFunc, runner leafWorker.IRunner) leafWorker.MiddlewareHandlerFunc {
		return func(ctx context.Context) error {
			if ctx == nil {
				ctx = context.Background()
			}

			builder, err := leafMandatory.NewMandatoryBuilder()
			if err != nil {
				return err
			}

			mandatory := buildMandatory(builder)
			ctx = leafMandatory.Context(ctx, mandatory)
			logRequest(ctx, runner.OperationName(), logger)

			start := leafTime.Now()
			err = next(ctx)
			duration := leafTime.Now().Sub(start)

			logResponse(ctx, runner.OperationName(), logger, duration, err)
			return err
		}
	}
}

func logRequest(ctx context.Context, operationName string, logger leafLogger.Logger) {
	requestMsg := fmt.Sprintf("[worker-log] - [REQUEST] Start %s", operationName)
	logger.Info(leafLogger.BuildMessage(ctx, requestMsg))
}

func logResponse(ctx context.Context, operationName string, logger leafLogger.Logger, duration time.Duration, err error) {
	responseMsg := fmt.Sprintf("[messaging-log] - [RESPONSE] - [%dms] Finish %s", duration.Milliseconds(), operationName)

	if err != nil {
		logger.Error(leafLogger.BuildMessage(ctx, responseMsg, leafLogger.WithAttr("error", err)))
	} else {
		logger.Info(leafLogger.BuildMessage(ctx, responseMsg))
	}
}
