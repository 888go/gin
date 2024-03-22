// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/http2"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func setupHTMLFiles(t *testing.T, mode string, tls bool, loadMethod func(*Engine)) *httptest.Server {
	X设置运行模式(mode)
	defer X设置运行模式(X常量_运行模式_测试)

	var router *Engine
	captureOutput(t, func() {
		router = X创建()
		router.X设置模板分隔符("{[{", "}]}")
		router.X设置Template模板函数(template.FuncMap{
			"formatAsDate": formatAsDate,
		})
		loadMethod(router)
		router.X绑定GET("/test", func(c *Context) {
			c.X输出html模板(http.StatusOK, "hello.tmpl", map[string]string{"name": "world"})
		})
		router.X绑定GET("/raw", func(c *Context) {
			c.X输出html模板(http.StatusOK, "raw.tmpl", map[string]any{
				"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
			})
		})
	})

	var ts *httptest.Server

	if tls {
		ts = httptest.NewTLSServer(router)
	} else {
		ts = httptest.NewServer(router)
	}

	return ts
}

func TestLoadHTMLGlobDebugMode(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_调试,
		false,
		func(router *Engine) {
			router.X加载HTML模板目录("./testdata/template/*")
		},
	)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/test", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}

func TestH2c(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Error(err)
	}
	r := X创建默认对象()
	r.X启用h2c支持 = true
	r.X绑定GET("/", func(c *Context) {
		c.X输出文本(200, "<h1>Hello world</h1>")
	})
	go func() {
		err := http.Serve(ln, r.X取主处理程序())
		if err != nil {
			t.Log(err)
		}
	}()
	defer ln.Close()

	url := "http://" + ln.Addr().String() + "/"

	httpClient := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(netw, addr)
			},
		},
	}

	res, err := httpClient.Get(url)
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}

func TestLoadHTMLGlobTestMode(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_测试,
		false,
		func(router *Engine) {
			router.X加载HTML模板目录("./testdata/template/*")
		},
	)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/test", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}

func TestLoadHTMLGlobReleaseMode(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_发布,
		false,
		func(router *Engine) {
			router.X加载HTML模板目录("./testdata/template/*")
		},
	)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/test", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}

func TestLoadHTMLGlobUsingTLS(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_调试,
		true,
		func(router *Engine) {
			router.X加载HTML模板目录("./testdata/template/*")
		},
	)
	defer ts.Close()

	// Use InsecureSkipVerify for avoiding `x509: certificate signed by unknown authority` error
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(fmt.Sprintf("%s/test", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}

func TestLoadHTMLGlobFromFuncMap(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_调试,
		false,
		func(router *Engine) {
			router.X加载HTML模板目录("./testdata/template/*")
		},
	)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/raw", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Date: 2017/07/01", string(resp))
}

func init() {
	X设置运行模式(X常量_运行模式_测试)
}

func TestCreateEngine(t *testing.T) {
	router := X创建()
	assert.Equal(t, "/", router.basePath)
	assert.Equal(t, router.engine, router)
	assert.Empty(t, router.Handlers)
}

func TestLoadHTMLFilesTestMode(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_测试,
		false,
		func(router *Engine) {
			router.X加载HTML模板文件("./testdata/template/hello.tmpl", "./testdata/template/raw.tmpl")
		},
	)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/test", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}

func TestLoadHTMLFilesDebugMode(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_调试,
		false,
		func(router *Engine) {
			router.X加载HTML模板文件("./testdata/template/hello.tmpl", "./testdata/template/raw.tmpl")
		},
	)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/test", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}

func TestLoadHTMLFilesReleaseMode(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_发布,
		false,
		func(router *Engine) {
			router.X加载HTML模板文件("./testdata/template/hello.tmpl", "./testdata/template/raw.tmpl")
		},
	)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/test", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}

func TestLoadHTMLFilesUsingTLS(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_测试,
		true,
		func(router *Engine) {
			router.X加载HTML模板文件("./testdata/template/hello.tmpl", "./testdata/template/raw.tmpl")
		},
	)
	defer ts.Close()

	// Use InsecureSkipVerify for avoiding `x509: certificate signed by unknown authority` error
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(fmt.Sprintf("%s/test", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}

func TestLoadHTMLFilesFuncMap(t *testing.T) {
	ts := setupHTMLFiles(
		t,
		X常量_运行模式_测试,
		false,
		func(router *Engine) {
			router.X加载HTML模板文件("./testdata/template/hello.tmpl", "./testdata/template/raw.tmpl")
		},
	)
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/raw", ts.URL))
	if err != nil {
		t.Error(err)
	}

	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Date: 2017/07/01", string(resp))
}

func TestAddRoute(t *testing.T) {
	router := X创建()
	router.addRoute("GET", "/", HandlersChain{func(_ *Context) {}})

	assert.Len(t, router.trees, 1)
	assert.NotNil(t, router.trees.get("GET"))
	assert.Nil(t, router.trees.get("POST"))

	router.addRoute("POST", "/", HandlersChain{func(_ *Context) {}})

	assert.Len(t, router.trees, 2)
	assert.NotNil(t, router.trees.get("GET"))
	assert.NotNil(t, router.trees.get("POST"))

	router.addRoute("POST", "/post", HandlersChain{func(_ *Context) {}})
	assert.Len(t, router.trees, 2)
}

func TestAddRouteFails(t *testing.T) {
	router := X创建()
	assert.Panics(t, func() { router.addRoute("", "/", HandlersChain{func(_ *Context) {}}) })
	assert.Panics(t, func() { router.addRoute("GET", "a", HandlersChain{func(_ *Context) {}}) })
	assert.Panics(t, func() { router.addRoute("GET", "/", HandlersChain{}) })

	router.addRoute("POST", "/post", HandlersChain{func(_ *Context) {}})
	assert.Panics(t, func() {
		router.addRoute("POST", "/post", HandlersChain{func(_ *Context) {}})
	})
}

func TestCreateDefaultRouter(t *testing.T) {
	router := X创建默认对象()
	assert.Len(t, router.Handlers, 2)
}

func TestNoRouteWithoutGlobalHandlers(t *testing.T) {
	var middleware0 HandlerFunc = func(c *Context) {}
	var middleware1 HandlerFunc = func(c *Context) {}

	router := X创建()

	router.X绑定404(middleware0)
	assert.Nil(t, router.Handlers)
	assert.Len(t, router.noRoute, 1)
	assert.Len(t, router.allNoRoute, 1)
	compareFunc(t, router.noRoute[0], middleware0)
	compareFunc(t, router.allNoRoute[0], middleware0)

	router.X绑定404(middleware1, middleware0)
	assert.Len(t, router.noRoute, 2)
	assert.Len(t, router.allNoRoute, 2)
	compareFunc(t, router.noRoute[0], middleware1)
	compareFunc(t, router.allNoRoute[0], middleware1)
	compareFunc(t, router.noRoute[1], middleware0)
	compareFunc(t, router.allNoRoute[1], middleware0)
}

func TestNoRouteWithGlobalHandlers(t *testing.T) {
	var middleware0 HandlerFunc = func(c *Context) {}
	var middleware1 HandlerFunc = func(c *Context) {}
	var middleware2 HandlerFunc = func(c *Context) {}

	router := X创建()
	router.X中间件(middleware2)

	router.X绑定404(middleware0)
	assert.Len(t, router.allNoRoute, 2)
	assert.Len(t, router.Handlers, 1)
	assert.Len(t, router.noRoute, 1)

	compareFunc(t, router.Handlers[0], middleware2)
	compareFunc(t, router.noRoute[0], middleware0)
	compareFunc(t, router.allNoRoute[0], middleware2)
	compareFunc(t, router.allNoRoute[1], middleware0)

	router.X中间件(middleware1)
	assert.Len(t, router.allNoRoute, 3)
	assert.Len(t, router.Handlers, 2)
	assert.Len(t, router.noRoute, 1)

	compareFunc(t, router.Handlers[0], middleware2)
	compareFunc(t, router.Handlers[1], middleware1)
	compareFunc(t, router.noRoute[0], middleware0)
	compareFunc(t, router.allNoRoute[0], middleware2)
	compareFunc(t, router.allNoRoute[1], middleware1)
	compareFunc(t, router.allNoRoute[2], middleware0)
}

func TestNoMethodWithoutGlobalHandlers(t *testing.T) {
	var middleware0 HandlerFunc = func(c *Context) {}
	var middleware1 HandlerFunc = func(c *Context) {}

	router := X创建()

	router.X绑定405(middleware0)
	assert.Empty(t, router.Handlers)
	assert.Len(t, router.noMethod, 1)
	assert.Len(t, router.allNoMethod, 1)
	compareFunc(t, router.noMethod[0], middleware0)
	compareFunc(t, router.allNoMethod[0], middleware0)

	router.X绑定405(middleware1, middleware0)
	assert.Len(t, router.noMethod, 2)
	assert.Len(t, router.allNoMethod, 2)
	compareFunc(t, router.noMethod[0], middleware1)
	compareFunc(t, router.allNoMethod[0], middleware1)
	compareFunc(t, router.noMethod[1], middleware0)
	compareFunc(t, router.allNoMethod[1], middleware0)
}

func TestRebuild404Handlers(t *testing.T) {
}

func TestNoMethodWithGlobalHandlers(t *testing.T) {
	var middleware0 HandlerFunc = func(c *Context) {}
	var middleware1 HandlerFunc = func(c *Context) {}
	var middleware2 HandlerFunc = func(c *Context) {}

	router := X创建()
	router.X中间件(middleware2)

	router.X绑定405(middleware0)
	assert.Len(t, router.allNoMethod, 2)
	assert.Len(t, router.Handlers, 1)
	assert.Len(t, router.noMethod, 1)

	compareFunc(t, router.Handlers[0], middleware2)
	compareFunc(t, router.noMethod[0], middleware0)
	compareFunc(t, router.allNoMethod[0], middleware2)
	compareFunc(t, router.allNoMethod[1], middleware0)

	router.X中间件(middleware1)
	assert.Len(t, router.allNoMethod, 3)
	assert.Len(t, router.Handlers, 2)
	assert.Len(t, router.noMethod, 1)

	compareFunc(t, router.Handlers[0], middleware2)
	compareFunc(t, router.Handlers[1], middleware1)
	compareFunc(t, router.noMethod[0], middleware0)
	compareFunc(t, router.allNoMethod[0], middleware2)
	compareFunc(t, router.allNoMethod[1], middleware1)
	compareFunc(t, router.allNoMethod[2], middleware0)
}

func compareFunc(t *testing.T, a, b any) {
	sf1 := reflect.ValueOf(a)
	sf2 := reflect.ValueOf(b)
	if sf1.Pointer() != sf2.Pointer() {
		t.Error("different functions")
	}
}

func TestListOfRoutes(t *testing.T) {
	router := X创建()
	router.X绑定GET("/favicon.ico", handlerTest1)
	router.X绑定GET("/", handlerTest1)
	group := router.X创建分组路由("/users")
	{
		group.X绑定GET("/", handlerTest2)
		group.X绑定GET("/:id", handlerTest1)
		group.X绑定POST("/:id", handlerTest2)
	}
	router.X绑定静态文件目录("/static", ".")

	list := router.X取路由数组()

	assert.Len(t, list, 7)
	assertRoutePresent(t, list, RouteInfo{
		X方法:  "GET",
		X路径:    "/favicon.ico",
		Handler: "^(.*/vendor/)?github.com/888go/gin.handlerTest1$",
	})
	assertRoutePresent(t, list, RouteInfo{
		X方法:  "GET",
		X路径:    "/",
		Handler: "^(.*/vendor/)?github.com/888go/gin.handlerTest1$",
	})
	assertRoutePresent(t, list, RouteInfo{
		X方法:  "GET",
		X路径:    "/users/",
		Handler: "^(.*/vendor/)?github.com/888go/gin.handlerTest2$",
	})
	assertRoutePresent(t, list, RouteInfo{
		X方法:  "GET",
		X路径:    "/users/:id",
		Handler: "^(.*/vendor/)?github.com/888go/gin.handlerTest1$",
	})
	assertRoutePresent(t, list, RouteInfo{
		X方法:  "POST",
		X路径:    "/users/:id",
		Handler: "^(.*/vendor/)?github.com/888go/gin.handlerTest2$",
	})
}

func TestEngineHandleContext(t *testing.T) {
	r := X创建()
	r.X绑定GET("/", func(c *Context) {
		c.X请求.URL.Path = "/v2"
		r.HandleContext底层方法(c)
	})
	v2 := r.X创建分组路由("/v2")
	{
		v2.X绑定GET("/", func(c *Context) {})
	}

	assert.NotPanics(t, func() {
		w := PerformRequest(r, "GET", "/")
		assert.Equal(t, 301, w.Code)
	})
}

func TestEngineHandleContextManyReEntries(t *testing.T) {
	expectValue := 10000

	var handlerCounter, middlewareCounter int64

	r := X创建()
	r.X中间件(func(c *Context) {
		atomic.AddInt64(&middlewareCounter, 1)
	})
	r.X绑定GET("/:count", func(c *Context) {
		countStr := c.X取API参数值("count")
		count, err := strconv.Atoi(countStr)
		assert.NoError(t, err)

		n, err := c.Writer.Write([]byte("."))
		assert.NoError(t, err)
		assert.Equal(t, 1, n)

		switch {
		case count > 0:
			c.X请求.URL.Path = "/" + strconv.Itoa(count-1)
			r.HandleContext底层方法(c)
		}
	}, func(c *Context) {
		atomic.AddInt64(&handlerCounter, 1)
	})

	assert.NotPanics(t, func() {
		w := PerformRequest(r, "GET", "/"+strconv.Itoa(expectValue-1)) // include 0 value
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, expectValue, w.Body.Len())
	})

	assert.Equal(t, int64(expectValue), handlerCounter)
	assert.Equal(t, int64(expectValue), middlewareCounter)
}

func TestPrepareTrustedCIRDsWith(t *testing.T) {
	r := X创建()

	// valid ipv4 cidr
	{
		expectedTrustedCIDRs := []*net.IPNet{parseCIDR("0.0.0.0/0")}
		err := r.X设置受信任代理([]string{"0.0.0.0/0"})

		assert.NoError(t, err)
		assert.Equal(t, expectedTrustedCIDRs, r.trustedCIDRs)
	}

	// invalid ipv4 cidr
	{
		err := r.X设置受信任代理([]string{"192.168.1.33/33"})

		assert.Error(t, err)
	}

	// valid ipv4 address
	{
		expectedTrustedCIDRs := []*net.IPNet{parseCIDR("192.168.1.33/32")}

		err := r.X设置受信任代理([]string{"192.168.1.33"})

		assert.NoError(t, err)
		assert.Equal(t, expectedTrustedCIDRs, r.trustedCIDRs)
	}

	// invalid ipv4 address
	{
		err := r.X设置受信任代理([]string{"192.168.1.256"})

		assert.Error(t, err)
	}

	// valid ipv6 address
	{
		expectedTrustedCIDRs := []*net.IPNet{parseCIDR("2002:0000:0000:1234:abcd:ffff:c0a8:0101/128")}
		err := r.X设置受信任代理([]string{"2002:0000:0000:1234:abcd:ffff:c0a8:0101"})

		assert.NoError(t, err)
		assert.Equal(t, expectedTrustedCIDRs, r.trustedCIDRs)
	}

	// invalid ipv6 address
	{
		err := r.X设置受信任代理([]string{"gggg:0000:0000:1234:abcd:ffff:c0a8:0101"})

		assert.Error(t, err)
	}

	// valid ipv6 cidr
	{
		expectedTrustedCIDRs := []*net.IPNet{parseCIDR("::/0")}
		err := r.X设置受信任代理([]string{"::/0"})

		assert.NoError(t, err)
		assert.Equal(t, expectedTrustedCIDRs, r.trustedCIDRs)
	}

	// invalid ipv6 cidr
	{
		err := r.X设置受信任代理([]string{"gggg:0000:0000:1234:abcd:ffff:c0a8:0101/129"})

		assert.Error(t, err)
	}

	// valid combination
	{
		expectedTrustedCIDRs := []*net.IPNet{
			parseCIDR("::/0"),
			parseCIDR("192.168.0.0/16"),
			parseCIDR("172.16.0.1/32"),
		}
		err := r.X设置受信任代理([]string{
			"::/0",
			"192.168.0.0/16",
			"172.16.0.1",
		})

		assert.NoError(t, err)
		assert.Equal(t, expectedTrustedCIDRs, r.trustedCIDRs)
	}

	// invalid combination
	{
		err := r.X设置受信任代理([]string{
			"::/0",
			"192.168.0.0/16",
			"172.16.0.256",
		})

		assert.Error(t, err)
	}

	// nil value
	{
		err := r.X设置受信任代理(nil)

		assert.Nil(t, r.trustedCIDRs)
		assert.Nil(t, err)
	}
}

func parseCIDR(cidr string) *net.IPNet {
	_, parsedCIDR, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println(err)
	}
	return parsedCIDR
}

func assertRoutePresent(t *testing.T, gotRoutes RoutesInfo, wantRoute RouteInfo) {
	for _, gotRoute := range gotRoutes {
		if gotRoute.X路径 == wantRoute.X路径 && gotRoute.X方法 == wantRoute.X方法 {
			assert.Regexp(t, wantRoute.Handler, gotRoute.Handler)
			return
		}
	}
	t.Errorf("route not found: %v", wantRoute)
}

func handlerTest1(c *Context) {}
func handlerTest2(c *Context) {}
