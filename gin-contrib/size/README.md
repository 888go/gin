# # 大小

[![运行测试](https://github.com/gin-contrib/size/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/size/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/size/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/size)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/size)](https://goreportcard.com/report/github.com/gin-contrib/size)
[![GoDoc](https://godoc.org/github.com/gin-contrib/size?status.svg)](https://godoc.org/github.com/gin-contrib/size)
[![加入Gitter聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

 Gin框架POST请求大小限制
## # 示例

```go
package main

import (
    "net/http"
    "log"

    "github.com/gin-contrib/size" // 包引入，用于限制请求大小
    "github.com/gin-gonic/gin"     // Gin Web框架包引入
)

func handler(ctx *gin.Context) { // 处理函数
    val := ctx.PostForm("b") // 从POST表单中获取名为“b”的值
    if len(ctx.Errors) > 0 { // 检查是否存在错误
        return
    }
    ctx.String(http.StatusOK, "got %s\n", val) // 如果无错误，返回HTTP状态码200，并在响应体中输出获取到的值
}

func main() {
    r := gin.Default() // 初始化一个Gin引擎实例
    r.Use(limits.RequestSizeLimiter(10)) // 使用中间件限制请求大小为10字节
    r.POST("/", handler) // 配置路由，处理POST请求根路径"/"，调用handler函数
    if err := r.Run(":8080"); err != nil { // 在8080端口运行Gin服务器
        log.Fatal(err) // 如果运行时发生错误，则输出错误信息并结束程序
    }
}
```
