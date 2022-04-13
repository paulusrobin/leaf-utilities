package leafWorker

import (
	"context"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"os"
	"time"
)

type (
	IRunner interface {
		OperationName() string
		Serve(sig chan os.Signal, logger leafLogger.Logger)
	}
	MiddlewareHandlerFunc func(ctx context.Context) error
	MiddlewareFunc        func(MiddlewareHandlerFunc, IRunner) MiddlewareHandlerFunc
	Runner                struct {
		operationName string
		interval      time.Duration
		fn            MiddlewareHandlerFunc
		middlewares   []MiddlewareFunc
	}
	Runners []IRunner
)

func applyMiddleware(runner IRunner, h MiddlewareHandlerFunc, middleware ...MiddlewareFunc) MiddlewareHandlerFunc {
	for i := 0; i < len(middleware); i++ {
		h = middleware[i](h, runner)
	}
	return h
}

/*
	===============
		Runner
	===============
*/
func NewRunner(operationName string, interval time.Duration, fn func(ctx context.Context) error, middlewares ...MiddlewareFunc) *Runner {
	return &Runner{
		operationName: operationName,
		interval:      interval,
		fn:            fn,
		middlewares:   middlewares,
	}
}

func (r Runner) OperationName() string {
	return r.operationName
}

func (r Runner) Serve(sig chan os.Signal, logger leafLogger.Logger) {
	go func() {
		r.run(logger)
		tick := time.Tick(r.interval)
		for {
			select {
			case <-tick:
				r.run(logger)
				break
			case <-sig:
				return
			}
		}
	}()
}

func (r Runner) run(logger leafLogger.Logger) {
	ctx := context.Background()
	r.fn = applyMiddleware(r, r.fn, r.middlewares...)
	if err := r.fn(ctx); err != nil {
		logger.StandardLogger().Warnf("[WORKER-SERVER] error on worker: %s", err.Error())
	}
}

func (r *Runners) Add(runner IRunner) {
	*r = append(*r, runner)
}
