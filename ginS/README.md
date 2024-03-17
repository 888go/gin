# # Gin 默认服务器

这是一个 Gin 的 API 实验。

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

翻译：

Gin 标准服务器

这是针对 Gin 的一个 API 测试示例。

```go
package main

import (
    "github.com/gin-gonic/gin" // 引入 Gin 框架
    "github.com/gin-gonic/gin/ginS" // 引入 Gin 相关包
)

func main() {
    ginS.GET("/", func(c *gin.Context) { // 定义 GET 请求的根路由处理函数
        c.String(200, "Hello World") // 返回 200 状态码及 "Hello World" 文本内容
    })
    ginS.Run() // 启动 Gin 服务器
}
```
