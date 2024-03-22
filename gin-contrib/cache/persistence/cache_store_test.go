package persistence

import (
	"math"
	"testing"
	"time"
)

type cacheFactory func(*testing.T, time.Duration) CacheStore

// 测试典型的缓存交互
func typicalGetSet(t *testing.T, newCache cacheFactory) {
	var err error
	cache := newCache(t, time.Hour)

	value := "foo"
	if err = cache.Set("value", value, DEFAULT); err != nil {
		t.Errorf("Error setting a value: %s", err)
	}

	value = ""
	err = cache.Get("value", &value)
	if err != nil {
		t.Errorf("Error getting a value: %s", err)
	}
	if value != "foo" {
		t.Errorf("Expected to get foo back, got %s", value)
	}
}

// 测试自增自减的情况
func incrDecr(t *testing.T, newCache cacheFactory) {
	var err error
	cache := newCache(t, time.Hour)

	// 正常的增/减操作。
	if err = cache.Set("int", 10, DEFAULT); err != nil {
		t.Errorf("Error setting int: %s", err)
	}
	newValue, err := cache.Increment("int", 50)
	if err != nil {
		t.Errorf("Error incrementing int: %s", err)
	}
	if newValue != 60 {
		t.Errorf("Expected 60, was %d", newValue)
	}

	if newValue, err = cache.Decrement("int", 50); err != nil {
		t.Errorf("Error decrementing: %s", err)
	}
	if newValue != 10 {
		t.Errorf("Expected 10, was %d", newValue)
	}

	// Increment wraparound
	newValue, err = cache.Increment("int", math.MaxUint64-5)
	if err != nil {
		t.Errorf("Error wrapping around: %s", err)
	}
	if newValue != 4 {
		t.Errorf("Expected wraparound 4, got %d", newValue)
	}

	// Decrement capped at 0
	newValue, err = cache.Decrement("int", 25)
	if err != nil {
		t.Errorf("Error decrementing below 0: %s", err)
	}
	if newValue != 0 {
		t.Errorf("Expected capped at 0, got %d", newValue)
	}
}

func expiration(t *testing.T, newCache cacheFactory) {
	// memcached 不支持小于 1 秒的过期时间。
	var err error
	cache := newCache(t, time.Second)
	// Test Set w/ DEFAULT
	value := 10
	if err := cache.Set("int", value, DEFAULT); err != nil {
		t.Errorf("wrong to set cache, but got: %s", err)
	}
	time.Sleep(2 * time.Second)
	err = cache.Get("int", &value)
	if err != ErrCacheMiss {
		t.Errorf("Expected CacheMiss, but got: %s", err)
	}

	// Test Set w/ short time
	if err := cache.Set("int", value, time.Second); err != nil {
		t.Errorf("wrong to set cache, but got: %s", err)
	}
	time.Sleep(2 * time.Second)
	err = cache.Get("int", &value)
	if err != ErrCacheMiss {
		t.Errorf("Expected CacheMiss, but got: %s", err)
	}

	// 测试集，包含更长的时间。
	if err := cache.Set("int", value, time.Hour); err != nil {
		t.Errorf("wrong to set cache, but got: %s", err)
	}
	time.Sleep(2 * time.Second)
	err = cache.Get("int", &value)
	if err != nil {
		t.Errorf("Expected to get the value, but got: %s", err)
	}

	// Test Set w/ forever.
	if err := cache.Set("int", value, FOREVER); err != nil {
		t.Errorf("wrong to set cache, but got: %s", err)
	}
	time.Sleep(2 * time.Second)
	err = cache.Get("int", &value)
	if err != nil {
		t.Errorf("Expected to get the value, but got: %s", err)
	}
}

func emptyCache(t *testing.T, newCache cacheFactory) {
	var err error
	cache := newCache(t, time.Hour)

	err = cache.Get("notexist", 0)
	if err == nil {
		t.Errorf("Error expected for non-existent key")
	}
	if err != ErrCacheMiss {
		t.Errorf("Expected ErrCacheMiss for non-existent key: %s", err)
	}

	err = cache.Delete("notexist")
	if err != ErrCacheMiss {
		t.Errorf("Expected ErrCacheMiss for non-existent key: %s", err)
	}

	_, err = cache.Increment("notexist", 1)
	if err != ErrCacheMiss {
		t.Errorf("Expected cache miss incrementing non-existent key: %s", err)
	}

	_, err = cache.Decrement("notexist", 1)
	if err != ErrCacheMiss {
		t.Errorf("Expected cache miss decrementing non-existent key: %s", err)
	}
}

func testReplace(t *testing.T, newCache cacheFactory) {
	var err error
	cache := newCache(t, time.Hour)

	// 在空缓存中替换。
	if err = cache.Replace("notexist", 1, FOREVER); err != ErrNotStored && err != ErrCacheMiss {
		t.Errorf("Replace in empty cache: expected ErrNotStored or ErrCacheMiss, got: %s", err)
	}

	// 设置值为1，然后将其替换为2
	if err = cache.Set("int", 1, time.Second); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if err = cache.Replace("int", 2, time.Second); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	var i int
	if err = cache.Get("int", &i); err != nil {
		t.Errorf("Unexpected error getting a replaced item: %s", err)
	}
	if i != 2 {
		t.Errorf("Expected 2, got %d", i)
	}

	// 等待其过期并尝试用3替换（但未能成功）。
	time.Sleep(2 * time.Second)
	if err = cache.Replace("int", 3, time.Second); err != ErrNotStored && err != ErrCacheMiss {
		t.Errorf("Expected ErrNotStored or ErrCacheMiss, got: %s", err)
	}
	if err = cache.Get("int", &i); err != ErrCacheMiss {
		t.Errorf("Expected cache miss, got: %s", err)
	}
}

func testAdd(t *testing.T, newCache cacheFactory) {
	var err error
	cache := newCache(t, time.Hour)
	// Add to an empty cache.
	if err = cache.Add("int", 1, time.Second); err != nil {
		t.Errorf("Unexpected error adding to empty cache: %s", err)
	}

	// 再次尝试添加。（失败）
	if err = cache.Add("int", 2, time.Second); err != ErrNotStored {
		t.Errorf("Expected ErrNotStored adding dupe to cache: %s", err)
	}

	// 等待它过期，然后再添加。
	time.Sleep(2 * time.Second)
	if err = cache.Add("int", 3, time.Second); err != nil {
		t.Errorf("Unexpected error adding to cache: %s", err)
	}

	// 获取并验证值。
	var i int
	if err = cache.Get("int", &i); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if i != 3 {
		t.Errorf("Expected 3, got: %d", i)
	}
}
