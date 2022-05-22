package leafRedis

import (
	"context"
	"encoding"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/paulusrobin/leaf-utilities/cache/cache"
	"time"
)

type (
	goredis struct {
		r   redis.UniversalClient
		cfg options
	}
)

func New(options ...Option) (leafCache.Redis, error) {
	var client redis.UniversalClient

	opt := defaultOption()
	for _, option := range options {
		option.Apply(&opt)
	}

	client = redis.NewUniversalClient(opt.universalOption())

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return &goredis{r: client, cfg: opt}, nil
}

func (c *goredis) Ping(ctx context.Context) error {
	if _, err := c.r.Ping().Result(); err != nil {
		return err
	}
	return nil
}

func (c *goredis) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}

	return nil
}

func (c *goredis) Keys(ctx context.Context, pattern string) ([]string, error) {
	if err := check(c); err != nil {
		return []string{}, err
	}
	return c.r.Keys(pattern).Result()
}

func (c *goredis) Get(ctx context.Context, key string, data encoding.BinaryUnmarshaler) error {
	if err := check(c); err != nil {
		return err
	}

	val, err := c.r.Get(key).Result()

	if err == redis.Nil {
		return err
	}

	if err != nil {
		return err
	}

	if err := data.(encoding.BinaryUnmarshaler).UnmarshalBinary([]byte(val)); err != nil {
		return err
	}

	return nil
}

func (c *goredis) setWithExpiration(ctx context.Context, key string, value encoding.BinaryMarshaler, duration time.Duration) error {
	if err := check(c); err != nil {
		return err
	}

	if _, err := c.r.Set(key, value, duration).Result(); err != nil {
		return err
	}
	return nil
}

func (c *goredis) Set(ctx context.Context, key string, value encoding.BinaryMarshaler) error {
	return c.setWithExpiration(ctx, key, value, 0)
}

func (c *goredis) SetWithExpiration(ctx context.Context, key string, value encoding.BinaryMarshaler, duration time.Duration) error {
	return c.setWithExpiration(ctx, key, value, duration)
}

func (c *goredis) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	if err := check(c); err != nil {
		return nil, err
	}

	val, err := c.r.HMGet(key, fields...).Result()
	if err == redis.Nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c *goredis) HMSet(ctx context.Context, key string, value map[string]interface{}) error {
	redisQuery := ""
	for k, v := range value {
		redisQuery += fmt.Sprintf("%s %v ", k, v)
	}

	if err := check(c); err != nil {
		return err
	}

	if _, err := c.r.HMSet(key, value).Result(); err != nil {
		return err
	}
	return nil
}

func (c *goredis) HMSetWithExpiration(ctx context.Context, key string, value map[string]interface{}, ttl time.Duration) error {
	redisQuery := ""
	for k, v := range value {
		redisQuery += fmt.Sprintf("%s %v ", k, v)
	}

	if err := check(c); err != nil {
		return err
	}

	if _, err := c.r.HMSet(key, value).Result(); err != nil {
		return err
	}

	if _, err := c.r.Expire(key, ttl).Result(); err != nil {
		c.r.Del(key)
		return err
	}
	return nil
}

func (c *goredis) HGet(ctx context.Context, key, field string, response encoding.BinaryUnmarshaler) error {
	if err := check(c); err != nil {
		return err
	}

	val, err := c.r.HGet(key, field).Result()
	if err == redis.Nil {
		return err
	}

	if err != nil {
		return err
	}

	if err := response.(encoding.BinaryUnmarshaler).UnmarshalBinary([]byte(val)); err != nil {
		return err
	}

	return nil
}

func (c *goredis) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	if err := check(c); err != nil {
		return nil, err
	}

	val, err := c.r.HGetAll(key).Result()
	if err == redis.Nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c *goredis) HSet(ctx context.Context, key, field string, value interface{}) error {
	if err := check(c); err != nil {
		return err
	}

	if _, err := c.r.HSet(key, field, value).Result(); err != nil {
		return err
	}
	return nil
}

func (c *goredis) HSetWithExpiration(ctx context.Context, key, field string, value interface{}, ttl time.Duration) error {
	if err := check(c); err != nil {
		return err
	}

	if _, err := c.r.HSet(key, field, value).Result(); err != nil {
		return err
	}
	if _, err := c.r.Expire(key, ttl).Result(); err != nil {
		c.r.Del(key)
		return err
	}
	return nil
}

func (c *goredis) MSet(ctx context.Context, data map[string]interface{}) error {
	if err := check(c); err != nil {
		return err
	}

	var pairs []interface{}
	for k, v := range data {
		pairs = append(pairs, k, v)
	}
	_, err := c.r.MSet(pairs...).Result()
	if err == redis.Nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *goredis) MGet(ctx context.Context, key []string) ([]interface{}, error) {
	if err := check(c); err != nil {
		return nil, err
	}

	val, err := c.r.MGet(key...).Result()
	if err == redis.Nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c *goredis) Remove(ctx context.Context, key string) error {
	if err := check(c); err != nil {
		return err
	}

	if _, err := c.r.Del(key).Result(); err != nil {
		return err
	}

	return nil
}

func (c *goredis) RemoveByPattern(ctx context.Context, pattern string, countPerLoop int64) error {
	if err := check(c); err != nil {
		return err
	}

	iteration := 1
	for {
		keys, _, err := c.r.Scan(0, pattern, countPerLoop).Result()
		if err != nil {
			return err
		}

		if len(keys) == 0 {
			break
		}

		if _, err := c.r.Del(keys...).Result(); err != nil {
			return err
		}

		iteration++
	}

	return nil
}

func (c *goredis) FlushDatabase(ctx context.Context) error {
	if err := check(c); err != nil {
		return err
	}

	if _, err := c.r.FlushDB().Result(); err != nil {
		return err
	}

	return nil
}

func (c *goredis) FlushAll(ctx context.Context) error {
	if err := check(c); err != nil {
		return err
	}

	if _, err := c.r.FlushAll().Result(); err != nil {
		return err
	}

	return nil
}

func check(c *goredis) error {
	if c.r == nil {
		return fmt.Errorf("redis client is not connected")
	}

	return nil
}
