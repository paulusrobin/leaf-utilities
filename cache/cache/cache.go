package leafCache

import "context"

type Cache interface {
	Ping(ctx context.Context) error
	Close() error
	Remove(ctx context.Context, key string) error
}
