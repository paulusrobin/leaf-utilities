package leafCache

import (
	"context"
	"encoding"
	"time"
)

type Redis interface {
	Cache
	Keys(ctx context.Context, pattern string) ([]string, error)

	Get(ctx context.Context, key string, data encoding.BinaryUnmarshaler) error
	Set(ctx context.Context, key string, value encoding.BinaryMarshaler) error
	SetWithExpiration(ctx context.Context, key string, value encoding.BinaryMarshaler, duration time.Duration) error

	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	HMSet(ctx context.Context, key string, value map[string]interface{}) error
	HMSetWithExpiration(ctx context.Context, key string, value map[string]interface{}, ttl time.Duration) error
	HGet(ctx context.Context, key string, field string, response encoding.BinaryUnmarshaler) error
	HSet(ctx context.Context, key string, field string, value interface{}) error
	HSetWithExpiration(ctx context.Context, key string, field string, value interface{}, ttl time.Duration) error
	MSet(ctx context.Context, data map[string]interface{}) error
	MGet(ctx context.Context, key []string) ([]interface{}, error)

	Remove(ctx context.Context, key string) error
	RemoveByPattern(ctx context.Context, pattern string, countPerLoop int64) error
	FlushDatabase(ctx context.Context) error
	FlushAll(ctx context.Context) error
}
