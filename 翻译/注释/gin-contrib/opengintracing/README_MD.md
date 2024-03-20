
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
# 追踪

[![运行测试](https://github.com/gin-contrib/opengintracing/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/opengintracing/actions/workflows/go.yml)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/opengintracing)](https://goreportcard.com/report/github.com/gin-contrib/opengintracing)
[![GoDoc](https://godoc.org/github.com/gin-contrib/opengintracing?status.png)](https://pkg.go.dev/github.com/gin-contrib/opengintracing)
[![许可协议: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

使用 opentracing 规范追踪请求，下载命令：

```bash
go get -u github.com/gin-contrib/opengintracing
```

欲了解更多信息，请查看 [opentracing/opentracing-go](https://github.com/opentracing/opentracing-go)。

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
# 用法示例

假设您有这样的架构设计：

![示例架构](example_architecture.png)

要开始追踪请求，您需要执行以下操作：

1. 在“API Gateway”上：

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

// 为转发到服务1的请求启动span，并注入headers
       app.POST("/service1",
           opengintracing.NewSpan(trace, "转发至服务1"),
           opengintracing.InjectToHeaders(trace, true),
           service1handler)

// 为转发到服务2的请求启动span，并注入headers
       app.POST("/service2",
           opengintracing.NewSpan(trace, "转发至服务2"),
           opengintracing.InjectToHeaders(trace, true),
           service2handler)
       ...
   }
   ```

2. 在“服务1”和“服务2”中：

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

// 从请求头中获取并继承自“API Gateway”的span
       app.POST("",
           opengintracing.SpanFromHeaders(trace, "操作", refFunc, true),
// 如果要在其他服务中继续追踪，请勿忘记注入headers
           opengintracing.InjectToHeaders(trace, true),
           handler)
       ...
   }
   ```

3. 同时，请确保将“服务1”中的headers传递给“服务3”。

4. 在“服务3”上，注入到headers不是必需的：

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

// 从请求头中获取并作为子span创建
       app.POST("",
           opengintracing.SpanFromHeaders(trace, "操作", refFunc, true),
           handler)
       ...
   }
   ```

以上代码示例展示了如何在各个服务中通过OpenTracing实现请求追踪。在API Gateway处创建并注入追踪信息到请求头，然后在后续的服务中继承或关联这些追踪信息以完成整个请求链路的追踪。

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
* [ ] 考虑添加包含SpanFromContext的示例
* [ ] 添加可构建示例（需要一个简单的日志跟踪器）

# <翻译结束>

