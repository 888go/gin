# # Rollbar

[![运行测试](https://github.com/gin-contrib/rollbar/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/rollbar/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/rollbar/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/rollbar)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/rollbar)](https://goreportcard.com/report/github.com/gin-contrib/rollbar)
[![GoDoc](https://godoc.org/github.com/gin-contrib/rollbar?status.svg)](https://godoc.org/github.com/gin-contrib/rollbar)
[![加入聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

这是一个用于集成[Rollbar](https://rollbar.com/)错误监控的中间件。它使用了[rollbar-go](https://github.com/rollbar/rollbar-go) SDK，可以报告错误并记录消息。
## # 使用方法

下载并安装：

```sh
go get github.com/gin-contrib/rollbar
```

在代码中导入：

```go
import "github.com/gin-contrib/rollbar"
```
## # 示例

```go
package main

import (
    "log"

    "github.com/gin-contrib/rollbar"
    "github.com/gin-gonic/gin"

    roll "github.com/rollbar/rollbar-go"
)

func main() {
    roll.SetToken("MY_TOKEN")
// roll.SetEnvironment("production") // 默认为 "development"

    r := gin.Default()
    r.Use(rollbar.Recovery(true))

    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
```

这段代码是Go语言编写的，主要内容如下：

1. 引入所需包，包括日志记录(log)、gin框架以及Rollbar的 Gin中间件。

2. 设置Rollbar的访问令牌(`SetToken`)以便将错误信息发送到Rollbar。这里用"MY_TOKEN"作为占位符，实际使用时应替换为实际的Rollbar访问令牌。

3. 可选地设置环境变量（`SetEnvironment`），默认为"development"，注释已将其屏蔽。

4. 初始化一个默认配置的gin路由器实例 `r`。

5. 使用Rollbar提供的Recovery中间件，当发生panic时会捕获并报告到Rollbar，参数设为`true`表示启用该功能。

6. 最后，尝试在8080端口运行web服务器，如果启动过程中出现错误，则通过log.Fatal记录并终止程序。
