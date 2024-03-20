
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

这个示例展示了如何将不同的路由各自放在独立的文件中，并以类似这样的有序方式将它们组合在一起：

```go
func 获取路由() {
	v1 := 路由器.Group("/v1")
	添加用户路由(v1)
	添加Ping路由(v1)

	v2 := 路由器.Group("/v2")
	添加Ping路由(v2)
}
```

（翻译说明：上述代码是Go语言编写的，对于变量和函数名称采用了意译的方式。具体含义如下）

此示例展示如何将不同路由分别组织在各自的文件中，并以一种有序结构进行组合，例如：

```go
// 定义获取路由函数
func getRoutes() {
// 创建一个以"/v1"为前缀的路由分组v1
	v1 := router.Group("/v1")
// 向v1分组中添加用户相关的路由规则
	addUserRoutes(v1)
// 向v1分组中添加Ping相关的路由规则
	addPingRoutes(v1)

// 创建一个以"/v2"为前缀的路由分组v2
	v2 := router.Group("/v2")
// 向v2分组中添加Ping相关的路由规则
	addPingRoutes(v2)
}
```

# <翻译结束>

