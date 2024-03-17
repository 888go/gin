# - 针对问题：
  - 在新建问题前，请先使用搜索工具。
  - 如果您发现了一个bug，请提供源代码和commit sha。
  - 查看现有问题并提供反馈或对其作出回应。
## # 描述

<!-- 问题的描述 -->
## # 如何复现

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/hello/:name", func(c *gin.Context) {
		c.String(200, "你好, %s", c.Param("name"))
	})
	g.Run(":9000")
}
```

这段代码展示了如何编写一个最小的可编译示例以体现问题。在该示例中，我们使用了 Gin 框架创建了一个 Web 服务器。服务器运行在 9000 端口，并提供了一个 GET 请求处理函数。当访问 "/hello/:name" 路由时，服务器会返回一个包含参数 "name" 值的消息，格式为 "你好, {name}"，其中状态码为 200。
## # 预期

```markdown
$ curl http://localhost:9000/hello/world
你好，世界
```

（此处为翻译结果，假设“Hello world”翻译为“你好，世界”）

这段内容的含义是：对于`curl`命令的预期执行结果是，当在终端中执行以下命令：

```shell
$ curl http://localhost:9000/hello/world
```

将会返回输出结果：“Hello world”，翻译成中文即为“你好，世界”。
## # 实际结果

```markdown
$ curl -i http://localhost:9000/hello/world
<YOUR RESULT>
```

（此处“<YOUR RESULT>”表示通过命令行获取的实际输出结果）
## # 环境

- Go版本:
- Gin版本（或提交引用）:
- 操作系统:
