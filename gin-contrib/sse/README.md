# # Server-Sent 事件

[![Go 参考](https://pkg.go.dev/badge/github.com/gin-contrib/sse.svg)](https://pkg.go.dev/github.com/gin-contrib/sse)
[![构建状态](https://github.com/gin-contrib/sse/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/sse/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/sse/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/sse)
[![Go 报告卡](https://goreportcard.com/badge/github.com/gin-contrib/sse)](https://goreportcard.com/report/github.com/gin-contrib/sse)

服务器发送事件（SSE）是一种技术，通过HTTP连接，浏览器能自动接收来自服务器的更新。W3C已将Server-Sent Events的EventSource API作为HTML5规范的一部分[1]进行了标准化，详情参见：[http://www.w3.org/TR/2009/WD-eventsource-20091029/](http://www.w3.org/TR/2009/WD-eventsource-20091029/)。

- [阅读HTML5Rocks团队关于SSE的精彩介绍](http://www.html5rocks.com/en/tutorials/eventsource/basics/)
- [浏览器支持情况](http://caniuse.com/#feat=eventsource)
## # 示例代码

```go
import "github.com/gin-contrib/sse"

func httpHandler(w http.ResponseWriter, req *http.Request) {
// 数据可以是基本类型，如字符串、整数或浮点数
    sse.Encode(w, sse.Event{
        Event: "message",
        Data:  "some data\nmore data",
    })

// 也可以是复杂类型，如映射（map）、结构体或切片
    sse.Encode(w, sse.Event{
        Id:    "124",
        Event: "message",
        Data: map[string]interface{}{
            "user":    "manu",
            "date":    time.Now().Unix(),
            "content": "hi!",
        },
    })
}
```

该段代码的 SSE 事件输出示例：

```
event: message
data: some data\nmore data

id: 124
event: message
data: {"content":"hi!","date":1431540810,"user":"manu"}
```

这段 Go 代码导入了 "github.com/gin-contrib/sse" 包，并定义了一个名为 `httpHandler` 的 HTTP 处理函数。在该函数中，使用了 `sse.Encode` 方法将两种不同类型的 Server-Sent Events (SSE) 发送给客户端。一种是包含简单字符串数据的事件，另一种是包含复杂数据结构（一个映射）的事件。在输出示例中，展示了这两种事件以 SSE 格式序列化后的样子。
## # 内容类型

```go
fmt.Println(sse.ContentType)
```
```
text/event-stream
```

翻译为：

内容类型变量 `sse.ContentType` 输出的值为：
```
text/event-stream
```
## # 解码支持

即将推出客户端实现的SSE（Server-Sent Events，服务器发送事件）功能。
