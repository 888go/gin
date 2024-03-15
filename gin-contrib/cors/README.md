# # CORS gin 中间件

[![运行测试](https://github.com/gin-contrib/cors/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/cors/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/cors/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/cors)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/cors)](https://goreportcard.com/report/github.com/gin-contrib/cors)
[![GoDoc](https://godoc.org/github.com/gin-contrib/cors?status.svg)](https://godoc.org/github.com/gin-contrib/cors)

启用 CORS 支持的 Gin 中间件/处理器。
## Usage

### # 开始使用

下载并安装：

```sh
go get github.com/gin-contrib/cors
```

在代码中导入：

```go
import "github.com/gin-contrib/cors"
```

#
## # 规范示例

```go
package main

import (
  "time"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
)

func main() {
  // 创建默认路由
  router := gin.Default()

  // 为 https://foo.com 和 https://github.com 域名设置 CORS，允许：
  // - PUT 和 PATCH 方法
  // - Origin 请求头
  // - 凭证共享
  // - 预检请求缓存12小时
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

  // 运行服务器
  router.Run()
}
```

#
## # 使用 DefaultConfig 作为起点

```go
func main() {
  // 初始化 Gin 路由器
  router := gin.Default()

  // - 默认情况下不允许任何来源
  // - 允许 GET、POST、PUT、HEAD 方法
  // - 禁用凭据共享
  // - 预检请求缓存12小时
  config := cors.DefaultConfig()
  
  // 允许特定来源
  config.AllowOrigins = []string{"http://google.com"}
  // 可以添加多个允许的来源
  // config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}

  // 若要允许所有来源，需设置 AllowAllOrigins 为 true
  // config.AllowAllOrigins = true

  // 使用 CORS 配置中间件
  router.Use(cors.New(config))

  // 运行服务器
  router.Run()
}
```

注意：虽然 Default() 方法默认允许所有来源，但 DefaultConfig() 并不默认允许所有来源，您仍需要通过设置 AllowAllOrigins 为 true 来实现这一目的。
## # `Default()` 允许所有来源

```go
func main() {
  // 创建一个Gin路由器实例
  router := gin.Default()
  
  // 相当于以下配置：
  // config := cors.DefaultConfig()
  // config.AllowAllOrigins = true
  // router.Use(cors.New(config))

  // 使用默认配置，允许所有来源
  router.Use(cors.Default())

  // 运行服务器
  router.Run()
}
```

允许所有来源将禁用Gin为客户端设置cookie的能力。在处理凭据时，不要允许所有来源。
