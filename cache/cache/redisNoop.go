package leafCache

import (
	"context"
	"encoding"
	"time"
)

type noopRedis struct{}

func (n noopRedis) Ping(ctx context.Context) error { return nil }

func (n noopRedis) Close() error { return nil }

func (n noopRedis) Keys(ctx context.Context, pattern string) ([]string, error) {
	return make([]string, 0), nil
}

func (n noopRedis) Get(ctx context.Context, key string, data encoding.BinaryUnmarshaler) error {
	return nil
}

func (n noopRedis) Set(ctx context.Context, key string, value encoding.BinaryMarshaler) error {
	return nil
}

func (n noopRedis) SetWithExpiration(ctx context.Context, key string, value encoding.BinaryMarshaler, duration time.Duration) error {
	return nil
}

func (n noopRedis) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return make(map[string]string), nil
}

func (n noopRedis) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return make([]interface{}, 0), nil
}

func (n noopRedis) HMSet(ctx context.Context, key string, value map[string]interface{}) error {
	return nil
}

func (n noopRedis) HMSetWithExpiration(ctx context.Context, key string, value map[string]interface{}, ttl time.Duration) error {
	return nil
}

func (n noopRedis) HGet(ctx context.Context, key string, field string, response encoding.BinaryUnmarshaler) error {
	return nil
}

func (n noopRedis) HSet(ctx context.Context, key string, field string, value interface{}) error {
	return nil
}

func (n noopRedis) HSetWithExpiration(ctx context.Context, key string, field string, value interface{}, ttl time.Duration) error {
	return nil
}

func (n noopRedis) MSet(ctx context.Context, data map[string]interface{}) error {
	return nil
}

func (n noopRedis) MGet(ctx context.Context, key []string) ([]interface{}, error) {
	return make([]interface{}, 0), nil
}

func (n noopRedis) Remove(ctx context.Context, key string) error { return nil }

func (n noopRedis) RemoveByPattern(ctx context.Context, pattern string, countPerLoop int64) error {
	return nil
}

func (n noopRedis) FlushDatabase(ctx context.Context) error { return nil }

func (n noopRedis) FlushAll(ctx context.Context) error { return nil }

func NoopRedis() Redis {
	return &noopRedis{}
}
