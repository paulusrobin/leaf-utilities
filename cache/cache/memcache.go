package leafCache

import (
	"context"
	"math"
	"time"
)

const Infinite = math.MaxInt32
const Forever = 87660 * time.Hour

type Memcache interface {
	Cache

	Len(ctx context.Context) int
	Size(ctx context.Context) uintptr
	Keys(ctx context.Context) []string

	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)

	Remove(ctx context.Context, key string) error
	Truncate(ctx context.Context) error
}
