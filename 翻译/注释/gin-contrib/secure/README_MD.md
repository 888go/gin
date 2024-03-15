
<原文开始>
Secure

[![Run Tests](https://github.com/gin-contrib/secure/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/secure/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/secure/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/secure)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/secure)](https://goreportcard.com/report/github.com/gin-contrib/secure)
[![GoDoc](https://godoc.org/github.com/gin-contrib/secure?status.svg)](https://godoc.org/github.com/gin-contrib/secure)

Secure middleware for [Gin](https://github.com/gin-gonic/gin/) framework.


<原文结束>

# <翻译开始>
# 安全

[![运行测试](https://github.com/gin-contrib/secure/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/secure/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/secure/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/secure)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/secure)](https://goreportcard.com/report/github.com/gin-contrib/secure)
[![GoDoc](https://godoc.org/github.com/gin-contrib/secure?status.svg)](https://godoc.org/github.com/gin-contrib/secure)

为 [Gin](https://github.com/gin-gonic/gin/) 框架提供的安全中间件。

# <翻译结束>


<原文开始>
Example

See the [example1](example/code1/example.go), [example2](example/code2/example.go).

DefaultConfig returns a Configuration with strict security settings

[embedmd]:
<原文结束>

# <翻译开始>
# 示例

请参阅[example1](example/code1/example.go)，[example2](example/code2/example.go)。

DefaultConfig 返回一个具有严格安全设置的 Configuration

[embedmd]:

# <翻译结束>


<原文开始>
(secure.go go /func DefaultConfig/ /^}$/)
```go
func DefaultConfig() Config {
	return Config{
		SSLRedirect:           true,
		IsDevelopment:         false,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}
}
```

[embedmd]:
<原文结束>

# <翻译开始>
# ```go
func 默认配置() 配置 {
	return 配置{
		SSL重定向:            true,
		是否开发环境:         false,
		STS有效期(秒):        315360000,
		STS包含子域名:       true,
		禁止框架加载:         true,
		防止内容类型嗅探:     true,
		浏览器XSS过滤:        true,
		内容安全策略:        "default-src 'self'",
		IE禁止下载:           true,
		SSL代理头信息:       map[string]string{"X-Forwarded-Proto": "https"},
	}
}
```

[embedmd]: (此处为注释或标记，表示引用MD格式的代码块结束)

# <翻译结束>


<原文开始>
(example/code1/example.go go)
```go
package main

import (
	"log"

	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(secure.New(secure.Config{
		AllowedHosts:          []string{"example.com", "ssl.example.com"},
		SSLRedirect:           true,
		SSLHost:               "ssl.example.com",
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Listen and Server in 0.0.0.0:8080
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
```

<原文结束>

# <翻译开始>
# 

# <翻译结束>

