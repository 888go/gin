package persistence

import (
	"net"
	"testing"
	"time"
)

// 这些测试需要在本地主机11211端口（默认设置）上运行的memcached服务
const testServer = "localhost:11211"

var newMemcachedStore = func(t *testing.T, defaultExpiration time.Duration) CacheStore {
	c, err := net.Dial("tcp", testServer)
	if err == nil {
		_, _ = c.Write([]byte("flush_all\r\n"))
		c.Close()
		return NewMemcachedStore([]string{testServer}, defaultExpiration)
	}
	t.Errorf("couldn't connect to memcached on %s", testServer)
	t.FailNow()
	panic("")
}


// ff:
// t:

// ff:
// t:
func TestMemcachedCache_TypicalGetSet(t *testing.T) {
	typicalGetSet(t, newMemcachedStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedCache_IncrDecr(t *testing.T) {
	incrDecr(t, newMemcachedStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedCache_Expiration(t *testing.T) {
	expiration(t, newMemcachedStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedCache_EmptyCache(t *testing.T) {
	emptyCache(t, newMemcachedStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedCache_Replace(t *testing.T) {
	testReplace(t, newMemcachedStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedCache_Add(t *testing.T) {
	testAdd(t, newMemcachedStore)
}
