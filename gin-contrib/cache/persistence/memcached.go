package persistence

import (
	"time"
	
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/888go/gin/gin-contrib/cache/utils"
)

// MemcachedStore表示具有memcached持久性的缓存
type MemcachedStore struct {
	*memcache.Client
	defaultExpiration time.Duration
}

// NewMemcachedStore返回一个MemcachedStore

// ff:
// defaultExpiration:
// hostList:

// ff:
// defaultExpiration:
// hostList:

// ff:
// defaultExpiration:
// hostList:

// ff:
// defaultExpiration:
// hostList:

// ff:
// defaultExpiration:
// hostList:
func NewMemcachedStore(hostList []string, defaultExpiration time.Duration) *MemcachedStore {
	return &MemcachedStore{memcache.New(hostList...), defaultExpiration}
}

// 设置方法请参见CacheStore界面

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:
func (c *MemcachedStore) Set(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Set, key, value, expires)
}

// 添加(见CacheStore接口)

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:
func (c *MemcachedStore) Add(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Add, key, value, expires)
}

// 替换(参见CacheStore接口)

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:

// ff:
// expires:
// value:
// key:
func (c *MemcachedStore) Replace(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Replace, key, value, expires)
}

// 获取(参见CacheStore接口)

// ff:
// value:
// key:

// ff:
// value:
// key:

// ff:
// value:
// key:

// ff:
// value:
// key:

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

// 删除(参见CacheStore界面)

// ff:
// key:

// ff:
// key:

// ff:
// key:

// ff:
// key:

// ff:
// key:
func (c *MemcachedStore) Delete(key string) error {
	return convertMemcacheError(c.Client.Delete(key))
}

// 增量(见CacheStore接口)

// ff:
// delta:
// key:

// ff:
// delta:
// key:

// ff:
// delta:
// key:

// ff:
// delta:
// key:

// ff:
// delta:
// key:
func (c *MemcachedStore) Increment(key string, delta uint64) (uint64, error) {
	newValue, err := c.Client.Increment(key, delta)
	return newValue, convertMemcacheError(err)
}

// 递减(见CacheStore接口)

// ff:
// delta:
// key:

// ff:
// delta:
// key:

// ff:
// delta:
// key:

// ff:
// delta:
// key:

// ff:
// delta:
// key:
func (c *MemcachedStore) Decrement(key string, delta uint64) (uint64, error) {
	newValue, err := c.Client.Decrement(key, delta)
	return newValue, convertMemcacheError(err)
}

// 刷新(见CacheStore接口)

// ff:

// ff:

// ff:

// ff:

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
