package leafMQ

import (
	"context"
)

type Publisher interface {
	Ping(ctx context.Context) error
	Publish(ctx context.Context, topic string, msg Message) error
	Close() error
}
