package cache

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
	
	"github.com/888go/gin/gin-contrib/cache/persistence"
	"github.com/888go/gin"
)

const (
	CACHE_MIDDLEWARE_KEY = "gincontrib.cache"
)

var (
	PageCachePrefix = "gincontrib.page.cache"
)

type responseCache struct {
	Status int
	Header http.Header
	Data   []byte
}

// RegisterResponseCacheGob用encoding/gob包注册responseCache类型

// ff:

// ff:

// ff:

// ff:

// ff:
func RegisterResponseCacheGob() {
	gob.Register(responseCache{})
}

type cachedWriter struct {
	gin.ResponseWriter
	status  int
	written bool
	store   persistence.CacheStore
	expire  time.Duration
	key     string
}

var _ gin.ResponseWriter = &cachedWriter{}

// CreateKey为给定字符串创建包特定的键

// ff:
// u:

// ff:
// u:

// ff:
// u:

// ff:
// u:

// ff:
// u:
func CreateKey(u string) string {
	return urlEscape(PageCachePrefix, u)
}

func urlEscape(prefix string, u string) string {
	key := url.QueryEscape(u)
	if len(key) > 200 {
		h := sha1.New()
		_, _ = io.WriteString(h, u)
		key = string(h.Sum(nil))
	}
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(":")
	buffer.WriteString(key)
	return buffer.String()
}

func newCachedWriter(store persistence.CacheStore, expire time.Duration, writer gin.ResponseWriter, key string) *cachedWriter {
	return &cachedWriter{writer, 0, false, store, expire, key}
}


// ff:
// code:

// ff:
// code:

// ff:
// code:

// ff:
// code:

// ff:
// code:
func (w *cachedWriter) WriteHeader(code int) {
	w.status = code
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}


// ff:

// ff:

// ff:

// ff:

// ff:
func (w *cachedWriter) Status() int {
	return w.ResponseWriter.Status()
}


// ff:

// ff:

// ff:

// ff:

// ff:
func (w *cachedWriter) Written() bool {
	return w.ResponseWriter.Written()
}


// ff:
// data:

// ff:
// data:

// ff:
// data:

// ff:
// data:

// ff:
// data:
func (w *cachedWriter) Write(data []byte) (int, error) {
	ret, err := w.ResponseWriter.Write(data)
	if err == nil {
		store := w.store
		var cache responseCache
		if err := store.Get(w.key, &cache); err == nil {
			data = append(cache.Data, data...)
		}

// 缓存状态码&lt的响应;300
		if w.Status() < 300 {
			val := responseCache{
				w.Status(),
				w.Header(),
				data,
			}
			err = store.Set(w.key, val, w.expire)
// 如果err != nil{需要logger}
		}
	}
	return ret, err
}


// ff:
// err:
// n:
// data:

// ff:
// err:
// n:
// data:

// ff:
// err:
// n:
// data:

// ff:
// err:
// n:
// data:

// ff:
// err:
// n:
// data:
func (w *cachedWriter) WriteString(data string) (n int, err error) {
	ret, err := w.ResponseWriter.WriteString(data)
	//cache responses with a status code < 300
	if err == nil && w.Status() < 300 {
		store := w.store
		val := responseCache{
			w.Status(),
			w.Header(),
			[]byte(data),
		}
		_ = store.Set(w.key, val, w.expire)
	}
	return ret, err
}

// 缓存的中间件

// ff:
// store:

// ff:
// store:

// ff:
// store:

// ff:
// store:

// ff:
// store:
func Cache(store *persistence.CacheStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(CACHE_MIDDLEWARE_KEY, store)
		c.Next()
	}
}


// ff:
// expire:
// store:

// ff:
// expire:
// store:

// ff:
// expire:
// store:

// ff:
// expire:
// store:

// ff:
// expire:
// store:
func SiteCache(store persistence.CacheStore, expire time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cache responseCache
		url := c.Request.URL
		key := CreateKey(url.RequestURI())
		if err := store.Get(key, &cache); err != nil {
			c.Next()
		} else {
			c.Writer.WriteHeader(cache.Status)
			for k, vals := range cache.Header {
				for _, v := range vals {
					c.Writer.Header().Set(k, v)
				}
			}
			_, _ = c.Writer.Write(cache.Data)
		}
	}
}

// CachePage装饰

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:
func CachePage(store persistence.CacheStore, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cache responseCache
		url := c.Request.URL
		key := CreateKey(url.RequestURI())
		if err := store.Get(key, &cache); err != nil {
			if err != persistence.ErrCacheMiss {
				log.Println(err.Error())
			}
// 取代的作家
			writer := newCachedWriter(store, expire, c.Writer, key)
			c.Writer = writer
			handle(c)

// 删除已终止上下文的缓存
			if c.IsAborted() {
				_ = store.Delete(key)
			}
		} else {
			c.Writer.WriteHeader(cache.Status)
			for k, vals := range cache.Header {
				for _, v := range vals {
					c.Writer.Header().Set(k, v)
				}
			}
			_, _ = c.Writer.Write(cache.Data)
		}
	}
}

// CachePageWithoutQuery增加忽略GET查询参数的功能

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:
func CachePageWithoutQuery(store persistence.CacheStore, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cache responseCache
		key := CreateKey(c.Request.URL.Path)
		if err := store.Get(key, &cache); err != nil {
			if err != persistence.ErrCacheMiss {
				log.Println(err.Error())
			}
// 取代的作家
			writer := newCachedWriter(store, expire, c.Writer, key)
			c.Writer = writer
			handle(c)
		} else {
			c.Writer.WriteHeader(cache.Status)
			for k, vals := range cache.Header {
				for _, v := range vals {
					c.Writer.Header().Set(k, v)
				}
			}
			_, _ = c.Writer.Write(cache.Data)
		}
	}
}

// CachePageAtomic装饰

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:
func CachePageAtomic(store persistence.CacheStore, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	var m sync.Mutex
	p := CachePage(store, expire, handle)
	return func(c *gin.Context) {
		m.Lock()
		defer m.Unlock()
		p(c)
	}
}


// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:

// ff:
// handle:
// expire:
// store:
func CachePageWithoutHeader(store persistence.CacheStore, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cache responseCache
		url := c.Request.URL
		key := CreateKey(url.RequestURI())
		if err := store.Get(key, &cache); err != nil {
			if err != persistence.ErrCacheMiss {
				log.Println(err.Error())
			}
// 取代的作家
			writer := newCachedWriter(store, expire, c.Writer, key)
			c.Writer = writer
			handle(c)

// 删除已终止上下文的缓存
			if c.IsAborted() {
				_ = store.Delete(key)
			}
		} else {
			c.Writer.WriteHeader(cache.Status)
			_, _ = c.Writer.Write(cache.Data)
		}
	}
}
