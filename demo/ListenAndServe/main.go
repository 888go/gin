package main

import (
	"github.com/888go/gin"
	"net/http"
	"time"
)

// 参考链接：https://gin-gonic.com/zh-cn/docs/examples/custom-http-config/
func main() {
	router := gin类.X创建默认对象()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
