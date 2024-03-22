// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

import (
	"net/http"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func init() {
	X设置运行模式(X常量_运行模式_测试)
}

func TestRouterGroupBasic(t *testing.T) {
	router := X创建()
	group := router.X创建分组路由("/hola", func(c *Context) {})
	group.X中间件(func(c *Context) {})

	assert.Len(t, group.Handlers, 2)
	assert.Equal(t, "/hola", group.X取路由基础路径())
	assert.Equal(t, router, group.engine)

	group2 := group.X创建分组路由("manu")
	group2.X中间件(func(c *Context) {}, func(c *Context) {})

	assert.Len(t, group2.Handlers, 4)
	assert.Equal(t, "/hola/manu", group2.X取路由基础路径())
	assert.Equal(t, router, group2.engine)
}

func TestRouterGroupBasicHandle(t *testing.T) {
	performRequestInGroup(t, http.MethodGet)
	performRequestInGroup(t, http.MethodPost)
	performRequestInGroup(t, http.MethodPut)
	performRequestInGroup(t, http.MethodPatch)
	performRequestInGroup(t, http.MethodDelete)
	performRequestInGroup(t, http.MethodHead)
	performRequestInGroup(t, http.MethodOptions)
}

func performRequestInGroup(t *testing.T, method string) {
	router := X创建()
	v1 := router.X创建分组路由("v1", func(c *Context) {})
	assert.Equal(t, "/v1", v1.X取路由基础路径())

	login := v1.X创建分组路由("/login/", func(c *Context) {}, func(c *Context) {})
	assert.Equal(t, "/v1/login/", login.X取路由基础路径())

	handler := func(c *Context) {
		c.X输出文本(http.StatusBadRequest, "the method was %s and index %d", c.X请求.Method, c.index)
	}

	switch method {
	case http.MethodGet:
		v1.X绑定GET("/test", handler)
		login.X绑定GET("/test", handler)
	case http.MethodPost:
		v1.X绑定POST("/test", handler)
		login.X绑定POST("/test", handler)
	case http.MethodPut:
		v1.X绑定PUT("/test", handler)
		login.X绑定PUT("/test", handler)
	case http.MethodPatch:
		v1.X绑定PATCH("/test", handler)
		login.X绑定PATCH("/test", handler)
	case http.MethodDelete:
		v1.X绑定DELETE("/test", handler)
		login.X绑定DELETE("/test", handler)
	case http.MethodHead:
		v1.X绑定HEAD("/test", handler)
		login.X绑定HEAD("/test", handler)
	case http.MethodOptions:
		v1.X绑定OPTIONS("/test", handler)
		login.X绑定OPTIONS("/test", handler)
	default:
		panic("unknown method")
	}

	w := PerformRequest(router, method, "/v1/login/test")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "the method was "+method+" and index 3", w.Body.String())

	w = PerformRequest(router, method, "/v1/test")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "the method was "+method+" and index 1", w.Body.String())
}

func TestRouterGroupInvalidStatic(t *testing.T) {
	router := X创建()
	assert.Panics(t, func() {
		router.X绑定静态文件目录("/path/:param", "/")
	})

	assert.Panics(t, func() {
		router.X绑定静态文件目录("/path/*param", "/")
	})
}

func TestRouterGroupInvalidStaticFile(t *testing.T) {
	router := X创建()
	assert.Panics(t, func() {
		router.X绑定静态单文件("/path/:param", "favicon.ico")
	})

	assert.Panics(t, func() {
		router.X绑定静态单文件("/path/*param", "favicon.ico")
	})
}

func TestRouterGroupInvalidStaticFileFS(t *testing.T) {
	router := X创建()
	assert.Panics(t, func() {
		router.X绑定静态单文件FS("/path/:param", "favicon.ico", Dir(".", false))
	})

	assert.Panics(t, func() {
		router.X绑定静态单文件FS("/path/*param", "favicon.ico", Dir(".", false))
	})
}

func TestRouterGroupTooManyHandlers(t *testing.T) {
	const (
		panicValue = "too many handlers"
		maximumCnt = abortIndex
	)
	router := X创建()
	handlers1 := make([]HandlerFunc, maximumCnt-1)
	router.X中间件(handlers1...)

	handlers2 := make([]HandlerFunc, maximumCnt+1)
	assert.PanicsWithValue(t, panicValue, func() {
		router.X中间件(handlers2...)
	})
	assert.PanicsWithValue(t, panicValue, func() {
		router.X绑定GET("/", handlers2...)
	})
}

func TestRouterGroupBadMethod(t *testing.T) {
	router := X创建()
	assert.Panics(t, func() {
		router.X绑定(http.MethodGet, "/")
	})
	assert.Panics(t, func() {
		router.X绑定(" GET", "/")
	})
	assert.Panics(t, func() {
		router.X绑定("GET ", "/")
	})
	assert.Panics(t, func() {
		router.X绑定("", "/")
	})
	assert.Panics(t, func() {
		router.X绑定("PO ST", "/")
	})
	assert.Panics(t, func() {
		router.X绑定("1GET", "/")
	})
	assert.Panics(t, func() {
		router.X绑定("PATCh", "/")
	})
}

func TestRouterGroupPipeline(t *testing.T) {
	router := X创建()
	testRoutesInterface(t, router)

	v1 := router.X创建分组路由("/v1")
	testRoutesInterface(t, v1)
}

func testRoutesInterface(t *testing.T, r IRoutes) {
	handler := func(c *Context) {}
	assert.Equal(t, r, r.X中间件(handler))

	assert.Equal(t, r, r.X绑定(http.MethodGet, "/handler", handler))
	assert.Equal(t, r, r.X绑定Any("/any", handler))
	assert.Equal(t, r, r.X绑定GET("/", handler))
	assert.Equal(t, r, r.X绑定POST("/", handler))
	assert.Equal(t, r, r.X绑定DELETE("/", handler))
	assert.Equal(t, r, r.X绑定PATCH("/", handler))
	assert.Equal(t, r, r.X绑定PUT("/", handler))
	assert.Equal(t, r, r.X绑定OPTIONS("/", handler))
	assert.Equal(t, r, r.X绑定HEAD("/", handler))
	assert.Equal(t, r, r.Match([]string{http.MethodPut, http.MethodPatch}, "/match", handler))

	assert.Equal(t, r, r.X绑定静态单文件("/file", "."))
	assert.Equal(t, r, r.X绑定静态单文件FS("/static2", ".", Dir(".", false)))
	assert.Equal(t, r, r.X绑定静态文件目录("/static", "."))
	assert.Equal(t, r, r.X绑定静态文件目录FS("/static2", Dir(".", false)))
}
