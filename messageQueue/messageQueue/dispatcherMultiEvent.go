package leafMQ

import (
	"context"
	"sync"
)

type multiEventDispatcher struct {
	handlers      map[string]HandlerFunc
	errorHandlers map[string]ErrorHandlerFunc
	middlewares   []MiddlewareFunc
	sync.Mutex
}

func (d *multiEventDispatcher) AddHandler(handler HandlerFunc, errorHandler ErrorHandlerFunc, msgType ...string) {
	length := len(msgType)
	for i := 0; i < length; i++ {
		d.handlers[msgType[i]] = handler
		d.errorHandlers[msgType[i]] = errorHandler
	}
}

func (d *multiEventDispatcher) Dispatch(dto DispatchDTO, middlewareFunc ...MiddlewareFunc) error {
	if dto.MsgType == "" {
		dto.Log.StandardLogger().Warnf("[DISPATCHER-MULTI-EVENT] receive with no type [%v][%v] %v", dto.Msg.GetID(), dto.RequestID, string(dto.Msg.Data))
		return MissingHandler
	}

	dto.Log.StandardLogger().Debugf("[DISPATCHER-MULTI-EVENT] receive [%s][%v][%v] %v", dto.MsgType, dto.Msg.GetID(), dto.RequestID, string(dto.Msg.Data))
	middlewareFunc = append(middlewareFunc, d.middlewares...)
	dispatch := applyMiddleware(d.dispatch, middlewareFunc...)
	return dispatch(context.Background(), dto)
}

func (d *multiEventDispatcher) Use(middlewareFunc ...MiddlewareFunc) {
	d.Lock()
	defer d.Unlock()
	d.middlewares = append(d.middlewares, middlewareFunc...)
}

func NewMultiEventDispatcher() Dispatcher {
	return &multiEventDispatcher{
		handlers:      make(map[string]HandlerFunc),
		errorHandlers: make(map[string]ErrorHandlerFunc),
		middlewares:   make([]MiddlewareFunc, 0),
	}
}

func (d *multiEventDispatcher) handlerExist(msgType string) bool {
	if _, found := d.handlers[msgType]; found {
		return true
	}

	if _, found := d.errorHandlers[msgType]; found {
		return true
	}

	return false
}

func (d *multiEventDispatcher) dispatch(ctx context.Context, dto DispatchDTO) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if dto.Type == Handle {
		return d.handle(ctx, dto.MsgType, dto.Msg)
	}
	return d.onError(ctx, dto.MsgType, dto.Msg, dto.Err)
}

func (d *multiEventDispatcher) handle(ctx context.Context, msgType string, msg Message) error {
	if !d.handlerExist(msgType) {
		return MissingHandler
	}
	return d.handlers[msgType](ctx, msg)
}

func (d *multiEventDispatcher) onError(ctx context.Context, msgType string, msg Message, err error) error {
	if !d.handlerExist(msgType) {
		return MissingHandler
	}
	d.errorHandlers[msgType](ctx, msg, err)
	return nil
}
