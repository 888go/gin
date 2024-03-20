package persistence

import (
	"time"
	
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/888go/gin/gin-contrib/cache/utils"
)

// MemcachedStore 表示使用 Memcached 作为持久化存储的缓存
type MemcachedStore struct {
	*memcache.Client
	defaultExpiration time.Duration
}

// NewMemcachedStore 返回一个 MemcachedStore 实例

// ff:
// defaultExpiration:
// hostList:
func NewMemcachedStore(hostList []string, defaultExpiration time.Duration) *MemcachedStore {
	return &MemcachedStore{memcache.New(hostList...), defaultExpiration}
}

// Set（参见 CacheStore 接口）

// ff:
// expires:
// value:
// key:
func (c *MemcachedStore) Set(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Set, key, value, expires)
}

// Add （参见 CacheStore 接口）

// ff:
// expires:
// value:
// key:
func (c *MemcachedStore) Add(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Add, key, value, expires)
}

// Replace（参见 CacheStore 接口）

// ff:
// expires:
// value:
// key:
func (c *MemcachedStore) Replace(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Replace, key, value, expires)
}

// Get（参见 CacheStore 接口）

// ff:
// value:
// key:
func (c *MemcachedStore) Get(key string, value interface{}) error {
	item, err := c.Client.Get(key)
	if err != nil {
		return convertMemcacheError(err)
	}
	return utils.Deserialize(item.Value, value)
}

// Delete（参考 CacheStore 接口）

// ff:
// key:
func (c *MemcachedStore) Delete(key string) error {
	return convertMemcacheError(c.Client.Delete(key))
}

// 自增（参见 CacheStore 接口）

// ff:
// delta:
// key:
func (c *MemcachedStore) Increment(key string, delta uint64) (uint64, error) {
	newValue, err := c.Client.Increment(key, delta)
	return newValue, convertMemcacheError(err)
}

// 减量（参考 CacheStore 接口）

// ff:
// delta:
// key:
func (c *MemcachedStore) Decrement(key string, delta uint64) (uint64, error) {
	newValue, err := c.Client.Decrement(key, delta)
	return newValue, convertMemcacheError(err)
}

// Flush（参考 CacheStore 接口）

// ff:
func (c *MemcachedStore) Flush() error {
	return ErrNotSupport
}

func (c *MemcachedStore) invoke(storeFn func(*memcache.Client, *memcache.Item) error,
	key string, value interface{}, expire time.Duration) error {

	switch expire {
	case DEFAULT:
		expire = c.defaultExpiration
	case FOREVER:
		expire = time.Duration(0)
	}

	b, err := utils.Serialize(value)
	if err != nil {
		return err
	}
	return convertMemcacheError(storeFn(c.Client, &memcache.Item{
		Key:        key,
		Value:      b,
		Expiration: int32(expire / time.Second),
	}))
}

func convertMemcacheError(err error) error {
	switch err {
	case nil:
		return nil
	case memcache.ErrCacheMiss:
		return ErrCacheMiss
	case memcache.ErrNotStored:
		return ErrNotStored
	}

	return err
}
