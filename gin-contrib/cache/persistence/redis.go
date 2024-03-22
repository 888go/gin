package persistence

import (
	"time"
	
	"github.com/888go/gin/gin-contrib/cache/utils"
	"github.com/gomodule/redigo/redis"
)

// RedisStore代表了使用Redis进行持久化的缓存
type RedisStore struct {
	pool              *redis.Pool
	defaultExpiration time.Duration
}

// NewRedisCache 返回一个 RedisStore
// 由于 redigo 目前还不支持分片/集群，因此 hostList 中目前只能包含一个主机地址
func NewRedisCache(host string, password string, defaultExpiration time.Duration) *RedisStore {
	var pool = &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			// redis协议可能应该设置为可配置的
			c, err := redis.Dial("tcp", host, redis.DialConnectTimeout(10*time.Second))
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			} else {
				// check with PING
				if _, err := c.Do("PING"); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// 自定义连接测试方法
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			// 不需要每次都检查连接。
			if time.Since(t) < 30*time.Second {
				return nil
			}

			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}
	return &RedisStore{pool, defaultExpiration}
}

// NewRedisCacheWithPool 使用提供的连接池返回一个 RedisStore
// 在 redigo 支持分片/集群之前，hostList 中将只包含一个主机地址
func NewRedisCacheWithPool(pool *redis.Pool, defaultExpiration time.Duration) *RedisStore {
	return &RedisStore{pool, defaultExpiration}
}

// Set（参见 CacheStore 接口）
func (c *RedisStore) Set(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	defer conn.Close()
	return c.invoke(conn.Do, key, value, expires)
}

// Add （参见 CacheStore 接口）
func (c *RedisStore) Add(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	defer conn.Close()
	if exists(conn, key) {
		return ErrNotStored
	}
	return c.invoke(conn.Do, key, value, expires)
}

// Replace（参见 CacheStore 接口）
func (c *RedisStore) Replace(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	defer conn.Close()
	if !exists(conn, key) {
		return ErrNotStored
	}
	err := c.invoke(conn.Do, key, value, expires)
	if value == nil {
		return ErrNotStored
	}

	return err

}

// Get（参见 CacheStore 接口）
func (c *RedisStore) Get(key string, ptrValue interface{}) error {
	conn := c.pool.Get()
	defer conn.Close()
	raw, err := conn.Do("GET", key)
	if raw == nil {
		return ErrCacheMiss
	}
	item, err := redis.Bytes(raw, err)
	if err != nil {
		return err
	}
	return utils.Deserialize(item, ptrValue)
}

func exists(conn redis.Conn, key string) bool {
	retval, _ := redis.Bool(conn.Do("EXISTS", key))
	return retval
}

// Delete（参考 CacheStore 接口）
func (c *RedisStore) Delete(key string) error {
	conn := c.pool.Get()
	defer conn.Close()
	if !exists(conn, key) {
		return ErrCacheMiss
	}
	_, err := conn.Do("DEL", key)
	return err
}

// 自增（参见 CacheStore 接口）
func (c *RedisStore) Increment(key string, delta uint64) (uint64, error) {
	conn := c.pool.Get()
	defer conn.Close()
// 根据缓存契约，在自增之前检查是否存在。
// Redis 会自动创建键，但我们不希望这样。因为我们需要自己而不是通过原生的 INCRBY（Redis 不支持自增后循环）来执行自增操作，所以我们获取值并以这种方式进行存在性检查，以尽量减少对 Redis 的调用。
	val, err := conn.Do("GET", key)
	if val == nil {
		return 0, ErrCacheMiss
	}
	if err == nil {
		currentVal, err := redis.Int64(val, nil)
		if err != nil {
			return 0, err
		}
		sum := currentVal + int64(delta)
		_, err = conn.Do("SET", key, sum)
		if err != nil {
			return 0, err
		}
		return uint64(sum), nil
	}

	return 0, err
}

// 减量（参考 CacheStore 接口）
func (c *RedisStore) Decrement(key string, delta uint64) (newValue uint64, err error) {
	conn := c.pool.Get()
	defer conn.Close()
// 按照缓存契约，在递增前检查是否存在。
// Redis 会自动创建键，但我们不希望这样，因此需要调用 exists 函数。
	if !exists(conn, key) {
		return 0, ErrCacheMiss
	}
// Decrement contract 表示你只能减到0
// 因此，我们获取当前值，如果减少的量大于该值，
// 则将值置为0
	currentVal, err := redis.Int64(conn.Do("GET", key))
	if err == nil && delta > uint64(currentVal) {
		tempint, err := redis.Int64(conn.Do("DECRBY", key, currentVal))
		return uint64(tempint), err
	}
	tempint, err := redis.Int64(conn.Do("DECRBY", key, delta))
	return uint64(tempint), err
}

// Flush（参考 CacheStore 接口）
func (c *RedisStore) Flush() error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("FLUSHALL")
	return err
}

func (c *RedisStore) invoke(f func(string, ...interface{}) (interface{}, error),
	key string, value interface{}, expires time.Duration) error {

	switch expires {
	case DEFAULT:
		expires = c.defaultExpiration
	case FOREVER:
		expires = time.Duration(0)
	}

	b, err := utils.Serialize(value)
	if err != nil {
		return err
	}

	if expires > 0 {
		_, err := f("SETEX", key, int32(expires/time.Second), b)
		return err
	}

	_, err = f("SET", key, b)
	return err

}
