# Gin Quick Start

## Contents

- [Build Tags](#build-tags)
  - [Build with json replacement](#build-with-json-replacement)
  - [Build without `MsgPack` rendering feature](#build-without-msgpack-rendering-feature)
- [API Examples](#api-examples)
  - [Using GET, POST, PUT, PATCH, DELETE and OPTIONS](#using-get-post-put-patch-delete-and-options)
  - [Parameters in path](#parameters-in-path)
  - [Querystring parameters](#querystring-parameters)
  - [Multipart/Urlencoded Form](#multiparturlencoded-form)
  - [Another example: query + post form](#another-example-query--post-form)
  - [Map as querystring or postform parameters](#map-as-querystring-or-postform-parameters)
  - [Upload files](#upload-files)
    - [Single file](#single-file)
    - [Multiple files](#multiple-files)
  - [Grouping routes](#grouping-routes)
  - [Blank Gin without middleware by default](#blank-gin-without-middleware-by-default)
  - [Using middleware](#using-middleware)
  - [Custom Recovery behavior](#custom-recovery-behavior)
  - [How to write log file](#how-to-write-log-file)
  - [Custom Log Format](#custom-log-format)
  - [Controlling Log output coloring](#controlling-log-output-coloring)
  - [Model binding and validation](#model-binding-and-validation)
  - [Custom Validators](#custom-validators)
  - [Only Bind Query String](#only-bind-query-string)
  - [Bind Query String or Post Data](#bind-query-string-or-post-data)
  - [Bind Uri](#bind-uri)
  - [Bind Header](#bind-header)
  - [Bind HTML checkboxes](#bind-html-checkboxes)
  - [Multipart/Urlencoded binding](#multiparturlencoded-binding)
  - [XML, JSON, YAML, TOML and ProtoBuf rendering](#xml-json-yaml-toml-and-protobuf-rendering)
    - [SecureJSON](#securejson)
    - [JSONP](#jsonp)
    - [AsciiJSON](#asciijson)
    - [PureJSON](#purejson)
  - [Serving static files](#serving-static-files)
  - [Serving data from file](#serving-data-from-file)
  - [Serving data from reader](#serving-data-from-reader)
  - [HTML rendering](#html-rendering)
    - [Custom Template renderer](#custom-template-renderer)
    - [Custom Delimiters](#custom-delimiters)
    - [Custom Template Funcs](#custom-template-funcs)
  - [Multitemplate](#multitemplate)
  - [Redirects](#redirects)
  - [Custom Middleware](#custom-middleware)
  - [Using BasicAuth() middleware](#using-basicauth-middleware)
  - [Goroutines inside a middleware](#goroutines-inside-a-middleware)
  - [Custom HTTP configuration](#custom-http-configuration)
  - [Support Let's Encrypt](#support-lets-encrypt)
  - [Run multiple service using Gin](#run-multiple-service-using-gin)
  - [Graceful shutdown or restart](#graceful-shutdown-or-restart)
    - [Third-party packages](#third-party-packages)
    - [Manually](#manually)
  - [Build a single binary with templates](#build-a-single-binary-with-templates)
  - [Bind form-data request with custom struct](#bind-form-data-request-with-custom-struct)
  - [Try to bind body into different structs](#try-to-bind-body-into-different-structs)
  - [Bind form-data request with custom struct and custom tag](#bind-form-data-request-with-custom-struct-and-custom-tag)
  - [http2 server push](#http2-server-push)
  - [Define format for the log of routes](#define-format-for-the-log-of-routes)
  - [Set and get a cookie](#set-and-get-a-cookie)
- [Don't trust all proxies](#dont-trust-all-proxies)
- [Testing](#testing)

## Build tags

### # 通过JSON替换构建

Gin默认使用`encoding/json`作为JSON包，但您可以通过使用其他标签进行构建来自定义它。

[jsoniter](https://github.com/json-iterator/go)

```sh
go build -tags=jsoniter .
```

[go-json](https://github.com/goccy/go-json)

```sh
go build -tags=go_json .
```

[sonic](https://github.com/bytedance/sonic)（您必须确保您的CPU支持avx指令。）

```sh
$ go build -tags="sonic avx" .
```

#
## # 不启用`MsgPack`渲染功能构建

Gin 默认启用了 `MsgPack` 渲染功能。但是，您可以通过指定 `nomsgpack` 构建标签来禁用此功能。

```sh
go build -tags=nomsgpack .
```

这对于减少可执行文件的二进制大小非常有用。详情请参阅[详细信息](https://github.com/gin-gonic/gin/pull/1852)。
## # API 示例

你可以在[Gin示例仓库](https://github.com/gin-gonic/examples)中找到大量可以直接运行的例子。

#
## # 使用GET, POST, PUT, PATCH, DELETE和OPTIONS

```go
func main() {
// 创建一个gin路由器，带有默认中间件：
// 日志记录和恢复（防崩溃）中间件
  router := gin.Default()

  router.GET("/someGet", getting)
  router.POST("/somePost", posting)
  router.PUT("/somePut", putting)
  router.DELETE("/someDelete", deleting)
  router.PATCH("/somePatch", patching)
  router.HEAD("/someHead", head)
  router.OPTIONS("/someOptions", options)

// 默认情况下，它在:8080端口提供服务，除非定义了
// 环境变量PORT。
  router.Run()
// 若要硬编码端口，使用router.Run(":3000")
}
```

#
## # 路径中的参数

```go
func main() {
  router := gin.Default()

// 此处理器将匹配 /user/john，但不会匹配 /user/ 或 /user
  router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "你好 %s", name)
  })

// 然而，这个处理器将匹配 /user/john/ 和 /user/john/send
// 如果没有其他路由器匹配 /user/john，则会重定向到 /user/john/
  router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " 正在执行 " + action
    c.String(http.StatusOK, message)
  })

// 对于每个匹配的请求，Context 将持有路由定义
  router.POST("/user/:name/*action", func(c *gin.Context) {
    b := c.FullPath() == "/user/:name/*action" // true
    c.String(http.StatusOK, "%t", b)
  })

// 此处理器将为 /user/groups 添加一个新的路由。
// 精确路由会在参数路由之前解析，无论它们定义的顺序如何。
// 以 /user/groups 开头的路由永远不会被解释为 /user/:name/... 路由
  router.GET("/user/groups", func(c *gin.Context) {
    c.String(http.StatusOK, "可用的组是 [...]")
  })

  router.Run(":8080")
}
```

#
## # 查询字符串参数

```go
func main() {
  // 初始化路由
  router := gin.Default()

// 使用现有底层请求对象解析查询字符串参数。
// 当请求的URL匹配：/welcome?firstname=Jane&lastname=Doe 时，该请求将会响应
  router.GET("/welcome", func(c *gin.Context) {
    // 获取查询参数"firstname"的值，若不存在则默认为"Guest"
    firstname := c.DefaultQuery("firstname", "Guest")
    // 快捷方式获取查询参数"lastname"的值，等同于 c.Request.URL.Query().Get("lastname")
    lastname := c.Query("lastname")

    // 返回状态码200，并向客户端发送消息 "Hello %s %s"，其中%s分别替换为firstname和lastname的值
    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
  })

  // 运行服务器在8080端口
  router.Run(":8080")
}
```
## # 多部分/Urlencoded 表单

```go
func main() {
  router := gin.Default()

  router.POST("/form_post", func(c *gin.Context) {
    // 获取表单中名为"message"的字段值
    message := c.PostForm("message")
    // 获取表单中名为"nick"的字段值，如果不存在则默认为"anonymous"
    nick := c.DefaultPostForm("nick", "anonymous")

    // 返回JSON响应，状态码为200（HTTP.StatusOK）
    c.JSON(http.StatusOK, gin.H{
      "status":  "posted",
      "message": message,
      "nick":    nick,
    })
  })

  // 启动服务器监听8080端口
  router.Run(":8080")
}
```
## # 另一个示例：查询 + POST 表单

```sh
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
```

```go
func main() {
  router := gin.Default()

  router.POST("/post", func(c *gin.Context) {

    id := c.Query("id")
    page := c.DefaultQuery("page", "0")
    name := c.PostForm("name")
    message := c.PostForm("message")

    fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
  })
  router.Run(":8080")
}
```

运行上述代码后，当向 `/post` 发送一个如上所示的 POST 请求时，输出将是：

```sh
id: 1234; page: 1; name: manu; message: this_is_great
```
## # 将查询字符串或post表单参数映射

```sh
POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded

names[first]=thinkerou&names[second]=tianou
```

```go
func main() {
  router := gin.Default()

  router.POST("/post", func(c *gin.Context) {

    // 获取查询字符串中的ids参数并映射为map
    ids := c.QueryMap("ids")
    // 获取post表单中的names参数并映射为map
    names := c.PostFormMap("names")

    fmt.Printf("ids: %v; names: %v", ids, names)
  })
  router.Run(":8080")
}
```

```sh
// 输出结果
ids: map[b:hello a:1234]; names: map[second:tianou first:thinkerou]
```

#
## Upload files

#### # 单文件上传

引用问题 [#774](https://github.com/gin-gonic/gin/issues/774) 和详细示例代码 [example code](https://github.com/gin-gonic/examples/tree/master/upload-file/single)。

`file.Filename` **不应**被信任。参见MDN上的 [`Content-Disposition`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition#Directives) 以及问题 [#1693](https://github.com/gin-gonic/gin/issues/1693)。

> 文件名始终是可选的，应用程序不应盲目使用：应剥离路径信息，并根据服务器文件系统规则进行转换。

```go
func main() {
    router := gin.Default()
    // 设置multipart表单的较小内存限制（默认为32 MiB）
    router.MaxMultipartMemory = 8 << 20 // 8 MiB
    router.POST("/upload", func(c *gin.Context) {
        // 单个文件
        file, _ := c.FormFile("file")
        log.Println(file.Filename)

        // 将文件上传到特定的目标位置(dst)
        c.SaveUploadedFile(file, dst)

        c.String(http.StatusOK, fmt.Sprintf("'%s' 已上传!", file.Filename))
    })
    router.Run(":8080")
}
```

如何使用 `curl` 进行上传：

```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/appleboy/test.zip" \
  -H "Content-Type: multipart/form-data"
```

##
## # 多个文件

请参阅[示例代码](https://github.com/gin-gonic/examples/tree/master/upload-file/multiple)的详细信息。

```go
func main() {
  router := gin.Default()
// 设置multipart表单的较低内存限制（默认为32 MiB）
  router.MaxMultipartMemory = 8 << 20 // 8 MiB
  router.POST("/upload", func(c *gin.Context) {
// 多部分表单
    form, _ := c.MultipartForm()
    files := form.File["upload[]"]

    for _, file := range files {
      log.Println(file.Filename)

      // 将文件上传到特定的目标路径。
      c.SaveUploadedFile(file, dst)
    }
    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
  })
  router.Run(":8080")
}
```

如何使用`curl`：

```bash
curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/appleboy/test1.zip" \
  -F "upload[]=@/Users/appleboy/test2.zip" \
  -H "Content-Type: multipart/form-data"
```

#
## # 分组路由

```go
func main() {
  router := gin.Default()

// 简单分组：v1
  v1 := router.Group("/v1")
  {
    v1.POST("/login", loginEndpoint)
    v1.POST("/submit", submitEndpoint)
    v1.POST("/read", readEndpoint)
  }

// 简单分组：v2
  v2 := router.Group("/v2")
  {
    v2.POST("/login", loginEndpoint)
    v2.POST("/submit", submitEndpoint)
    v2.POST("/read", readEndpoint)
  }

  router.Run(":8080")
}
```

#
## # 默认情况下，Blank Gin 不带中间件

使用方法：

```go
r := gin.New()
```

而非

```go
// 默认已附加 Logger 和 Recovery 中间件
r := gin.Default()
```

#
## # 使用中间件

```go
func main() {
  // 默认情况下创建一个不带任何中间件的路由器
  r := gin.New()

  // 全局中间件
  // Logger 中间件将日志写入 gin.DefaultWriter，即使你设置 GIN_MODE=release。默认情况下 gin.DefaultWriter = os.Stdout
  r.Use(gin.Logger())

  // Recovery 中间件从任何 panic 恢复，并在发生 panic 时写入一个 500 状态码。
  r.Use(gin.Recovery())

  // 路由级别中间件，你可以添加任意多的中间件。
  r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

  // 授权分组
  // authorized := r.Group("/", AuthRequired())
  // 完全等同于：
  authorized := r.Group("/")
  // 分组级别的中间件！在这个例子中，我们在 "authorized" 分组中仅使用自定义创建的
  // AuthRequired() 中间件。
  authorized.Use(AuthRequired())
  {
    authorized.POST("/login", loginEndpoint)
    authorized.POST("/submit", submitEndpoint)
    authorized.POST("/read", readEndpoint)

    // 嵌套分组
    testing := authorized.Group("testing")
    // 访问 0.0.0.0:8080/testing/analytics
    testing.GET("/analytics", analyticsEndpoint)
  }

  // 监听并服务在 0.0.0.0:8080
  r.Run(":8080")
}
```

#
## # 自定义恢复行为

```go
func main() {
    // 创建一个默认不包含任何中间件的路由器
    r := gin.New()

    // 全局中间件
    // Logger 中间件会将日志写入 gin.DefaultWriter，即使你设置 GIN_MODE=release 也是如此。
    // 默认情况下 gin.DefaultWriter = os.Stdout
    r.Use(gin.Logger())

    // Recovery 中间件可以从任何 panic 恢复，并在发生 panic 时返回 500 状态码。
    r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
        if err, ok := recovered.(string); ok {
            c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
        }
        c.AbortWithStatus(http.StatusInternalServerError)
    }))

    r.GET("/panic", func(c *gin.Context) {
        // 使用字符串触发 panic —— 自定义中间件可以将此信息保存到数据库或报告给用户
        panic("foo")
    })

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "ohai")
    })

    // 监听并服务 0.0.0.0:8080
    r.Run(":8080")
}
```

#
## How to write log file

```go
func main() {
  // Disable Console Color, you don't need console color when writing the logs to file.
  gin.DisableConsoleColor()

  // Logging to a file.
  f, _ := os.Create("gin.log")
  gin.DefaultWriter = io.MultiWriter(f)

  // Use the following code if you need to write the logs to file and console at the same time.
  // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

聽 聽router.Run(":8080")
}
```

### Custom Log Format

```go
func main() {
  router := gin.New()

  // LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
  // By default gin.DefaultWriter = os.Stdout
  router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

    // your custom format
    return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
        param.ClientIP,
        param.TimeStamp.Format(time.RFC1123),
        param.Method,
        param.Path,
        param.Request.Proto,
        param.StatusCode,
        param.Latency,
        param.Request.UserAgent(),
        param.ErrorMessage,
    )
  }))
  router.Use(gin.Recovery())

  router.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```

Sample Output

```sh
::1 - [Fri, 07 Dec 2018 17:04:38 JST] "GET /ping HTTP/1.1 200 122.767碌s "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Safari/537.36" "
```

### # 控制日志输出颜色

默认情况下，根据检测到的TTY，应在控制台上对日志输出进行着色。

从不为日志着色：

```go
func main() {
// 禁用日志的颜色
  gin.DisableConsoleColor()

// 创建一个gin路由器，并使用默认中间件：
// 日志记录和恢复（防崩溃）中间件
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```

始终为日志着色：

```go
func main() {
// 强制日志的颜色
  gin.ForceConsoleColor()

// 创建一个gin路由器，并使用默认中间件：
// 日志记录和恢复（防崩溃）中间件
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```
## Model binding and validation

To bind a request body into a type, use model binding. We currently support binding of JSON, XML, YAML, TOML and standard form values (foo=bar&boo=baz).

Gin uses [**go-playground/validator/v10**](https://github.com/go-playground/validator) for validation. Check the full docs on tags usage [here](https://pkg.go.dev/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags).

Note that you need to set the corresponding binding tag on all fields you want to bind. For example, when binding from JSON, set `json:"fieldname"`.

Also, Gin provides two sets of methods for binding:

- **Type** - Must bind
  - **Methods** - `Bind`, `BindJSON`, `BindXML`, `BindQuery`, `BindYAML`, `BindHeader`, `BindTOML`
  - **Behavior** - These methods use `MustBindWith` under the hood. If there is a binding error, the request is aborted with `c.AbortWithError(400, err).SetType(ErrorTypeBind)`. This sets the response status code to 400 and the `Content-Type` header is set to `text/plain; charset=utf-8`. Note that if you try to set the response code after this, it will result in a warning `[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422`. If you wish to have greater control over the behavior, consider using the `ShouldBind` equivalent method.
- **Type** - Should bind
  - **Methods** - `ShouldBind`, `ShouldBindJSON`, `ShouldBindXML`, `ShouldBindQuery`, `ShouldBindYAML`, `ShouldBindHeader`, `ShouldBindTOML`,
  - **Behavior** - These methods use `ShouldBindWith` under the hood. If there is a binding error, the error is returned and it is the developer's responsibility to handle the request and error appropriately.

When using the Bind-method, Gin tries to infer the binder depending on the Content-Type header. If you are sure what you are binding, you can use `MustBindWith` or `ShouldBindWith`.

You can also specify that specific fields are required. If a field is decorated with `binding:"required"` and has an empty value when binding, an error will be returned.

```go
// Binding from JSON
type Login struct {
  User     string `form:"user" json:"user" xml:"user"  binding:"required"`
  Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
  router := gin.Default()

  // Example for binding JSON ({"user": "manu", "password": "123"})
  router.POST("/loginJSON", func(c *gin.Context) {
    var json Login
    if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if json.User != "manu" || json.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // Example for binding XML (
  //  <?xml version="1.0" encoding="UTF-8"?>
  //  <root>
  //    <user>manu</user>
  //    <password>123</password>
  //  </root>)
  router.POST("/loginXML", func(c *gin.Context) {
    var xml Login
    if err := c.ShouldBindXML(&xml); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if xml.User != "manu" || xml.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // Example for binding a HTML form (user=manu&password=123)
  router.POST("/loginForm", func(c *gin.Context) {
    var form Login
    // This will infer what binder to use depending on the content-type header.
    if err := c.ShouldBind(&form); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if form.User != "manu" || form.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // Listen and serve on 0.0.0.0:8080
  router.Run(":8080")
}
```

Sample request

```sh
$ curl -v -X POST \
  http://localhost:8080/loginJSON \
  -H 'content-type: application/json' \
  -d '{ "user": "manu" }'
> POST /loginJSON HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.51.0
> Accept: */*
> content-type: application/json
> Content-Length: 18
>
* upload completely sent off: 18 out of 18 bytes
< HTTP/1.1 400 Bad Request
< Content-Type: application/json; charset=utf-8
< Date: Fri, 04 Aug 2017 03:51:31 GMT
< Content-Length: 100
<
{"error":"Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
```

Skip validate: when running the above example using the above the `curl` command, it returns error. Because the example use `binding:"required"` for `Password`. If use `binding:"-"` for `Password`, then it will not return error when running the above example again.

### Custom Validators

It is also possible to register custom validators. See the [example code](https://github.com/gin-gonic/examples/tree/master/custom-validation/server.go).

```go
package main

import (
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
  "github.com/go-playground/validator/v10"
)

// Booking contains binded and validated data.
type Booking struct {
  CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
  CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
  date, ok := fl.Field().Interface().(time.Time)
  if ok {
    today := time.Now()
    if today.After(date) {
      return false
    }
  }
  return true
}

func main() {
  route := gin.Default()

  if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    v.RegisterValidation("bookabledate", bookableDate)
  }

  route.GET("/bookable", getBookable)
  route.Run(":8085")
}

func getBookable(c *gin.Context) {
  var b Booking
  if err := c.ShouldBindWith(&b, binding.Query); err == nil {
    c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
  } else {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
}
```

```console
$ curl "localhost:8085/bookable?check_in=2030-04-16&check_out=2030-04-17"
{"message":"Booking dates are valid!"}

$ curl "localhost:8085/bookable?check_in=2030-03-10&check_out=2030-03-09"
{"error":"Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"}

$ curl "localhost:8085/bookable?check_in=2000-03-09&check_out=2000-03-10"
{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookabledate' tag"}%
```

[Struct level validations](https://github.com/go-playground/validator/releases/tag/v8.7) can also be registered this way.
See the [struct-lvl-validation example](https://github.com/gin-gonic/examples/tree/master/struct-lvl-validations) to learn more.

### # 仅绑定查询字符串

`ShouldBindQuery` 函数仅绑定查询参数而不绑定 post 数据。有关详细信息，请参阅 [此处](https://github.com/gin-gonic/gin/issues/742#issuecomment-315953017)。

```go
package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

type Person struct {
  Name    string `form:"name"`
  Address string `form:"address"`
}

func main() {
  route := gin.Default()
  route.Any("/testing", startPage)
  route.Run(":8085")
}

func startPage(c *gin.Context) {
  var person Person
  if c.ShouldBindQuery(&person) == nil {
    log.Println("====== 仅通过查询字符串绑定 ======")
    log.Println(person.Name)
    log.Println(person.Address)
  }
  c.String(http.StatusOK, "Success")
}
```

#
## # 查询字符串或 POST 数据绑定

查看[详细信息](https://github.com/gin-gonic/gin/issues/742#issuecomment-264681292)。

```go
package main

import (
  "log"
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
)

type Person struct {
  Name       string    `form:"name"`
  Address    string    `form:"address"`
  Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
  CreateTime time.Time `form:"createTime" time_format:"unixNano"`
  UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
  router := gin.Default()
  router.GET("/testing", startPage)
  router.Run(":8085")
}

func startPage(c *gin.Context) {
  var person Person
// 如果是 `GET` 请求，仅使用 `Form` 绑定引擎（查询参数）。
// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后使用 `Form`（表单数据）进行绑定。
// 查看更多内容请访问：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
  if c.ShouldBind(&person) == nil {
    log.Println(person.Name)
    log.Println(person.Address)
    log.Println(person.Birthday)
    log.Println(person.CreateTime)
    log.Println(person.UnixTime)
  }

  c.String(http.StatusOK, "Success")
}
```

通过以下命令进行测试：

```sh
curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
```
## # URI 绑定

查看[详细信息](https://github.com/gin-gonic/gin/issues/846)。

```go
package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type Person struct {
  ID   string `uri:"id" binding:"required,uuid"`
  Name string `uri:"name" binding:"required"`
}

func main() {
  router := gin.Default()
  router.GET("/:name/:id", func(c *gin.Context) {
    var person Person
    if err := c.ShouldBindUri(&person); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
      return
    }
    c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
  })
  router.Run(":8088")
}
```

使用以下命令进行测试：

```sh
curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
curl -v localhost:8088/thinkerou/not-uuid
```
## # 标题：绑定请求头

```go
package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

// 定义测试请求头结构体
type testHeader struct {
  Rate   int    `header:"Rate"` // 绑定Rate字段到请求头的"Rate"键值
  Domain string `header:"Domain"` // 绑定Domain字段到请求头的"Domain"键值
}

func main() {
  // 初始化 Gin 框架
  r := gin.Default()

  // 设置GET路由处理函数
  r.GET("/", func(c *gin.Context) {
    // 创建一个testHeader实例
    h := testHeader{}

    // 尝试从请求头中绑定数据到h
    if err := c.ShouldBindHeader(&h); err != nil {
      // 绑定失败则返回错误信息
      c.JSON(http.StatusOK, err)
    } else {
      // 打印已绑定的数据
      fmt.Printf("%#v\n", h)

      // 返回JSON格式响应，包含已绑定的Rate和Domain信息
      c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
    }
  })

  // 运行服务器
  r.Run()

// 客户端示例请求：
// curl -H "rate:300" -H "domain:music" 127.0.0.1:8080/
// 预期输出：
// {"Domain":"music","Rate":300}
}
```
## # 绑定HTML复选框

参见[详细信息](https://github.com/gin-gonic/gin/issues/129#issuecomment-124260092)

main.go

```go
...

// 定义表单结构体
type myForm struct {
    Colors []string `form:"colors[]"`
}

...

// 处理表单的函数
func formHandler(c *gin.Context) {
    // 初始化表单实例
    var fakeForm myForm
    // 绑定请求数据到表单结构体
    c.ShouldBind(&fakeForm)
    // 返回JSON响应，包含选中的颜色
    c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}

...

```

form.html

```html
<!-- 表单内容 -->
<form action="/" method="POST">
    <p>选择一些颜色</p>
    <label for="red">红色</label>
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">绿色</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">蓝色</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit" value="提交">
</form>
```

结果：

```json
{"color":["red","green","blue"]}
```
## # 多部分/Urlencoded 绑定

```go
type ProfileForm struct {
  Name    string                `form:"name" binding:"required"`
  Avatar  *multipart.FileHeader `form:"avatar" binding:"required"`

// 或者，如果是多个文件
// Avatars []*multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
  router := gin.Default()
  router.POST("/profile", func(c *gin.Context) {
// 可以通过显式绑定声明来绑定多部分表单：
// c.ShouldBindWith(&form, binding.Form)
// 或者可以简单地使用自动绑定 ShouldBind 方法：
    var form ProfileForm
// 在这种情况下，将会自动选择合适的绑定方式
    if err := c.ShouldBind(&form); err != nil {
      c.String(http.StatusBadRequest, "bad request")
      return
    }

    err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
    if err != nil {
      c.String(http.StatusInternalServerError, "unknown error")
      return
    }

// db.Save(&form)

    c.String(http.StatusOK, "ok")
  })
  router.Run(":8080")
}
```

使用以下命令测试：

```sh
curl -X POST -v --form name=user --form "avatar=@./avatar.png" http://localhost:8080/profile
```

#
## XML, JSON, YAML, TOML and ProtoBuf rendering

```go
func main() {
  r := gin.Default()

  // gin.H is a shortcut for map[string]any
  r.GET("/someJSON", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/moreJSON", func(c *gin.Context) {
    // You also can use a struct
    var msg struct {
      Name    string `json:"user"`
      Message string
      Number  int
    }
    msg.Name = "Lena"
    msg.Message = "hey"
    msg.Number = 123
    // Note that msg.Name becomes "user" in the JSON
    // Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
    c.JSON(http.StatusOK, msg)
  })

  r.GET("/someXML", func(c *gin.Context) {
    c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/someYAML", func(c *gin.Context) {
    c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/someTOML", func(c *gin.Context) {
    c.TOML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/someProtoBuf", func(c *gin.Context) {
    reps := []int64{int64(1), int64(2)}
    label := "test"
    // The specific definition of protobuf is written in the testdata/protoexample file.
    data := &protoexample.Test{
      Label: &label,
      Reps:  reps,
    }
    // Note that data becomes binary data in the response
    // Will output protoexample.Test protobuf serialized data
    c.ProtoBuf(http.StatusOK, data)
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

#### SecureJSON

Using SecureJSON to prevent json hijacking. Default prepends `"while(1),"` to response body if the given struct is array values.

```go
func main() {
  r := gin.Default()

  // You can also use your own secure json prefix
  // r.SecureJsonPrefix(")]}',\n")

  r.GET("/someJSON", func(c *gin.Context) {
    names := []string{"lena", "austin", "foo"}

    // Will output  :   while(1);["lena","austin","foo"]
    c.SecureJSON(http.StatusOK, names)
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

#### JSONP

Using JSONP to request data from a server  in a different domain. Add callback to response body if the query parameter callback exists.

```go
func main() {
  r := gin.Default()

  r.GET("/JSONP", func(c *gin.Context) {
    data := gin.H{
      "foo": "bar",
    }

    //callback is x
    // Will output  :   x({\"foo\":\"bar\"})
    c.JSONP(http.StatusOK, data)
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")

        // client
        // curl http://127.0.0.1:8080/JSONP?callback=x
}
```

#### AsciiJSON

Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.

```go
func main() {
  r := gin.Default()

  r.GET("/someJSON", func(c *gin.Context) {
    data := gin.H{
      "lang": "GO璇█",
      "tag":  "<br>",
    }

    // will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
    c.AsciiJSON(http.StatusOK, data)
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

#### # PureJSON

通常情况下，JSON会将特殊的HTML字符替换为它们的Unicode实体，例如 `<` 会变为 `\u003c`。如果你想直接编码这些字符，你可以使用PureJSON。

此功能在Go 1.6及更低版本中不可用。

```go
func main() {
  r := gin.Default()

// 服务端提供Unicode实体
  r.GET("/json", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

// 服务端提供原始字符
  r.GET("/purejson", func(c *gin.Context) {
    c.PureJSON(http.StatusOK, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

// 监听并服务在0.0.0.0:8080端口
  r.Run(":8080")
}
```

#
## # 服务静态文件

```go
func main() {
  // 初始化路由
  router := gin.Default()

  // 设置静态文件目录，访问路径为"/assets"，本地文件夹为"./assets"
  router.Static("/assets", "./assets")

  // 设置静态文件系统，访问路径为"/more_static"，使用http.Dir指定文件系统目录为"my_file_system"
  router.StaticFS("/more_static", http.Dir("my_file_system"))

  // 直接服务单个静态文件，访问路径为"/favicon.ico"，对应本地文件为"./resources/favicon.ico"
  router.StaticFile("/favicon.ico", "./resources/favicon.ico")

  // 通过文件系统服务单个静态文件，访问路径为"/more_favicon.ico"，文件名为"more_favicon.ico"，在指定的文件系统目录"http.Dir("my_file_system")"中查找
  router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))

  // 监听并服务0.0.0.0:8080端口
  router.Run(":8080")
}
```

#
## # 从文件中提供数据

```go
func main() {
  router := gin.Default()

  // 从本地文件系统提供文件
  router.GET("/local/file", func(c *gin.Context) {
    c.File("local/file.go")
  })

  // 从自定义文件系统提供文件
  var fs http.FileSystem = // ...
  router.GET("/fs/file", func(c *gin.Context) {
    c.FileFromFS("fs/file.go", fs)
  })
}
```

#
## # 从读取器提供数据

```go
func main() {
  router := gin.Default()
  router.GET("/someDataFromReader", func(c *gin.Context) {
    // 发送GET请求获取数据
    response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
    if err != nil || response.StatusCode != http.StatusOK {
      // 如果发生错误或返回状态码不是200，设置服务不可用状态码并返回
      c.Status(http.StatusServiceUnavailable)
      return
    }

    // 获取响应体并设置关闭函数
    reader := response.Body
    defer reader.Close()

    // 获取内容长度和内容类型
    contentLength := response.ContentLength
    contentType := response.Header.Get("Content-Type")

    // 添加额外的HTTP头信息
    extraHeaders := map[string]string{
      "Content-Disposition": `attachment; filename="gopher.png"`,
    }

    // 使用DataFromReader方法将数据从读取器发送到客户端，并附带额外的头部信息
    c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
  })
  
  // 启动服务器监听8080端口
  router.Run(":8080")
}
```

#
## # HTML 渲染

使用 LoadHTMLGlob() 或 LoadHTMLFiles()

```go
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*") // 或者：router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
  router.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "主网站",
    })
  })
  router.Run(":8080")
}
```

templates/index.tmpl

```html
<html>
  <h1>
    {{ .title }}
  </h1>
</html>
```

在不同目录下使用相同名称的模板

```go
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/**/*")
  router.GET("/posts/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
      "title": "文章",
    })
  })
  router.GET("/users/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
      "title": "用户",
    })
  })
  router.Run(":8080")
}
```

templates/posts/index.tmpl

```html
{{ define "posts/index.tmpl" }}
<html><h1>
  {{ .title }}
</h1>
<p>使用 posts/index.tmpl</p>
</html>
{{ end }}
```

templates/users/index.tmpl

```html
{{ define "users/index.tmpl" }}
<html><h1>
  {{ .title }}
</h1>
<p>使用 users/index.tmpl</p>
</html>
{{ end }}
```

##
## # 自定义模板渲染

您也可以使用自己的 HTML 模板渲染器。

```go
import "html/template"

func main() {
  // 初始化路由
  router := gin.Default()

  // 加载并设置 HTML 模板
  html := template.Must(template.ParseFiles("file1", "file2"))
  router.SetHTMLTemplate(html)

  // 运行服务器在 8080 端口
  router.Run(":8080")
}
```

##
## # 自定义分隔符

您可以使用自定义分隔符

```go
r := gin.Default()
r.Delims("{[{", "}]}")
r.LoadHTMLGlob("/path/to/templates")
```

##
## # 自定义模板函数

请参阅详细示例代码：[https://github.com/gin-gonic/examples/tree/master/template](https://github.com/gin-gonic/examples/tree/master/template)。

main.go

```go
import (
    "fmt"
    "html/template"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
    year, month, day := t.Date()
    return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
    router := gin.Default()
    router.Delims("{[{", "}]}") // 设置模板分隔符
    router.SetFuncMap(template.FuncMap{
        "formatAsDate": formatAsDate, // 注册自定义模板函数
    })
    router.LoadHTMLFiles("./testdata/template/raw.tmpl")

    router.GET("/raw", func(c *gin.Context) {
        c.HTML(http.StatusOK, "raw.tmpl", gin.H{
            "now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
        })
    })

    router.Run(":8080")
}
```

raw.tmpl

```html
Date: {[{.now | formatAsDate}]}
```

结果：

```sh
Date: 2017/07/01
```
## # 多模板

Gin 默认只允许使用一个 html.Template。若要使用类似 Go 1.6 中的 `block template` 功能，请查看 [多模板渲染](https://github.com/gin-contrib/multitemplate)。
## # 重定向

发出HTTP重定向很容易，支持内部和外部位置。

```go
r.GET("/test", func(c *gin.Context) {
  c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
})
```

从POST发出HTTP重定向。参考问题：[#444](https://github.com/gin-gonic/gin/issues/444)

```go
r.POST("/test", func(c *gin.Context) {
  c.Redirect(http.StatusFound, "/foo")
})
```

发出路由重定向，像下面这样使用`HandleContext`。

```go
r.GET("/test", func(c *gin.Context) {
    c.Request.URL.Path = "/test2"
    r.HandleContext(c)
})
r.GET("/test2", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"hello": "world"})
})
```

#
## # 自定义中间件

```go
func Logger() gin.HandlerFunc {
  return func(c *gin.Context) {
    t := time.Now()

// 设置示例变量
    c.Set("example", "12345")

// 请求之前

    c.Next()

// 请求之后
    latency := time.Since(t)
    log.Print(latency)

// 访问我们发送的状态
    status := c.Writer.Status()
    log.Println(status)
  }
}

func main() {
  r := gin.New()
  r.Use(Logger())

  r.GET("/test", func(c *gin.Context) {
    example := c.MustGet("example").(string)

// 将打印: "12345"
    log.Println(example)
  })

// 监听并服务在 0.0.0.0:8080
  r.Run(":8080")
}
```

#
## # 使用 BasicAuth() 中间件

```go
// 模拟一些私密数据
var secrets = gin.H{
  "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
  "austin": gin.H{"email": "austin@example.com", "phone": "666"},
  "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
  r := gin.Default()

// 使用 gin.BasicAuth() 中间件的分组
// gin.Accounts 是一个快捷方式，用于 map[string]string
  authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
    "foo":    "bar",
    "austin": "1234",
    "lena":   "hello2",
    "manu":   "4321",
  }))

// 访问 "/admin/secrets" 端点
// 在浏览器中输入 "localhost:8080/admin/secrets"
  authorized.GET("/secrets", func(c *gin.Context) {
// 获取用户信息，该信息由 BasicAuth 中间件设置
    user := c.MustGet(gin.AuthUserKey).(string)
    if secret, ok := secrets[user]; ok {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
    } else {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": "无秘密信息 :("})
    }
  })

// 监听并服务在 0.0.0.0:8080
  r.Run(":8080")
}
```

#
## # 在中间件或处理器内部启动新的 Goroutines 时，你不 **应该** 在其中使用原始上下文，而必须使用只读副本。

```go
func main() {
  r := gin.Default()

  r.GET("/long_async", func(c *gin.Context) {
    // 创建一个副本以供 Goroutine 内部使用
    cCp := c.Copy()
    go func() {
      // 模拟一个耗时任务，使用 time.Sleep()，持续5秒
      time.Sleep(5 * time.Second)

      // 注意这里你正在使用复制的上下文 "cCp"，这一点很重要
      log.Println("完成！路径为：" + cCp.Request.URL.Path)
    }()
  })

  r.GET("/long_sync", func(c *gin.Context) {
    // 模拟一个耗时任务，使用 time.Sleep()，持续5秒
    time.Sleep(5 * time.Second)

    // 由于我们没有使用 Goroutine，所以不需要复制上下文
    log.Println("完成！路径为：" + c.Request.URL.Path)
  })

  // 监听并服务在 0.0.0.0:8080
  r.Run(":8080")
}
```

#
## # 自定义HTTP配置

直接使用`http.ListenAndServe()`，如下所示：

```go
func main() {
  router := gin.Default()
  http.ListenAndServe(":8080", router)
}
```

或者

```go
func main() {
  router := gin.Default()

  s := &http.Server{
    Addr:           ":8080",
    Handler:        router,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  s.ListenAndServe()
}
```

#
## # 支持 Let's Encrypt

单行配置 Let's Encrypt HTTPS 服务器的示例。

```go
package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/autotls"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

// Ping 处理器
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
```

自定义 autocert 管理器的示例。

```go
package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/autotls"
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/acme/autocert"
)

func main() {
  r := gin.Default()

// Ping 处理器
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  m := autocert.Manager{
    Prompt:     autocert.AcceptTOS, // 接受服务条款
    HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"), // 主机白名单策略
    Cache:      autocert.DirCache("/var/www/.cache"), // 缓存证书到指定目录
  }

  log.Fatal(autotls.RunWithManager(r, &m))
}
```

#
## # 使用 Gin 运行多个服务

参见 [问题](https://github.com/gin-gonic/gin/issues/346)，并尝试以下示例：

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "欢迎来到服务器01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "欢迎来到服务器02",
			},
		)
	})

	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := server02.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
```

#
## # 优雅地关闭或重启

您可以采用几种方法来实现优雅的关闭或重启。您可以利用专门为该目的而构建的第三方包，也可以手动使用内置包中的函数和方法来完成同样的操作。

##
## # 第三方包

我们可以使用[fvbock/endless](https://github.com/fvbock/endless) 替换默认的 `ListenAndServe`。有关更多详细信息，请参阅问题 [#296](https://github.com/gin-gonic/gin/issues/296)。

```go
router := gin.Default()
router.GET("/", handler)
// [...]
endless.ListenAndServe(":4242", router)
```

替代方案：

* [grace](https://github.com/facebookgo/grace)：用于 Go 服务器的优雅重启与零停机部署。
* [graceful](https://github.com/tylerb/graceful)：Graceful 是一个 Go 包，可实现 http.Handler 服务器的优雅关闭。
* [manners](https://github.com/braintree/manners)：一个有礼貌的 Go HTTP 服务器，能够优雅地关闭。
## # 手动操作

如果您使用的是 Go 1.8 或更高版本，可能不需要使用那些库。考虑使用 `http.Server` 内置的 [Shutdown()](https://pkg.go.dev/net/http#Server.Shutdown) 方法进行优雅关闭。下面的示例展示了其用法，并且我们在此处（<https://github.com/gin-gonic/examples/tree/master/graceful-shutdown>）提供了更多使用 gin 的示例。

```go
// +build go1.8

package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
    time.Sleep(5 * time.Second)
    c.String(http.StatusOK, "欢迎来到 Gin Server")
  })

  srv := &http.Server{
    Addr:    ":8080",
    Handler: router,
  }

// 在一个 goroutine 中初始化服务器，以便在下面处理优雅关闭时不会阻塞
  go func() {
    if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
      log.Printf("监听错误: %s\n", err)
    }
  }()

// 等待中断信号以在 5 秒超时的情况下优雅地关闭服务器
  quit := make(chan os.Signal)
// kill (不带参数) 默认发送 syscall.SIGTERM
// kill -2 是 syscall.SIGINT
// kill -9 是 syscall.SIGKILL 但无法捕获，所以不需要添加
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  log.Println("正在关闭服务器...")

// 使用上下文通知服务器它有 5 秒钟的时间来完成当前正在处理的请求
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  if err := srv.Shutdown(ctx); err != nil {
    log.Fatal("服务器被迫强制关闭:", err)
  }

  log.Println("服务器退出")
}
```

#
## # 构建包含模板的单个二进制文件

您可以通过使用 [embed](https://pkg.go.dev/embed) 包将服务器构建为包含模板的单个二进制文件。

```go
package main

import (
  "embed"
  "html/template"
  "net/http"

  "github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS

func main() {
  router := gin.Default()
  templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
  router.SetHTMLTemplate(templ)

// 示例：/public/assets/images/example.png
  router.StaticFS("/public", http.FS(f))

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "主网站",
    })
  })

  router.GET("/foo", func(c *gin.Context) {
    c.HTML(http.StatusOK, "bar.tmpl", gin.H{
      "title": "Foo 网站",
    })
  })

  router.GET("favicon.ico", func(c *gin.Context) {
    file, _ := f.ReadFile("assets/favicon.ico")
    c.Data(
      http.StatusOK,
      "image/x-icon",
      file,
    )
  })

  router.Run(":8080")
}

```

在 `https://github.com/gin-gonic/examples/tree/master/assets-in-binary/example02` 目录中查看完整示例。
## Bind form-data request with custom struct

The follow example using custom struct:

```go
type StructA struct {
    FieldA string `form:"field_a"`
}

type StructB struct {
    NestedStruct StructA
    FieldB string `form:"field_b"`
}

type StructC struct {
    NestedStructPointer *StructA
    FieldC string `form:"field_c"`
}

type StructD struct {
    NestedAnonyStruct struct {
        FieldX string `form:"field_x"`
    }
    FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
    var b StructB
    c.Bind(&b)
    c.JSON(http.StatusOK, gin.H{
        "a": b.NestedStruct,
        "b": b.FieldB,
    })
}

func GetDataC(c *gin.Context) {
    var b StructC
    c.Bind(&b)
    c.JSON(http.StatusOK, gin.H{
        "a": b.NestedStructPointer,
        "c": b.FieldC,
    })
}

func GetDataD(c *gin.Context) {
    var b StructD
    c.Bind(&b)
    c.JSON(http.StatusOK, gin.H{
        "x": b.NestedAnonyStruct,
        "d": b.FieldD,
    })
}

func main() {
    r := gin.Default()
    r.GET("/getb", GetDataB)
    r.GET("/getc", GetDataC)
    r.GET("/getd", GetDataD)

    r.Run()
}
```

Using the command `curl` command result:

```sh
$ curl "http://localhost:8080/getb?field_a=hello&field_b=world"
{"a":{"FieldA":"hello"},"b":"world"}
$ curl "http://localhost:8080/getc?field_a=hello&field_c=world"
{"a":{"FieldA":"hello"},"c":"world"}
$ curl "http://localhost:8080/getd?field_x=hello&field_d=world"
{"d":"world","x":{"FieldX":"hello"}}
```

### Try to bind body into different structs

The normal methods for binding request body consumes `c.Request.Body` and they
cannot be called multiple times.

```go
type formA struct {
  Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
  Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
  objA := formA{}
  objB := formB{}
  // This c.ShouldBind consumes c.Request.Body and it cannot be reused.
  if errA := c.ShouldBind(&objA); errA == nil {
    c.String(http.StatusOK, `the body should be formA`)
  // Always an error is occurred by this because c.Request.Body is EOF now.
  } else if errB := c.ShouldBind(&objB); errB == nil {
    c.String(http.StatusOK, `the body should be formB`)
  } else {
    ...
  }
}
```

For this, you can use `c.ShouldBindBodyWith`.

```go
func SomeHandler(c *gin.Context) {
  objA := formA{}
  objB := formB{}
  // This reads c.Request.Body and stores the result into the context.
  if errA := c.ShouldBindBodyWith(&objA, binding.Form); errA == nil {
    c.String(http.StatusOK, `the body should be formA`)
  // At this time, it reuses body stored in the context.
  } else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
    c.String(http.StatusOK, `the body should be formB JSON`)
  // And it can accepts other formats
  } else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
    c.String(http.StatusOK, `the body should be formB XML`)
  } else {
    ...
  }
}
```

1. `c.ShouldBindBodyWith` stores body into the context before binding. This has
a slight impact to performance, so you should not use this method if you are
enough to call binding at once.
2. This feature is only needed for some formats -- `JSON`, `XML`, `MsgPack`,
`ProtoBuf`. For other formats, `Query`, `Form`, `FormPost`, `FormMultipart`,
can be called by `c.ShouldBind()` multiple times without any damage to
performance (See [#1341](https://github.com/gin-gonic/gin/pull/1341)).

### # 自定义结构体与自定义标签绑定表单数据请求

```go
const (
  customerTag = "url"
  defaultMemory = 32 << 20 // 默认内存大小为32MB
)

type customerBinding struct {}

func (customerBinding) Name() string {
  return "form" // 返回“form”作为绑定器名称
}

func (customerBinding) Bind(req *http.Request, obj any) error {
  if err := req.ParseForm(); err != nil {
    return err // 解析请求的表单数据，如有错误则返回
  }
  if err := req.ParseMultipartForm(defaultMemory); err != nil {
    if err != http.ErrNotMultipart {
      return err // 解析多部分表单数据，如非多部分表单或有其他错误则返回
    }
  }
  if err := binding.MapFormWithTag(obj, req.Form, customerTag); err != nil {
    return err // 将请求表单数据映射到对象，并使用自定义标签（customerTag）进行匹配，如有错误则返回
  }
  return validate(obj) // 对对象进行验证
}

func validate(obj any) error {
  if binding.Validator == nil {
    return nil // 如果验证器为空，则直接返回nil
  }
  return binding.Validator.ValidateStruct(obj) // 使用验证器对结构体进行验证并返回可能的错误
}

// 现在我们可以这样操作!!!
// FormA 是一个外部类型，我们无法修改其标签
type FormA struct {
  FieldA string `url:"field_a"` // 字段FieldA，其标签为"url:field_a"
}

func ListHandler(s *Service) func(ctx *gin.Context) {
  return func(ctx *gin.Context) {
    var urlBinding = customerBinding{} // 创建自定义绑定器实例
    var opt FormA                       // 创建FormA类型的变量opt
    err := ctx.MustBindWith(&opt, urlBinding) // 使用自定义绑定器将请求上下文中的表单数据绑定到opt变量
    if err != nil {
      // 处理绑定过程中可能出现的错误
    }
    // 绑定成功后执行的相关逻辑...
  }
}
```

#
## # http2 服务器推送

http.Pusher 功能仅在 **go1.8+** 版本中支持。详细信息请参阅 [golang 博客](https://go.dev/blog/h2push)。

```go
package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
	<title>Https Test</title>
	<script src="/assets/app.js"></script>
</head>
<body>
	<h1 style="color:red;">欢迎，Ginner！</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 进行服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("推送失败: %v", err)
			}
		}
		c.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})

	// 在 https://127.0.0.1:8080 上监听并服务
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
```

#
## # 定义路由日志的格式

默认的路由日志格式为：

```sh
[GIN-debug] POST   /foo                      --> main.main.func1 (3 handlers)
[GIN-debug] GET    /bar                      --> main.main.func2 (3 handlers)
[GIN-debug] GET    /status                   --> main.main.func3 (3 handlers)
```

如果你想以特定的格式（例如 JSON、键值对或其他格式）记录这些信息，可以通过设置 `gin.DebugPrintRouteFunc` 来自定义日志格式。

下面的示例中，我们使用标准 log 包来记录所有路由信息，但你可以根据需求选择其他日志工具。

```go
import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
    log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
  }

  r.POST("/foo", func(c *gin.Context) {
    c.JSON(http.StatusOK, "foo")
  })

  r.GET("/bar", func(c *gin.Context) {
    c.JSON(http.StatusOK, "bar")
  })

  r.GET("/status", func(c *gin.Context) {
    c.JSON(http.StatusOK, "ok")
  })

// 在 http://0.0.0.0:8080 上监听和提供服务
  r.Run()
}
```
## # 设置和获取 cookie

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  // 初始化路由
  router := gin.Default()

  // 设置 GET 请求处理函数
  router.GET("/cookie", func(c *gin.Context) {

      // 尝试从请求中获取名为 "gin_cookie" 的 cookie
      cookie, err := c.Cookie("gin_cookie")

      // 如果获取 cookie 时出现错误（如未设置）
      if err != nil {
          // 将 cookie 的默认值设为 "NotSet"
          cookie = "NotSet"

          // 设置名为 "gin_cookie"，值为 "test" 的 cookie，
          // 有效期为3600秒，路径为"/"，域名是 "localhost"，
          // 不允许通过 HTTPS 以外的协议传输，且仅允许在浏览器中发送（HTTP only）
          c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
      }

      // 打印 cookie 的当前值
      fmt.Printf("Cookie value: %s \n", cookie)
  })

  // 运行服务器
  router.Run()
}
```
## Don't trust all proxies

Gin lets you specify which headers to hold the real client IP (if any),
as well as specifying which proxies (or direct clients) you trust to
specify one of these headers.

Use function `SetTrustedProxies()` on your `gin.Engine` to specify network addresses
or network CIDRs from where clients which their request headers related to client
IP can be trusted. They can be IPv4 addresses, IPv4 CIDRs, IPv6 addresses or
IPv6 CIDRs.

**Attention:** Gin trust all proxies by default if you don't specify a trusted 
proxy using the function above, **this is NOT safe**. At the same time, if you don't
use any proxy, you can disable this feature by using `Engine.SetTrustedProxies(nil)`,
then `Context.ClientIP()` will return the remote address directly to avoid some
unnecessary computation.

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.SetTrustedProxies([]string{"192.168.1.2"})

  router.GET("/", func(c *gin.Context) {
    // If the client is 192.168.1.2, use the X-Forwarded-For
    // header to deduce the original client IP from the trust-
    // worthy parts of that header.
    // Otherwise, simply return the direct client IP
    fmt.Printf("ClientIP: %s\n", c.ClientIP())
  })
  router.Run()
}
```

**Notice:** If you are using a CDN service, you can set the `Engine.TrustedPlatform`
to skip TrustedProxies check, it has a higher priority than TrustedProxies. 
Look at the example below:

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  // Use predefined header gin.PlatformXXX
  router.TrustedPlatform = gin.PlatformGoogleAppEngine
  // Or set your own trusted request header for another trusted proxy service
  // Don't set it to any suspect request header, it's unsafe
  router.TrustedPlatform = "X-CDN-IP"

  router.GET("/", func(c *gin.Context) {
    // If you set TrustedPlatform, ClientIP() will resolve the
    // corresponding header and return IP directly
    fmt.Printf("ClientIP: %s\n", c.ClientIP())
  })
  router.Run()
}
```

## # 测试

`net/http/httptest` 包是进行 HTTP 测试的首选方式。

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })
  return r
}

func main() {
  r := setupRouter()
  r.Run(":8080")
}
```

上述代码示例的测试：

```go
package main

import (
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
  router := setupRouter()

  w := httptest.NewRecorder()
  req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
  router.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Equal(t, "pong", w.Body.String())
}
```

这段代码首先展示了一个使用 Gin 框架创建简易 HTTP 路由（GET 方法，访问路径为 `/ping`，返回内容为字符串 "pong"）的示例，并在 `main` 函数中运行服务器。

接下来是一个针对此路由的测试用例，通过 `httptest` 包创建模拟 HTTP 请求和响应。测试函数 `TestPingRoute` 中，我们设置了路由、创建了请求记录器、构造了 GET 请求并发送到路由处理器，然后使用 `assert` 包验证响应状态码应为 `http.StatusOK`，响应体内容应为 `"pong"`。
