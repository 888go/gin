package persistence

import (
	"time"
	
	"github.com/888go/gin/gin-contrib/cache/utils"
	"github.com/memcachier/mc/v3"
)

// MemcachedBinaryStore代表使用二进制协议的缓存，其持久化机制为Memcached
type MemcachedBinaryStore struct {
	*mc.Client
	defaultExpiration time.Duration
}

// NewMemcachedBinaryStore 返回一个 MemcachedBinaryStore
func NewMemcachedBinaryStore(hostList, username, password string, defaultExpiration time.Duration) *MemcachedBinaryStore {
	return &MemcachedBinaryStore{mc.NewMC(hostList, username, password), defaultExpiration}
}

// NewMemcachedBinaryStoreWithConfig 使用提供的配置返回一个 MemcachedBinaryStore 实例
func NewMemcachedBinaryStoreWithConfig(hostList, username, password string, defaultExpiration time.Duration, config *mc.Config) *MemcachedBinaryStore {
	return &MemcachedBinaryStore{mc.NewMCwithConfig(hostList, username, password, config), defaultExpiration}
}

// Set（参见 CacheStore 接口）
func (s *MemcachedBinaryStore) Set(key string, value interface{}, expires time.Duration) error {
	exp := s.getExpiration(expires)
	b, err := utils.Serialize(value)
	if err != nil {
		return err
	}
	_, err = s.Client.Set(key, string(b), 0, exp, 0)
	return convertMcError(err)
}

// Add （参见 CacheStore 接口）
func (s *MemcachedBinaryStore) Add(key string, value interface{}, expires time.Duration) error {
	exp := s.getExpiration(expires)
	b, err := utils.Serialize(value)
	if err != nil {
		return err
	}
	_, err = s.Client.Add(key, string(b), 0, exp)
	return convertMcError(err)
}

// Replace（参见 CacheStore 接口）
func (s *MemcachedBinaryStore) Replace(key string, value interface{}, expires time.Duration) error {
	exp := s.getExpiration(expires)
	b, err := utils.Serialize(value)
	if err != nil {
		return err
	}
	_, err = s.Client.Replace(key, string(b), 0, exp, 0)
	return convertMcError(err)
}

// Get（参见 CacheStore 接口）
func (s *MemcachedBinaryStore) Get(key string, value interface{}) error {
	val, _, _, err := s.Client.Get(key)
	if err != nil {
		return convertMcError(err)
	}
	return utils.Deserialize([]byte(val), value)
}

// Delete（参考 CacheStore 接口）
func (s *MemcachedBinaryStore) Delete(key string) error {
	return convertMcError(s.Client.Del(key))
}

// 自增（参见 CacheStore 接口）
func (s *MemcachedBinaryStore) Increment(key string, delta uint64) (uint64, error) {
	n, _, err := s.Client.Incr(key, delta, 0, 0xffffffff, 0)
	return n, convertMcError(err)
}

// 减量（参考 CacheStore 接口）
func (s *MemcachedBinaryStore) Decrement(key string, delta uint64) (uint64, error) {
	n, _, err := s.Client.Decr(key, delta, 0, 0xffffffff, 0)
	return n, convertMcError(err)
}

// Flush（参考 CacheStore 接口）
func (s *MemcachedBinaryStore) Flush() error {
	return convertMcError(s.Client.Flush(0))
}

// getExpiration 将gin-contrib/cache中以time.Duration形式表示的过期时间转换为有效的memcached过期时间，
// 过期时间要么表示为秒（小于30天），要么表示为Unix时间戳（大于30天）
func (s *MemcachedBinaryStore) getExpiration(expires time.Duration) uint32 {
	switch expires {
	case DEFAULT:
		expires = s.defaultExpiration
	case FOREVER:
		expires = time.Duration(0)
	}
	exp := uint32(expires.Seconds())
	if exp > 60*60*24*30 { // > 30 days
		exp += uint32(time.Now().Unix())
	}
	return exp
}

func convertMcError(err error) error {
	switch err {
	case nil:
		return nil
	case mc.ErrNotFound:
		return ErrCacheMiss
	case mc.ErrValueNotStored:
		return ErrNotStored
	case mc.ErrKeyExists:
		return ErrNotStored
	}
	return err
}
