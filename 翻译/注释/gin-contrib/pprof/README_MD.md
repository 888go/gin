
<原文开始>
pprof

[![Run Tests](https://github.com/gin-contrib/pprof/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/pprof/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/pprof/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/pprof)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/pprof)](https://goreportcard.com/report/github.com/gin-contrib/pprof)
[![GoDoc](https://godoc.org/github.com/gin-contrib/pprof?status.svg)](https://godoc.org/github.com/gin-contrib/pprof)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

gin pprof middleware

> Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.


<原文结束>

# <翻译开始>
# pprof

[![运行测试](https://github.com/gin-contrib/pprof/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/pprof/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/pprof/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/pprof)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/pprof)](https://goreportcard.com/report/github.com/gin-contrib/pprof)
[![GoDoc](https://godoc.org/github.com/gin-contrib/pprof?status.svg)](https://godoc.org/github.com/gin-contrib/pprof)
[![加入 Gitter 聊天](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

gin pprof 中间件

> pprof 包通过其HTTP服务器提供运行时分析数据，这些数据采用pprof可视化工具所期望的格式。

# <翻译结束>


<原文开始>
Start using it

Download and install it:

```bash
go get github.com/gin-contrib/pprof
```

Import it in your code:

```go
import "github.com/gin-contrib/pprof"
```

#
<原文结束>

# <翻译开始>
# 开始使用

下载并安装：

```bash
go get github.com/gin-contrib/pprof
```

在代码中导入：

```go
import "github.com/gin-contrib/pprof"
```

#

# <翻译结束>


<原文开始>
Example

```go
package main

import (
  "github.com/gin-contrib/pprof"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  pprof.Register(router)
  router.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 示例

```go
package main

import (
  "github.com/gin-contrib/pprof"
  "github.com/gin-gonic/gin"
)

func main() {
  // 初始化一个 Gin 路由器
  router := gin.Default()

  // 注册 pprof 工具到路由中
  pprof.Register(router)

  // 运行服务器在 8080 端口
  router.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
change default path prefix

```go
func main() {
  router := gin.Default()
  // default is "debug/pprof"
  pprof.Register(router, "dev/pprof")
  router.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 更改默认路径前缀

```go
func main() {
  // 初始化 Gin 路由器
  router := gin.Default()
  
  // 默认值为 "debug/pprof"，现在改为 "dev/pprof"
  pprof.Register(router, "dev/pprof")
  
  // 运行服务器在 8080 端口
  router.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
custom router group

```go
package main

import (
  "net/http"

  "github.com/gin-contrib/pprof"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  adminGroup := router.Group("/admin", func(c *gin.Context) {
    if c.Request.Header.Get("Authorization") != "foobar" {
      c.AbortWithStatus(http.StatusForbidden)
      return
    }
    c.Next()
  })
  pprof.RouteRegister(adminGroup, "pprof")
  router.Run(":8080")
}

```

#
<原文结束>

# <翻译开始>
# 自定义路由分组

```go
package main

import (
  "net/http"
  "github.com/gin-contrib/pprof"
  "github.com/gin-gonic/gin"
)

func main() {
  // 初始化默认的gin路由器
  router := gin.Default()

  // 创建一个名为adminGroup的路由分组，该分组路径前缀为"/admin"，并添加中间件进行权限校验
  adminGroup := router.Group("/admin", func(c *gin.Context) {
    // 检查请求头中的"Authorization"字段是否为"foobar"，若不是则终止请求，并返回403 Forbidden状态码
    if c.Request.Header.Get("Authorization") != "foobar" {
      c.AbortWithStatus(http.StatusForbidden)
      return
    }
    // 如果通过验证，则调用Next方法继续执行后续的中间件和处理函数
    c.Next()
  })

  // 将pprof路由注册到adminGroup中，路由名称为"pprof"
  pprof.RouteRegister(adminGroup, "pprof")

  // 运行服务器在8080端口
  router.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
Use the pprof tool

Then use the pprof tool to look at the heap profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```

Or to look at a 30-second CPU profile:

```bash
go tool pprof http://localhost:8080/debug/pprof/profile
```

Or to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```

Or to collect a 5-second execution trace:

```bash
wget http://localhost:8080/debug/pprof/trace?seconds=5
```

<原文结束>

# <翻译开始>
# 使用pprof工具

然后使用pprof工具查看堆内存分析：

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```

或者查看30秒的CPU分析：

```bash
go tool pprof http://localhost:8080/debug/pprof/profile
```

或者在你的程序中调用runtime.SetBlockProfileRate之后，查看goroutine阻塞分析：

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```

或者收集5秒的执行追踪：

```bash
wget http://localhost:8080/debug/pprof/trace?seconds=5
```

# <翻译结束>

