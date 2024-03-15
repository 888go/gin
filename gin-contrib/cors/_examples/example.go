package main

import (
	"time"
	
	"github.com/888go/gin/gin-contrib/cors"
	"github.com/888go/gin"
)

func main() {
	router := gin.Default()
// 为 https://foo.com 和 https://github.com 域名设置 CORS（跨域资源共享策略），允许：
// - PUT 和 PATCH 方法
// - Origin 请求头
// - 证书共享（即 Cookies 和 HTTP 认证信息）
// - 预检请求（OPTIONS 方法）的缓存时间为 12 小时
// 注意：在实际代码中，还需配合中间件或特定的 CORS 库来实现上述配置。
	router.Use(cors.New(cors.Config{
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
	router.Run()
}
