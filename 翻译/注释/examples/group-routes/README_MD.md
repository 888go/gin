
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
func getRoutes() {
v1 := router.Group("/v1")
addUserRoutes(v1)
addPingRoutes(v1)

v2 := router.Group("/v2")
addPingRoutes(v2)
}
```
# <翻译结束>

