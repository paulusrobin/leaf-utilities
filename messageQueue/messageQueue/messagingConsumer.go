package leafMQ

import (
	"context"
)

type Consumer interface {
	Ping(ctx context.Context) error
	Use(middlewareFunc ...MiddlewareFunc)
	Listen()
	Subscribe(topic string, dispatcher Dispatcher) error
	Close() error
}
