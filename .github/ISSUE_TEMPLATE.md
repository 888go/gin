# - 针对问题：
  - 在新建问题前，请先使用搜索工具。
  - 如果你发现了一个bug，请提供源代码和提交的SHA值。
  - 查阅现有问题，并提供反馈或对其作出回应。
## # 描述

<!-- 问题的描述 -->
## # 如何复现

<!-- 一个最小化的可编译代码示例，展示问题如下 -->
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/hello/:name", func(c *gin.Context) {
		c.String(200, "你好 %s", c.Param("name"))
	})
	g.Run(":9000")
}
```
这段代码是一个使用 Gin 框架创建的简单 Web 服务器。为了重现这个程序的行为，请按照以下步骤操作：

1. 确保已安装 Go 语言环境并在系统中配置好 GOPATH。
2. 在终端运行 `go get github.com/gin-gonic/gin` 来获取 Gin 框架。
3. 将上述代码粘贴到一个名为 `main.go` 的文件中。
4. 在终端进入包含 `main.go` 文件的目录，并运行 `go run main.go` 命令。
5. 一旦程序启动，它将在本地 9000 端口监听请求。访问 `http://localhost:9000/hello/yourname`（将 "yourname" 替换为您想要打招呼的名字），浏览器将显示 "Hello yourname"。
## # 预期结果

```shell
$ curl http://localhost:9000/hello/world
你好，世界
```

这段内容表示对`curl`命令执行后期望得到的结果。在这个示例中，当你在终端中运行`curl http://localhost:9000/hello/world`命令时，期望返回的内容是"Hello world"（译为“你好，世界”）。这表示当向本地服务器的9000端口发送一个GET请求到/hello/world路径时，服务器应返回“你好，世界”这个字符串作为响应。
## # 实际结果

```
$ curl -i http://localhost:9000/hello/world
<YOUR RESULT>
```

（此处的"YOUR RESULT"表示您通过执行上述命令后得到的实际输出结果）
## # 环境

- Go版本:
- Gin版本（或提交引用）:
- 操作系统：
