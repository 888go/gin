package persistence

import (
	"reflect"
	"time"
	
	"github.com/robfig/go-cache"
)

// InMemoryStore 表示使用内存持久化的缓存
type InMemoryStore struct {
	cache.Cache
}

// NewInMemoryStore 返回一个 InMemoryStore
func NewInMemoryStore(defaultExpiration time.Duration) *InMemoryStore {
	return &InMemoryStore{*cache.New(defaultExpiration, time.Minute)}
}

// Get（参见 CacheStore 接口）
func (c *InMemoryStore) Get(key string, value interface{}) error {
	val, found := c.Cache.Get(key)
	if !found {
		return ErrCacheMiss
	}

	v := reflect.ValueOf(value)
	if v.Type().Kind() == reflect.Ptr && v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))
		return nil
	}
	return ErrNotStored
}

// Set（参见 CacheStore 接口）
func (c *InMemoryStore) Set(key string, value interface{}, expires time.Duration) error {
	// 注释：go-cache 能够识别 DEFAULT 和 FOREVER 的值
	c.Cache.Set(key, value, expires)
	return nil
}

// Add （参见 CacheStore 接口）
func (c *InMemoryStore) Add(key string, value interface{}, expires time.Duration) error {
	err := c.Cache.Add(key, value, expires)
	if err == cache.ErrKeyExists {
		return ErrNotStored
	}
	return err
}

// Replace（参见 CacheStore 接口）
func (c *InMemoryStore) Replace(key string, value interface{}, expires time.Duration) error {
	if err := c.Cache.Replace(key, value, expires); err != nil {
		return ErrNotStored
	}
	return nil
}

// Delete（参考 CacheStore 接口）
func (c *InMemoryStore) Delete(key string) error {
	if found := c.Cache.Delete(key); !found {
		return ErrCacheMiss
	}
	return nil
}

// 自增（参见 CacheStore 接口）
func (c *InMemoryStore) Increment(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Increment(key, n)
	if err == cache.ErrCacheMiss {
		return 0, ErrCacheMiss
	}
	return newValue, err
}

// 减量（参考 CacheStore 接口）
func (c *InMemoryStore) Decrement(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Decrement(key, n)
	if err == cache.ErrCacheMiss {
		return 0, ErrCacheMiss
	}
	return newValue, err
}

// Flush（参考 CacheStore 接口）
func (c *InMemoryStore) Flush() error {
	c.Cache.Flush()
	return nil
}
