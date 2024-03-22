package main

import (
	"time"
	
	"github.com/888go/gin/gin-contrib/cors"
	"github.com/888go/gin"
)

func main() {
	router := gin类.X创建默认对象()
// 允许以下内容的CORS（跨源资源共享）策略，针对 https://foo.com 和 https://github.com 两个来源：
// - PUT 和 PATCH 请求方法
// - Origin 请求头
// - 证书共享（即允许带凭证的请求）
// - 预检请求（Preflight requests）缓存有效期为12小时
	router.X中间件(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.X监听()
}
