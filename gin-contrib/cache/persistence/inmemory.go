package persistence

import (
	"reflect"
	"time"
	
	"github.com/robfig/go-cache"
)

// InMemoryStore表示具有内存持久性的缓存
type InMemoryStore struct {
	cache.Cache
}

// NewInMemoryStore返回一个InMemoryStore

// ff:
// defaultExpiration:

// ff:
// defaultExpiration:

// ff:
// defaultExpiration:

// ff:
// defaultExpiration:

// ff:
// defaultExpiration:
func NewInMemoryStore(defaultExpiration time.Duration) *InMemoryStore {
	return &InMemoryStore{*cache.New(defaultExpiration, time.Minute)}
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
func (c *InMemoryStore) Set(key string, value interface{}, expires time.Duration) error {
// 注意:go-cache理解DEFAULT和FOREVER的值
	c.Cache.Set(key, value, expires)
	return nil
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
func (c *InMemoryStore) Add(key string, value interface{}, expires time.Duration) error {
	err := c.Cache.Add(key, value, expires)
	if err == cache.ErrKeyExists {
		return ErrNotStored
	}
	return err
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
func (c *InMemoryStore) Replace(key string, value interface{}, expires time.Duration) error {
	if err := c.Cache.Replace(key, value, expires); err != nil {
		return ErrNotStored
	}
	return nil
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
func (c *InMemoryStore) Delete(key string) error {
	if found := c.Cache.Delete(key); !found {
		return ErrCacheMiss
	}
	return nil
}

// 增量(见CacheStore接口)

// ff:
// n:
// key:

// ff:
// n:
// key:

// ff:
// n:
// key:

// ff:
// n:
// key:

// ff:
// n:
// key:
func (c *InMemoryStore) Increment(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Increment(key, n)
	if err == cache.ErrCacheMiss {
		return 0, ErrCacheMiss
	}
	return newValue, err
}

// 递减(见CacheStore接口)

// ff:
// n:
// key:

// ff:
// n:
// key:

// ff:
// n:
// key:

// ff:
// n:
// key:

// ff:
// n:
// key:
func (c *InMemoryStore) Decrement(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Decrement(key, n)
	if err == cache.ErrCacheMiss {
		return 0, ErrCacheMiss
	}
	return newValue, err
}

// 刷新(见CacheStore接口)

// ff:

// ff:

// ff:

// ff:

// ff:
func (c *InMemoryStore) Flush() error {
	c.Cache.Flush()
	return nil
}
