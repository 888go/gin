package persistence

import (
	"testing"
	"time"
)

var newInMemoryStore = func(_ *testing.T, defaultExpiration time.Duration) CacheStore {
	return NewInMemoryStore(defaultExpiration)
}

// 测试典型的缓存交互

// ff:
// t:

// ff:
// t:
func TestInMemoryCache_TypicalGetSet(t *testing.T) {
	typicalGetSet(t, newInMemoryStore)
}


// ff:
// t:

// ff:
// t:
func TestInMemoryCache_IncrDecr(t *testing.T) {
	incrDecr(t, newInMemoryStore)
}


// ff:
// t:

// ff:
// t:
func TestInMemoryCache_Expiration(t *testing.T) {
	expiration(t, newInMemoryStore)
}


// ff:
// t:

// ff:
// t:
func TestInMemoryCache_EmptyCache(t *testing.T) {
	emptyCache(t, newInMemoryStore)
}


// ff:
// t:

// ff:
// t:
func TestInMemoryCache_Replace(t *testing.T) {
	testReplace(t, newInMemoryStore)
}


// ff:
// t:

// ff:
// t:
func TestInMemoryCache_Add(t *testing.T) {
	testAdd(t, newInMemoryStore)
}
