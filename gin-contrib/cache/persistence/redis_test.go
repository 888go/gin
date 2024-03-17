package persistence

import (
	"net"
	"testing"
	"time"
)

// 这些测试需要在本地主机6379端口（默认端口）上运行的Redis服务器
const redisTestServer = "localhost:6379"

var newRedisStore = func(t *testing.T, defaultExpiration time.Duration) CacheStore {
	c, err := net.Dial("tcp", redisTestServer)
	if err == nil {
		_, _ = c.Write([]byte("flush_all\r\n"))
		c.Close()
		redisCache := NewRedisCache(redisTestServer, "", defaultExpiration)
		redisCache.Flush()
		return redisCache
	}
	t.Errorf("couldn't connect to redis on %s", redisTestServer)
	t.FailNow()
	panic("")
}


// ff:
// t:

// ff:
// t:
func TestRedisCache_TypicalGetSet(t *testing.T) {
	typicalGetSet(t, newRedisStore)
}


// ff:
// t:

// ff:
// t:
func TestRedisCache_IncrDecr(t *testing.T) {
	incrDecr(t, newRedisStore)
}


// ff:
// t:

// ff:
// t:
func TestRedisCache_Expiration(t *testing.T) {
	expiration(t, newRedisStore)
}


// ff:
// t:

// ff:
// t:
func TestRedisCache_EmptyCache(t *testing.T) {
	emptyCache(t, newRedisStore)
}


// ff:
// t:

// ff:
// t:
func TestRedisCache_Replace(t *testing.T) {
	testReplace(t, newRedisStore)
}


// ff:
// t:

// ff:
// t:
func TestRedisCache_Add(t *testing.T) {
	testAdd(t, newRedisStore)
}
