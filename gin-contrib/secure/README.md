# # 安全

[![运行测试](https://github.com/gin-contrib/secure/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/secure/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/secure/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/secure)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/secure)](https://goreportcard.com/report/github.com/gin-contrib/secure)
[![GoDoc](https://godoc.org/github.com/gin-contrib/secure?status.svg)](https://godoc.org/github.com/gin-contrib/secure)

这是为[Gin](https://github.com/gin-gonic/gin/)框架提供的安全中间件。
## # 示例

请参阅以下示例代码：[example1](example/code1/example.go)，[example2](example/code2/example.go)。

DefaultConfig 返回一个具有严格安全设置的 Configuration 对象

[embedmd]:
# # 以下是Go语言代码片段的翻译，内容为`secure.go`文件中名为`DefaultConfig`的函数定义：

```go
func DefaultConfig() Config {
	return Config{
		SSLRedirect:           true,  // 是否启用SSL重定向
		IsDevelopment:         false, // 是否处于开发模式
		STSSeconds:            315360000, // 安全传输层头部（STS）的最大有效期（单位：秒）
		STSIncludeSubdomains:  true,  // STS设置是否应用于子域名
		FrameDeny:             true,   // 是否禁止页面在iframe内展示
		ContentTypeNosniff:    true,  // 是否启用防止内容类型嗅探
		BrowserXssFilter:      true,  // 是否启用浏览器XSS过滤器
		ContentSecurityPolicy: "default-src 'self'", // 内容安全策略，默认源只允许同源请求
		IENoOpen:              true,  // 是否禁用IE浏览器的“打开方式”功能以防止下载恶意文件
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"}, // SSL代理头设置，将"X-Forwarded-Proto"头的值设为"https"
	}
}
```

这段代码定义了一个返回`Config`结构体实例的函数`DefaultConfig`，该结构体用于配置一系列的安全选项，如SSL重定向、STS安全策略、网页安全特性等。
# # ```go
package main

import (
	"log"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func main() {
// 初始化路由
	router := gin.Default()

// 使用安全中间件配置
	router.Use(secure.New(secure.Config{
		AllowedHosts:          []string{"example.com", "ssl.example.com"}, // 允许的主机名
		SSLRedirect:           true,                                       // 启用 HTTPS 重定向
		SSLHost:               "ssl.example.com",                          // SSL 主机名
		STSSeconds:            315360000,                                  // 设置 HSTS 头部的有效期（单位：秒）
		STSIncludeSubdomains:  true,                                       // 在所有子域名中启用 HSTS
		FrameDeny:             true,                                       // 禁止页面在 frame 中显示
		ContentTypeNosniff:    true,                                       // 阻止浏览器猜测内容类型
		BrowserXssFilter:      true,                                       // 启用浏览器 XSS 过滤器
		ContentSecurityPolicy: "default-src 'self'",                         // 设置内容安全策略
		IENoOpen:              true,                                       // 防止 IE 打开某些文件类型为下载而不是直接打开
		ReferrerPolicy:        "strict-origin-when-cross-origin",            // 设置referrer策略
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"}, // 识别来自代理服务器的HTTPS请求头
	}))

// 定义一个简单的GET接口
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong") // 返回HTTP 200状态码及字符串"pong"
	})

// 监听并服务在 0.0.0.0:8080
	if err := router.Run(); err != nil {
		log.Fatal(err) // 如果运行过程中出现错误，输出错误信息并终止程序
	}
}
```

这段代码是用 Go 语言编写的 Gin 框架的一个 Web 应用示例。它创建了一个基本的 HTTP 路由，并应用了 secure 中间件以增强安全性，包括 HTTPS 重定向、HSTS 支持等。同时定义了一个简单的 GET 请求处理函数 `/ping`，用于返回 "pong" 字符串。最后，应用监听在 0.0.0.0:8080 端口上启动服务。
