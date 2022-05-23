package leafViper

import (
	"context"
	"fmt"
	leafCache "github.com/paulusrobin/leaf-utilities/cache/cache"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	leafFeatureFlag "github.com/paulusrobin/leaf-utilities/featureFlag/featureFlag"
)

type (
	backup struct {
		// Redis is used for backup if feature flag remote server is not working
		redis leafCache.Redis
	}
)

func (b backup) Get(key string) (map[string]interface{}, error) {
	var result map[string]interface{}
	if b.redis != nil {
		data, err := b.redis.HGetAll(context.TODO(), key)
		if err != nil {
			return nil, err
		}
		result = make(map[string]interface{}, len(data))

		for k, v := range data {
			temp := map[string]interface{}{}
			result[k] = v
			if json.Unmarshal([]byte(v), &temp) == nil {
				result[k] = temp
			}
		}
	}

	return result, nil
}

func (b backup) Set(key string, data map[string]interface{}) error {
	if b.redis != nil {
		if err := b.redis.HMSet(context.TODO(), key, marshallMap(data)); err != nil {
			return err
		}
	}
	return fmt.Errorf("no backup provided")
}

func NewRedisBackup(redis leafCache.Redis) (leafFeatureFlag.Backup, error) {
	return &backup{
		redis: redis,
	}, nil
}
