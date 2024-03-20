
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// Content-Type MIME of the most common data formats.
<原文结束>

# <翻译开始>
// Content-Type MIME 是最常见的数据格式的 MIME 类型。
# <翻译结束>


<原文开始>
// BodyBytesKey indicates a default body bytes key.
<原文结束>

# <翻译开始>
// BodyBytesKey 表示一个默认的正文字节键。
# <翻译结束>


<原文开始>
// ContextKey is the key that a Context returns itself for.
<原文结束>

# <翻译开始>
// ContextKey 是一个键，用于在 Context 中返回其自身。
# <翻译结束>


<原文开始>
// abortIndex represents a typical value used in abort functions.
<原文结束>

# <翻译开始>
// abortIndex 表示在中止函数中常用的一个典型值。
# <翻译结束>


<原文开始>
// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
<原文结束>

# <翻译开始>
// Context 是 gin 中最重要的部分。它允许我们在中间件之间传递变量，管理流程，验证请求的 JSON，并例如渲染 JSON 响应。
# <翻译结束>


<原文开始>
// This mutex protects Keys map.
<原文结束>

# <翻译开始>
// 这个互斥锁保护了Keys映射。
# <翻译结束>


<原文开始>
// Keys is a key/value pair exclusively for the context of each request.
<原文结束>

# <翻译开始>
// Keys 是一组键值对，它在每个请求的上下文中具有唯一性。
# <翻译结束>


<原文开始>
// Errors is a list of errors attached to all the handlers/middlewares who used this context.
<原文结束>

# <翻译开始>
// Errors 是一个错误列表，其中包含了所有使用了此上下文的处理器/中间件所附加的错误。
# <翻译结束>


<原文开始>
// Accepted defines a list of manually accepted formats for content negotiation.
<原文结束>

# <翻译开始>
// Accepted 定义了一个手动接受的内容协商格式列表。
# <翻译结束>


<原文开始>
// queryCache caches the query result from c.Request.URL.Query().
<原文结束>

# <翻译开始>
// queryCache 对从 c.Request.URL.Query() 获取的查询结果进行缓存。
# <翻译结束>


<原文开始>
	// formCache caches c.Request.PostForm, which contains the parsed form data from POST, PATCH,
	// or PUT body parameters.
<原文结束>

# <翻译开始>
// formCache 对 c.Request.PostForm 进行缓存，其中包含从 POST、PATCH 或 PUT 请求体参数解析得到的表单数据。
# <翻译结束>


<原文开始>
	// SameSite allows a server to define a cookie attribute making it impossible for
	// the browser to send this cookie along with cross-site requests.
<原文结束>

# <翻译开始>
// SameSite 允许服务器定义一个 cookie 属性，使得浏览器无法在跨站请求中携带此 cookie。
# <翻译结束>


<原文开始>
// Copy returns a copy of the current context that can be safely used outside the request's scope.
// This has to be used when the context has to be passed to a goroutine.
<原文结束>

# <翻译开始>
// Copy 返回当前上下文的副本，该副本可以在请求范围之外安全使用。
// 当需要将上下文传递给一个goroutine时，必须使用此方法。
# <翻译结束>


<原文开始>
// HandlerName returns the main handler's name. For example if the handler is "handleGetUsers()",
// this function will return "main.handleGetUsers".
<原文结束>

# <翻译开始>
// HandlerName 返回主处理程序的名称。例如，如果处理程序是 "handleGetUsers()"，
// 该函数将返回 "main.handleGetUsers"。
// 例如:
// 如果处理程序为“handleGetUsers()”，则此函数将返回“main.handleGetUsers”
// 包名为"github.com/888go/gin",返回如下:
// github.com/888go/gin.handleGetUsers
# <翻译结束>


<原文开始>
// HandlerNames returns a list of all registered handlers for this context in descending order,
// following the semantics of HandlerName()
<原文结束>

# <翻译开始>
// HandlerNames 返回与此上下文关联的已注册处理程序的降序列表，遵循HandlerName()的语义
// 返回数组参考如下:
// 0 = {string} "github.com/888go/gin.TestContextHandlerNames.func1"
// 1 = {string} "github.com/888go/gin.handlerNameTest"
// 2 = {string} "github.com/888go/gin.TestContextHandlerNames.func2"
// 3 = {string} "github.com/888go/gin.handlerNameTest2"
# <翻译结束>


<原文开始>
// Handler returns the main handler.
<原文结束>

# <翻译开始>
// Handler 返回主处理程序。
# <翻译结束>


<原文开始>
// FullPath returns a matched route full path. For not found routes
// returns an empty string.
//
//	router.GET("/user/:id", func(c *gin.Context) {
//	    c.FullPath() == "/user/:id" // true
//	})
<原文结束>

# <翻译开始>
// FullPath 返回已匹配路由的完整路径。对于未找到的路由，返回一个空字符串。
//
// 示例：
//   router.GET("/user/:id", func(c *gin.Context) {
//       c.FullPath() == "/user/:id" // 将会返回 true
//   })
# <翻译结束>


<原文开始>
// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in GitHub.
<原文结束>

# <翻译开始>
// Next 应仅在中间件内部使用。
// 它在调用处理程序内部执行链中待处理的后续处理程序。
// 参考 GitHub 上的示例。
# <翻译结束>


<原文开始>
// IsAborted returns true if the current context was aborted.
<原文结束>

# <翻译开始>
// IsAborted 返回当前上下文是否已中止。
# <翻译结束>


<原文开始>
// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
<原文结束>

# <翻译开始>
// Abort 阻止待处理的处理器被调用。请注意，这不会停止当前处理器。
// 假设你有一个授权中间件用于验证当前请求是否已授权。
// 如果授权失败（例如，密码不匹配），则调用 Abort 来确保该请求的剩余处理器不会被调用。
# <翻译结束>


<原文开始>
// AbortWithStatus calls `Abort()` and writes the headers with the specified status code.
// For example, a failed attempt to authenticate a request could use: context.AbortWithStatus(401).
<原文结束>

# <翻译开始>
// AbortWithStatus 方法调用 `Abort()`，并使用指定的状态码写入头部信息。
// 例如，在尝试验证请求失败时，可以这样使用：context.AbortWithStatus(401)。
// 
// 注释翻译成中文如下：
// 
// AbortWithStatus 函数会调用 `Abort()` 函数，并携带特定状态码设置响应头。
// 举例来说，如果尝试验证请求失败，可以采用如下的方式：context.AbortWithStatus(401)。
# <翻译结束>


<原文开始>
// AbortWithStatusJSON calls `Abort()` and then `JSON` internally.
// This method stops the chain, writes the status code and return a JSON body.
// It also sets the Content-Type as "application/json".
<原文结束>

# <翻译开始>
// AbortWithStatusJSON 在内部调用`Abort()`和`JSON`方法。
// 该方法中断执行链，写入状态码并返回一个JSON格式的响应体。
// 同时将Content-Type设置为"application/json"。
# <翻译结束>


<原文开始>
// AbortWithError calls `AbortWithStatus()` and `Error()` internally.
// This method stops the chain, writes the status code and pushes the specified error to `c.Errors`.
// See Context.Error() for more details.
<原文结束>

# <翻译开始>
// AbortWithError 在内部调用 `AbortWithStatus()` 和 `Error()`。
// 该方法停止执行链，写入状态码并将指定错误推送到 `c.Errors`。
// 有关更多详细信息，请参阅 Context.Error()。
# <翻译结束>


<原文开始>
// Error attaches an error to the current context. The error is pushed to a list of errors.
// It's a good idea to call Error for each error that occurred during the resolution of a request.
// A middleware can be used to collect all the errors and push them to a database together,
// print a log, or append it in the HTTP response.
// Error will panic if err is nil.
<原文结束>

# <翻译开始>
// Error 将错误附着到当前上下文中。该错误会被推送到错误列表中。
// 在请求解析过程中，对于发生的每个错误调用 Error 是一个好主意。
// 可以使用中间件来收集所有错误，并将它们一起推送到数据库、打印日志或将其添加到HTTP响应中。
// 如果err为nil，Error将会触发panic。
# <翻译结束>


<原文开始>
// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
<原文结束>

# <翻译开始>
// Set 用于为当前上下文独占存储一个新的键值对。
// 如果之前未使用过，它还会初始化 c.Keys。
# <翻译结束>


<原文开始>
// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
<原文结束>

# <翻译开始>
// Get 方法根据给定的键返回其对应的值，即：(value, true)。
// 若该值不存在，则返回 (nil, false)。
# <翻译结束>


<原文开始>
// MustGet returns the value for the given key if it exists, otherwise it panics.
<原文结束>

# <翻译开始>
// MustGet 返回给定键对应的值，如果该键存在。否则，函数会触发panic异常。
# <翻译结束>


<原文开始>
// GetString returns the value associated with the key as a string.
<原文结束>

# <翻译开始>
// GetString 方法返回与键关联的值，以字符串形式。
# <翻译结束>


<原文开始>
// GetBool returns the value associated with the key as a boolean.
<原文结束>

# <翻译开始>
// GetBool返回与key关联的值，将其转化为布尔类型。
# <翻译结束>


<原文开始>
// GetInt returns the value associated with the key as an integer.
<原文结束>

# <翻译开始>
// GetInt 通过键返回与其关联的整数值。
# <翻译结束>


<原文开始>
// GetInt64 returns the value associated with the key as an integer.
<原文结束>

# <翻译开始>
// GetInt64 以整数形式返回与键关联的值。
# <翻译结束>


<原文开始>
// GetUint returns the value associated with the key as an unsigned integer.
<原文结束>

# <翻译开始>
// GetUint 返回与键关联的值，以无符号整数形式。
# <翻译结束>


<原文开始>
// GetUint64 returns the value associated with the key as an unsigned integer.
<原文结束>

# <翻译开始>
// GetUint64返回与key关联的值，将其转化为无符号整数。
# <翻译结束>


<原文开始>
// GetFloat64 returns the value associated with the key as a float64.
<原文结束>

# <翻译开始>
// GetFloat64 通过key返回关联的float64类型的值。
# <翻译结束>


<原文开始>
// GetTime returns the value associated with the key as time.
<原文结束>

# <翻译开始>
// GetTime 函数通过键返回其关联的时间值。
# <翻译结束>


<原文开始>
// GetDuration returns the value associated with the key as a duration.
<原文结束>

# <翻译开始>
// GetDuration返回与键关联的值，其类型为持续时间。
# <翻译结束>


<原文开始>
// GetStringSlice returns the value associated with the key as a slice of strings.
<原文结束>

# <翻译开始>
// GetStringSlice 函数返回与键关联的值，该值为字符串切片。
# <翻译结束>


<原文开始>
// GetStringMap returns the value associated with the key as a map of interfaces.
<原文结束>

# <翻译开始>
// GetStringMap 返回与键关联的值，该值为接口映射（map）类型。
# <翻译结束>


<原文开始>
// GetStringMapString returns the value associated with the key as a map of strings.
<原文结束>

# <翻译开始>
// GetStringMapString返回与键关联的值，该值为字符串映射（map）类型。
# <翻译结束>


<原文开始>
// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
<原文结束>

# <翻译开始>
// GetStringMapStringSlice 返回与键关联的值，该值为字符串到字符串切片的映射。
# <翻译结束>


<原文开始>
// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
//
//	router.GET("/user/:id", func(c *gin.Context) {
//	    // a GET request to /user/john
//	    id := c.Param("id") // id == "/john"
//	    // a GET request to /user/john/
//	    id := c.Param("id") // id == "/john/"
//	})
<原文结束>

# <翻译开始>
// Param 返回URL参数的值。
// 这是c.Params.ByName(key)的一个快捷方式。
//
// 示例：
// 使用router.GET设置路由处理函数，访问"/user/:id"时，
// ```go
// router.GET("/user/:id", func(c *gin.Context) {
//     // 当发送一个GET请求到/user/john
//     id := c.Param("id") // 这时id的值为"john"
//     // 当发送一个GET请求到/user/john/
//     id := c.Param("id") // 这时id的值为"john/"
// })
// ```
// 注：在上述示例中，":id"是一个动态参数，其值会被解析并存储到c.Param("id")中。
# <翻译结束>


<原文开始>
// AddParam adds param to context and
// replaces path param key with given value for e2e testing purposes
// Example Route: "/user/:id"
// AddParam("id", 1)
// Result: "/user/1"
<原文结束>

# <翻译开始>
// AddParam 将参数添加到上下文，并为了端到端测试的目的，用给定的值替换路径参数键
// 示例路由："/user/:id"
// AddParam("id", 1)
// 结果："/user/1"
# <翻译结束>


<原文开始>
// Query returns the keyed url query value if it exists,
// otherwise it returns an empty string `("")`.
// It is shortcut for `c.Request.URL.Query().Get(key)`
//
//	    GET /path?id=1234&name=Manu&value=
//		   c.Query("id") == "1234"
//		   c.Query("name") == "Manu"
//		   c.Query("value") == ""
//		   c.Query("wtf") == ""
<原文结束>

# <翻译开始>
// Query方法返回键所对应的URL查询值，如果该值存在，则返回该值，否则返回一个空字符串 `("")`。
// 这是 `c.Request.URL.Query().Get(key)` 的快捷方式。
//
//    GET /path?id=1234&name=Manu&value=
//       c.Query("id") 返回 "1234"
//       c.Query("name") 返回 "Manu"
//       c.Query("value") 返回 ""
//       c.Query("wtf") 返回 ""
# <翻译结束>


<原文开始>
// DefaultQuery returns the keyed url query value if it exists,
// otherwise it returns the specified defaultValue string.
// See: Query() and GetQuery() for further information.
//
//	GET /?name=Manu&lastname=
//	c.DefaultQuery("name", "unknown") == "Manu"
//	c.DefaultQuery("id", "none") == "none"
//	c.DefaultQuery("lastname", "none") == ""
<原文结束>

# <翻译开始>
// DefaultQuery 返回键值对形式的URL查询参数的值，如果该参数存在，则返回其值；否则返回指定的defaultValue字符串。
// 有关更多详细信息，请参阅：Query() 和 GetQuery()。
//
// 示例：
// 请求 GET /?name=Manu&lastname=
// c.DefaultQuery("name", "unknown") 将返回 "Manu"
// c.DefaultQuery("id", "none") 将返回 "none"
// c.DefaultQuery("lastname", "none") 将返回 ""
# <翻译结束>


<原文开始>
// GetQuery is like Query(), it returns the keyed url query value
// if it exists `(value, true)` (even when the value is an empty string),
// otherwise it returns `("", false)`.
// It is shortcut for `c.Request.URL.Query().Get(key)`
//
//	GET /?name=Manu&lastname=
//	("Manu", true) == c.GetQuery("name")
//	("", false) == c.GetQuery("id")
//	("", true) == c.GetQuery("lastname")
<原文结束>

# <翻译开始>
// GetQuery 方法类似于 Query()，当给定键的 URL 查询值存在时，它返回该查询值及其对应的布尔值 `(value, true)`（即使该值是一个空字符串）；
// 否则，它返回 `("", false)`。此方法是 `c.Request.URL.Query().Get(key)` 的快捷方式。
//
// 示例：
// 请求 GET /?name=Manu&lastname=
// ("Manu", true) 等价于 c.GetQuery("name")
// ("", false) 等价于 c.GetQuery("id")
// ("", true) 等价于 c.GetQuery("lastname")
# <翻译结束>


<原文开始>
// QueryArray returns a slice of strings for a given query key.
// The length of the slice depends on the number of params with the given key.
<原文结束>

# <翻译开始>
// QueryArray 函数针对给定的查询键返回一个字符串切片。
// 返回切片的长度取决于具有该键的参数的数量。
# <翻译结束>


<原文开始>
// GetQueryArray returns a slice of strings for a given query key, plus
// a boolean value whether at least one value exists for the given key.
<原文结束>

# <翻译开始>
// GetQueryArray 返回给定查询键的字符串切片，以及
// 一个布尔值，表示该键是否存在至少一个值。
# <翻译结束>


<原文开始>
// QueryMap returns a map for a given query key.
<原文结束>

# <翻译开始>
// QueryMap 根据给定的查询键返回一个映射（map）。
# <翻译结束>


<原文开始>
// GetQueryMap returns a map for a given query key, plus a boolean value
// whether at least one value exists for the given key.
<原文结束>

# <翻译开始>
// GetQueryMap 为给定的查询键返回一个映射（map），同时返回一个布尔值，
// 表示该键是否存在至少一个值。
# <翻译结束>


<原文开始>
// PostForm returns the specified key from a POST urlencoded form or multipart form
// when it exists, otherwise it returns an empty string `("")`.
<原文结束>

# <翻译开始>
// PostForm 返回从 POST 请求中 urlencoded 表单或 multipart 表单获取的指定键值，如果该键存在，则返回其对应的值；否则返回空字符串 `("")`。
# <翻译结束>


<原文开始>
// DefaultPostForm returns the specified key from a POST urlencoded form or multipart form
// when it exists, otherwise it returns the specified defaultValue string.
// See: PostForm() and GetPostForm() for further information.
<原文结束>

# <翻译开始>
// DefaultPostForm 函数在 POST 请求的 urlencoded 表单或 multipart 表单中查找指定键的值，
// 如果该键存在，则返回对应的值，否则返回指定的 defaultValue 字符串。
// 有关更多信息，请参阅 PostForm() 和 GetPostForm() 函数。
# <翻译结束>


<原文开始>
// GetPostForm is like PostForm(key). It returns the specified key from a POST urlencoded
// form or multipart form when it exists `(value, true)` (even when the value is an empty string),
// otherwise it returns ("", false).
// For example, during a PATCH request to update the user's email:
//
//	    email=mail@example.com  -->  ("mail@example.com", true) := GetPostForm("email") // set email to "mail@example.com"
//		   email=                  -->  ("", true) := GetPostForm("email") // set email to ""
//	                            -->  ("", false) := GetPostForm("email") // do nothing with email
<原文结束>

# <翻译开始>
// 以下是将给定的Go注释翻译成中文：
// 
// GetPostForm 类似于 PostForm(key)。当存在时，它从POST urlencoded表单或multipart表单中返回指定键的值 `(value, true)`（即使该值为空字符串），
// 否则返回 ("", false)。
// 例如，在进行PATCH请求以更新用户邮箱时：
//
//	    email=mail@example.com  -->  ("mail@example.com", true) := GetPostForm("email") // 将邮箱设置为 "mail@example.com"
//		   email=                  -->  ("", true) := GetPostForm("email") // 将邮箱设置为空字符串
//	                            -->  ("", false) := GetPostForm("email") // 对邮箱不做任何处理
# <翻译结束>


<原文开始>
// PostFormArray returns a slice of strings for a given form key.
// The length of the slice depends on the number of params with the given key.
<原文结束>

# <翻译开始>
// PostFormArray 为给定的表单键返回一个字符串切片。
// 切片的长度取决于具有该键的参数的数量。
# <翻译结束>


<原文开始>
// GetPostFormArray returns a slice of strings for a given form key, plus
// a boolean value whether at least one value exists for the given key.
<原文结束>

# <翻译开始>
// GetPostFormArray 针对给定表单键返回一个字符串切片，以及
// 一个布尔值，表示该键是否存在至少一个值。
# <翻译结束>


<原文开始>
// PostFormMap returns a map for a given form key.
<原文结束>

# <翻译开始>
// PostFormMap 为给定的表单键返回一个映射（map）。
# <翻译结束>


<原文开始>
// GetPostFormMap returns a map for a given form key, plus a boolean value
// whether at least one value exists for the given key.
<原文结束>

# <翻译开始>
// GetPostFormMap 为给定的表单键返回一个映射，同时返回一个布尔值，
// 表示是否存在至少一个为此给定键的值。
# <翻译结束>


<原文开始>
// get is an internal method and returns a map which satisfies conditions.
<原文结束>

# <翻译开始>
// get 是一个内部方法，它返回一个满足特定条件的地图（map）。
# <翻译结束>


<原文开始>
// FormFile returns the first file for the provided form key.
<原文结束>

# <翻译开始>
// FormFile返回提供的表单键所对应的第一个文件。
# <翻译结束>


<原文开始>
// MultipartForm is the parsed multipart form, including file uploads.
<原文结束>

# <翻译开始>
// MultipartForm 是已解析的多部分表单，包括文件上传。
# <翻译结束>


<原文开始>
// SaveUploadedFile uploads the form file to specific dst.
<原文结束>

# <翻译开始>
// SaveUploadedFile 将表单文件上传到指定的dst。
# <翻译结束>


<原文开始>
// Bind checks the Method and Content-Type to select a binding engine automatically,
// Depending on the "Content-Type" header different bindings are used, for example:
//
//	"application/json" --> JSON binding
//	"application/xml"  --> XML binding
//
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// It writes a 400 error and sets Content-Type header "text/plain" in the response if input is not valid.
<原文结束>

# <翻译开始>
// Bind 会根据 Method 和 Content-Type 自动选择绑定引擎，
// 根据 "Content-Type" 头部的不同，使用不同的绑定方式，例如：
//
//	"application/json" --> JSON 绑定
//	"application/xml"  --> XML 绑定
//
// 若 Content-Type 为 "application/json"，它将把请求体解析为 JSON，同时可将 XML 视为 JSON 输入进行处理。
// 它会将 json 数据解码到指定的结构体指针中。
// 如果输入无效，则在响应中写入 400 错误，并设置 Content-Type 头部为 "text/plain"。
# <翻译结束>


<原文开始>
// BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
<原文结束>

# <翻译开始>
// BindJSON 是一个快捷方式，等同于 c.MustBindWith(obj, binding.JSON)。
# <翻译结束>


<原文开始>
// BindXML is a shortcut for c.MustBindWith(obj, binding.BindXML).
<原文结束>

# <翻译开始>
// BindXML 是一个快捷方式，用于 c.MustBindWith(obj, binding.BindXML)。 
// 
// 更详细的翻译：
// 
// BindXML 是一个便捷方法，它等同于调用 c.MustBindWith(obj, binding.BindXML)。
// 其中，c 通常代表上下文（Context），obj 代表要绑定的对象，binding.BindXML 表示使用 XML 绑定方式进行数据绑定。这个方法会确保 XML 数据成功绑定到对象上，如果绑定失败，则会触发 panic。
# <翻译结束>


<原文开始>
// BindQuery is a shortcut for c.MustBindWith(obj, binding.Query).
<原文结束>

# <翻译开始>
// BindQuery 是一个快捷方式，用于 c.MustBindWith(obj, binding.Query)。
// 如果解析错误,它将使用HTTP 400中止请求
// BindQuery 函数只绑定 url 查询参数而忽略 post 数据。参阅详细信息:
// https://gin-gonic.com/zh-cn/docs/examples/only-bind-query-string/
# <翻译结束>


<原文开始>
// BindYAML is a shortcut for c.MustBindWith(obj, binding.YAML).
<原文结束>

# <翻译开始>
// BindYAML 是一个快捷方式，等同于 c.MustBindWith(obj, binding.YAML)。
// 注意: 如果解析错误,它将使用HTTP 400中止请求
# <翻译结束>


<原文开始>
// BindTOML is a shortcut for c.MustBindWith(obj, binding.TOML).
<原文结束>

# <翻译开始>
// BindTOML 是一个快捷方式，用于 c.MustBindWith(obj, binding.TOML)。
// 注意: 如果解析错误,它将使用HTTP 400中止请求
# <翻译结束>


<原文开始>
// BindHeader is a shortcut for c.MustBindWith(obj, binding.Header).
<原文结束>

# <翻译开始>
// BindHeader 是一个快捷方式，等同于 c.MustBindWith(obj, binding.Header)。
// 注意: 如果解析错误,它将使用HTTP 400中止请求
# <翻译结束>


<原文开始>
// BindUri binds the passed struct pointer using binding.Uri.
// It will abort the request with HTTP 400 if any error occurs.
<原文结束>

# <翻译开始>
// BindUri通过binding.Uri将传递的结构体指针进行绑定。
// 如果发生任何错误，它将使用HTTP 400中止请求。
# <翻译结束>


<原文开始>
// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
<原文结束>

# <翻译开始>
// MustBindWith 使用指定的绑定引擎绑定传入的结构体指针。
// 如果在执行过程中出现任何错误，它将终止请求并返回HTTP状态码400。
// 请参阅binding包以获取更多信息。
# <翻译结束>


<原文开始>
// ShouldBind checks the Method and Content-Type to select a binding engine automatically,
// Depending on the "Content-Type" header different bindings are used, for example:
//
//	"application/json" --> JSON binding
//	"application/xml"  --> XML binding
//
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// Like c.Bind() but this method does not set the response status code to 400 or abort if input is not valid.
<原文结束>

# <翻译开始>
// ShouldBind 会根据 Method（请求方法）和 Content-Type（内容类型）自动选择一个绑定引擎，
// 根据 "Content-Type" 头部的不同，采用不同的绑定方式，例如：
//
//	"application/json" --> JSON 绑定
//	"application/xml"  --> XML 绑定
//
// 若 Content-Type 为 "application/json"，它将把请求体当作 JSON 解析，并使用 JSON 或 XML 作为 JSON 输入。
// 它会将解析后的 json 数据解码到指定的结构体指针中。
// 类似于 c.Bind() 方法，但该方法在输入无效时不会将响应状态码设置为 400 或终止执行。
//
// 注意: c.ShouldBind***方法不能多次被调用, 如果绑定类型为" JSON, XML, MsgPack, ProtoBuf", 第一次绑定之后 c.Request.Body会设置成EOF, 如果需要多次绑定, 可以使用c.ShouldBindBodyWith
# <翻译结束>


<原文开始>
// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
<原文结束>

# <翻译开始>
// ShouldBindJSON 是 c.ShouldBindWith(obj, binding.JSON) 的快捷方式。
//
// 注意: c.ShouldBind***方法不能多次被调用, 如果绑定类型为" JSON, XML, MsgPack, ProtoBuf", 第一次绑定之后 c.Request.Body会设置成EOF, 如果需要多次绑定, 可以使用c.ShouldBindBodyWith
# <翻译结束>


<原文开始>
// ShouldBindXML is a shortcut for c.ShouldBindWith(obj, binding.XML).
<原文结束>

# <翻译开始>
// ShouldBindXML 是 c.ShouldBindWith(obj, binding.XML) 的快捷方式。
//
// 注意: c.ShouldBind***方法不能多次被调用, 如果绑定类型为" JSON, XML, MsgPack, ProtoBuf", 第一次绑定之后 c.Request.Body会设置成EOF, 如果需要多次绑定, 可以使用c.ShouldBindBodyWith
# <翻译结束>


<原文开始>
// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
<原文结束>

# <翻译开始>
// ShouldBindQuery 是一个快捷方式，用于 c.ShouldBindWith(obj, binding.Query)。
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
// ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据。参阅详细信息:
// https://gin-gonic.com/zh-cn/docs/examples/only-bind-query-string/
# <翻译结束>


<原文开始>
// ShouldBindYAML is a shortcut for c.ShouldBindWith(obj, binding.YAML).
<原文结束>

# <翻译开始>
// ShouldBindYAML 是 c.ShouldBindWith(obj, binding.YAML) 的快捷方式。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
# <翻译结束>


<原文开始>
// ShouldBindTOML is a shortcut for c.ShouldBindWith(obj, binding.TOML).
<原文结束>

# <翻译开始>
// ShouldBindTOML 是 c.ShouldBindWith(obj, binding.TOML) 的快捷方式。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
# <翻译结束>


<原文开始>
// ShouldBindHeader is a shortcut for c.ShouldBindWith(obj, binding.Header).
<原文结束>

# <翻译开始>
// ShouldBindHeader 是一个快捷方式，用于 c.ShouldBindWith(obj, binding.Header)。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
# <翻译结束>


<原文开始>
// ShouldBindUri binds the passed struct pointer using the specified binding engine.
<原文结束>

# <翻译开始>
// ShouldBindUri 使用指定的绑定引擎，将传入的结构体指针进行绑定。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
# <翻译结束>


<原文开始>
// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
<原文结束>

# <翻译开始>
// ShouldBindWith 使用指定的绑定引擎绑定传入的结构体指针。
// 请参阅binding包。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
# <翻译结束>


<原文开始>
// ShouldBindBodyWith is similar with ShouldBindWith, but it stores the request
// body into the context, and reuse when it is called again.
//
// NOTE: This method reads the body before binding. So you should use
// ShouldBindWith for better performance if you need to call only once.
<原文结束>

# <翻译开始>
// ShouldBindBodyWith 与 ShouldBindWith 类似，但它会将请求体存储到上下文中，并在再次调用时重用。
//
// 注意：此方法在绑定前读取请求体。因此，如果你只需要调用一次，为了获得更好的性能，你应该使用 ShouldBindWith。
# <翻译结束>


<原文开始>
// ClientIP implements one best effort algorithm to return the real client IP.
// It calls c.RemoteIP() under the hood, to check if the remote IP is a trusted proxy or not.
// If it is it will then try to parse the headers defined in Engine.RemoteIPHeaders (defaulting to [X-Forwarded-For, X-Real-Ip]).
// If the headers are not syntactically valid OR the remote IP does not correspond to a trusted proxy,
// the remote IP (coming from Request.RemoteAddr) is returned.
<原文结束>

# <翻译开始>
// ClientIP 实现了一种尽力而为的算法，用于返回真实的客户端 IP 地址。
// 在底层，它调用 c.RemoteIP() 来检查远程 IP 是否为可信代理。
// 如果是可信代理，则尝试解析 Engine.RemoteIPHeaders 中定义的头部（默认为 [X-Forwarded-For, X-Real-Ip]）。
// 如果这些头部格式不合法 或者 远程 IP 不对应于一个可信代理，
// 则返回来自 Request.RemoteAddr 的远程 IP 地址。
# <翻译结束>


<原文开始>
// Check if we're running on a trusted platform, continue running backwards if error
<原文结束>

# <翻译开始>
// 检查我们是否在受信任的平台上运行，如果有错误则继续向后执行
# <翻译结束>


<原文开始>
// Developers can define their own header of Trusted Platform or use predefined constants
<原文结束>

# <翻译开始>
// 开发者可以定义自己的可信平台头文件，也可以使用预定义的常量
# <翻译结束>


<原文开始>
// Legacy "AppEngine" flag
<原文结束>

# <翻译开始>
// "AppEngine"老版本标志
# <翻译结束>


<原文开始>
	// It also checks if the remoteIP is a trusted proxy or not.
	// In order to perform this validation, it will see if the IP is contained within at least one of the CIDR blocks
	// defined by Engine.SetTrustedProxies()
<原文结束>

# <翻译开始>
// 它还会检查 remoteIP 是否为可信代理。
// 为了执行此验证，它会查看该 IP 是否至少包含在由 Engine.SetTrustedProxies() 方法定义的一个 CIDR 块中。
# <翻译结束>


<原文开始>
// RemoteIP parses the IP from Request.RemoteAddr, normalizes and returns the IP (without the port).
<原文结束>

# <翻译开始>
// RemoteIP 从 Request.RemoteAddr 解析 IP，进行规范化处理并返回不含端口号的 IP 地址。
# <翻译结束>


<原文开始>
// ContentType returns the Content-Type header of the request.
<原文结束>

# <翻译开始>
// ContentType 返回请求的 Content-Type 头部信息。
# <翻译结束>


<原文开始>
// IsWebsocket returns true if the request headers indicate that a websocket
// handshake is being initiated by the client.
<原文结束>

# <翻译开始>
// IsWebsocket 返回一个布尔值，如果请求头表明客户端正在进行websocket握手，则返回true。
# <翻译结束>


<原文开始>
// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
<原文结束>

# <翻译开始>
// bodyAllowedForStatus 是 http 包中未导出函数 bodyAllowedForStatus 的复制版本。
# <翻译结束>


<原文开始>
// Status sets the HTTP response code.
<原文结束>

# <翻译开始>
// Status 设置 HTTP 响应代码。
# <翻译结束>


<原文开始>
// Header is an intelligent shortcut for c.Writer.Header().Set(key, value).
// It writes a header in the response.
// If value == "", this method removes the header `c.Writer.Header().Del(key)`
<原文结束>

# <翻译开始>
// Header 是一个智能快捷方式，用于 c.Writer.Header().Set(key, value)。
// 它在响应中写入一个头信息。
// 如果 value 等于 "", 则此方法会删除相应头信息：`c.Writer.Header().Del(key)`。
# <翻译结束>


<原文开始>
// GetHeader returns value from request headers.
<原文结束>

# <翻译开始>
// GetHeader 从请求头中返回值。
# <翻译结束>


<原文开始>
// GetRawData returns stream data.
<原文结束>

# <翻译开始>
// GetRawData 返回原始数据流。
# <翻译结束>


<原文开始>
// SetSameSite with cookie
<原文结束>

# <翻译开始>
// SetSameSite 设置 cookie 的同站属性
# <翻译结束>


<原文开始>
// SetCookie adds a Set-Cookie header to the ResponseWriter's headers.
// The provided cookie must have a valid Name. Invalid cookies may be
// silently dropped.
<原文结束>

# <翻译开始>
// SetCookie 向 ResponseWriter 的头信息中添加一个 Set-Cookie 头部。提供的 cookie 必须具有有效的名称。不合法的 cookie 可能会被悄悄丢弃。
# <翻译结束>


<原文开始>
// Cookie returns the named cookie provided in the request or
// ErrNoCookie if not found. And return the named cookie is unescaped.
// If multiple cookies match the given name, only one cookie will
// be returned.
<原文结束>

# <翻译开始>
// Cookie返回请求中提供的指定名称的cookie，如果未找到，则返回ErrNoCookie错误。同时返回的指定名称的cookie是经过解码的。
// 如果多个cookie与给定名称匹配，则只返回一个cookie。
# <翻译结束>


<原文开始>
// Render writes the response headers and calls render.Render to render data.
<原文结束>

# <翻译开始>
// Render方法会写入响应头并调用render.Render来渲染数据。
# <翻译结束>


<原文开始>
// Pushing error to c.Errors
<原文结束>

# <翻译开始>
// 将错误推送到c.Errors
# <翻译结束>


<原文开始>
// HTML renders the HTTP template specified by its file name.
// It also updates the HTTP code and sets the Content-Type as "text/html".
// See http://golang.org/doc/articles/wiki/
<原文结束>

# <翻译开始>
// HTML 根据其文件名渲染 HTTP 模板。
// 同时，它还会更新 HTTP 状态码，并将 Content-Type 设置为 "text/html"。
// 详情参见：http://golang.org/doc/articles/wiki/
# <翻译结束>


<原文开始>
// IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.
// It also sets the Content-Type as "application/json".
// WARNING: we recommend using this only for development purposes since printing pretty JSON is
// more CPU and bandwidth consuming. Use Context.JSON() instead.
<原文结束>

# <翻译开始>
// IndentedJSON 将给定的结构体序列化为美观的 JSON（缩进+换行符）并写入响应体中。
// 同时，它还会将 Content-Type 设置为 "application/json"。
// 警告：我们建议仅在开发目的下使用此方法，因为打印美观的 JSON 会消耗更多的 CPU 和带宽。请改用 Context.JSON()。
# <翻译结束>


<原文开始>
// SecureJSON serializes the given struct as Secure JSON into the response body.
// Default prepends "while(1)," to response body if the given struct is array values.
// It also sets the Content-Type as "application/json".
<原文结束>

# <翻译开始>
// SecureJSON将给定的结构体作为安全的JSON序列化到响应体中。
// 默认情况下，如果给定的结构体是数组值，则会在响应体前缀添加 "while(1),"。
// 同时，它还会将Content-Type设置为"application/json"。
# <翻译结束>


<原文开始>
// JSONP serializes the given struct as JSON into the response body.
// It adds padding to response body to request data from a server residing in a different domain than the client.
// It also sets the Content-Type as "application/javascript".
<原文结束>

# <翻译开始>
// JSONP将给定的结构体以JSON格式序列化到响应体中。
// 它在响应体中添加填充，以便从与客户端不同域的服务器请求数据。
// 同时，它还将Content-Type设置为"application/javascript"。
# <翻译结束>


<原文开始>
// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
<原文结束>

# <翻译开始>
// JSON将给定的结构体以JSON格式序列化到响应体中。
// 同时，它还将Content-Type设置为"application/json"。
# <翻译结束>


<原文开始>
// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
<原文结束>

# <翻译开始>
// AsciiJSON 将给定的结构体按 JSON 格式序列化，并以 ASCII 字符串形式写入响应体中。
// 同时，它还会将 Content-Type 设置为 "application/json"。
# <翻译结束>


<原文开始>
// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
<原文结束>

# <翻译开始>
// PureJSON 将给定的结构体作为 JSON 序列化到响应体中。
// 与 JSON 不同，PureJSON 不会将特殊 HTML 字符替换为它们的 Unicode 实体。
# <翻译结束>


<原文开始>
// XML serializes the given struct as XML into the response body.
// It also sets the Content-Type as "application/xml".
<原文结束>

# <翻译开始>
// XML将给定的结构体作为XML序列化到响应体中。
// 同时，它还会将Content-Type设置为"application/xml"。
# <翻译结束>


<原文开始>
// YAML serializes the given struct as YAML into the response body.
<原文结束>

# <翻译开始>
// YAML 将给定的结构体以 YAML 格式序列化并写入响应体中。
# <翻译结束>


<原文开始>
// TOML serializes the given struct as TOML into the response body.
<原文结束>

# <翻译开始>
// TOML将给定的结构体序列化为TOML格式，并写入响应体中。
# <翻译结束>


<原文开始>
// ProtoBuf serializes the given struct as ProtoBuf into the response body.
<原文结束>

# <翻译开始>
// ProtoBuf将给定的结构体作为ProtoBuf序列化到响应体中。
# <翻译结束>


<原文开始>
// String writes the given string into the response body.
<原文结束>

# <翻译开始>
// String将给定的字符串写入响应体中。
# <翻译结束>


<原文开始>
// Redirect returns an HTTP redirect to the specific location.
<原文结束>

# <翻译开始>
// Redirect 返回一个HTTP重定向到特定位置。
# <翻译结束>


<原文开始>
// Data writes some data into the body stream and updates the HTTP code.
<原文结束>

# <翻译开始>
// Data 将一些数据写入主体流并更新HTTP状态码。
# <翻译结束>


<原文开始>
// DataFromReader writes the specified reader into the body stream and updates the HTTP code.
<原文结束>

# <翻译开始>
// DataFromReader 将指定读取器的内容写入主体流，并更新HTTP状态码。
# <翻译结束>


<原文开始>
// File writes the specified file into the body stream in an efficient way.
<原文结束>

# <翻译开始>
// File 以高效的方式将指定的文件写入正文流中。
# <翻译结束>


<原文开始>
// FileFromFS writes the specified file from http.FileSystem into the body stream in an efficient way.
<原文结束>

# <翻译开始>
// FileFromFS 以高效的方式将指定的文件从 http.FileSystem 写入到 body 流中。
# <翻译结束>


<原文开始>
// FileAttachment writes the specified file into the body stream in an efficient way
// On the client side, the file will typically be downloaded with the given filename
<原文结束>

# <翻译开始>
// FileAttachment 以高效的方式将指定文件写入主体流
// 在客户端，该文件通常会以给定的文件名下载
# <翻译结束>


<原文开始>
// SSEvent writes a Server-Sent Event into the body stream.
<原文结束>

# <翻译开始>
// SSEvent 将一个服务器发送事件写入到主体数据流中。
# <翻译结束>


<原文开始>
// Stream sends a streaming response and returns a boolean
// indicates "Is client disconnected in middle of stream"
<原文结束>

# <翻译开始>
// Stream 发送一个流式响应，并返回一个布尔值
// 表示“在流传输过程中客户端是否已断开连接”
# <翻译结束>


<原文开始>
// Negotiate contains all negotiations data.
<原文结束>

# <翻译开始>
// Negotiate 包含所有协商数据。
# <翻译结束>


<原文开始>
// Negotiate calls different Render according to acceptable Accept format.
<原文结束>

# <翻译开始>
// Negotiate 根据可接受的 Accept 格式调用不同的 Render 方法。
# <翻译结束>


<原文开始>
// NegotiateFormat returns an acceptable Accept format.
<原文结束>

# <翻译开始>
// NegotiateFormat 返回一个可接受的 Accept 格式。
# <翻译结束>


<原文开始>
			// According to RFC 2616 and RFC 2396, non-ASCII characters are not allowed in headers,
			// therefore we can just iterate over the string without casting it into []rune
<原文结束>

# <翻译开始>
// 根据RFC 2616和RFC 2396的规定，非ASCII字符在头部中是不允许出现的，
// 因此我们可以在不将其转换为[]rune的情况下直接遍历该字符串。
# <翻译结束>


<原文开始>
// SetAccepted sets Accept header data.
<原文结束>

# <翻译开始>
// SetAccepted 设置 Accept 头部数据。
# <翻译结束>


<原文开始>
// hasRequestContext returns whether c.Request has Context and fallback.
<原文结束>

# <翻译开始>
// hasRequestContext 返回 c.Request 是否包含 Context 以及回退机制。
# <翻译结束>


<原文开始>
// Deadline returns that there is no deadline (ok==false) when c.Request has no Context.
<原文结束>

# <翻译开始>
// Deadline 返回当 c.Request 没有 Context 时，表示没有截止时间（ok==false）。
# <翻译结束>


<原文开始>
// Done returns nil (chan which will wait forever) when c.Request has no Context.
<原文结束>

# <翻译开始>
// 当c.Request没有Context时，Done返回nil（表示一个将永远等待的通道）。
# <翻译结束>


<原文开始>
// Err returns nil when c.Request has no Context.
<原文结束>

# <翻译开始>
// Err在c.Request没有Context时返回nil。
# <翻译结束>


<原文开始>
// Value returns the value associated with this context for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
<原文结束>

# <翻译开始>
// Value 方法返回与该上下文关联的键key所对应的值，如果该键没有关联任何值，则返回nil。对同一键连续调用Value方法将返回相同的结果。
# <翻译结束>

