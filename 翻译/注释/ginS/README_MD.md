
<原文开始>
Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```

<原文结束>

# <翻译开始>
# Gin 默认服务器

这是一个针对 Gin 的 API 实验。

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
)

func main() {
// 定义根路由，当访问根路径时返回 "Hello World"
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
// 运行 Gin 服务器
	ginS.Run()
}
```

这段代码是使用 Gin 框架创建一个简单的 HTTP 服务器，当访问服务器的根 URL（"/"）时，会返回 HTTP 状态码 200 以及内容 "Hello World"。

# <翻译结束>

