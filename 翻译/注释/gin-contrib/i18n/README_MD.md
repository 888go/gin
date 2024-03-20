
<原文开始>
i18n

[![Run Tests](https://github.com/gin-contrib/i18n/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/i18n/actions/workflows/go.yml)
[![CodeQL](https://github.com/gin-contrib/i18n/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/gin-contrib/i18n/actions/workflows/codeql-analysis.yml)
[![codecov](https://codecov.io/gh/gin-contrib/i18n/branch/master/graph/badge.svg?token=QNMN3KM28Y)](https://codecov.io/gh/gin-contrib/i18n)
[![GoDoc](https://godoc.org/github.com/gin-contrib/i18n?status.svg)](https://godoc.org/github.com/gin-contrib/i18n)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/i18n)](https://goreportcard.com/report/github.com/gin-contrib/i18n)


<原文结束>

# <翻译开始>
# i18n

[![运行测试](https://github.com/gin-contrib/i18n/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/i18n/actions/workflows/go.yml)
[![CodeQL](https://github.com/gin-contrib/i18n/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/gin-contrib/i18n/actions/workflows/codeql-analysis.yml)
[![codecov](https://codecov.io/gh/gin-contrib/i18n/branch/master/graph/badge.svg?token=QNMN3KM28Y)](https://codecov.io/gh/gin-contrib/i18n)
[![GoDoc](https://godoc.org/github.com/gin-contrib/i18n?status.svg)](https://godoc.org/github.com/gin-contrib/i18n)
[![Go 项目评分](https://goreportcard.com/badge/github.com/gin-contrib/i18n)](https://goreportcard.com/report/github.com/gin-contrib/i18n)

以上内容是关于 Gin 框架贡献的 i18n（国际化）组件的README信息摘要，包含如下内容：

1. 运行测试：展示该项目在GitHub Actions上运行测试的状态。
2. CodeQL：展示项目的CodeQL代码扫描分析状态。
3. codecov：显示项目代码覆盖率的图形徽章，点击可查看详细覆盖率报告。
4. GoDoc：指向该项目在GoDoc上的API文档链接。
5. Go Report Card：显示该项目的Go语言质量报告，并提供了质量报告的访问链接。

# <翻译结束>


<原文开始>
Usage

Download and install it:

```sh
go get github.com/gin-contrib/i18n
```

Import it in your code:

```go
import ginI18n "github.com/gin-contrib/i18n"
```

Canonical example:

```go
package main

import (
  "log"
  "net/http"

  ginI18n "github.com/gin-contrib/i18n"
  "github.com/gin-gonic/gin"
  "github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
  // new gin engine
  gin.SetMode(gin.ReleaseMode)
  router := gin.New()

  // apply i18n middleware
  router.Use(ginI18n.Localize())

  router.GET("/", func(ctx *gin.Context) {
    ctx.String(http.StatusOK, ginI18n.MustGetMessage(ctx, "welcome"))
  })

  router.GET("/:name", func(ctx *gin.Context) {
    ctx.String(http.StatusOK, ginI18n.MustGetMessage(
      ctx,
      &i18n.LocalizeConfig{
        MessageID: "welcomeWithName",
        TemplateData: map[string]string{
          "name": ctx.Param("name"),
        },
      }))
  })

  if err := router.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Bundle

```go
package main

import (
  "encoding/json"
  "log"
  "net/http"

  ginI18n "github.com/gin-contrib/i18n"
  "github.com/gin-gonic/gin"
  "github.com/nicksnyder/go-i18n/v2/i18n"
  "golang.org/x/text/language"
)

func main() {
  // new gin engine
  gin.SetMode(gin.ReleaseMode)
  router := gin.New()

  // apply i18n middleware
  router.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
    RootPath:         "./testdata/localizeJSON",
    AcceptLanguage:   []language.Tag{language.German, language.English},
    DefaultLanguage:  language.English,
    UnmarshalFunc:    json.Unmarshal,
    FormatBundleFile: "json",
  })))

  router.GET("/", func(ctx *gin.Context) {
    ctx.String(http.StatusOK, ginI18n.MustGetMessage(ctx, "welcome"))
  })

  router.GET("/:name", func(ctx *gin.Context) {
    ctx.String(http.StatusOK, ginI18n.MustGetMessage(
      ctx,
      &i18n.LocalizeConfig{
        MessageID: "welcomeWithName",
        TemplateData: map[string]string{
          "name": ctx.Param("name"),
        },
      }))
  })

  if err := router.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Get Language Handler

```go
package main

import (
  "log"
  "net/http"

  ginI18n "github.com/gin-contrib/i18n"
  "github.com/gin-gonic/gin"
  "github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
  // new gin engine
  gin.SetMode(gin.ReleaseMode)
  router := gin.New()

  // apply i18n middleware
  router.Use(ginI18n.Localize(
    ginI18n.WithGetLngHandle(
      func(context *gin.Context, defaultLng string) string {
        lng := context.Query("lng")
        if lng == "" {
          return defaultLng
        }
        return lng
      },
    ),
  ))

  router.GET("/", func(ctx *gin.Context) {
    ctx.String(http.StatusOK, ginI18n.MustGetMessage(ctx, "welcome"))
  })

  router.GET("/:name", func(ctx *gin.Context) {
    ctx.String(http.StatusOK, ginI18n.MustGetMessage(
      ctx,
      &i18n.LocalizeConfig{
        MessageID: "welcomeWithName",
        TemplateData: map[string]string{
          "name": ctx.Param("name"),
        },
      }))
  })

  if err := router.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```


<原文结束>

# <翻译开始>
# 

# <翻译结束>


<原文开始>
License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.

<原文结束>

# <翻译开始>
# 许可证

本项目采用 MIT 许可证。您可以在 [LICENSE](LICENSE) 文件中查看完整的许可证文本。

# <翻译结束>

