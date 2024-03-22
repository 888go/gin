package pprof

import (
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/888go/gin"
)

func Test_getPrefix(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"default value", nil, "/debug/pprof"},
		{"test user input value", []string{"test/pprof"}, "test/pprof"},
		{"test user input value", []string{"test/pprof", "pprof"}, "test/pprof"},
	}
	for _, tt := range tests {
		if got := getPrefix(tt.args...); got != tt.want {
			t.Errorf("%q. getPrefix() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestRegisterAndRouteRegister(t *testing.T) {
	bearerToken := "Bearer token"
	gin类.X设置运行模式(gin类.X常量_运行模式_发布)
	r := gin类.X创建()
	Register(r)
	adminGroup := r.X创建分组路由("/admin", func(c *gin类.Context) {
		if c.X请求.Header.Get("Authorization") != bearerToken {
			c.X停止并带状态码(http.StatusForbidden)
			return
		}
		c.X中间件继续()
	})
	RouteRegister(adminGroup, "pprof")

	req, _ := http.NewRequest(http.MethodGet, "/debug/pprof/", nil)
	rw := httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if expected, got := http.StatusOK, rw.Code; expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}

	req, _ = http.NewRequest(http.MethodGet, "/admin/pprof/", nil)
	rw = httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if expected, got := http.StatusForbidden, rw.Code; expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}

	req, _ = http.NewRequest(http.MethodGet, "/admin/pprof/", nil)
	req.Header.Set("Authorization", bearerToken)
	rw = httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if expected, got := http.StatusOK, rw.Code; expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}
}
