// 版权所有 ? 2017 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中查阅。

package gin类

import (
	"html/template"
	"net/http"
	"os"
	"testing"
)

func BenchmarkOneRoute(B *testing.B) {
	router := X创建()
	router.X绑定GET("/ping", func(c *Context) {})
	runRequest(B, router, "GET", "/ping")
}

func BenchmarkRecoveryMiddleware(B *testing.B) {
	router := X创建()
	router.X中间件(Recovery())
	router.X绑定GET("/", func(c *Context) {})
	runRequest(B, router, "GET", "/")
}

func BenchmarkLoggerMiddleware(B *testing.B) {
	router := X创建()
	router.X中间件(LoggerWithWriter(newMockWriter()))
	router.X绑定GET("/", func(c *Context) {})
	runRequest(B, router, "GET", "/")
}

func BenchmarkManyHandlers(B *testing.B) {
	router := X创建()
	router.X中间件(Recovery(), LoggerWithWriter(newMockWriter()))
	router.X中间件(func(c *Context) {})
	router.X中间件(func(c *Context) {})
	router.X绑定GET("/ping", func(c *Context) {})
	runRequest(B, router, "GET", "/ping")
}

func Benchmark5Params(B *testing.B) {
	DefaultWriter = os.Stdout
	router := X创建()
	router.X中间件(func(c *Context) {})
	router.X绑定GET("/param/:param1/:params2/:param3/:param4/:param5", func(c *Context) {})
	runRequest(B, router, "GET", "/param/path/to/parameter/john/12345")
}

func BenchmarkOneRouteJSON(B *testing.B) {
	router := X创建()
	data := struct {
		Status string `json:"status"`
	}{"ok"}
	router.X绑定GET("/json", func(c *Context) {
		c.X输出JSON(http.StatusOK, data)
	})
	runRequest(B, router, "GET", "/json")
}

func BenchmarkOneRouteHTML(B *testing.B) {
	router := X创建()
	t := template.Must(template.New("index").Parse(`
		<html><body><h1>{{.}}</h1></body></html>`))
	router.X设置Template模板(t)

	router.X绑定GET("/html", func(c *Context) {
		c.X输出html模板(http.StatusOK, "index", "hola")
	})
	runRequest(B, router, "GET", "/html")
}

func BenchmarkOneRouteSet(B *testing.B) {
	router := X创建()
	router.X绑定GET("/ping", func(c *Context) {
		c.X设置值("key", "value")
	})
	runRequest(B, router, "GET", "/ping")
}

func BenchmarkOneRouteString(B *testing.B) {
	router := X创建()
	router.X绑定GET("/text", func(c *Context) {
		c.X输出文本(http.StatusOK, "this is a plain text")
	})
	runRequest(B, router, "GET", "/text")
}

func BenchmarkManyRoutesFist(B *testing.B) {
	router := X创建()
	router.X绑定Any("/ping", func(c *Context) {})
	runRequest(B, router, "GET", "/ping")
}

func BenchmarkManyRoutesLast(B *testing.B) {
	router := X创建()
	router.X绑定Any("/ping", func(c *Context) {})
	runRequest(B, router, "OPTIONS", "/ping")
}

func Benchmark404(B *testing.B) {
	router := X创建()
	router.X绑定Any("/something", func(c *Context) {})
	router.X绑定404(func(c *Context) {})
	runRequest(B, router, "GET", "/ping")
}

func Benchmark404Many(B *testing.B) {
	router := X创建()
	router.X绑定GET("/", func(c *Context) {})
	router.X绑定GET("/path/to/something", func(c *Context) {})
	router.X绑定GET("/post/:id", func(c *Context) {})
	router.X绑定GET("/view/:id", func(c *Context) {})
	router.X绑定GET("/favicon.ico", func(c *Context) {})
	router.X绑定GET("/robots.txt", func(c *Context) {})
	router.X绑定GET("/delete/:id", func(c *Context) {})
	router.X绑定GET("/user/:id/:mode", func(c *Context) {})

	router.X绑定404(func(c *Context) {})
	runRequest(B, router, "GET", "/viewfake")
}

type mockWriter struct {
	headers http.Header
}

func newMockWriter() *mockWriter {
	return &mockWriter{
		http.Header{},
	}
}

func (m *mockWriter) Header() (h http.Header) {
	return m.headers
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockWriter) WriteHeader(int) {}

func runRequest(B *testing.B, r *Engine, method, path string) {
	// create fake request
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}
	w := newMockWriter()
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		r.ServeHTTP(w, req)
	}
}
