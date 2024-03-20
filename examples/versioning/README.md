# # 版本控制

这是一个在gin-gonic框架中使用自定义中间件组实现的端点版本控制（例如 `/v1/path`）示例。
## # 如何运行？

1）执行命令：`go run main.go`

2）在`http://localhost:8080`测试API。

- 若要在v1版本中查看用户列表，路径应为`http://localhost:8080/v1/users`。同样地，所有其他路由也都可以正常使用。
