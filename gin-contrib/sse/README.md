# # Server-Sent 事件

[![Go 参考](https://pkg.go.dev/badge/github.com/gin-contrib/sse.svg)](https://pkg.go.dev/github.com/gin-contrib/sse)
[![构建状态](https://github.com/gin-contrib/sse/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/sse/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/sse/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/sse)
[![Go 报告卡](https://goreportcard.com/badge/github.com/gin-contrib/sse)](https://goreportcard.com/report/github.com/gin-contrib/sse)

服务器发送事件（SSE）是一种技术，通过 HTTP 连接，浏览器可以从服务器自动接收更新。由 W3C 标准化组织作为 HTML5 的一部分标准化了 Server-Sent Events 的 EventSource API，参见[W3C 官方文档](http://www.w3.org/TR/2009/WD-eventsource-20091029/)。

- [阅读 HTML5Rocks 提供的优秀 SSE 入门教程](http://www.html5rocks.com/en/tutorials/eventsource/basics/)
- [浏览器兼容性](http://caniuse.com/#feat=eventsource)
## # 示例代码

```go
import "github.com/gin-contrib/sse"

func httpHandler(w http.ResponseWriter, req *http.Request) {
    // 数据可以是基本类型，如字符串、整数或浮点数
    sse.Encode(w, sse.Event{
        Event: "message",
        Data:  "some data\nmore data",
    })

    // 也可以是复杂类型，如映射、结构体或切片
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

对应的SSE（Server-Sent Events）格式输出：

```
event: message
data: some data\nmore data

id: 124
event: message
data: {"content":"hi!","date":1431540810,"user":"manu"}
```
## # 内容类型

```go
fmt.Println(sse.ContentType)
```
```
text/event-stream
```

翻译：
内容类型

```go
fmt.Println(sse.ContentType)
```
```
文本/事件流
```
## # 解码支持

客户端实现的SSE即将推出。
