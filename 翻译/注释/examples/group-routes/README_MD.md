
<原文开始>
Group routes

This example shows how to group different routes in their own files and group them together in a orderly manner like this:

```go
func getRoutes() {
	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)
}
```

<原文结束>

# <翻译开始>
# 分组路由

此示例展示了如何将不同的路由分别放在各自的文件中，并以类似这样的有序方式将它们组合在一起：

```go
func 获取Routes() {
	v1 := 路由器.Group("/v1")
	添加用户路由(v1)
	添加Ping路由(v1)

	v2 := 路由器.Group("/v2")
	添加Ping路由(v2)
}
```

翻译后的内容（调整为更符合中文习惯的表达）：

该示例说明了如何将不同路由各自存放在独立的文件中，然后按照如下所示的有序方式进行分组整合：

```go
func 获取路由() {
	v1 := 路由器.NewGroup("/v1")
	添加用户相关路由(v1)
	添加Ping相关路由(v1)

	v2 := 路由器.NewGroup("/v2")
	添加Ping相关路由(v2)
}
```

# <翻译结束>

