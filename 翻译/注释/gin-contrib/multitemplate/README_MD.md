
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

这是一个自定义的 HTML 渲染器，支持多模板，即同时使用多个 `*template.Template`。

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

查看 [example/simple/example.go](example/simple/example.go)

```go
package main

import (
  "github.com/gin-contrib/multitemplate"
  "github.com/gin-gonic/gin"
)

// 创建自定义渲染器函数
func createMyRender() multitemplate.Renderer {
  r := multitemplate.NewRenderer()
  // 添加多个模板文件，构成"index"模版
  r.AddFromFiles("index", "templates/base.html", "templates/index.html")
  // 添加多个模板文件，构成"article"模版
  r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
  return r
}

func main() {
  // 初始化 Gin 路由器
  router := gin.Default()
  // 设置 HTML 渲染器为自定义的多模板渲染器
  router.HTMLRender = createMyRender()

  // 定义路由与处理函数
  router.GET("/", func(c *gin.Context) {
    // 使用 "index" 模板并传入参数，返回 HTML 响应
    c.HTML(200, "index", gin.H{
      "title": "Html5 Template Engine",
    })
  })

  router.GET("/article", func(c *gin.Context) {
    // 使用 "article" 模板并传入参数，返回 HTML 响应
    c.HTML(200, "article", gin.H{
      "title": "Html5 Article Engine",
    })
  })

  // 启动服务器监听 8080 端口
  router.Run(":8080")
}
```

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
# 高级示例

[近似模拟 html/template 继承](https://elithrar.github.io/article/approximating-html-template-inheritance/)

查看 [example/advanced/example.go](example/advanced/example.go)

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
      "title": "Html5 文章引擎",
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

这段代码是一个使用 Gin 框架和 multitemplate 库实现的 web 应用程序，它模拟了类似 html/template 的模板继承功能。`loadTemplates` 函数会加载指定目录（templatesDir）下的布局（layouts）和包含文件（includes），并构建一个模板渲染器。在主函数中，设置路由 GET 请求处理函数以渲染不同的 HTML 页面，并将渲染器设置到 Gin 路由器上，最后启动服务器监听 8080 端口。

# <翻译结束>

