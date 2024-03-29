package persistence

import (
	"testing"
	"time"
)

var newInMemoryStore = func(_ *testing.T, defaultExpiration time.Duration) CacheStore {
	return NewInMemoryStore(defaultExpiration)
}

// 测试典型的缓存交互
func TestInMemoryCache_TypicalGetSet(t *testing.T) {
	typicalGetSet(t, newInMemoryStore)
}

func TestInMemoryCache_IncrDecr(t *testing.T) {
	incrDecr(t, newInMemoryStore)
}

func TestInMemoryCache_Expiration(t *testing.T) {
	expiration(t, newInMemoryStore)
}

func TestInMemoryCache_EmptyCache(t *testing.T) {
	emptyCache(t, newInMemoryStore)
}

func TestInMemoryCache_Replace(t *testing.T) {
	testReplace(t, newInMemoryStore)
}

func TestInMemoryCache_Add(t *testing.T) {
	testAdd(t, newInMemoryStore)
}
