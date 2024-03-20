
<原文开始>
Contents

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


<原文结束>

# <翻译开始>
# 

# <翻译结束>


<原文开始>
Build with json replacement

Gin uses `encoding/json` as default json package but you can change it by build from other tags.

[jsoniter](https://github.com/json-iterator/go)

```sh
go build -tags=jsoniter .
```

[go-json](https://github.com/goccy/go-json)

```sh
go build -tags=go_json .
```

[sonic](https://github.com/bytedance/sonic) (you have to ensure that your cpu support avx instruction.)

```sh
$ go build -tags="sonic avx" .
```

#
<原文结束>

# <翻译开始>
# 使用json替换构建

Gin默认使用`encoding/json`作为其json包，但您可以通过使用其他标签来自定义构建时进行更改。

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

# <翻译结束>


<原文开始>
Build without `MsgPack` rendering feature

Gin enables `MsgPack` rendering feature by default. But you can disable this feature by specifying `nomsgpack` build tag.

```sh
go build -tags=nomsgpack .
```

This is useful to reduce the binary size of executable files. See the [detail information](https://github.com/gin-gonic/gin/pull/1852).


<原文结束>

# <翻译开始>
# 不启用`MsgPack`渲染功能构建

Gin 默认启用了`MsgPack`渲染功能。但是，你可以通过指定`nomsgpack`构建标签来禁用此功能。

```sh
go build -tags=nomsgpack .
```

这对于减小可执行文件的二进制大小非常有用。有关详细信息，请参阅[此处](https://github.com/gin-gonic/gin/pull/1852)。

# <翻译结束>


<原文开始>
API Examples

You can find a number of ready-to-run examples at [Gin examples repository](https://github.com/gin-gonic/examples).

#
<原文结束>

# <翻译开始>
# API示例

你可以在[Gin示例仓库](https://github.com/gin-gonic/examples)找到多个现成的、可直接运行的例子。

# <翻译结束>


<原文开始>
Using GET, POST, PUT, PATCH, DELETE and OPTIONS

```go
func main() {
  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()

  router.GET("/someGet", getting)
  router.POST("/somePost", posting)
  router.PUT("/somePut", putting)
  router.DELETE("/someDelete", deleting)
  router.PATCH("/somePatch", patching)
  router.HEAD("/someHead", head)
  router.OPTIONS("/someOptions", options)

  // By default it serves on :8080 unless a
  // PORT environment variable was defined.
  router.Run()
  // router.Run(":3000") for a hard coded port
}
```

#
<原文结束>

# <翻译开始>
# 使用GET、POST、PUT、PATCH、DELETE和OPTIONS方法

```go
func main() {
// 创建一个gin路由器，其中包含默认中间件：
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
// 若要硬编码指定端口，则使用router.Run(":3000")
}
```

#

# <翻译结束>


<原文开始>
Parameters in path

```go
func main() {
  router := gin.Default()

  // This handler will match /user/john but will not match /user/ or /user
  router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })

  // However, this one will match /user/john/ and also /user/john/send
  // If no other routers match /user/john, it will redirect to /user/john/
  router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
  })

  // For each matched request Context will hold the route definition
  router.POST("/user/:name/*action", func(c *gin.Context) {
    b := c.FullPath() == "/user/:name/*action" // true
    c.String(http.StatusOK, "%t", b)
  })

  // This handler will add a new router for /user/groups.
  // Exact routes are resolved before param routes, regardless of the order they were defined.
  // Routes starting with /user/groups are never interpreted as /user/:name/... routes
  router.GET("/user/groups", func(c *gin.Context) {
    c.String(http.StatusOK, "The available groups are [...]")
  })

  router.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 路径中的参数

```go
func main() {
  router := gin.Default()

// 这个处理器将匹配 /user/john，但不会匹配 /user/ 或 /user
  router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "你好，%s", name)
  })

// 然而，这个处理器将匹配 /user/john/ 和 /user/john/send
// 如果没有其他处理器匹配 /user/john，它将重定向到 /user/john/
  router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " 正在执行 " + action
    c.String(http.StatusOK, message)
  })

// 对于每个匹配的请求，Context 将持有路由定义
  router.POST("/user/:name/*action", func(c *gin.Context) {
    b := c.FullPath() == "/user/:name/*action" // 结果为 true
    c.String(http.StatusOK, "%t", b)
  })

// 这个处理器将添加一个新的路由 /user/groups。
// 精确路由会在参数路由之前解析，无论它们定义的顺序如何。
// 以 /user/groups 开头的路由永远不会被解释为 /user/:name/... 路由
  router.GET("/user/groups", func(c *gin.Context) {
    c.String(http.StatusOK, "可用的组包括 [...]")
  })

  router.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
Querystring parameters

```go
func main() {
  router := gin.Default()

  // Query string parameters are parsed using the existing underlying request object.
  // The request responds to an url matching:  /welcome?firstname=Jane&lastname=Doe
  router.GET("/welcome", func(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
  })
  router.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 查询字符串参数

```go
func main() {
// 初始化路由
  router := gin.Default()

// 使用已存在的底层请求对象解析查询字符串参数。
// 当请求的URL匹配：/welcome?firstname=Jane&lastname=Doe 时，该请求会得到响应
  router.GET("/welcome", func(c *gin.Context) {
// 获取查询参数“firstname”，若不存在则默认为 "Guest"
    firstname := c.DefaultQuery("firstname", "Guest")
// 快捷获取查询参数“lastname”
    lastname := c.Query("lastname") // 相当于 c.Request.URL.Query().Get("lastname")

// 返回状态码200，并向客户端发送消息
    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
  })

// 启动服务器监听8080端口
  router.Run(":8080")
}
```

# 简介

这段代码是用Go语言编写的，使用Gin框架创建一个Web服务器。在`main`函数中，首先初始化了一个Gin路由器实例。

接下来定义了一个GET类型的路由处理函数，该函数用于处理访问路径"/welcome"的HTTP GET请求。在该函数内部，通过调用`c.DefaultQuery`和`c.Query`方法分别获取请求的查询字符串参数"firstname"和"lastname"，其中"firstname"参数如果在请求中未提供，则默认值设为"Guest"。

最后，将包含问候信息（包括从查询字符串获取的名字）的字符串以HTTP 200状态码返回给客户端。

最后启动服务器并监听8080端口，等待接收和处理HTTP请求。

# <翻译结束>


<原文开始>
Multipart/Urlencoded Form

```go
func main() {
  router := gin.Default()

  router.POST("/form_post", func(c *gin.Context) {
    message := c.PostForm("message")
    nick := c.DefaultPostForm("nick", "anonymous")

    c.JSON(http.StatusOK, gin.H{
      "status":  "posted",
      "message": message,
      "nick":    nick,
    })
  })
  router.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 多部分/URL 编码表单

```go
func main() {
// 初始化 Gin 框架路由
  router := gin.Default()

// 处理 POST 请求到 "/form_post" 路径
  router.POST("/form_post", func(c *gin.Context) {
// 从请求体中获取名为 "message" 的表单字段值
    message := c.PostForm("message")

// 从请求体中获取名为 "nick" 的表单字段值，如果未提供，则默认为 "anonymous"
    nick := c.DefaultPostForm("nick", "anonymous")

// 返回 JSON 格式的响应，HTTP 状态码为 200（OK）
    c.JSON(http.StatusOK, gin.H{
      "status":  "posted", // 状态为已发布
      "message": message,  // 返回接收到的消息内容
      "nick":    nick,     // 返回昵称或默认昵称
    })
  })

// 启动服务器监听 8080 端口
  router.Run(":8080")
}
```

# <翻译结束>


<原文开始>
Another example: query + post form

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

```sh
id: 1234; page: 1; name: manu; message: this_is_great
```

#
<原文结束>

# <翻译开始>
# 另一个示例：查询 + POST 表单

```sh
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
```

```go
func main() {
  router := gin.Default()

// 处理POST请求到"/post"路由
  router.POST("/post", func(c *gin.Context) {

// 从查询参数中获取id和page值
    id := c.Query("id")
    page := c.DefaultQuery("page", "0")

// 从POST表单中获取name和message值
    name := c.PostForm("name")
    message := c.PostForm("message")

// 打印获取到的参数值
    fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
  })

// 运行服务器在8080端口
  router.Run(":8080")
}
```

```sh
// 控制台输出结果
id: 1234; page: 1; name: manu; message: this_is_great
```

# <翻译结束>


<原文开始>
Map as querystring or postform parameters

```sh
POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded

names[first]=thinkerou&names[second]=tianou
```

```go
func main() {
  router := gin.Default()

  router.POST("/post", func(c *gin.Context) {

    ids := c.QueryMap("ids")
    names := c.PostFormMap("names")

    fmt.Printf("ids: %v; names: %v", ids, names)
  })
  router.Run(":8080")
}
```

```sh
ids: map[b:hello a:1234]; names: map[second:tianou first:thinkerou]
```

#
<原文结束>

# <翻译开始>
# 将映射作为查询字符串或POST表单参数

```sh
POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded

names[first]=thinkerou&names[second]=tianou
```

```go
func main() {
  router := gin.Default()

// 注册处理POST请求的路由
  router.POST("/post", func(c *gin.Context) {

// 从查询字符串中获取并解析ids参数为映射
    ids := c.QueryMap("ids")

// 从POST表单中获取并解析names参数为映射
    names := c.PostFormMap("names")

// 打印解析得到的ids和names映射
    fmt.Printf("ids: %v; names: %v", ids, names)
  })

// 运行服务器在8080端口
  router.Run(":8080")
}
```

```sh
// 输出结果
ids: map[b:hello a:1234]; names: map[second:tianou first:thinkerou]
```

此示例展示了一个使用Gin框架处理HTTP POST请求的场景，其中请求包含查询字符串和POST表单数据。通过`c.QueryMap`和`c.PostFormMap`方法分别获取并解析这些参数为映射类型，并最终输出到控制台。

# <翻译结束>


<原文开始>
Single file

References issue [#774](https://github.com/gin-gonic/gin/issues/774) and detail [example code](https://github.com/gin-gonic/examples/tree/master/upload-file/single).

`file.Filename` **SHOULD NOT** be trusted. See [`Content-Disposition` on MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition#Directives) and [#1693](https://github.com/gin-gonic/gin/issues/1693)

> The filename is always optional and must not be used blindly by the application: path information should be stripped, and conversion to the server file system rules should be done.

```go
func main() {
  router := gin.Default()
  // Set a lower memory limit for multipart forms (default is 32 MiB)
  router.MaxMultipartMemory = 8 << 20  // 8 MiB
  router.POST("/upload", func(c *gin.Context) {
    // Single file
    file, _ := c.FormFile("file")
    log.Println(file.Filename)

    // Upload the file to specific dst.
    c.SaveUploadedFile(file, dst)

    c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
  })
  router.Run(":8080")
}
```

How to `curl`:

```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/appleboy/test.zip" \
  -H "Content-Type: multipart/form-data"
```

##
<原文结束>

# <翻译开始>
# 单文件上传

引用问题 [#774](https://github.com/gin-gonic/gin/issues/774) 和详细示例代码：[example code](https://github.com/gin-gonic/examples/tree/master/upload-file/single)。

**不应**盲目信任 `file.Filename`。参见MDN上的 [`Content-Disposition`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition#Directives) 和相关问题 [#1693](https://github.com/gin-gonic/gin/issues/1693)：

> 文件名始终是可选的，应用程序不应盲目使用它：应去除路径信息，并根据服务器文件系统规则进行转换。

```go
func main() {
  router := gin.Default()
// 设置multipart表单的较低内存限制（默认为32 MiB）
  router.MaxMultipartMemory = 8 << 20  // 设置为8 MiB
  router.POST("/upload", func(c *gin.Context) {
// 单个文件上传
    file, _ := c.FormFile("file")
    log.Println(file.Filename)

// 将文件上传到特定的目标位置(dst)
    c.SaveUploadedFile(file, dst)

    c.String(http.StatusOK, fmt.Sprintf("'%s' 已上传!", file.Filename))
  })
  router.Run(":8080")
}
```

如何使用 `curl` 进行测试：

```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/appleboy/test.zip" \
  -H "Content-Type: multipart/form-data"
```

##

# <翻译结束>


<原文开始>
Multiple files

See the detail [example code](https://github.com/gin-gonic/examples/tree/master/upload-file/multiple).

```go
func main() {
  router := gin.Default()
  // Set a lower memory limit for multipart forms (default is 32 MiB)
  router.MaxMultipartMemory = 8 << 20  // 8 MiB
  router.POST("/upload", func(c *gin.Context) {
    // Multipart form
    form, _ := c.MultipartForm()
    files := form.File["upload[]"]

    for _, file := range files {
      log.Println(file.Filename)

      // Upload the file to specific dst.
      c.SaveUploadedFile(file, dst)
    }
    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
  })
  router.Run(":8080")
}
```

How to `curl`:

```bash
curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/appleboy/test1.zip" \
  -F "upload[]=@/Users/appleboy/test2.zip" \
  -H "Content-Type: multipart/form-data"
```

#
<原文结束>

# <翻译开始>
# 多文件上传

查看详细[示例代码](https://github.com/gin-gonic/examples/tree/master/upload-file/multiple)。

```go
func main() {
// 初始化路由
  router := gin.Default()
// 设置multipart表单的内存限制（默认为32 MiB）
  router.MaxMultipartMemory = 8 << 20  // 设置为8 MiB
// 处理POST请求到"/upload"
  router.POST("/upload", func(c *gin.Context) {
// 获取multipart表单
    form, _ := c.MultipartForm()
// 获取表单中名为"upload[]"的所有文件
    files := form.File["upload[]"]

// 遍历所有文件
    for _, file := range files {
      log.Println(file.Filename)  // 输出文件名

// 将上传的文件保存到特定的目标路径(dst)
      c.SaveUploadedFile(file, dst)
    }
// 返回状态码和已上传文件数量的消息
    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
  })
// 启动服务器监听8080端口
  router.Run(":8080")
}
```

如何使用`curl`进行上传：

```bash
curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/appleboy/test1.zip" \
  -F "upload[]=@/Users/appleboy/test2.zip" \
  -H "Content-Type: multipart/form-data"
```

# <翻译结束>


<原文开始>
Grouping routes

```go
func main() {
  router := gin.Default()

  // Simple group: v1
  v1 := router.Group("/v1")
  {
    v1.POST("/login", loginEndpoint)
    v1.POST("/submit", submitEndpoint)
    v1.POST("/read", readEndpoint)
  }

  // Simple group: v2
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
<原文结束>

# <翻译开始>
# 以下是Go语言代码，内容为使用gin框架对路由进行分组：

```go
func main() {
// 初始化gin路由器
  router := gin.Default()

// 简单的路由分组：v1版本
  v1 := router.Group("/v1")
  {
// 在v1版本下定义POST请求路由
    v1.POST("/login", loginEndpoint)   // 登录接口
    v1.POST("/submit", submitEndpoint) // 提交接口
    v1.POST("/read", readEndpoint)     // 阅读接口
  }

// 简单的路由分组：v2版本
  v2 := router.Group("/v2")
  {
// 在v2版本下定义POST请求路由
    v2.POST("/login", loginEndpoint)   // 登录接口
    v2.POST("/submit", submitEndpoint) // 提交接口
    v2.POST("/read", readEndpoint)     // 阅读接口
  }

// 启动服务器监听8080端口
  router.Run(":8080")
}
```

这段代码中，通过gin框架创建了两个路由分组：v1和v2。在每个分组下分别定义了三个POST类型的HTTP接口，分别是登录、提交和阅读功能，并且针对不同的版本（v1和v2）提供了相同的接口。最后启动服务器监听8080端口等待客户端连接。

# <翻译结束>


<原文开始>
Blank Gin without middleware by default

Use

```go
r := gin.New()
```

instead of

```go
// Default With the Logger and Recovery middleware already attached
r := gin.Default()
```

#
<原文结束>

# <翻译开始>
# 默认情况下，Blank Gin 不带中间件

使用方式：

```go
r := gin.New()
```

而非

```go
// 默认情况下已附加 Logger 和 Recovery 中间件
r := gin.Default()
```

#

# <翻译结束>


<原文开始>
Using middleware

```go
func main() {
  // Creates a router without any middleware by default
  r := gin.New()

  // Global middleware
  // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
  // By default gin.DefaultWriter = os.Stdout
  r.Use(gin.Logger())

  // Recovery middleware recovers from any panics and writes a 500 if there was one.
  r.Use(gin.Recovery())

  // Per route middleware, you can add as many as you desire.
  r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

  // Authorization group
  // authorized := r.Group("/", AuthRequired())
  // exactly the same as:
  authorized := r.Group("/")
  // per group middleware! in this case we use the custom created
  // AuthRequired() middleware just in the "authorized" group.
  authorized.Use(AuthRequired())
  {
    authorized.POST("/login", loginEndpoint)
    authorized.POST("/submit", submitEndpoint)
    authorized.POST("/read", readEndpoint)

    // nested group
    testing := authorized.Group("testing")
    // visit 0.0.0.0:8080/testing/analytics
    testing.GET("/analytics", analyticsEndpoint)
  }

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 使用中间件

```go
func main() {
// 默认情况下创建一个不包含任何中间件的路由器
  r := gin.New()

// 全局中间件
// Logger 中间件将会把日志写入 gin.DefaultWriter，即使你设置了 GIN_MODE=release。默认情况下 gin.DefaultWriter = os.Stdout
  r.Use(gin.Logger())

// Recovery 中间件从任何 panic 恢复，并在发生 panic 时返回 500 状态码。
  r.Use(gin.Recovery())

// 路由级别的中间件，你可以根据需要添加任意多个。
  r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

// 授权分组
// authorized := r.Group("/", AuthRequired())
// 这与下面的代码完全等效：
  authorized := r.Group("/")
// 分组级别中间件！在这个例子中，我们在 "authorized" 分组中仅使用自定义创建的
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

// 在 0.0.0.0:8080 监听并提供服务
  r.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
Custom Recovery behavior

```go
func main() {
  // Creates a router without any middleware by default
  r := gin.New()

  // Global middleware
  // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
  // By default gin.DefaultWriter = os.Stdout
  r.Use(gin.Logger())

  // Recovery middleware recovers from any panics and writes a 500 if there was one.
  r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
    if err, ok := recovered.(string); ok {
      c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
    }
    c.AbortWithStatus(http.StatusInternalServerError)
  }))

  r.GET("/panic", func(c *gin.Context) {
    // panic with a string -- the custom middleware could save this to a database or report it to the user
    panic("foo")
  })

  r.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "ohai")
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 自定义恢复行为

```go
func main() {
// 创建一个默认不包含任何中间件的路由器
  r := gin.New()

// 全局中间件
// Logger 中间件会将日志写入 gin.DefaultWriter，即使你设置 GIN_MODE=release。默认情况下 gin.DefaultWriter = os.Stdout
  r.Use(gin.Logger())

// Recovery 中间件用于从任何 panic 恢复，并在发生 panic 时返回一个 500 状态码。
  r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
    if err, ok := recovered.(string); ok {
      c.String(http.StatusInternalServerError, fmt.Sprintf("错误: %s", err))
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

// 监听并绑定到 0.0.0.0:8080 进行服务
  r.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
Controlling Log output coloring

By default, logs output on console should be colorized depending on the detected TTY.

Never colorize logs:

```go
func main() {
  // Disable log's color
  gin.DisableConsoleColor()

  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```

Always colorize logs:

```go
func main() {
  // Force log's color
  gin.ForceConsoleColor()

  // Creates a gin router with default middleware:
  // logger and recovery (crash-free) middleware
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 控制日志输出颜色

默认情况下，根据检测到的 TTY，日志在控制台的输出应进行着色。

从不为日志着色：

```go
func main() {
// 禁用日志的颜色
  gin.DisableConsoleColor()

// 创建一个gin路由器，包含默认中间件：
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
// 强制日志着色
  gin.ForceConsoleColor()

// 创建一个gin路由器，包含默认中间件：
// 日志记录和恢复（防崩溃）中间件
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```

# <翻译结束>


<原文开始>
Model binding and validation

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

#
<原文结束>

# <翻译开始>
# 

# <翻译结束>


<原文开始>
Custom Validators

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

#
<原文结束>

# <翻译开始>
# 

# <翻译结束>


<原文开始>
Only Bind Query String

`ShouldBindQuery` function only binds the query params and not the post data. See the [detail information](https://github.com/gin-gonic/gin/issues/742#issuecomment-315953017).

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
    log.Println("====== Only Bind By Query String ======")
    log.Println(person.Name)
    log.Println(person.Address)
  }
  c.String(http.StatusOK, "Success")
}

```

#
<原文结束>

# <翻译开始>
# 仅绑定查询字符串

`ShouldBindQuery` 函数仅绑定查询参数，而不绑定 post 数据。有关详细信息，请参阅 [此处](https://github.com/gin-gonic/gin/issues/742#issuecomment-315953017)。

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
  router := gin.Default()
  router.Any("/testing", startPage)
  router.Run(":8085")
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

# <翻译结束>


<原文开始>
Bind Query String or Post Data

See the [detail information](https://github.com/gin-gonic/gin/issues/742#issuecomment-264681292).

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
  route := gin.Default()
  route.GET("/testing", startPage)
  route.Run(":8085")
}

func startPage(c *gin.Context) {
  var person Person
  // If `GET`, only `Form` binding engine (`query`) used.
  // If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
  // See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
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

Test it with:

```sh
curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
```

#
<原文结束>

# <翻译开始>
# 绑定查询字符串或Post数据

请参阅[详细信息](https://github.com/gin-gonic/gin/issues/742#issuecomment-264681292)。

```go
package main

import (
  "log"
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
)

type Person struct {
  Name       string    `form:"name"`        // 表单字段"Name"
  Address    string    `form:"address"`     // 表单字段"Address"
  Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"` // 表单字段"Birthday"，时间格式为"2006-01-02"，使用UTC时间
  CreateTime time.Time `form:"createTime" time_format:"unixNano"` // 表单字段"CreateTime"，时间格式为纳秒级Unix时间戳
  UnixTime   time.Time `form:"unixTime" time_format:"unix"`      // 表单字段"UnixTime"，时间格式为秒级Unix时间戳
}

func main() {
  route := gin.Default()
  route.GET("/testing", startPage)
  route.Run(":8085")
}

func startPage(c *gin.Context) {
  var person Person
// 如果是GET请求，仅使用`Form`绑定引擎（查询参数）。
// 如果是POST请求，首先检查`content-type`是否为JSON或XML，然后使用`Form`（表单数据）进行绑定。
// 更多详情见：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
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

通过以下命令测试：

```sh
curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
```

# <翻译结束>


<原文开始>
Bind Uri

See the [detail information](https://github.com/gin-gonic/gin/issues/846).

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type Person struct {
  ID string `uri:"id" binding:"required,uuid"`
  Name string `uri:"name" binding:"required"`
}

func main() {
  route := gin.Default()
  route.GET("/:name/:id", func(c *gin.Context) {
    var person Person
    if err := c.ShouldBindUri(&person); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
      return
    }
    c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
  })
  route.Run(":8088")
}
```

Test it with:

```sh
curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
curl -v localhost:8088/thinkerou/not-uuid
```

#
<原文结束>

# <翻译开始>
# 绑定 Uri

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

# 结束

# <翻译结束>


<原文开始>
Bind Header

```go
package main

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)

type testHeader struct {
  Rate   int    `header:"Rate"`
  Domain string `header:"Domain"`
}

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    h := testHeader{}

    if err := c.ShouldBindHeader(&h); err != nil {
      c.JSON(http.StatusOK, err)
    }

    fmt.Printf("%#v\n", h)
    c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
  })

  r.Run()

// client
// curl -H "rate:300" -H "domain:music" 127.0.0.1:8080/
// output
// {"Domain":"music","Rate":300}
}
```

#
<原文结束>

# <翻译开始>
# 标题：绑定Header

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// 定义一个测试Header结构体
type testHeader struct {
    Rate   int    `header:"Rate"`    // 绑定到HTTP请求头的Rate字段
    Domain string `header:"Domain"` // 绑定到HTTP请求头的Domain字段
}

func main() {
// 初始化Gin框架
    r := gin.Default()

// 配置路由GET /
    r.GET("/", func(c *gin.Context) {
// 创建testHeader实例
        h := testHeader{}

// 尝试从请求头中解析并绑定数据到h
        if err := c.ShouldBindHeader(&h); err != nil {
// 如果绑定失败，返回错误信息
            c.JSON(http.StatusOK, err)
        }

// 输出已绑定的数据
        fmt.Printf("%#v\n", h)

// 返回JSON响应，包含已绑定的Rate和Domain值
        c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
    })

// 启动服务器
    r.Run()

// 客户端示例命令
// 使用curl发送带有Rate和Domain头部信息的请求
// curl -H "rate:300" -H "domain:music" 127.0.0.1:8080/

// 预期输出结果
// {"Domain":"music","Rate":300}
}
```

这段代码定义了一个Go语言程序，使用Gin框架创建一个Web服务。在该服务中，我们定义了一个`testHeader`结构体来绑定HTTP请求头中的特定字段（Rate和Domain）。在"/"路由对应的处理函数中，我们尝试从请求头中提取并绑定数据到结构体实例，并将绑定成功后的数据以JSON格式返回给客户端。同时，提供了一个示例curl命令，演示如何向服务发送带有相应请求头的HTTP请求以及预期的输出结果。

# <翻译结束>


<原文开始>
Bind HTML checkboxes

See the [detail information](https://github.com/gin-gonic/gin/issues/129#issuecomment-124260092)

main.go

```go
...

type myForm struct {
    Colors []string `form:"colors[]"`
}

...

func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)
    c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}

...

```

form.html

```html
<form action="/" method="POST">
    <p>Check some colors</p>
    <label for="red">Red</label>
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit">
</form>
```

result:

```json
{"color":["red","green","blue"]}
```

#
<原文结束>

# <翻译开始>
# 绑定HTML复选框

参考详细信息：[https://github.com/gin-gonic/gin/issues/129#issuecomment-124260092](https://github.com/gin-gonic/gin/issues/129#issuecomment-124260092)

main.go 文件内容：

```go
...

// 定义表单结构体
type myForm struct {
    Colors []string `form:"colors[]"`
}

...

// 表单处理函数
func formHandler(c *gin.Context) {
// 创建表单实例
    var fakeForm myForm
// 绑定请求数据到表单实例
    c.ShouldBind(&fakeForm)
// 返回JSON响应
    c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}

...

```

form.html 文件内容：

```html
<!-- HTML表单 -->
<form action="/" method="POST">
    <p>选择一些颜色</p>
    <label for="red">红色</label>
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">绿色</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">蓝色</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit">
</form>
```

结果示例（JSON格式）：

```json
{"color":["red","green","blue"]}
```

# <翻译结束>


<原文开始>
Multipart/Urlencoded binding

```go
type ProfileForm struct {
  Name   string                `form:"name" binding:"required"`
  Avatar *multipart.FileHeader `form:"avatar" binding:"required"`

  // or for multiple files
  // Avatars []*multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
  router := gin.Default()
  router.POST("/profile", func(c *gin.Context) {
    // you can bind multipart form with explicit binding declaration:
    // c.ShouldBindWith(&form, binding.Form)
    // or you can simply use autobinding with ShouldBind method:
    var form ProfileForm
    // in this case proper binding will be automatically selected
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

Test it with:

```sh
curl -X POST -v --form name=user --form "avatar=@./avatar.png" http://localhost:8080/profile
```

#
<原文结束>

# <翻译开始>
# 多部分/Urlencoded 绑定

```go
// 定义表单结构体
type ProfileForm struct {
  Name   string                `form:"name" binding:"required"` // 名称，必填项
  Avatar *multipart.FileHeader `form:"avatar" binding:"required"` // 头像文件，必填项

// 或者，如果需要上传多个文件
// Avatars []*multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
  router := gin.Default()

// 注册POST路由处理函数
  router.POST("/profile", func(c *gin.Context) {
// 显式声明绑定 multipart 表单：
// c.ShouldBindWith(&form, binding.Form)
// 或者直接使用自动绑定 ShouldBind 方法：
    var form ProfileForm
// 在这种情况下，将会自动选择合适的绑定方式
    if err := c.ShouldBind(&form); err != nil {
      c.String(http.StatusBadRequest, "bad request") // 如果绑定失败，返回400错误
      return
    }

// 保存上传的文件
    err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
    if err != nil {
      c.String(http.StatusInternalServerError, "unknown error") // 文件保存失败，返回500错误
      return
    }

// db.Save(&form) // 假设将表单数据保存到数据库

    c.String(http.StatusOK, "ok") // 成功处理请求，返回200状态码和"ok"
  })

  router.Run(":8080") // 运行服务器在8080端口
}

// 测试命令示例
```

测试代码：

```sh
curl -X POST -v --form name=user --form "avatar=@./avatar.png" http://localhost:8080/profile
```

通过以上示例，您可以在Go语言中使用Gin框架处理包含文件上传的多部分表单提交。在`ProfileForm`结构体中定义了表单字段，并通过`ShouldBind`方法进行表单数据绑定。同时，`SaveUploadedFile`方法用于将上传的文件头信息关联的实际文件内容保存到服务器指定位置。最后提供了一个使用curl命令测试该接口的示例。

# <翻译结束>


<原文开始>
XML, JSON, YAML, TOML and ProtoBuf rendering

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

##
<原文结束>

# <翻译开始>
XML, JSON, YAML, TOML and ProtoBuf rendering

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

##
# <翻译结束>


<原文开始>
SecureJSON

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

##
<原文结束>

# <翻译开始>
SecureJSON

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

##
# <翻译结束>


<原文开始>
JSONP

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

##
<原文结束>

# <翻译开始>
JSONP

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

##
# <翻译结束>


<原文开始>
PureJSON

Normally, JSON replaces special HTML characters with their unicode entities, e.g. `<` becomes  `\u003c`. If you want to encode such characters literally, you can use PureJSON instead.
This feature is unavailable in Go 1.6 and lower.

```go
func main() {
  r := gin.Default()

  // Serves unicode entities
  r.GET("/json", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

  // Serves literal characters
  r.GET("/purejson", func(c *gin.Context) {
    c.PureJSON(http.StatusOK, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

  // listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 纯JSON

通常情况下，JSON会将特殊的HTML字符替换为它们的Unicode实体，例如`<`会变成 `\u003c`。如果你想直接编码这样的字符，可以使用PureJSON。

此功能在Go 1.6及更低版本中不可用。

```go
func main() {
  r := gin.Default()

// 服务端返回Unicode实体
  r.GET("/json", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

// 服务端返回原始字符
  r.GET("/purejson", func(c *gin.Context) {
    c.PureJSON(http.StatusOK, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

// 在0.0.0.0:8080监听并提供服务
  r.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
Serving static files

```go
func main() {
  router := gin.Default()
  router.Static("/assets", "./assets")
  router.StaticFS("/more_static", http.Dir("my_file_system"))
  router.StaticFile("/favicon.ico", "./resources/favicon.ico")
  router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))
  
  // Listen and serve on 0.0.0.0:8080
  router.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 服务静态文件

```go
func main() {
// 创建默认的Gin路由器实例
  router := gin.Default()

// 设置"/assets"路由，提供"./assets"目录下的静态文件服务
  router.Static("/assets", "./assets")

// 设置"/more_static"路由，从"my_file_system"目录下提供静态文件服务
  router.StaticFS("/more_static", http.Dir("my_file_system"))

// 设置"/favicon.ico"路由，提供"./resources/favicon.ico"文件作为 favicon
  router.StaticFile("/favicon.ico", "./resources/favicon.ico")

// 设置"/more_favicon.ico"路由，从"my_file_system"目录下的"more_favicon.ico"文件作为 favicon
  router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))

// 监听并服务在 0.0.0.0:8080 端口
  router.Run(":8080")
}
```

这个Go代码示例展示了如何使用Gin框架来服务静态文件。通过`router.Static()`和`router.StaticFS()`方法，我们可以指定目录以服务其中的静态资源（如CSS、JavaScript、图片等）。而`router.StaticFile()`和`router.StaticFileFS()`方法则用于指定单个静态文件，这里主要用于设置网站图标（favicon）。最后，程序将在0.0.0.0:8080端口启动并监听HTTP请求。

# <翻译结束>


<原文开始>
Serving data from file

```go
func main() {
  router := gin.Default()

  router.GET("/local/file", func(c *gin.Context) {
    c.File("local/file.go")
  })

  var fs http.FileSystem = // ...
  router.GET("/fs/file", func(c *gin.Context) {
    c.FileFromFS("fs/file.go", fs)
  })
}

```

#
<原文结束>

# <翻译开始>
# 从文件中提供数据

```go
func main() {
// 初始化路由
  router := gin.Default()

// 处理GET请求，返回本地文件
  router.GET("/local/file", func(c *gin.Context) {
// 将本地文件"local/file.go"发送给客户端
    c.File("local/file.go")
  })

// 声明一个http.FileSystem类型的变量fs（待初始化或赋值）
  var fs http.FileSystem = // ...

// 处理GET请求，返回文件系统中的文件
  router.GET("/fs/file", func(c *gin.Context) {
// 使用指定的文件系统fs，将"fs/file.go"文件发送给客户端
    c.FileFromFS("fs/file.go", fs)
  })
}
```

这段代码使用Gin框架创建了一个Web服务器，定义了两个处理GET请求的路由：

1. 当用户访问"/local/file"时，服务器会读取并返回本地目录下的"local/file.go"文件内容。
2. 当用户访问"/fs/file"时，服务器会通过自定义的http.FileSystem接口fs读取并返回"fs/file.go"文件内容。

# <翻译结束>


<原文开始>
Serving data from reader

```go
func main() {
  router := gin.Default()
  router.GET("/someDataFromReader", func(c *gin.Context) {
    response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
    if err != nil || response.StatusCode != http.StatusOK {
      c.Status(http.StatusServiceUnavailable)
      return
    }

    reader := response.Body
     defer reader.Close()
    contentLength := response.ContentLength
    contentType := response.Header.Get("Content-Type")

    extraHeaders := map[string]string{
      "Content-Disposition": `attachment; filename="gopher.png"`,
    }

    c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
  })
  router.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 从reader提供数据

```go
func main() {
// 初始化路由
  router := gin.Default()

// 定义GET请求处理函数，路径为"/someDataFromReader"
  router.GET("/someDataFromReader", func(c *gin.Context) {
// 发起HTTP GET请求获取远程资源
    response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
    
// 检查错误或非200状态码
    if err != nil || response.StatusCode != http.StatusOK {
// 若有错误或非200状态码，返回503（服务不可用）
      c.Status(http.StatusServiceUnavailable)
      return
    }

// 获取响应体并作为reader
    reader := response.Body
    defer reader.Close() // 在函数结束时关闭reader

// 获取响应内容长度和类型
    contentLength := response.ContentLength
    contentType := response.Header.Get("Content-Type")

// 设置额外的头部信息
    extraHeaders := map[string]string{
      "Content-Disposition": `attachment; filename="gopher.png"`,
    }

// 使用reader向客户端发送数据，并附带额外的头部信息
    c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
  })

// 启动服务器监听8080端口
  router.Run(":8080")
}
```

这段Go代码实现了一个基于Gin框架的web服务器，当接收到对`/someDataFromReader`的GET请求时，它会发起一个HTTP GET请求去获取指定URL的资源。获取成功后，将资源内容通过HTTP响应以流的方式返回给客户端，并设置了自定义的Content-Type和Content-Disposition头部信息。

# <翻译结束>


<原文开始>
HTML rendering

Using LoadHTMLGlob() or LoadHTMLFiles()

```go
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")
  //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
  router.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Main website",
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

Using templates with same name in different directories

```go
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/**/*")
  router.GET("/posts/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
      "title": "Posts",
    })
  })
  router.GET("/users/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
      "title": "Users",
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
<p>Using posts/index.tmpl</p>
</html>
{{ end }}
```

templates/users/index.tmpl

```html
{{ define "users/index.tmpl" }}
<html><h1>
  {{ .title }}
</h1>
<p>Using users/index.tmpl</p>
</html>
{{ end }}
```

##
<原文结束>

# <翻译开始>
# HTML 渲染

使用 LoadHTMLGlob() 或 LoadHTMLFiles()

```go
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*") // 或者使用：router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
  router.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "主网站",
    })
  })
  router.Run(":8080")
}
```

templates/index.tmpl 模板文件

```html
<html>
  <h1>
    {{ .title }}
  </h1>
</html>
```

在不同目录中使用相同名称的模板

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

templates/posts/index.tmpl 模板文件

```html
{{ define "posts/index.tmpl" }}
<html><h1>
  {{ .title }}
</h1>
<p>正在使用 posts/index.tmpl</p>
</html>
{{ end }}
```

templates/users/index.tmpl 模板文件

```html
{{ define "users/index.tmpl" }}
<html><h1>
  {{ .title }}
</h1>
<p>正在使用 users/index.tmpl</p>
</html>
{{ end }}
```

##

# <翻译结束>


<原文开始>
Custom Template renderer

You can also use your own html template render

```go
import "html/template"

func main() {
  router := gin.Default()
  html := template.Must(template.ParseFiles("file1", "file2"))
  router.SetHTMLTemplate(html)
  router.Run(":8080")
}
```

##
<原文结束>

# <翻译开始>
# 自定义模板渲染

您也可以使用自己的HTML模板渲染器。

```go
import "html/template"

func main() {
// 初始化路由
  router := gin.Default()

// 必须解析并加载模板文件（"file1"和"file2"）
  html := template.Must(template.ParseFiles("file1", "file2"))

// 设置自定义的HTML模板到路由中
  router.SetHTMLTemplate(html)

// 运行服务器在8080端口
  router.Run(":8080")
}
```

##

# <翻译结束>


<原文开始>
Custom Delimiters

You may use custom delims

```go
  r := gin.Default()
  r.Delims("{[{", "}]}")
  r.LoadHTMLGlob("/path/to/templates")
```

##
<原文结束>

# <翻译开始>
# 自定义分隔符

您可以使用自定义分隔符

```go
  r := gin.Default()
  r.Delims("{[{", "}]}")
  r.LoadHTMLGlob("/path/to/templates")
```

##

# <翻译结束>


<原文开始>
Custom Template Funcs

See the detail [example code](https://github.com/gin-gonic/examples/tree/master/template).

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
  router.Delims("{[{", "}]}")
  router.SetFuncMap(template.FuncMap{
      "formatAsDate": formatAsDate,
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

Result:

```sh
Date: 2017/07/01
```

#
<原文结束>

# <翻译开始>
# 自定义模板函数

请参阅详细示例代码：[https://github.com/gin-gonic/examples/tree/master/template](https://github.com/gin-gonic/examples/tree/master/template)。

main.go 文件内容：

```go
import (
    "fmt"
    "html/template"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// 定义格式化日期为字符串的函数
func formatAsDate(t time.Time) string {
    year, month, day := t.Date()
    return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
// 初始化 Gin 框架路由
    router := gin.Default()
// 设置模板标签分隔符
    router.Delims("{[{", "}]}")
    
// 设置自定义模板函数映射
    router.SetFuncMap(template.FuncMap{
        "formatAsDate": formatAsDate,
    })
// 加载 HTML 模板文件
    router.LoadHTMLFiles("./testdata/template/raw.tmpl")

// 注册 GET 请求处理函数
    router.GET("/raw", func(c *gin.Context) {
// 使用模板渲染响应，并传递参数
        c.HTML(http.StatusOK, "raw.tmpl", gin.H{
            "now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
        })
    })

// 启动服务器监听8080端口
    router.Run(":8080")
}
```

raw.tmpl 模板文件内容：

```html
Date: {[{.now | formatAsDate}]}
```

运行结果：

```sh
Date: 2017/07/01
```

# <翻译结束>


<原文开始>
Multitemplate

Gin allow by default use only one html.Template. Check [a multitemplate render](https://github.com/gin-contrib/multitemplate) for using features like go 1.6 `block template`.

#
<原文结束>

# <翻译开始>
# 多模板

Gin默认情况下只允许使用一个html.Template。若要使用类似Go 1.6的`block template`等功能，请查看[a multitemplate render](https://github.com/gin-contrib/multitemplate)。

#

# <翻译结束>


<原文开始>
Redirects

Issuing a HTTP redirect is easy. Both internal and external locations are supported.

```go
r.GET("/test", func(c *gin.Context) {
  c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
})
```

Issuing a HTTP redirect from POST. Refer to issue: [#444](https://github.com/gin-gonic/gin/issues/444)

```go
r.POST("/test", func(c *gin.Context) {
  c.Redirect(http.StatusFound, "/foo")
})
```

Issuing a Router redirect, use `HandleContext` like below.

``` go
r.GET("/test", func(c *gin.Context) {
    c.Request.URL.Path = "/test2"
    r.HandleContext(c)
})
r.GET("/test2", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"hello": "world"})
})
```

#
<原文结束>

# <翻译开始>
# 重定向

发出HTTP重定向非常简单，同时支持内部和外部位置。

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

执行路由器重定向，可以像下面这样使用`HandleContext`。

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

# <翻译结束>


<原文开始>
Custom Middleware

```go
func Logger() gin.HandlerFunc {
  return func(c *gin.Context) {
    t := time.Now()

    // Set example variable
    c.Set("example", "12345")

    // before request

    c.Next()

    // after request
    latency := time.Since(t)
    log.Print(latency)

    // access the status we are sending
    status := c.Writer.Status()
    log.Println(status)
  }
}

func main() {
  r := gin.New()
  r.Use(Logger())

  r.GET("/test", func(c *gin.Context) {
    example := c.MustGet("example").(string)

    // it would print: "12345"
    log.Println(example)
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 自定义中间件

```go
// 定义 Logger 中间件函数
func Logger() gin.HandlerFunc {
  return func(c *gin.Context) {
// 获取当前时间
    t := time.Now()

// 设置示例变量
    c.Set("example", "12345")

// 请求前的操作

// 调用下一个中间件或处理请求
    c.Next()

// 请求后的操作
    latency := time.Since(t)
    log.Print(latency)  // 打印延迟时间

// 访问即将发送的状态码
    status := c.Writer.Status()
    log.Println(status)  // 打印状态码
  }
}

func main() {
// 创建 Gin 应用实例
  r := gin.New()
// 使用自定义的 Logger 中间件
  r.Use(Logger())

// 定义 GET 请求路由及其处理函数
  r.GET("/test", func(c *gin.Context) {
// 从上下文中获取示例变量
    example := c.MustGet("example").(string)

// 将打印: "12345"
    log.Println(example)
  })

// 在 0.0.0.0:8080 端口监听并提供服务
  r.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
Using BasicAuth() middleware

```go
// simulate some private data
var secrets = gin.H{
  "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
  "austin": gin.H{"email": "austin@example.com", "phone": "666"},
  "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
  r := gin.Default()

  // Group using gin.BasicAuth() middleware
  // gin.Accounts is a shortcut for map[string]string
  authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
    "foo":    "bar",
    "austin": "1234",
    "lena":   "hello2",
    "manu":   "4321",
  }))

  // /admin/secrets endpoint
  // hit "localhost:8080/admin/secrets
  authorized.GET("/secrets", func(c *gin.Context) {
    // get user, it was set by the BasicAuth middleware
    user := c.MustGet(gin.AuthUserKey).(string)
    if secret, ok := secrets[user]; ok {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
    } else {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
    }
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 使用 BasicAuth() 中间件

```go
// 模拟一些私密数据
var secrets = gin.H{
  "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
  "austin": gin.H{"email": "austin@example.com", "phone": "666"},
  "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
  r := gin.Default()

// 使用 gin.BasicAuth() 中间件的 Group
// gin.Accounts 是一个快捷方式，用于创建 map[string]string
  authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
    "foo":    "bar",
    "austin": "1234",
    "lena":   "hello2",
    "manu":   "4321",
  }))

// 访问 "/admin/secrets" 端点
// 在本地访问 "localhost:8080/admin/secrets"
  authorized.GET("/secrets", func(c *gin.Context) {
// 获取用户信息，该信息由 BasicAuth 中间件设置
    user := c.MustGet(gin.AuthUserKey).(string)
    if secret, exists := secrets[user]; exists {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
    } else {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
    }
  })

// 监听并服务于 0.0.0.0:8080
  r.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
Goroutines inside a middleware

When starting new Goroutines inside a middleware or handler, you **SHOULD NOT** use the original context inside it, you have to use a read-only copy.

```go
func main() {
  r := gin.Default()

  r.GET("/long_async", func(c *gin.Context) {
    // create copy to be used inside the goroutine
    cCp := c.Copy()
    go func() {
      // simulate a long task with time.Sleep(). 5 seconds
      time.Sleep(5 * time.Second)

      // note that you are using the copied context "cCp", IMPORTANT
      log.Println("Done! in path " + cCp.Request.URL.Path)
    }()
  })

  r.GET("/long_sync", func(c *gin.Context) {
    // simulate a long task with time.Sleep(). 5 seconds
    time.Sleep(5 * time.Second)

    // since we are NOT using a goroutine, we do not have to copy the context
    log.Println("Done! in path " + c.Request.URL.Path)
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

#
<原文结束>

# <翻译开始>
# 在中间件或处理器内部启动新的 Goroutines 时，你不 **应该** 在其中使用原始上下文，而必须使用只读副本。

```go
func main() {
  r := gin.Default()

  r.GET("/long_async", func(c *gin.Context) {
// 创建一个用于在 Goroutine 中使用的副本
    cCp := c.Copy()
    go func() {
// 模拟一个耗时任务，用 time.Sleep() 暂停5秒
      time.Sleep(5 * time.Second)

// 注意这里你正在使用复制的上下文 "cCp"，这一点非常重要
      log.Println("完成！在路径 " + cCp.Request.URL.Path)
    }()
  })

  r.GET("/long_sync", func(c *gin.Context) {
// 模拟一个耗时任务，用 time.Sleep() 暂停5秒
    time.Sleep(5 * time.Second)

// 由于我们没有使用 Goroutine，所以无需复制上下文
    log.Println("完成！在路径 " + c.Request.URL.Path)
  })

// 监听并服务于 0.0.0.0:8080
  r.Run(":8080")
}
```

#

# <翻译结束>


<原文开始>
Custom HTTP configuration

Use `http.ListenAndServe()` directly, like this:

```go
func main() {
  router := gin.Default()
  http.ListenAndServe(":8080", router)
}
```

or

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
<原文结束>

# <翻译开始>
# 自定义HTTP配置

直接使用`http.ListenAndServe()`，如下所示：

```go
func main() {
// 创建一个 Gin 路由器实例
  router := gin.Default()

// 直接通过 http.ListenAndServe() 启动服务
  http.ListenAndServe(":8080", router)
}
```

或者

```go
func main() {
// 创建一个 Gin 路由器实例
  router := gin.Default()

// 创建一个 http.Server 结构体实例并进行配置
  s := &http.Server{
    Addr:           ":8080", // 设置监听地址为 8080 端口
    Handler:        router,  // 设置处理器为 Gin 路由器
    ReadTimeout:    10 * time.Second, // 设置读超时时间为 10 秒
    WriteTimeout:   10 * time.Second, // 设置写超时时间为 10 秒
    MaxHeaderBytes: 1 << 20, // 设置最大头信息字节数为 1MB
  }

// 使用配置好的 Server 实例启动服务
  s.ListenAndServe()
}
```

#（注：该部分内容到此为止）

# <翻译结束>


<原文开始>
Support Let's Encrypt

example for 1-line LetsEncrypt HTTPS servers.

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

  // Ping handler
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
```

example for custom autocert manager.

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

  // Ping handler
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  m := autocert.Manager{
    Prompt:     autocert.AcceptTOS,
    HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
    Cache:      autocert.DirCache("/var/www/.cache"),
  }

  log.Fatal(autotls.RunWithManager(r, &m))
}
```

#
<原文结束>

# <翻译开始>
# 支持 Let's Encrypt

以下是单行 Let's Encrypt HTTPS 服务器的示例。

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

以下是自定义 autocert 管理器的示例。

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
    Cache:      autocert.DirCache("/var/www/.cache"), // 缓存目录
  }

  log.Fatal(autotls.RunWithManager(r, &m))
}
```

#

# <翻译结束>


<原文开始>
Run multiple service using Gin

See the [question](https://github.com/gin-gonic/gin/issues/346) and try the following example:

```go
package main

import (
  "log"
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
  "golang.org/x/sync/errgroup"
)

var (
  g errgroup.Group
)

func router01() http.Handler {
  e := gin.New()
  e.Use(gin.Recovery())
  e.GET("/", func(c *gin.Context) {
    c.JSON(
      http.StatusOK,
      gin.H{
        "code":  http.StatusOK,
        "error": "Welcome server 01",
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
        "error": "Welcome server 02",
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
<原文结束>

# <翻译开始>
# 

# <翻译结束>


<原文开始>
Graceful shutdown or restart

There are a few approaches you can use to perform a graceful shutdown or restart. You can make use of third-party packages specifically built for that, or you can manually do the same with the functions and methods from the built-in packages.

##
<原文结束>

# <翻译开始>
# 优雅地关闭或重启

您可以采用几种方法来实现优雅的关闭或重启操作。您可以利用专为此目的而构建的第三方包，或者您也可以使用内置包中的函数和方法手动完成相同的操作。

##

# <翻译结束>


<原文开始>
Third-party packages

We can use [fvbock/endless](https://github.com/fvbock/endless) to replace the default `ListenAndServe`. Refer to issue [#296](https://github.com/gin-gonic/gin/issues/296) for more details.

```go
router := gin.Default()
router.GET("/", handler)
// [...]
endless.ListenAndServe(":4242", router)
```

Alternatives:

* [grace](https://github.com/facebookgo/grace): Graceful restart & zero downtime deploy for Go servers.
* [graceful](https://github.com/tylerb/graceful): Graceful is a Go package enabling graceful shutdown of an http.Handler server.
* [manners](https://github.com/braintree/manners): A polite Go HTTP server that shuts down gracefully.

##
<原文结束>

# <翻译开始>
# 第三方包

我们可以使用[fvbock/endless](https://github.com/fvbock/endless)来替代默认的`ListenAndServe`。更多详情请参考问题[#296](https://github.com/gin-gonic/gin/issues/296)。

```go
router := gin.Default()
router.GET("/", handler)
// [...]
endless.ListenAndServe(":4242", router)
```

可选方案：

* [grace](https://github.com/facebookgo/grace)：为Go服务器提供优雅重启和零停机部署。
* [graceful](https://github.com/tylerb/graceful)：Graceful是一个Go语言包，用于支持http.Handler服务器的优雅关闭。
* [manners](https://github.com/braintree/manners)：一个礼貌的Go HTTP服务器，能够实现优雅地关闭。

# <翻译结束>


<原文开始>
Manually

In case you are using Go 1.8 or a later version, you may not need to use those libraries. Consider using `http.Server`'s built-in [Shutdown()](https://pkg.go.dev/net/http#Server.Shutdown) method for graceful shutdowns. The example below describes its usage, and we've got more examples using gin [here](https://github.com/gin-gonic/examples/tree/master/graceful-shutdown).

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
    c.String(http.StatusOK, "Welcome Gin Server")
  })

  srv := &http.Server{
    Addr:    ":8080",
    Handler: router,
  }

  // Initializing the server in a goroutine so that
  // it won't block the graceful shutdown handling below
  go func() {
    if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
      log.Printf("listen: %s\n", err)
    }
  }()

  // Wait for interrupt signal to gracefully shutdown the server with
  // a timeout of 5 seconds.
  quit := make(chan os.Signal)
  // kill (no param) default send syscall.SIGTERM
  // kill -2 is syscall.SIGINT
  // kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  log.Println("Shutting down server...")

  // The context is used to inform the server it has 5 seconds to finish
  // the request it is currently handling
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  if err := srv.Shutdown(ctx); err != nil {
    log.Fatal("Server forced to shutdown:", err)
  }

  log.Println("Server exiting")
}
```

#
<原文结束>

# <翻译开始>
# 手动

如果你正在使用 Go 1.8 或更高版本，你可能不需要使用那些库。考虑使用 `http.Server` 内置的 [Shutdown()](https://pkg.go.dev/net/http#Server.Shutdown) 方法进行优雅关闭。下面的示例描述了其用法，我们还提供了更多使用 gin 的示例 [在这里](https://github.com/gin-gonic/examples/tree/master/graceful-shutdown)。

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
        c.String(http.StatusOK, "欢迎来到 Gin 服务器")
    })

    srv := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

// 在一个 goroutine 中初始化服务器，以便它不会阻塞下面优雅关闭的处理过程
    go func() {
        if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Printf("监听错误: %s\n", err)
        }
    }()

// 等待中断信号以在 5 秒超时后优雅地关闭服务器
    quit := make(chan os.Signal)
// kill（无参数）默认发送 syscall.SIGTERM
// kill -2 是 syscall.SIGINT
// kill -9 是 syscall.SIGKILL，但无法捕获，所以无需添加
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("正在关闭服务器...")

// 使用上下文来告知服务器它有 5 秒钟的时间完成当前正在处理的请求
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("服务器被迫关闭:", err)
    }

    log.Println("服务器退出")
}
```

#

# <翻译结束>


<原文开始>
Build a single binary with templates

You can build a server into a single binary containing templates by using the [embed](https://pkg.go.dev/embed) package.

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

  // example: /public/assets/images/example.png
  router.StaticFS("/public", http.FS(f))

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Main website",
    })
  })

  router.GET("/foo", func(c *gin.Context) {
    c.HTML(http.StatusOK, "bar.tmpl", gin.H{
      "title": "Foo website",
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

See a complete example in the `https://github.com/gin-gonic/examples/tree/master/assets-in-binary/example02` directory.

#
<原文结束>

# <翻译开始>
# 构建包含模板的单个二进制文件

您可以使用 [embed](https://pkg.go.dev/embed) 包将服务器构建为包含模板的单个二进制文件。

```go
package main

import (
  "embed"
  "html/template"
  "net/http"

  "github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS // 使用embed.FS嵌入资源文件

func main() {
  router := gin.Default()
// 从嵌入的文件系统中解析模板
  templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
  router.SetHTMLTemplate(templ)

// 示例：/public/assets/images/example.png
  router.StaticFS("/public", http.FS(f)) // 设置静态文件服务

// 定义路由与处理函数
  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "主网站",
    })
  })

  router.GET("/foo", func(c *gin.Context) {
    c.HTML(http.StatusOK, "bar.tmpl", gin.H{
      "title": "Foo网站",
    })
  })

// 从嵌入的资源中提供favicon.ico
  router.GET("favicon.ico", func(c *gin.Context) {
    file, _ := f.ReadFile("assets/favicon.ico")
    c.Data(
      http.StatusOK,
      "image/x-icon",
      file,
    )
  })

  router.Run(":8080") // 启动服务器监听8080端口
}

// 查看完整示例，请访问 https://github.com/gin-gonic/examples/tree/master/assets-in-binary/example02 目录。
```

此代码示例展示了如何在Go语言中利用`embed`包将模板和静态资源文件嵌入到单个二进制文件中，并通过 Gin 框架设置路由，实现从嵌入资源中提供模板和静态文件的功能。

# <翻译结束>


<原文开始>
Bind form-data request with custom struct

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

#
<原文结束>

# <翻译开始>
# 

# <翻译结束>


<原文开始>
Try to bind body into different structs

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

#
<原文结束>

# <翻译开始>
# 尝试将请求体绑定到不同的结构体

通常用于绑定请求体的方法会消耗 `c.Request.Body`，并且无法多次调用。

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
// 这里的 c.ShouldBind 会消耗 c.Request.Body，并且不能再重用。
  if errA := c.ShouldBind(&objA); errA == nil {
    c.String(http.StatusOK, `请求体应为 formA 格式`)
// 由于此时 c.Request.Body 已经是 EOF（文件结束符），所以此处总会报错
  } else if errB := c.ShouldBind(&objB); errB == nil {
    c.String(http.StatusOK, `请求体应为 formB 格式`)
  } else {
    ...
  }
}
```

为此，可以使用 `c.ShouldBindBodyWith` 方法。

```go
func SomeHandler(c *gin.Context) {
  objA := formA{}
  objB := formB{}
// 此处读取 c.Request.Body 并将结果存储在上下文中。
  if errA := c.ShouldBindBodyWith(&objA, binding.Form); errA == nil {
    c.String(http.StatusOK, `请求体应为 formA 格式`)
// 此时，它重用了存储在上下文中的请求体。
  } else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
    c.String(http.StatusOK, `请求体应为 formB JSON 格式`)
// 同时，它也支持其他格式
  } else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
    c.String(http.StatusOK, `请求体应为 formB XML 格式`)
  } else {
    ...
  }
}
```

1. `c.ShouldBindBodyWith` 在绑定前将请求体存储到上下文中，这会对性能造成轻微影响。因此，如果你只需要一次性调用绑定方法，不应该使用此方法。
2. 只有针对某些格式——`JSON`、`XML`、`MsgPack`、`ProtoBuf` 才需要这个特性。对于其他格式如 `Query`、`Form`、`FormPost`、`FormMultipart`，可以通过多次调用 `c.ShouldBind()` 而不会对性能造成损害（参见 [#1341](https://github.com/gin-gonic/gin/pull/1341)）。

# <翻译结束>


<原文开始>
Bind form-data request with custom struct and custom tag

```go
const (
  customerTag = "url"
  defaultMemory = 32 << 20
)

type customerBinding struct {}

func (customerBinding) Name() string {
  return "form"
}

func (customerBinding) Bind(req *http.Request, obj any) error {
  if err := req.ParseForm(); err != nil {
    return err
  }
  if err := req.ParseMultipartForm(defaultMemory); err != nil {
    if err != http.ErrNotMultipart {
      return err
    }
  }
  if err := binding.MapFormWithTag(obj, req.Form, customerTag); err != nil {
    return err
  }
  return validate(obj)
}

func validate(obj any) error {
  if binding.Validator == nil {
    return nil
  }
  return binding.Validator.ValidateStruct(obj)
}

// Now we can do this!!!
// FormA is an external type that we can't modify it's tag
type FormA struct {
  FieldA string `url:"field_a"`
}

func ListHandler(s *Service) func(ctx *gin.Context) {
  return func(ctx *gin.Context) {
    var urlBinding = customerBinding{}
    var opt FormA
    err := ctx.MustBindWith(&opt, urlBinding)
    if err != nil {
      ...
    }
    ...
  }
}
```

#
<原文结束>

# <翻译开始>
# 在Go语言中，通过自定义结构体和标签绑定表单数据请求

```go
// 定义常量
const (
  customerTag = "url" // 自定义标签名称
  defaultMemory = 32 << 20 // 默认内存大小
)

// 定义customerBinding结构体
type customerBinding struct {}

// 定义customerBinding的Name方法
func (customerBinding) Name() string {
  return "form" // 返回“form”作为请求绑定类型
}

// 定义customerBinding的Bind方法，用于将请求与对象进行绑定
func (customerBinding) Bind(req *http.Request, obj any) error {
// 解析请求的表单数据
  if err := req.ParseForm(); err != nil {
    return err
  }
  
// 解析multipart表单数据，并分配默认内存大小
  if err := req.ParseMultipartForm(defaultMemory); err != nil {
    if err != http.ErrNotMultipart { // 如果错误不是非multipart形式，则返回错误
      return err
    }
  }

// 使用自定义标签将表单数据映射到对象
  if err := binding.MapFormWithTag(obj, req.Form, customerTag); err != nil {
    return err
  }

// 对对象进行验证
  return validate(obj)
}

// 验证函数
func validate(obj any) error {
  if binding.Validator == nil { // 如果没有设置验证器，则直接返回nil
    return nil
  }
  return binding.Validator.ValidateStruct(obj) // 对结构体执行验证
}

// 外部定义的不可修改其标签的FormA类型
type FormA struct {
  FieldA string `url:"field_a"` // 使用自定义标签“url”
}

// 定义处理列表请求的处理器函数
func ListHandler(s *Service) func(ctx *gin.Context) {
  return func(ctx *gin.Context) {
// 创建一个customerBinding实例
    var urlBinding = customerBinding{}
    
// 创建FormA类型的变量opt
    var opt FormA
    
// 使用自定义绑定器将请求上下文中的表单数据绑定到opt
    err := ctx.MustBindWith(&opt, urlBinding)
    if err != nil {
// 若出现错误，则处理错误...
    }
    
// 继续处理其他逻辑...
  }
}
```

这个代码示例展示了一个自定义的表单数据绑定方法，使用了`customerBinding`结构体及自定义标签`url`。在实际应用中，可以通过`customerBinding`将HTTP请求中的表单数据绑定到带有特定标签（如`url:"field_a"`）的结构体字段上，并对其进行验证。在`ListHandler`函数中，我们演示了如何利用这个自定义绑定器来处理 Gin 框架中的 HTTP 请求。

# <翻译结束>


<原文开始>
http2 server push

http.Pusher is supported only **go1.8+**. See the [golang blog](https://go.dev/blog/h2push) for detail information.

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
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
  r := gin.Default()
  r.Static("/assets", "./assets")
  r.SetHTMLTemplate(html)

  r.GET("/", func(c *gin.Context) {
    if pusher := c.Writer.Pusher(); pusher != nil {
      // use pusher.Push() to do server push
      if err := pusher.Push("/assets/app.js", nil); err != nil {
        log.Printf("Failed to push: %v", err)
      }
    }
    c.HTML(http.StatusOK, "https", gin.H{
      "status": "success",
    })
  })

  // Listen and Server in https://127.0.0.1:8080
  r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
```

#
<原文结束>

# <翻译开始>
# http2 服务器推送

http.Pusher 功能仅在 **go1.8+** 版本中支持。详细信息请参阅 [golang 官方博客](https://go.dev/blog/h2push)。

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
  <title>Https 测试</title>
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
// 使用 pusher.Push() 实现服务器推送
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

# <翻译结束>


<原文开始>
Define format for the log of routes

The default log of routes is:

```sh
[GIN-debug] POST   /foo                      --> main.main.func1 (3 handlers)
[GIN-debug] GET    /bar                      --> main.main.func2 (3 handlers)
[GIN-debug] GET    /status                   --> main.main.func3 (3 handlers)
```

If you want to log this information in given format (e.g. JSON, key values or something else), then you can define this format with `gin.DebugPrintRouteFunc`.
In the example below, we log all routes with standard log package but you can use another log tools that suits of your needs.

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

  // Listen and Server in http://0.0.0.0:8080
  r.Run()
}
```

#
<原文结束>

# <翻译开始>
# 定义路由日志的格式

默认的路由日志是：

```sh
[GIN-debug] POST   /foo                      --> main.main.func1 (3 handlers)
[GIN-debug] GET    /bar                      --> main.main.func2 (3 handlers)
[GIN-debug] GET    /status                   --> main.main.func3 (3 handlers)
```

如果你想以特定的格式（例如JSON、键值对或其他格式）记录这些信息，可以通过设置`gin.DebugPrintRouteFunc`来自定义该格式。

在下面的例子中，我们使用标准的日志包来记录所有路由信息，但你可以根据自己的需求选择其他的日志工具。

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

// 监听并服务在 http://0.0.0.0:8080
  r.Run()
}
```

# 结束

# <翻译结束>


<原文开始>
Set and get a cookie

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/cookie", func(c *gin.Context) {

      cookie, err := c.Cookie("gin_cookie")

      if err != nil {
          cookie = "NotSet"
          c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
      }

      fmt.Printf("Cookie value: %s \n", cookie)
  })

  router.Run()
}
```


<原文结束>

# <翻译开始>
# 设置和获取一个cookie

```go
import (
    "fmt"

    "github.com/gin-gonic/gin"
)

func main() {
// 初始化路由
    router := gin.Default()

// 定义GET请求路由处理函数
    router.GET("/cookie", func(c *gin.Context) {

// 获取名为"gin_cookie"的cookie值
        cookie, err := c.Cookie("gin_cookie")

// 如果获取时出现错误（即cookie未设置）
        if err != nil {
// 将cookie值设为"NotSet"
            cookie = "NotSet"

// 设置名为"gin_cookie"的cookie，有效期为3600秒，路径为"/"，域名是"localhost"，不进行安全传输（false），且标记为HTTP only（true）
            c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
        }

// 输出cookie的当前值
        fmt.Printf("Cookie value: %s \n", cookie)
    })

// 运行服务器
    router.Run()
}
```

# <翻译结束>


<原文开始>
Don't trust all proxies

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


<原文结束>

# <翻译开始>
# 不要信任所有代理

Gin 允许你指定哪些头部包含真实的客户端 IP 地址（如果有），以及指定你信任的代理（或直接客户端），以设置这些头部之一。

在你的 `gin.Engine` 上使用函数 `SetTrustedProxies()` 来指定网络地址或网络CIDR，以便信任来自这些地方的客户端请求头中与客户端IP相关的信息。它们可以是IPv4地址、IPv4 CIDR、IPv6地址或IPv6 CIDR。

**注意：** 如果你不使用 `SetTrustedProxies()` 函数指定可信代理，Gin 默认会信任所有代理，**这并不安全**。同时，如果你不使用任何代理，可以通过 `Engine.SetTrustedProxies(nil)` 禁用此功能，此时 `Context.ClientIP()` 将直接返回远程地址，以避免不必要的计算。

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.SetTrustedProxies([]string{"192.168.1.2"})

  router.GET("/", func(c *gin.Context) {
// 如果客户端是192.168.1.2，则从可信赖部分使用 X-Forwarded-For 头部推断原始客户端IP。
// 否则，直接返回直接客户端IP。
    fmt.Printf("ClientIP: %s\n", c.ClientIP())
  })
  router.Run()
}
```

**提示：** 如果你正在使用CDN服务，可以设置 `Engine.TrustedPlatform` 以跳过 TrustedProxies 检查，它具有高于 TrustedProxies 的优先级。请参考以下示例：

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
// 使用预定义头部 gin.PlatformXXX
  router.TrustedPlatform = gin.PlatformGoogleAppEngine
// 或为其他可信代理服务设置你自己的可信请求头部
// 不要将其设置为任何可疑请求头部，这是不安全的
  router.TrustedPlatform = "X-CDN-IP"

  router.GET("/", func(c *gin.Context) {
// 如果设置了 TrustedPlatform，ClientIP() 会解析相应的头部并直接返回IP
    fmt.Printf("ClientIP: %s\n", c.ClientIP())
  })
  router.Run()
}
```

# <翻译结束>


<原文开始>
Testing

The `net/http/httptest` package is preferable way for HTTP testing.

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

Test for code example above:

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

<原文结束>

# <翻译开始>
# 测试

`net/http/httptest` 包是进行 HTTP 测试的首选方式。

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

// 设置路由函数
func setupRouter() *gin.Engine {
  r := gin.Default()
// 添加GET请求处理函数，访问路径为"/ping"
  r.GET("/ping", func(c *gin.Context) {
// 返回状态码200和消息"pong"
    c.String(http.StatusOK, "pong")
  })
  return r
}

func main() {
// 初始化并运行路由
  r := setupRouter()
  r.Run(":8080")
}
```

针对上述代码示例的测试：

```go
package main

import (
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/stretchr/testify/assert"
)

// 测试PingRoute
func TestPingRoute(t *testing.T) {
// 创建并初始化路由
  router := setupRouter()

// 创建响应记录器（模拟HTTP响应）
  w := httptest.NewRecorder()
// 创建GET请求实例，请求路径为"/ping"
  req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
// 调用ServeHTTP方法处理请求
  router.ServeHTTP(w, req)

// 断言响应状态码应为200
  assert.Equal(t, http.StatusOK, w.Code)
// 断言响应体内容应为"pong"
  assert.Equal(t, "pong", w.Body.String())
}
```

# <翻译结束>

