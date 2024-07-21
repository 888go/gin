
<原文开始>
Multitemplate

[![Run Tests](https://github.com/gin-contrib/multitemplate/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/multitemplate/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/multitemplate/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/multitemplate)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/multitemplate)](https://goreportcard.com/report/github.com/gin-contrib/multitemplate)
[![GoDoc](https://godoc.org/github.com/gin-contrib/multitemplate?status.svg)](https://godoc.org/github.com/gin-contrib/multitemplate)

This is a custom HTML render to support multi templates, ie. more than one `*template.Template`.


<原文结束>

# <翻译开始>
# 多模板

[![运行测试](https://github.com/gin-contrib/multitemplate/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/multitemplate/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/multitemplate/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/multitemplate)
[![Go 项目报告卡](https://goreportcard.com/badge/github.com/gin-contrib/multitemplate)](https://goreportcard.com/report/github.com/gin-contrib/multitemplate)
[![GoDoc](https://godoc.org/github.com/gin-contrib/multitemplate?status.svg)](https://godoc.org/github.com/gin-contrib/multitemplate)

这是一个自定义的HTML渲染器，用于支持多个模板，即不止一个`*template.Template`。

# <翻译结束>


<原文开始>
Start using it

Download and install it:

```sh
go get github.com/gin-contrib/multitemplate
```

Import it in your code:

```go
import "github.com/gin-contrib/multitemplate"
```

#
<原文结束>

# <翻译开始>
# 开始使用

下载并安装：

```sh
go get github.com/gin-contrib/multitemplate
```

在代码中导入：

```go
import "github.com/gin-contrib/multitemplate"
```

#

# <翻译结束>


<原文开始>
Simple example

See [example/simple/example.go](example/simple/example.go)

```go
package main

import (
  "github.com/gin-contrib/multitemplate"
  "github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
  r := multitemplate.NewRenderer()
  r.AddFromFiles("index", "templates/base.html", "templates/index.html")
  r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
  return r
}

func main() {
  router := gin.Default()
  router.HTMLRender = createMyRender()
  router.GET("/", func(c *gin.Context) {
    c.HTML(200, "index", gin.H{
      "title": "Html5 Template Engine",
    })
  })
  router.GET("/article", func(c *gin.Context) {
    c.HTML(200, "article", gin.H{
      "title": "Html5 Article Engine",
    })
  })
  router.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 简单示例

参见 [example/simple/example.go](example/simple/example.go)

```go
package main

import (
  "github.com/gin-contrib/multitemplate"
  "github.com/gin-gonic/gin"
)

// 创建自定义渲染器函数
func createMyRender() multitemplate.Renderer {
  r := multitemplate.NewRenderer()
// 添加“index”模板，引用"templates/base.html"和"templates/index.html"
  r.AddFromFiles("index", "templates/base.html", "templates/index.html")
// 添加“article”模板，引用"templates/base.html"、"templates/index.html"和"templates/article.html"
  r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
  return r
}

func main() {
// 初始化默认的Gin路由器
  router := gin.Default()
// 设置HTML渲染器为自定义的渲染器
  router.HTMLRender = createMyRender()
// 定义GET请求根路由处理函数
  router.GET("/", func(c *gin.Context) {
// 使用渲染器返回200状态码及“index”模板，并传递参数
    c.HTML(200, "index", gin.H{
      "title": "Html5 Template Engine",
    })
  })
// 定义GET请求/article路由处理函数
  router.GET("/article", func(c *gin.Context) {
// 使用渲染器返回200状态码及“article”模板，并传递参数
    c.HTML(200, "article", gin.H{
      "title": "Html5 Article Engine",
    })
  })
// 运行服务器在8080端口
  router.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
Advanced example

[Approximating html/template Inheritance](https://elithrar.github.io/article/approximating-html-template-inheritance/)

See [example/advanced/example.go](example/advanced/example.go)

```go
package main

import (
  "path/filepath"

  "github.com/gin-contrib/multitemplate"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.HTMLRender = loadTemplates("./templates")
  router.GET("/", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
      "title": "Welcome!",
    })
  })
  router.GET("/article", func(c *gin.Context) {
    c.HTML(200, "article.html", gin.H{
      "title": "Html5 Article Engine",
    })
  })

  router.Run(":8080")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
  r := multitemplate.NewRenderer()

  layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
  if err != nil {
    panic(err.Error())
  }

  includes, err := filepath.Glob(templatesDir + "/includes/*.html")
  if err != nil {
    panic(err.Error())
  }

  // Generate our templates map from our layouts/ and includes/ directories
  for _, include := range includes {
    layoutCopy := make([]string, len(layouts))
    copy(layoutCopy, layouts)
    files := append(layoutCopy, include)
    r.AddFromFiles(filepath.Base(include), files...)
  }
  return r
}
```

<原文结束>

# <翻译开始>
# 进阶示例

[近似实现 html/template 继承](https:  //elithrar.github.io/article/approximating-html-template-inheritance/)

参见 [example/advanced/example.go](example/advanced/example.go)

```go
package main

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.HTMLRender = loadTemplates("./templates")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "欢迎!",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article.html", gin.H{
			"title": "Html5文章引擎",
		})
	})

	router.Run(":8080")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

/includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

  // 从 layouts/ 和 includes/ 目录生成模板映射
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
```

这段代码展示了一个 Gin 框架的进阶示例，其中使用了 `multitemplate` 包来处理 HTML 模板，并模拟实现了类似 html/template 的继承功能。`loadTemplates` 函数用于从指定目录加载布局（layouts）和包含文件（includes），并构建一个模板渲染器。在主函数中设置路由时，根据不同的路径返回不同页面的 HTML 响应，其中标题参数可自定义。最后运行服务器在 8080 端口监听请求。

# <翻译结束>

