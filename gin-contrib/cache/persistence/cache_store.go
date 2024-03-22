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

// CacheStore 是缓存后端的接口
type CacheStore interface {
// Get 从缓存中检索一个项目。返回该项目或nil，以及一个布尔值，表示是否找到了该键。
	Get(key string, value interface{}) error

	// Set 将一个项设置到缓存中，替换任何已存在的项。
	Set(key string, value interface{}, expire time.Duration) error

// Add 将一个项添加到缓存中，但只有在给定键下尚未存在项，或者已存在的项已过期时才会添加。否则返回错误。
	Add(key string, value interface{}, expire time.Duration) error

// Replace仅当缓存键已存在时设置新的值。如果不存在，则返回错误。
	Replace(key string, data interface{}, expire time.Duration) error

	// Delete 从缓存中移除一个项目。如果键不在缓存中，则不执行任何操作。
	Delete(key string) error

	// Increment 函数对一个实数进行增加操作，并在值不是实数时返回错误
	Increment(key string, data uint64) (uint64, error)

	// Decrement 函数对一个实数进行减一操作，如果该值不是实数，则返回错误
	Decrement(key string, data uint64) (uint64, error)

	// Flush 清空缓存中的所有项。
	Flush() error
}
