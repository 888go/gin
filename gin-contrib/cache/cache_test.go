package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	
	"github.com/888go/gin/gin-contrib/cache/persistence"
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin类.X设置运行模式(gin类.X常量_运行模式_测试)
}

func TestCache(t *testing.T) {
	//TODO:unit test
}

func TestWrite(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin类.CreateTestContext(w)

	store := persistence.NewInMemoryStore(60 * time.Second)
	writer := newCachedWriter(store, time.Second*3, c.Writer, "mykey")
	c.Writer = writer

	c.Writer.WriteHeader(204)
	c.Writer.WriteHeaderNow()
	_, _ = c.Writer.Write([]byte("foo"))
	assert.Equal(t, 204, c.Writer.Status())
	assert.Equal(t, "foo", w.Body.String())
	assert.True(t, c.Writer.Written())
}

func TestCachePage(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_ping", CachePage(store, time.Second*3, func(c *gin类.Context) {
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_ping", router)
	w2 := performRequest("GET", "/cache_ping", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageExpire(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_ping", CachePage(store, time.Second, func(c *gin类.Context) {
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_ping", router)
	time.Sleep(time.Second * 2)
	w2 := performRequest("GET", "/cache_ping", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageAtomic(t *testing.T) {
// memoryDelayStore 是 InMemoryStore 的一个包装器，
// 旨在通过延迟写入模拟数据竞争（data race）场景
	store := newDelayStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/atomic", CachePageAtomic(store, time.Second*5, func(c *gin类.Context) {
		c.X输出文本(200, "OK")
	}))

	outp := make(chan string, 10)

	for i := 0; i < 5; i++ {
		go func() {
			resp := performRequest("GET", "/atomic", router)
			outp <- resp.Body.String()
		}()
	}
	time.Sleep(time.Millisecond * 500)
	for i := 0; i < 5; i++ {
		go func() {
			resp := performRequest("GET", "/atomic", router)
			outp <- resp.Body.String()
		}()
	}
	time.Sleep(time.Millisecond * 500)

	for i := 0; i < 10; i++ {
		v := <-outp
		assert.Equal(t, "OK", v)
	}
}

func TestCachePageWithoutHeader(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_ping", CachePageWithoutHeader(store, time.Second*3, func(c *gin类.Context) {
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_ping", router)
	w2 := performRequest("GET", "/cache_ping", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.NotNil(t, w1.Header()["Content-Type"])
	assert.Nil(t, w2.Header()["Content-Type"])
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithoutHeaderExpire(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_ping", CachePage(store, time.Second, func(c *gin类.Context) {
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_ping", router)
	time.Sleep(time.Second * 2)
	w2 := performRequest("GET", "/cache_ping", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.NotNil(t, w1.Header()["Content-Type"])
	assert.NotNil(t, w2.Header()["Content-Type"])
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCacheHtmlFile(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X加载HTML模板文件("example/template.html")
	router.X绑定GET("/cache_html", CachePage(store, time.Second*3, func(c *gin类.Context) {
		c.X输出html模板(http.StatusOK, "template.html", gin类.H{"values": fmt.Sprint(time.Now().UnixNano())})
	}))

	w1 := performRequest("GET", "/cache_html", router)
	w2 := performRequest("GET", "/cache_html", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCacheHtmlFileExpire(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X加载HTML模板文件("example/template.html")
	router.X绑定GET("/cache_html", CachePage(store, time.Second*1, func(c *gin类.Context) {
		c.X输出html模板(http.StatusOK, "template.html", gin类.H{"values": fmt.Sprint(time.Now().UnixNano())})
	}))

	w1 := performRequest("GET", "/cache_html", router)
	time.Sleep(time.Second * 2)
	w2 := performRequest("GET", "/cache_html", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageAborted(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_aborted", CachePage(store, time.Second*3, func(c *gin类.Context) {
		c.X停止并带状态码且返回JSON(200, map[string]int64{"time": time.Now().UnixNano()})
	}))

	w1 := performRequest("GET", "/cache_aborted", router)
	time.Sleep(time.Millisecond * 500)
	w2 := performRequest("GET", "/cache_aborted", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePage400(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_400", CachePage(store, time.Second*3, func(c *gin类.Context) {
		c.X输出文本(400, fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_400", router)
	time.Sleep(time.Millisecond * 500)
	w2 := performRequest("GET", "/cache_400", router)

	assert.Equal(t, 400, w1.Code)
	assert.Equal(t, 400, w2.Code)
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithoutHeaderAborted(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_aborted", CachePage(store, time.Second*3, func(c *gin类.Context) {
		c.X停止并带状态码且返回JSON(200, map[string]int64{"time": time.Now().UnixNano()})
	}))

	w1 := performRequest("GET", "/cache_aborted", router)
	time.Sleep(time.Millisecond * 500)
	w2 := performRequest("GET", "/cache_aborted", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.NotNil(t, w1.Header()["Content-Type"])
	assert.NotNil(t, w2.Header()["Content-Type"])
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithoutHeader400(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_400", CachePage(store, time.Second*3, func(c *gin类.Context) {
		c.X输出文本(400, fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_400", router)
	time.Sleep(time.Millisecond * 500)
	w2 := performRequest("GET", "/cache_400", router)

	assert.Equal(t, 400, w1.Code)
	assert.Equal(t, 400, w2.Code)
	assert.NotNil(t, w1.Header()["Content-Type"])
	assert.NotNil(t, w2.Header()["Content-Type"])
	assert.NotEqual(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageStatus207(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_207", CachePage(store, time.Second*3, func(c *gin类.Context) {
		c.X输出文本(207, fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_207", router)
	time.Sleep(time.Millisecond * 500)
	w2 := performRequest("GET", "/cache_207", router)

	assert.Equal(t, 207, w1.Code)
	assert.Equal(t, 207, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestCachePageWithoutQuery(t *testing.T) {
	store := persistence.NewInMemoryStore(60 * time.Second)

	router := gin类.X创建()
	router.X绑定GET("/cache_without_query", CachePageWithoutQuery(store, time.Second*3, func(c *gin类.Context) {
		c.X输出文本(200, "pong "+fmt.Sprint(time.Now().UnixNano()))
	}))

	w1 := performRequest("GET", "/cache_without_query?foo=1", router)
	w2 := performRequest("GET", "/cache_without_query?foo=2", router)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, w1.Body.String(), w2.Body.String())
}

func TestRegisterResponseCacheGob(t *testing.T) {
	RegisterResponseCacheGob()
	r := responseCache{Status: 200, Data: []byte("test")}
	mCache := new(bytes.Buffer)
	encCache := gob.NewEncoder(mCache)
	err := encCache.Encode(r)
	assert.Nil(t, err)

	var decodedResp responseCache
	pCache := bytes.NewBuffer(mCache.Bytes())
	decCache := gob.NewDecoder(pCache)
	err = decCache.Decode(&decodedResp)
	assert.Nil(t, err)

}
func performRequest(method, target string, router *gin类.Engine) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}

type memoryDelayStore struct {
	*persistence.InMemoryStore
}

func newDelayStore(defaultExpiration time.Duration) *memoryDelayStore {
	v := &memoryDelayStore{}
	v.InMemoryStore = persistence.NewInMemoryStore(defaultExpiration)
	return v
}

func (c *memoryDelayStore) Set(key string, value interface{}, expires time.Duration) error {
	time.Sleep(time.Millisecond * 3)
	return c.InMemoryStore.Set(key, value, expires)
}

func (c *memoryDelayStore) Add(key string, value interface{}, expires time.Duration) error {
	time.Sleep(time.Millisecond * 3)
	return c.InMemoryStore.Add(key, value, expires)
}
