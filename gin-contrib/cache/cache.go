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

// RegisterResponseCacheGob 注册 responseCache 类型到 encoding/gob 包中
func RegisterResponseCacheGob() {
	gob.Register(responseCache{})
}

type cachedWriter struct {
	gin类.ResponseWriter
	status  int
	written bool
	store   persistence.CacheStore
	expire  time.Duration
	key     string
}

var _ gin类.ResponseWriter = &cachedWriter{}

// CreateKey 为给定的字符串创建一个特定于包的密钥
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

func newCachedWriter(store persistence.CacheStore, expire time.Duration, writer gin类.ResponseWriter, key string) *cachedWriter {
	return &cachedWriter{writer, 0, false, store, expire, key}
}

func (w *cachedWriter) WriteHeader(code int) {
	w.status = code
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

func (w *cachedWriter) Status() int {
	return w.ResponseWriter.Status()
}

func (w *cachedWriter) Written() bool {
	return w.ResponseWriter.Written()
}

func (w *cachedWriter) Write(data []byte) (int, error) {
	ret, err := w.ResponseWriter.Write(data)
	if err == nil {
		store := w.store
		var cache responseCache
		if err := store.Get(w.key, &cache); err == nil {
			data = append(cache.Data, data...)
		}

		//缓存状态码小于300的响应
		if w.Status() < 300 {
			val := responseCache{
				w.Status(),
				w.Header(),
				data,
			}
			err = store.Set(w.key, val, w.expire)
// 如果err不为nil {
//   // 需要使用日志记录器
// }
		}
	}
	return ret, err
}

func (w *cachedWriter) WriteString(data string) (n int, err error) {
	ret, err := w.ResponseWriter.WriteString(data)
	//缓存状态码小于300的响应
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

// Cache Middleware
func Cache(store *persistence.CacheStore) gin类.HandlerFunc {
	return func(c *gin类.Context) {
		c.X设置值(CACHE_MIDDLEWARE_KEY, store)
		c.X中间件继续()
	}
}

func SiteCache(store persistence.CacheStore, expire time.Duration) gin类.HandlerFunc {
	return func(c *gin类.Context) {
		var cache responseCache
		url := c.X请求.URL
		key := CreateKey(url.RequestURI())
		if err := store.Get(key, &cache); err != nil {
			c.X中间件继续()
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

// CachePage Decorator
func CachePage(store persistence.CacheStore, expire time.Duration, handle gin类.HandlerFunc) gin类.HandlerFunc {
	return func(c *gin类.Context) {
		var cache responseCache
		url := c.X请求.URL
		key := CreateKey(url.RequestURI())
		if err := store.Get(key, &cache); err != nil {
			if err != persistence.ErrCacheMiss {
				log.Println(err.Error())
			}
			// replace writer
			writer := newCachedWriter(store, expire, c.Writer, key)
			c.Writer = writer
			handle(c)

			// 清除已中止上下文的缓存
			if c.X是否已停止() {
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

// CachePageWithoutQuery 添加忽略GET请求参数的能力。
func CachePageWithoutQuery(store persistence.CacheStore, expire time.Duration, handle gin类.HandlerFunc) gin类.HandlerFunc {
	return func(c *gin类.Context) {
		var cache responseCache
		key := CreateKey(c.X请求.URL.Path)
		if err := store.Get(key, &cache); err != nil {
			if err != persistence.ErrCacheMiss {
				log.Println(err.Error())
			}
			// replace writer
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

// CachePageAtomic 装饰器
func CachePageAtomic(store persistence.CacheStore, expire time.Duration, handle gin类.HandlerFunc) gin类.HandlerFunc {
	var m sync.Mutex
	p := CachePage(store, expire, handle)
	return func(c *gin类.Context) {
		m.Lock()
		defer m.Unlock()
		p(c)
	}
}

func CachePageWithoutHeader(store persistence.CacheStore, expire time.Duration, handle gin类.HandlerFunc) gin类.HandlerFunc {
	return func(c *gin类.Context) {
		var cache responseCache
		url := c.X请求.URL
		key := CreateKey(url.RequestURI())
		if err := store.Get(key, &cache); err != nil {
			if err != persistence.ErrCacheMiss {
				log.Println(err.Error())
			}
			// replace writer
			writer := newCachedWriter(store, expire, c.Writer, key)
			c.Writer = writer
			handle(c)

			// 清除已中止上下文的缓存
			if c.X是否已停止() {
				_ = store.Delete(key)
			}
		} else {
			c.Writer.WriteHeader(cache.Status)
			_, _ = c.Writer.Write(cache.Data)
		}
	}
}
