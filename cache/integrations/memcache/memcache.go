package leafMemcache

import (
	"context"
	"fmt"
	"github.com/enricodg/leaf-utilities/cache/cache"
	"reflect"
	"sync"
	"time"
)

type (
	cache struct {
		option option
		data   map[string]interface{}
		timer  map[string]*time.Timer
		queue  []string
		size   uintptr
		sync.Mutex
	}
)

func (c *cache) Ping(ctx context.Context) error {
	c.option.logger.StandardLogger().Debug("[MEMORY-CACHE] ping success")
	return nil
}

func (c *cache) Close() error {
	c.option.logger.StandardLogger().Info("[MEMORY-CACHE] closing resource")
	return nil
}

func (c *cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	c.Lock()
	defer c.Unlock()

	newItemSize := c.getSize(value)
	if newItemSize > uintptr(c.option.maxEntrySize) {
		return fmt.Errorf("entry value too large")
	}

	for {
		if c.size+newItemSize > uintptr(c.option.maxEntriesInWindow) {
			c.forceRemove(c.queue[0], fmt.Errorf("window too large").Error())
		} else {
			break
		}
	}
	if len(c.data) >= c.option.maxEntriesKey {
		c.forceRemove(c.queue[0], fmt.Errorf("too many keys invcoked").Error())
	}

	c.data[key] = value
	c.queue = append(c.queue, key)
	c.size += newItemSize

	if _, ok := c.timer[key]; !ok {
		timer := time.AfterFunc(ttl, func() {
			_ = c.Remove(context.Background(), key)
			c.timer[key].Stop()
			delete(c.timer, key)
		})
		c.timer[key] = timer
	} else {
		c.timer[key].Reset(ttl)
	}

	return nil
}

func (c *cache) Get(ctx context.Context, key string) (interface{}, error) {
	c.Lock()
	defer c.Unlock()
	if value, ok := c.data[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("key %s is not found", key)
}

func (c *cache) Remove(ctx context.Context, key string) error {
	data, err := c.Get(ctx, key)
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()
	delete(c.data, key)
	c.removeQueue()

	if c.option.onRemove != nil {
		c.option.onRemove(key, data)
	}
	return nil
}

func (c *cache) Truncate(ctx context.Context) error {
	c.Lock()
	defer c.Unlock()

	for key := range c.data {
		delete(c.data, key)
		c.removeQueue()
	}
	return nil
}

func (c *cache) Len(ctx context.Context) int {
	c.Lock()
	defer c.Unlock()
	return len(c.data)
}

func (c *cache) Size(ctx context.Context) uintptr {
	c.Lock()
	defer c.Unlock()
	return c.size
}

func (c *cache) Keys(ctx context.Context) []string {
	return c.queue
}

func New(options ...Option) leafCache.Memcache {
	o := defaultOption()

	for _, opt := range options {
		opt.Apply(&o)
	}
	return &cache{
		option: o,
		data:   make(map[string]interface{}),
		timer:  make(map[string]*time.Timer),
		queue:  make([]string, 0),
		size:   uintptr(0),
		Mutex:  sync.Mutex{},
	}
}

func (c *cache) forceRemove(key string, reason string) {
	delete(c.data, key)
	c.removeQueue()

	if c.option.onRemoveWithReason != nil {
		c.option.onRemoveWithReason(key, reason)
	}
}

func (c *cache) removeQueue() {
	if len(c.queue) > 1 {
		sizeItemToRemove := c.getSize(c.queue[0])
		c.queue = c.queue[1:len(c.queue)]
		c.size -= sizeItemToRemove
	} else {
		c.queue = make([]string, 0)
		c.size = 0
	}
}

func (c *cache) getSize(T interface{}) uintptr {
	return reflect.TypeOf(T).Size()
}
