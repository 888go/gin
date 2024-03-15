
<原文开始>
tracing

[![Run Tests](https://github.com/gin-contrib/opengintracing/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/opengintracing/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/opengintracing)](https://goreportcard.com/report/github.com/gin-contrib/opengintracing)
[![GoDoc](https://godoc.org/github.com/gin-contrib/opengintracing?status.png)](https://pkg.go.dev/github.com/gin-contrib/opengintracing)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

trace requests using opentracing specification, Download:

```bash
go get -u github.com/gin-contrib/opengintracing
```

See [opentracing/opentracing-go](https://github.com/opentracing/opentracing-go) for more information.


<原文结束>

# <翻译开始>
# 跟踪

[![运行测试](https://github.com/gin-contrib/opengintracing/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/opengintracing/actions/workflows/go.yml)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/opengintracing)](https://goreportcard.com/report/github.com/gin-contrib/opengintracing)
[![GoDoc](https://godoc.org/github.com/gin-contrib/opengintracing?status.png)](https://pkg.go.dev/github.com/gin-contrib/opengintracing)
[![许可协议: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

使用 opentracing 规范追踪请求，下载命令：

```bash
go get -u github.com/gin-contrib/opengintracing
```

更多相关信息请参阅 [opentracing/opentracing-go](https://github.com/opentracing/opentracing-go)。

# <翻译结束>


<原文开始>
Usage

For example you have architecture like this

![Example architecture](example_architecture.png)

To start requests tracing you have to:

* On "API Gateway": start span, inject headers and pass it to services

```go
package main
import (
    ...
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/opengintracing"
    "github.com/opentracing/opentracing-go"
    ...
)

var trace = /* setup tracer */

func main() {
    ...
    app := gin.Default()

    app.POST("/service1",
        opengintracing.NewSpan(trace, "forward to service 1"),
        opengintracing.InjectToHeaders(trace, true),
        service1handler)
    app.POST("/service2",
        opengintracing.NewSpan(trace, "forward to service 2"),
        opengintracing.InjectToHeaders(trace, true),
        service2handler)
    ...
}
```

* On "Service 1", "Service 2" start span inherited from "API Gateway"`s span

```go
package main
import (
    ...
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/opengintracing"
    "github.com/opentracing/opentracing-go"
    ...
)

var trace = /* setup tracer */

func main() {
    ...
    refFunc := opentracing.FollowsFrom
    app := gin.Default()

    app.POST("",
        opengintracing.SpanFromHeaders(trace, "operation", refFunc, true),
        // don`t forget to inject if you want continue tracing in other service
        opengintracing.InjectToHeaders(trace, true),
        handler)
    ...
}
```

Also don`t forget to forward headers from "Service 1" to "Service 3"

* On "Service 3" injecting to headers is not required

```go
package main
import (
    ...
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/opengintracing"
    "github.com/opentracing/opentracing-go"
    ...
)

var trace = /* setup tracer */

func main() {
    ...
    refFunc := opentracing.ChildOf
    app := gin.Default()

    app.POST("",
        opengintracing.SpanFromHeaders(trace, "operation", refFunc, true),
        handler)
    ...
}
```


<原文结束>

# <翻译开始>
# 用法

例如，您有这样的架构：

![](example_architecture.png)

要开始请求追踪，请按照以下步骤操作：

* 在“API Gateway”上：启动跨度（span），注入头部信息并将它们传递给服务

```go
package main
import (
    ...
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/opengintracing"
    "github.com/opentracing/opentracing-go"
    ...
)

var trace = /* 设置追踪器 */

func main() {
    ...
    app := gin.Default()

    app.POST("/service1",
        opengintracing.NewSpan(trace, "转发到服务 1"),
        opengintracing.InjectToHeaders(trace, true),
        service1handler)
    app.POST("/service2",
        opengintracing.NewSpan(trace, "转发到服务 2"),
        opengintracing.InjectToHeaders(trace, true),
        service2handler)
    ...
}
```

* 在“服务 1”和“服务 2”上，从“API Gateway”的跨度继承并启动新的跨度

```go
package main
import (
    ...
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/opengintracing"
    "github.com/opentracing/opentracing-go"
    ...
)

var trace = /* 设置追踪器 */

func main() {
    ...
    refFunc := opentracing.FollowsFrom
    app := gin.Default()

    app.POST("",
        opengintracing.SpanFromHeaders(trace, "操作", refFunc, true),
        // 如果要在其他服务中继续追踪，请不要忘记注入
        opengintracing.InjectToHeaders(trace, true),
        handler)
    ...
}
```

同时，别忘了将“服务 1”中的头部信息转发给“服务 3”

* 在“服务 3”上，不需要注入到头部信息

```go
package main
import (
    ...
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/opengintracing"
    "github.com/opentracing/opentracing-go"
    ...
)

var trace = /* 设置追踪器 */

func main() {
    ...
    refFunc := opentracing.ChildOf
    app := gin.Default()

    app.POST("",
        opengintracing.SpanFromHeaders(trace, "操作", refFunc, true),
        handler)
    ...
}
```

# <翻译结束>


<原文开始>
TODO

* [x] add code sample
* [ ] maybe add sample with SpanFromContext
* [ ] add buildable example (needed simple logging tracer)

<原文结束>

# <翻译开始>
# 待办事项：

* [x] 添加代码示例
* [ ] 考虑添加使用 SpanFromContext 的示例
* [ ] 添加可构建示例（需要一个简单的日志跟踪器）

# <翻译结束>

