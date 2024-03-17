package persistence

import (
	"testing"
	"time"
	
	"github.com/memcachier/mc/v3"
)

// 这些测试需要在本地主机11211端口（默认设置）上运行的memcached服务
const localhost = "localhost:11211"

var newMcStore = func(t *testing.T, defaultExpiration time.Duration) CacheStore {
	mcStore := NewMemcachedBinaryStore(localhost, "", "", defaultExpiration)
	err := mcStore.Flush()
	if err == nil {
		return mcStore
	}
	t.Errorf("Failed to connect to memcached on %s with %s", localhost, err)
	t.FailNow()
	panic("")
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinary_TypicalGetSet(t *testing.T) {
	typicalGetSet(t, newMcStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinary_IncrDecr(t *testing.T) {
	incrDecr(t, newMcStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinary_Expiration(t *testing.T) {
	expiration(t, newMcStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinary_EmptyCache(t *testing.T) {
	emptyCache(t, newMcStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinary_Replace(t *testing.T) {
	testReplace(t, newMcStore)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinary_Add(t *testing.T) {
	testAdd(t, newMcStore)
}

var newMcStoreWithConfig = func(t *testing.T, defaultExpiration time.Duration) CacheStore {
	config := mc.DefaultConfig()
	config.PoolSize = 2
	mcStore := NewMemcachedBinaryStoreWithConfig(localhost, "", "", defaultExpiration, config)
	err := mcStore.Flush()
	if err == nil {
		return mcStore
	}
	t.Errorf("Failed to connect to memcached on %s with %s", localhost, err)
	t.FailNow()
	panic("")
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinaryWithConfig_TypicalGetSet(t *testing.T) {
	typicalGetSet(t, newMcStoreWithConfig)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinaryWithConfig_IncrDecr(t *testing.T) {
	incrDecr(t, newMcStoreWithConfig)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinaryWithConfig_Expiration(t *testing.T) {
	expiration(t, newMcStoreWithConfig)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinaryWithConfig_EmptyCache(t *testing.T) {
	emptyCache(t, newMcStoreWithConfig)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinaryWithConfig_Replace(t *testing.T) {
	testReplace(t, newMcStoreWithConfig)
}


// ff:
// t:

// ff:
// t:
func TestMemcachedBinaryWithConfig_Add(t *testing.T) {
	testAdd(t, newMcStoreWithConfig)
}
