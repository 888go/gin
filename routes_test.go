// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

type header struct {
	Key   string
	Value string
}

// PerformRequest 用于测试 gin 路由器。
func PerformRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func testRouteOK(method string, t *testing.T) {
	passed := false
	passedAny := false
	r := X创建()
	r.X绑定Any("/test2", func(c *Context) {
		passedAny = true
	})
	r.X绑定(method, "/test", func(c *Context) {
		passed = true
	})

	w := PerformRequest(r, method, "/test")
	assert.True(t, passed)
	assert.Equal(t, http.StatusOK, w.Code)

	PerformRequest(r, method, "/test2")
	assert.True(t, passedAny)
}

// TestSingleRouteOK 测试 POST 路由是否被正确调用。
func testRouteNotOK(method string, t *testing.T) {
	passed := false
	router := X创建()
	router.X绑定(method, "/test_2", func(c *Context) {
		passed = true
	})

	w := PerformRequest(router, method, "/test")

	assert.False(t, passed)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestSingleRouteOK 测试 POST 路由是否被正确调用。
func testRouteNotOK2(method string, t *testing.T) {
	passed := false
	router := X创建()
	router.HandleMethodNotAllowed = true
	var methodRoute string
	if method == http.MethodPost {
		methodRoute = http.MethodGet
	} else {
		methodRoute = http.MethodPost
	}
	router.X绑定(methodRoute, "/test", func(c *Context) {
		passed = true
	})

	w := PerformRequest(router, method, "/test")

	assert.False(t, passed)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestRouterMethod(t *testing.T) {
	router := X创建()
	router.X绑定PUT("/hey2", func(c *Context) {
		c.X输出文本(http.StatusOK, "sup2")
	})

	router.X绑定PUT("/hey", func(c *Context) {
		c.X输出文本(http.StatusOK, "called")
	})

	router.X绑定PUT("/hey3", func(c *Context) {
		c.X输出文本(http.StatusOK, "sup3")
	})

	w := PerformRequest(router, http.MethodPut, "/hey")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "called", w.Body.String())
}

func TestRouterGroupRouteOK(t *testing.T) {
	testRouteOK(http.MethodGet, t)
	testRouteOK(http.MethodPost, t)
	testRouteOK(http.MethodPut, t)
	testRouteOK(http.MethodPatch, t)
	testRouteOK(http.MethodHead, t)
	testRouteOK(http.MethodOptions, t)
	testRouteOK(http.MethodDelete, t)
	testRouteOK(http.MethodConnect, t)
	testRouteOK(http.MethodTrace, t)
}

func TestRouteNotOK(t *testing.T) {
	testRouteNotOK(http.MethodGet, t)
	testRouteNotOK(http.MethodPost, t)
	testRouteNotOK(http.MethodPut, t)
	testRouteNotOK(http.MethodPatch, t)
	testRouteNotOK(http.MethodHead, t)
	testRouteNotOK(http.MethodOptions, t)
	testRouteNotOK(http.MethodDelete, t)
	testRouteNotOK(http.MethodConnect, t)
	testRouteNotOK(http.MethodTrace, t)
}

func TestRouteNotOK2(t *testing.T) {
	testRouteNotOK2(http.MethodGet, t)
	testRouteNotOK2(http.MethodPost, t)
	testRouteNotOK2(http.MethodPut, t)
	testRouteNotOK2(http.MethodPatch, t)
	testRouteNotOK2(http.MethodHead, t)
	testRouteNotOK2(http.MethodOptions, t)
	testRouteNotOK2(http.MethodDelete, t)
	testRouteNotOK2(http.MethodConnect, t)
	testRouteNotOK2(http.MethodTrace, t)
}

func TestRouteRedirectTrailingSlash(t *testing.T) {
	router := X创建()
	router.X重定向固定路径 = false
	router.X重定向尾部斜杠 = true
	router.X绑定GET("/path", func(c *Context) {})
	router.X绑定GET("/path2/", func(c *Context) {})
	router.X绑定POST("/path3", func(c *Context) {})
	router.X绑定PUT("/path4/", func(c *Context) {})

	w := PerformRequest(router, http.MethodGet, "/path/")
	assert.Equal(t, "/path", w.Header().Get("Location"))
	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path2")
	assert.Equal(t, "/path2/", w.Header().Get("Location"))
	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	w = PerformRequest(router, http.MethodPost, "/path3/")
	assert.Equal(t, "/path3", w.Header().Get("Location"))
	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)

	w = PerformRequest(router, http.MethodPut, "/path4")
	assert.Equal(t, "/path4/", w.Header().Get("Location"))
	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path")
	assert.Equal(t, http.StatusOK, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path2/")
	assert.Equal(t, http.StatusOK, w.Code)

	w = PerformRequest(router, http.MethodPost, "/path3")
	assert.Equal(t, http.StatusOK, w.Code)

	w = PerformRequest(router, http.MethodPut, "/path4/")
	assert.Equal(t, http.StatusOK, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path2", header{Key: "X-Forwarded-Prefix", Value: "/api"})
	assert.Equal(t, "/api/path2/", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path2/", header{Key: "X-Forwarded-Prefix", Value: "/api/"})
	assert.Equal(t, 200, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path/", header{Key: "X-Forwarded-Prefix", Value: "../../api#?"})
	assert.Equal(t, "/api/path", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path/", header{Key: "X-Forwarded-Prefix", Value: "../../api"})
	assert.Equal(t, "/api/path", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path2", header{Key: "X-Forwarded-Prefix", Value: "../../api"})
	assert.Equal(t, "/api/path2/", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path2", header{Key: "X-Forwarded-Prefix", Value: "/../../api"})
	assert.Equal(t, "/api/path2/", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path/", header{Key: "X-Forwarded-Prefix", Value: "api/../../"})
	assert.Equal(t, "//path", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path/", header{Key: "X-Forwarded-Prefix", Value: "api/../../../"})
	assert.Equal(t, "/path", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path2", header{Key: "X-Forwarded-Prefix", Value: "../../gin-gonic.com"})
	assert.Equal(t, "/gin-goniccom/path2/", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path2", header{Key: "X-Forwarded-Prefix", Value: "/../../gin-gonic.com"})
	assert.Equal(t, "/gin-goniccom/path2/", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path/", header{Key: "X-Forwarded-Prefix", Value: "https://gin-gonic.com/#"})
	assert.Equal(t, "https/gin-goniccom/https/gin-goniccom/path", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path/", header{Key: "X-Forwarded-Prefix", Value: "#api"})
	assert.Equal(t, "api/api/path", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path/", header{Key: "X-Forwarded-Prefix", Value: "/nor-mal/#?a=1"})
	assert.Equal(t, "/nor-mal/a1/path", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path/", header{Key: "X-Forwarded-Prefix", Value: "/nor-mal/%2e%2e/"})
	assert.Equal(t, "/nor-mal/2e2e/path", w.Header().Get("Location"))
	assert.Equal(t, 301, w.Code)

	router.X重定向尾部斜杠 = false

	w = PerformRequest(router, http.MethodGet, "/path/")
	assert.Equal(t, http.StatusNotFound, w.Code)
	w = PerformRequest(router, http.MethodGet, "/path2")
	assert.Equal(t, http.StatusNotFound, w.Code)
	w = PerformRequest(router, http.MethodPost, "/path3/")
	assert.Equal(t, http.StatusNotFound, w.Code)
	w = PerformRequest(router, http.MethodPut, "/path4")
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestRouteRedirectFixedPath(t *testing.T) {
	router := X创建()
	router.X重定向固定路径 = true
	router.X重定向尾部斜杠 = false

	router.X绑定GET("/path", func(c *Context) {})
	router.X绑定GET("/Path2", func(c *Context) {})
	router.X绑定POST("/PATH3", func(c *Context) {})
	router.X绑定POST("/Path4/", func(c *Context) {})

	w := PerformRequest(router, http.MethodGet, "/PATH")
	assert.Equal(t, "/path", w.Header().Get("Location"))
	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	w = PerformRequest(router, http.MethodGet, "/path2")
	assert.Equal(t, "/Path2", w.Header().Get("Location"))
	assert.Equal(t, http.StatusMovedPermanently, w.Code)

	w = PerformRequest(router, http.MethodPost, "/path3")
	assert.Equal(t, "/PATH3", w.Header().Get("Location"))
	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)

	w = PerformRequest(router, http.MethodPost, "/path4")
	assert.Equal(t, "/Path4/", w.Header().Get("Location"))
	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
}

// TestContextParamsGet 测试从URL中解析参数的功能。
func TestRouteParamsByName(t *testing.T) {
	name := ""
	lastName := ""
	wild := ""
	router := X创建()
	router.X绑定GET("/test/:name/:last_name/*wild", func(c *Context) {
		name = c.X参数.ByName("name")
		lastName = c.X参数.ByName("last_name")
		var ok bool
		wild, ok = c.X参数.Get("wild")

		assert.True(t, ok)
		assert.Equal(t, name, c.X取API参数值("name"))
		assert.Equal(t, lastName, c.X取API参数值("last_name"))

		assert.Empty(t, c.X取API参数值("wtf"))
		assert.Empty(t, c.X参数.ByName("wtf"))

		wtf, ok := c.X参数.Get("wtf")
		assert.Empty(t, wtf)
		assert.False(t, ok)
	})

	w := PerformRequest(router, http.MethodGet, "/test/john/smith/is/super/great")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "john", name)
	assert.Equal(t, "smith", lastName)
	assert.Equal(t, "/is/super/great", wild)
}

// TestContextParamsGet 测试即使存在额外的斜杠，也能从 URL 中解析出参数。
func TestRouteParamsByNameWithExtraSlash(t *testing.T) {
	name := ""
	lastName := ""
	wild := ""
	router := X创建()
	router.X删除多余斜杠 = true
	router.X绑定GET("/test/:name/:last_name/*wild", func(c *Context) {
		name = c.X参数.ByName("name")
		lastName = c.X参数.ByName("last_name")
		var ok bool
		wild, ok = c.X参数.Get("wild")

		assert.True(t, ok)
		assert.Equal(t, name, c.X取API参数值("name"))
		assert.Equal(t, lastName, c.X取API参数值("last_name"))

		assert.Empty(t, c.X取API参数值("wtf"))
		assert.Empty(t, c.X参数.ByName("wtf"))

		wtf, ok := c.X参数.Get("wtf")
		assert.Empty(t, wtf)
		assert.False(t, ok)
	})

	w := PerformRequest(router, http.MethodGet, "//test//john//smith//is//super//great")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "john", name)
	assert.Equal(t, "smith", lastName)
	assert.Equal(t, "/is/super/great", wild)
}

// TestRouteParamsNotEmpty 测试即使在初始化上下文（已在先前请求中发生）后注册了带有参数/通配符的路由，上下文参数也会被设置。
func TestRouteParamsNotEmpty(t *testing.T) {
	name := ""
	lastName := ""
	wild := ""
	router := X创建()

	w := PerformRequest(router, http.MethodGet, "/test/john/smith/is/super/great")

	assert.Equal(t, http.StatusNotFound, w.Code)

	router.X绑定GET("/test/:name/:last_name/*wild", func(c *Context) {
		name = c.X参数.ByName("name")
		lastName = c.X参数.ByName("last_name")
		var ok bool
		wild, ok = c.X参数.Get("wild")

		assert.True(t, ok)
		assert.Equal(t, name, c.X取API参数值("name"))
		assert.Equal(t, lastName, c.X取API参数值("last_name"))

		assert.Empty(t, c.X取API参数值("wtf"))
		assert.Empty(t, c.X参数.ByName("wtf"))

		wtf, ok := c.X参数.Get("wtf")
		assert.Empty(t, wtf)
		assert.False(t, ok)
	})

	w = PerformRequest(router, http.MethodGet, "/test/john/smith/is/super/great")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "john", name)
	assert.Equal(t, "smith", lastName)
	assert.Equal(t, "/is/super/great", wild)
}

// TestHandleStaticFile - 确保静态文件处理正常
func TestRouteStaticFile(t *testing.T) {
	// SETUP file
	testRoot, _ := os.Getwd()
	f, err := os.CreateTemp(testRoot, "")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(f.Name())
	_, err = f.WriteString("Gin Web Framework")
	assert.NoError(t, err)
	f.Close()

	dir, filename := filepath.Split(f.Name())

	// SETUP gin
	router := X创建()
	router.X绑定静态文件目录("/using_static", dir)
	router.X绑定静态单文件("/result", f.Name())

	w := PerformRequest(router, http.MethodGet, "/using_static/"+filename)
	w2 := PerformRequest(router, http.MethodGet, "/result")

	assert.Equal(t, w, w2)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Gin Web Framework", w.Body.String())
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))

	w3 := PerformRequest(router, http.MethodHead, "/using_static/"+filename)
	w4 := PerformRequest(router, http.MethodHead, "/result")

	assert.Equal(t, w3, w4)
	assert.Equal(t, http.StatusOK, w3.Code)
}

// TestHandleStaticFile - 确保静态文件处理正常
func TestRouteStaticFileFS(t *testing.T) {
	// SETUP file
	testRoot, _ := os.Getwd()
	f, err := os.CreateTemp(testRoot, "")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(f.Name())
	_, err = f.WriteString("Gin Web Framework")
	assert.NoError(t, err)
	f.Close()

	dir, filename := filepath.Split(f.Name())
	// SETUP gin
	router := X创建()
	router.X绑定静态文件目录("/using_static", dir)
	router.X绑定静态单文件FS("/result_fs", filename, Dir(dir, false))

	w := PerformRequest(router, http.MethodGet, "/using_static/"+filename)
	w2 := PerformRequest(router, http.MethodGet, "/result_fs")

	assert.Equal(t, w, w2)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Gin Web Framework", w.Body.String())
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))

	w3 := PerformRequest(router, http.MethodHead, "/using_static/"+filename)
	w4 := PerformRequest(router, http.MethodHead, "/result_fs")

	assert.Equal(t, w3, w4)
	assert.Equal(t, http.StatusOK, w3.Code)
}

// TestHandleStaticDir - 确保根目录/子目录正确处理
func TestRouteStaticListingDir(t *testing.T) {
	router := X创建()
	router.X绑定静态文件目录FS("/", Dir("./", true))

	w := PerformRequest(router, http.MethodGet, "/")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "gin.go")
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

// TestHandleHeadToDir - 确保根目录/子目录处理正确
func TestRouteStaticNoListing(t *testing.T) {
	router := X创建()
	router.X绑定静态文件目录("/", "./")

	w := PerformRequest(router, http.MethodGet, "/")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.NotContains(t, w.Body.String(), "gin.go")
}

func TestRouterMiddlewareAndStatic(t *testing.T) {
	router := X创建()
	static := router.X创建分组路由("/", func(c *Context) {
		c.Writer.Header().Add("Last-Modified", "Mon, 02 Jan 2006 15:04:05 MST")
		c.Writer.Header().Add("Expires", "Mon, 02 Jan 2006 15:04:05 MST")
		c.Writer.Header().Add("X-GIN", "Gin Framework")
	})
	static.X绑定静态文件目录("/", "./")

	w := PerformRequest(router, http.MethodGet, "/gin.go")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "package gin")
// 当Go版本小于等于1.16时，Content-Type='text/plain; charset=utf-8'，
// 否则，Content-Type='text/x-go; charset=utf-8'
	assert.NotEqual(t, "", w.Header().Get("Content-Type"))
	assert.NotEqual(t, w.Header().Get("Last-Modified"), "Mon, 02 Jan 2006 15:04:05 MST")
	assert.Equal(t, "Mon, 02 Jan 2006 15:04:05 MST", w.Header().Get("Expires"))
	assert.Equal(t, "Gin Framework", w.Header().Get("x-GIN"))
}

func TestRouteNotAllowedEnabled(t *testing.T) {
	router := X创建()
	router.HandleMethodNotAllowed = true
	router.X绑定POST("/path", func(c *Context) {})
	w := PerformRequest(router, http.MethodGet, "/path")
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)

	router.X绑定405(func(c *Context) {
		c.X输出文本(http.StatusTeapot, "responseText")
	})
	w = PerformRequest(router, http.MethodGet, "/path")
	assert.Equal(t, "responseText", w.Body.String())
	assert.Equal(t, http.StatusTeapot, w.Code)
}

func TestRouteNotAllowedEnabled2(t *testing.T) {
	router := X创建()
	router.HandleMethodNotAllowed = true
	// 向trees添加一个methodTree
	router.addRoute(http.MethodPost, "/", HandlersChain{func(_ *Context) {}})
	router.X绑定GET("/path2", func(c *Context) {})
	w := PerformRequest(router, http.MethodPost, "/path2")
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestRouteNotAllowedDisabled(t *testing.T) {
	router := X创建()
	router.HandleMethodNotAllowed = false
	router.X绑定POST("/path", func(c *Context) {})
	w := PerformRequest(router, http.MethodGet, "/path")
	assert.Equal(t, http.StatusNotFound, w.Code)

	router.X绑定405(func(c *Context) {
		c.X输出文本(http.StatusTeapot, "responseText")
	})
	w = PerformRequest(router, http.MethodGet, "/path")
	assert.Equal(t, "404 page not found", w.Body.String())
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestRouterNotFoundWithRemoveExtraSlash(t *testing.T) {
	router := X创建()
	router.X删除多余斜杠 = true
	router.X绑定GET("/path", func(c *Context) {})
	router.X绑定GET("/", func(c *Context) {})

	testRoutes := []struct {
		route    string
		code     int
		location string
	}{
		{"/../path", http.StatusOK, ""},    // CleanPath
		{"/nope", http.StatusNotFound, ""}, // NotFound
	}
	for _, tr := range testRoutes {
		w := PerformRequest(router, "GET", tr.route)
		assert.Equal(t, tr.code, w.Code)
		if w.Code != http.StatusNotFound {
			assert.Equal(t, tr.location, fmt.Sprint(w.Header().Get("Location")))
		}
	}
}

func TestRouterNotFound(t *testing.T) {
	router := X创建()
	router.X重定向固定路径 = true
	router.X绑定GET("/path", func(c *Context) {})
	router.X绑定GET("/dir/", func(c *Context) {})
	router.X绑定GET("/", func(c *Context) {})

	testRoutes := []struct {
		route    string
		code     int
		location string
	}{
		{"/path/", http.StatusMovedPermanently, "/path"},   // TSR -/
		{"/dir", http.StatusMovedPermanently, "/dir/"},     // TSR +/
		{"/PATH", http.StatusMovedPermanently, "/path"},    // Fixed Case
		{"/DIR/", http.StatusMovedPermanently, "/dir/"},    // Fixed Case
		{"/PATH/", http.StatusMovedPermanently, "/path"},   // Fixed Case -/
		{"/DIR", http.StatusMovedPermanently, "/dir/"},     // Fixed Case +/
		{"/../path", http.StatusMovedPermanently, "/path"}, // Without CleanPath
		{"/nope", http.StatusNotFound, ""},                 // NotFound
	}
	for _, tr := range testRoutes {
		w := PerformRequest(router, http.MethodGet, tr.route)
		assert.Equal(t, tr.code, w.Code)
		if w.Code != http.StatusNotFound {
			assert.Equal(t, tr.location, fmt.Sprint(w.Header().Get("Location")))
		}
	}

	// 测试自定义未找到处理器
	var notFound bool
	router.X绑定404(func(c *Context) {
		c.X停止并带状态码(http.StatusNotFound)
		notFound = true
	})
	w := PerformRequest(router, http.MethodGet, "/nope")
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.True(t, notFound)

	// 测试非GET方法（期望返回307而非301）
	router.X绑定PATCH("/path", func(c *Context) {})
	w = PerformRequest(router, http.MethodPatch, "/path/")
	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
	assert.Equal(t, "map[Location:[/path]]", fmt.Sprint(w.Header()))

	// 测试特殊情况，其中前缀 "/" 对应的节点不存在
	router = X创建()
	router.X绑定GET("/a", func(c *Context) {})
	w = PerformRequest(router, http.MethodGet, "/")
	assert.Equal(t, http.StatusNotFound, w.Code)

	// 用于重现问题#2843的bug测试
	router = X创建()
	router.X绑定404(func(c *Context) {
		if c.X请求.RequestURI == "/login" {
			c.X输出文本(200, "login")
		}
	})
	router.X绑定GET("/logout", func(c *Context) {
		c.X输出文本(200, "logout")
	})
	w = PerformRequest(router, http.MethodGet, "/login")
	assert.Equal(t, "login", w.Body.String())
	w = PerformRequest(router, http.MethodGet, "/logout")
	assert.Equal(t, "logout", w.Body.String())
}

func TestRouterStaticFSNotFound(t *testing.T) {
	router := X创建()
	router.X绑定静态文件目录FS("/", http.FileSystem(http.Dir("/thisreallydoesntexist/")))
	router.X绑定404(func(c *Context) {
		c.X输出文本(404, "non existent")
	})

	w := PerformRequest(router, http.MethodGet, "/nonexistent")
	assert.Equal(t, "non existent", w.Body.String())

	w = PerformRequest(router, http.MethodHead, "/nonexistent")
	assert.Equal(t, "non existent", w.Body.String())
}

func TestRouterStaticFSFileNotFound(t *testing.T) {
	router := X创建()

	router.X绑定静态文件目录FS("/", http.FileSystem(http.Dir(".")))

	assert.NotPanics(t, func() {
		PerformRequest(router, http.MethodGet, "/nonexistent")
	})
}

// 用于重现问题 #1805 的 bug 测试
func TestMiddlewareCalledOnceByRouterStaticFSNotFound(t *testing.T) {
	router := X创建()

	// 中间件必须在每个请求中仅调用一次。
	middlewareCalledNum := 0
	router.X中间件(func(c *Context) {
		middlewareCalledNum++
	})

	router.X绑定静态文件目录FS("/", http.FileSystem(http.Dir("/thisreallydoesntexist/")))

	// First access
	PerformRequest(router, http.MethodGet, "/nonexistent")
	assert.Equal(t, 1, middlewareCalledNum)

	// Second access
	PerformRequest(router, http.MethodHead, "/nonexistent")
	assert.Equal(t, 2, middlewareCalledNum)
}

func TestRouteRawPath(t *testing.T) {
	route := X创建()
	route.X使用原始路径 = true

	route.X绑定POST("/project/:name/build/:num", func(c *Context) {
		name := c.X参数.ByName("name")
		num := c.X参数.ByName("num")

		assert.Equal(t, name, c.X取API参数值("name"))
		assert.Equal(t, num, c.X取API参数值("num"))

		assert.Equal(t, "Some/Other/Project", name)
		assert.Equal(t, "222", num)
	})

	w := PerformRequest(route, http.MethodPost, "/project/Some%2FOther%2FProject/build/222")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRouteRawPathNoUnescape(t *testing.T) {
	route := X创建()
	route.X使用原始路径 = true
	route.UnescapePathValues = false

	route.X绑定POST("/project/:name/build/:num", func(c *Context) {
		name := c.X参数.ByName("name")
		num := c.X参数.ByName("num")

		assert.Equal(t, name, c.X取API参数值("name"))
		assert.Equal(t, num, c.X取API参数值("num"))

		assert.Equal(t, "Some%2FOther%2FProject", name)
		assert.Equal(t, "333", num)
	})

	w := PerformRequest(route, http.MethodPost, "/project/Some%2FOther%2FProject/build/333")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRouteServeErrorWithWriteHeader(t *testing.T) {
	route := X创建()
	route.X中间件(func(c *Context) {
		c.X设置状态码(421)
		c.X中间件继续()
	})

	w := PerformRequest(route, http.MethodGet, "/NotFound")
	assert.Equal(t, 421, w.Code)
	assert.Equal(t, 0, w.Body.Len())
}

func TestRouteContextHoldsFullPath(t *testing.T) {
	router := X创建()

	// Test routes
	routes := []string{
		"/simple",
		"/project/:name",
		"/",
		"/news/home",
		"/news",
		"/simple-two/one",
		"/simple-two/one-two",
		"/project/:name/build/*params",
		"/project/:name/bui",
		"/user/:id/status",
		"/user/:id",
		"/user/:id/profile",
	}

	for _, route := range routes {
		actualRoute := route
		router.X绑定GET(route, func(c *Context) {
			// 对于每个已定义的路由，其上下文应包含完整的路径
			assert.Equal(t, actualRoute, c.X取路由路径())
			c.X停止并带状态码(http.StatusOK)
		})
	}

	for _, route := range routes {
		w := PerformRequest(router, http.MethodGet, route)
		assert.Equal(t, http.StatusOK, w.Code)
	}

	// Test not found
	router.X中间件(func(c *Context) {
		// 对于未找到的路由，其完整路径为空
		assert.Equal(t, "", c.X取路由路径())
	})

	w := PerformRequest(router, http.MethodGet, "/not-found")
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestEngineHandleMethodNotAllowedCornerCase(t *testing.T) {
	r := X创建()
	r.HandleMethodNotAllowed = true

	base := r.X创建分组路由("base")
	base.X绑定GET("/metrics", handlerTest1)

	v1 := base.X创建分组路由("v1")

	v1.X绑定GET("/:id/devices", handlerTest1)
	v1.X绑定GET("/user/:id/groups", handlerTest1)

	v1.X绑定GET("/orgs/:id", handlerTest1)
	v1.X绑定DELETE("/orgs/:id", handlerTest1)

	w := PerformRequest(r, "GET", "/base/v1/user/groups")
	assert.Equal(t, http.StatusNotFound, w.Code)
}
