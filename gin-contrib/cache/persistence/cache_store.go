package persistence

import (
	"errors"
	"time"
)

const (
	DEFAULT = time.Duration(0)
	FOREVER = time.Duration(-1)
)

var (
	PageCachePrefix = "gincontrib.page.cache"
	ErrCacheMiss    = errors.New("cache: key not found.")
	ErrNotStored    = errors.New("cache: not stored.")
	ErrNotSupport   = errors.New("cache: not support.")
)

// CacheStore是缓存后端接口
type CacheStore interface {
// Get从缓存中检索项
// 返回项或nil，以及指示是否找到键的bool值
	Get(key string, value interface{}) error

// Set将项设置到缓存，替换任何现有项
	Set(key string, value interface{}, expire time.Duration) error

// Add仅在给定键的项不存在或现有项已过期时向缓存添加项
// 否则返回错误
	Add(key string, value interface{}, expire time.Duration) error

// Replace仅在缓存键已经存在时才为该键设置新值
// 如果没有，则返回错误
	Replace(key string, data interface{}, expire time.Duration) error

// Delete从缓存中删除项
// 如果键不在缓存中，则不执行任何操作
	Delete(key string) error

// Increment对实数递增，如果值不是实数则返回error
	Increment(key string, data uint64) (uint64, error)

// 递减一个实数，如果值不是实数，则返回错误
	Decrement(key string, data uint64) (uint64, error)

// 刷新从缓存中删除所有项
	Flush() error
}
