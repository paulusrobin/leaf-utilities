package leafMQ

import (
	"context"
	"sync"
)

type singleEventDispatcher struct {
	handler      HandlerFunc
	errorHandler ErrorHandlerFunc
	middlewares  []MiddlewareFunc
	sync.Mutex
}

func (d *singleEventDispatcher) AddHandler(handler HandlerFunc, errorHandler ErrorHandlerFunc, msgType ...string) {
	d.handler = handler
	d.errorHandler = errorHandler
}

func (d *singleEventDispatcher) Dispatch(dto DispatchDTO, middlewareFunc ...MiddlewareFunc) error {
	dto.Log.StandardLogger().Debugf("[DISPATCHER-SINGLE-EVENT] receive [%v][%v] %v", dto.Msg.GetID(), dto.RequestID, string(dto.Msg.Data))
	middlewareFunc = append(middlewareFunc, d.middlewares...)
	dispatch := applyMiddleware(d.dispatch, middlewareFunc...)
	return dispatch(context.Background(), dto)
}

func (d *singleEventDispatcher) Use(middlewareFunc ...MiddlewareFunc) {
	d.Lock()
	defer d.Unlock()
	d.middlewares = append(d.middlewares, middlewareFunc...)
}

func NewSingleEventDispatcher() Dispatcher {
	return &singleEventDispatcher{
		middlewares: make([]MiddlewareFunc, 0),
	}
}

func (d *singleEventDispatcher) dispatch(ctx context.Context, dto DispatchDTO) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if dto.Type == Handle {
		return d.handle(ctx, dto.Msg)
	}
	return d.onError(ctx, dto.Msg, dto.Err)
}

func (d *singleEventDispatcher) handle(ctx context.Context, msg Message) error {
	if d.handler == nil {
		return MissingHandler
	}
	return d.handler(ctx, msg)
}

func (d *singleEventDispatcher) onError(ctx context.Context, msg Message, err error) error {
	if d.errorHandler == nil {
		return MissingHandler
	}
	d.errorHandler(ctx, msg, err)
	return nil
}
