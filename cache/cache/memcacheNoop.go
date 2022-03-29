package leafCache

import (
	"context"
	"time"
)

type noopMemcache struct{}

func (n noopMemcache) Ping(ctx context.Context) error { return nil }

func (n noopMemcache) Close() error { return nil }

func (n noopMemcache) Len(ctx context.Context) int { return 0 }

func (n noopMemcache) Size(ctx context.Context) uintptr { return 0 }

func (n noopMemcache) Keys(ctx context.Context) []string { return make([]string, 0) }

func (n noopMemcache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return nil
}

func (n noopMemcache) Get(ctx context.Context, key string) (interface{}, error) { return "", nil }

func (n noopMemcache) Remove(ctx context.Context, key string) error { return nil }

func (n noopMemcache) Truncate(ctx context.Context) error { return nil }

func NoopMemcache() Memcache {
	return &noopMemcache{}
}
