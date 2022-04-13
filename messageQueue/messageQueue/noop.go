package leafMQ

import (
	"context"
)

type (
	noopQueue struct {
		consumer  Consumer
		publisher Publisher
	}
)

func (n noopQueue) Ping(ctx context.Context) error { return nil }

func (n noopQueue) Close() error { return nil }

func (n noopQueue) Publish(ctx context.Context, topic string, msg Message) error { return nil }

func (n noopQueue) Use(middlewareFunc ...MiddlewareFunc) {}

func (n noopQueue) Listen() {}

func (n noopQueue) Subscribe(topic string, dispatcher Dispatcher) error { return nil }

func (n noopQueue) Publisher() Publisher { return n }

func (n noopQueue) Consumer() Consumer { return n }

func NoopQueue() MessageQueue {
	return &noopQueue{}
}
