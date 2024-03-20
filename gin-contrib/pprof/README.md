# # pprof

[![](https://github.com/gin-contrib/pprof/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/pprof/actions/workflows/go.yml)：运行测试
[![](https://codecov.io/gh/gin-contrib/pprof/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/pprof)：Codecov 代码覆盖率
[![](https://goreportcard.com/badge/github.com/gin-contrib/pprof)](https://goreportcard.com/report/github.com/gin-contrib/pprof)：Go 语言报告卡
[![](https://godoc.org/github.com/gin-contrib/pprof?status.svg)](https://godoc.org/github.com/gin-contrib/pprof)：GoDoc 文档
[![](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)：加入 Gitter 聊天室讨论

gin pprof 中间件

> pprof 包通过其 HTTP 服务器提供运行时性能分析数据，这些数据以 pprof 可视化工具期望的格式呈现。
## Usage

### # 开始使用它

下载并安装：

```bash
go get github.com/gin-contrib/pprof
```

在代码中导入它：

```go
import "github.com/gin-contrib/pprof"
```

#
## # 示例

```go
package main

import (
  "github.com/gin-contrib/pprof"
  "github.com/gin-gonic/gin"
)

func main() {
// 初始化路由
  router := gin.Default()

// 注册pprof到路由中
  pprof.Register(router)

// 在8080端口运行服务器
  router.Run(":8080")
}
```

#
## # 更改默认路径前缀

```go
func main() {
// 创建默认的gin路由器实例
  router := gin.Default()

// 默认是 "debug/pprof"，现在更改为 "dev/pprof"
  pprof.Register(router, "dev/pprof")

// 在8080端口运行服务器
  router.Run(":8080")
}
```

#
## # 自定义路由分组

```go
package main

import (
    "net/http"
    "github.com/gin-contrib/pprof"
    "github.com/gin-gonic/gin"
)

func main() {
// 初始化默认路由
    router := gin.Default()

// 创建名为adminGroup的分组路由，其路径前缀为"/admin"
    adminGroup := router.Group("/admin", func(c *gin.Context) {
// 在进入该分组路由处理函数之前进行权限校验
        if c.Request.Header.Get("Authorization") != "foobar" {
// 权限验证失败，则中止请求并返回403 Forbidden状态码
            c.AbortWithStatus(http.StatusForbidden)
            return
        }
// 权限验证通过，则调用下一个中间件或路由处理函数
        c.Next()
    })

// 将pprof路由注册到adminGroup分组中，路径前缀为"pprof"
    pprof.RouteRegister(adminGroup, "pprof")

// 启动服务器监听8080端口
    router.Run(":8080")
}
```

#
## # 使用pprof工具

接下来，使用pprof工具查看堆内存分析：

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```

或者查看30秒的CPU使用情况分析：

```bash
go tool pprof http://localhost:8080/debug/pprof/profile
```

或者在你的程序中调用runtime.SetBlockProfileRate后，查看goroutine阻塞分析：

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```

或者收集5秒钟的执行跟踪信息：

```bash
wget http://localhost:8080/debug/pprof/trace?seconds=5
```
