
<原文开始>
rollbar

[![Run Tests](https://github.com/gin-contrib/rollbar/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/rollbar/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/rollbar/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/rollbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/rollbar)](https://goreportcard.com/report/github.com/gin-contrib/rollbar)
[![GoDoc](https://godoc.org/github.com/gin-contrib/rollbar?status.svg)](https://godoc.org/github.com/gin-contrib/rollbar)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Middleware to integrate with [rollbar](https://rollbar.com/) error monitoring. It uses [rollbar-go](https://github.com/rollbar/rollbar-go) SDK that reports errors and logs messages.


<原文结束>

# <翻译开始>
# rollbar

[![运行测试](https://github.com/gin-contrib/rollbar/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/rollbar/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/rollbar/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/rollbar)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/rollbar)](https://goreportcard.com/report/github.com/gin-contrib/rollbar)
[![GoDoc](https://godoc.org/github.com/gin-contrib/rollbar?status.svg)](https://godoc.org/github.com/gin-contrib/rollbar)
[![加入聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

这是一个中间件，用于集成 [rollbar](https://rollbar.com/) 错误监控服务。它使用了 [rollbar-go](https://github.com/rollbar/rollbar-go) SDK，该SDK能够上报错误和记录消息。

# <翻译结束>


<原文开始>
Usage

Download and install it:

```sh
go get github.com/gin-contrib/rollbar
```

Import it in your code:

```go
import "github.com/gin-contrib/rollbar"
```


<原文结束>

# <翻译开始>
# 使用方法

下载并安装：

```sh
go get github.com/gin-contrib/rollbar
```

在代码中导入：

```go
import "github.com/gin-contrib/rollbar"
```

# <翻译结束>


<原文开始>
Example

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
  // roll.SetEnvironment("production") // defaults to "development"

  r := gin.Default()
  r.Use(rollbar.Recovery(true))

  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

<原文结束>

# <翻译开始>
# 示例

```go
package main

import (
  "log"

  "github.com/gin-contrib/rollbar"
  "github.com/gin-gonic/gin"

  roll "github.com/rollbar/rollbar-go"
)

func main() {
  roll.SetToken("我的令牌")
// roll.SetEnvironment("生产环境")// 默认为 "开发环境"

  r := gin.Default()
  r.Use(rollbar.Recovery(true))

  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

这段代码是Go语言编写的，主要内容是一个web应用的初始化和错误处理。首先导入所需包，包括gin框架以及rollbar用于错误收集和报告。

在main函数中：

1. 使用`roll.SetToken("我的令牌")`设置Rollbar服务的访问令牌。
2. 注释掉的部分表示设置运行环境，默认为"开发环境"，如果需要可以改为"生产环境"。
3. 初始化一个gin默认路由实例 `r := gin.Default()`。
4. 使用 `r.Use(rollbar.Recovery(true))` 将Rollbar的错误恢复中间件添加到路由中，以便在程序出现panic时能捕获并报告错误。
5. 最后，尝试运行服务器在8080端口，如果启动失败，则通过`log.Fatal(err)`输出错误信息并终止程序。

# <翻译结束>

