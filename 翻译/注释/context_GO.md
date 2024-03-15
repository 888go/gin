
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到
# <翻译结束>


<原文开始>
// Content-Type MIME of the most common data formats.
<原文结束>

# <翻译开始>
// 内容类型MIME最常用的数据格式
# <翻译结束>


<原文开始>
// BodyBytesKey indicates a default body bytes key.
<原文结束>

# <翻译开始>
// BodyBytesKey默认的体字节键
# <翻译结束>


<原文开始>
// ContextKey is the key that a Context returns itself for.
<原文结束>

# <翻译开始>
// ContextKey是Context返回自身的键
# <翻译结束>


<原文开始>
// abortIndex represents a typical value used in abort functions.
<原文结束>

# <翻译开始>
// abortIndex表示中止函数中使用的典型值
# <翻译结束>


<原文开始>
// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
<原文结束>

# <翻译开始>
// 环境是杜松子酒最重要的部分
// 例如，它允许我们在中间件之间传递变量、管理流、验证请求的JSON并呈现JSON响应
# <翻译结束>


<原文开始>
	// This mutex protects Keys map.
<原文结束>

# <翻译开始>
// 这个互斥锁保护键映射
# <翻译结束>


<原文开始>
	// Keys is a key/value pair exclusively for the context of each request.
<原文结束>

# <翻译开始>
// Keys是每个请求上下文专用的键/值对
# <翻译结束>


<原文开始>
	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
<原文结束>

# <翻译开始>
// Errors是附加到使用此上下文的所有处理程序/中间件的错误列表
# <翻译结束>


<原文开始>
	// Accepted defines a list of manually accepted formats for content negotiation.
<原文结束>

# <翻译开始>
// Accepted定义了一个手动接受的格式列表，用于内容协商
# <翻译结束>


<原文开始>
	// queryCache caches the query result from c.Request.URL.Query().
<原文结束>

# <翻译开始>
// queryCache缓存c.Request.URL.Query()的查询结果
# <翻译结束>


<原文开始>
	// formCache caches c.Request.PostForm, which contains the parsed form data from POST, PATCH,
	// or PUT body parameters.
<原文结束>

# <翻译开始>
// c.Request
// PostForm，它包含来自POST、PATCH或PUT主体参数的解析表单数据
# <翻译结束>


<原文开始>
	// SameSite allows a server to define a cookie attribute making it impossible for
	// the browser to send this cookie along with cross-site requests.
<原文结束>

# <翻译开始>
// SameSite允许服务器定义cookie属性，使浏览器无法将此cookie与跨站点请求一起发送
# <翻译结束>


<原文开始>
// Copy returns a copy of the current context that can be safely used outside the request's scope.
// This has to be used when the context has to be passed to a goroutine.
<原文结束>

# <翻译开始>
// Copy返回当前上下文的副本，该副本可在请求作用域之外安全地使用
// 当必须将上下文传递给程序时，必须使用此方法
# <翻译结束>


<原文开始>
// HandlerName returns the main handler's name. For example if the handler is "handleGetUsers()",
// this function will return "main.handleGetUsers".
<原文结束>

# <翻译开始>
// HandlerName返回主处理程序的名称
// 例如，如果处理程序为“handleGetUsers()”，则此函数将返回“main.handleGetUsers”
# <翻译结束>


<原文开始>
// HandlerNames returns a list of all registered handlers for this context in descending order,
// following the semantics of HandlerName()
<原文结束>

# <翻译开始>
// HandlerNames按照HandlerName()的语义，按降序返回此上下文的所有已注册处理程序的列表
# <翻译结束>


<原文开始>
// Handler returns the main handler.
<原文结束>

# <翻译开始>
// Handler返回主处理程序
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
// FullPath返回匹配的路由完整路径
// 对于未找到的路由返回一个空字符串
// router.GET("/user/:id"， func(c *gin.Context) {c. fullpath () == "/user/:id"真正})
# <翻译结束>


<原文开始>
// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in GitHub.
<原文结束>

# <翻译开始>
// Next应该只在中间件内部使用
// 它执行调用处理程序内部链中的挂起处理程序
// 参见GitHub中的示例
# <翻译结束>


<原文开始>
// IsAborted returns true if the current context was aborted.
<原文结束>

# <翻译开始>
// 如果当前上下文被中止，IsAborted返回true
# <翻译结束>


<原文开始>
// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
<原文结束>

# <翻译开始>
// Abort防止调用挂起的处理程序
// 注意，这不会停止当前处理程序
// 假设您有一个授权中间件，用于验证当前请求是否已授权
// 如果授权失败(例如:密码不匹配)，调用Abort以确保不调用此请求的其余处理程序
# <翻译结束>


<原文开始>
// AbortWithStatus calls `Abort()` and writes the headers with the specified status code.
// For example, a failed attempt to authenticate a request could use: context.AbortWithStatus(401).
<原文结束>

# <翻译开始>
// AbortWithStatus调用`Abort()`并写入带有指定状态码的头文件
// 例如，验证请求失败时可以使用:context.AbortWithStatus(401)
# <翻译结束>


<原文开始>
// AbortWithStatusJSON calls `Abort()` and then `JSON` internally.
// This method stops the chain, writes the status code and return a JSON body.
// It also sets the Content-Type as "application/json".
<原文结束>

# <翻译开始>
// AbortWithStatusJSON调用' Abort() '，然后在内部调用' JSON '
// 此方法停止链，编写状态代码并返回JSON主体
// 它还将Content-Type设置为“application/json”
# <翻译结束>


<原文开始>
// AbortWithError calls `AbortWithStatus()` and `Error()` internally.
// This method stops the chain, writes the status code and pushes the specified error to `c.Errors`.
// See Context.Error() for more details.
<原文结束>

# <翻译开始>
// AbortWithError在内部调用`AbortWithStatus()`和`Error()`
// 此方法停止链，写入状态码并将指定的错误推入' c.Errors '
// 有关详细信息，请参阅Context.Error()
# <翻译结束>


<原文开始>
// Error attaches an error to the current context. The error is pushed to a list of errors.
// It's a good idea to call Error for each error that occurred during the resolution of a request.
// A middleware can be used to collect all the errors and push them to a database together,
// print a log, or append it in the HTTP response.
// Error will panic if err is nil.
<原文结束>

# <翻译开始>
// Error将错误附加到当前上下文
// 错误被推入错误列表
// 对解析请求期间发生的每个错误调用Error是一个好主意
// 中间件可用于收集所有错误并将它们一起推送到数据库、打印日志或将其附加到HTTP响应中
// 如果err为nil, Error将出现Panic
# <翻译结束>


<原文开始>
// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
<原文结束>

# <翻译开始>
// Set用于存储专门用于此上下文的新键/值对
// 如果以前没有使用c.Keys，它也会延迟初始化它
# <翻译结束>


<原文开始>
// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
<原文结束>

# <翻译开始>
// Get返回给定键的值，即:(value, true)
// 如果值不存在，则返回(nil, false)
# <翻译结束>


<原文开始>
// MustGet returns the value for the given key if it exists, otherwise it panics.
<原文结束>

# <翻译开始>
// 如果给定的键存在，则必须返回该键的值，否则会产生Panic
# <翻译结束>


<原文开始>
// GetString returns the value associated with the key as a string.
<原文结束>

# <翻译开始>
// GetString以字符串的形式返回与键相关的值
# <翻译结束>


<原文开始>
// GetBool returns the value associated with the key as a boolean.
<原文结束>

# <翻译开始>
// GetBool返回与键相关联的值作为布尔值
# <翻译结束>


<原文开始>
// GetInt returns the value associated with the key as an integer.
<原文结束>

# <翻译开始>
// GetInt以整数形式返回与键相关的值
# <翻译结束>


<原文开始>
// GetInt64 returns the value associated with the key as an integer.
<原文结束>

# <翻译开始>
// GetInt64以整数形式返回与键关联的值
# <翻译结束>


<原文开始>
// GetUint returns the value associated with the key as an unsigned integer.
<原文结束>

# <翻译开始>
// GetUint以无符号整数的形式返回与键相关的值
# <翻译结束>


<原文开始>
// GetUint64 returns the value associated with the key as an unsigned integer.
<原文结束>

# <翻译开始>
// GetUint64以无符号整数的形式返回与键相关的值
# <翻译结束>


<原文开始>
// GetFloat64 returns the value associated with the key as a float64.
<原文结束>

# <翻译开始>
// GetFloat64返回与该键相关的值作为float64
# <翻译结束>


<原文开始>
// GetTime returns the value associated with the key as time.
<原文结束>

# <翻译开始>
// GetTime返回与键相关的值作为time
# <翻译结束>


<原文开始>
// GetDuration returns the value associated with the key as a duration.
<原文结束>

# <翻译开始>
// GetDuration以持续时间的形式返回与键相关的值
# <翻译结束>


<原文开始>
// GetStringSlice returns the value associated with the key as a slice of strings.
<原文结束>

# <翻译开始>
// GetStringSlice以字符串切片的形式返回与键相关的值
# <翻译结束>


<原文开始>
// GetStringMap returns the value associated with the key as a map of interfaces.
<原文结束>

# <翻译开始>
// GetStringMap以接口映射的形式返回与键相关的值
# <翻译结束>


<原文开始>
// GetStringMapString returns the value associated with the key as a map of strings.
<原文结束>

# <翻译开始>
// GetStringMapString以字符串映射的形式返回与键相关的值
# <翻译结束>


<原文开始>
// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
<原文结束>

# <翻译开始>
// GetStringMapStringSlice返回与键相关的值，作为到字符串切片的映射
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
// 参数返回URL参数的值
// 它是c. param . byname (key) router.GET("/user/:id"， func(c *gin.Context) {GET请求/user/john id:= c. param ("id") id == "/john"一个GET请求到/user/john/ id:= c.参数("id") id == "/john/"}）
# <翻译结束>


<原文开始>
// AddParam adds param to context and
// replaces path param key with given value for e2e testing purposes
// Example Route: "/user/:id"
// AddParam("id", 1)
// Result: "/user/1"
<原文结束>

# <翻译开始>
// AddParam将参数添加到上下文中，并用给定的值替换路径参数键，用于端到端测试
// 示例Route: "/user/:id"AddParam("id"， 1) Result: "/user/1"
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
// Query如果存在则返回键控url查询值，否则返回空字符串' ("") '
// 这是快捷方式的' c.Request.URL.Query().Get(key) ' GET /path?id=1234&name= manual &value= c.Query("id") == "1234"c.Query("name") == " manual "c.Query("value") == "c.查询("wtf") == ";
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
// 如果存在，则返回键控url查询值，否则返回指定的defaultValue字符串
// 更多信息请参见:Query()和GetQuery()
// GET / ?name=姓名&lastname= c.DefaultQuery("name"， "unknown") ==姓名"c.DefaultQuery("id"， "none") == "none"c.DefaultQuery("lastname"， "none") == "
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
// GetQuery类似于Query()，如果存在' (value, true) '(即使值是空字符串)，它返回键控url查询值，否则它返回' (""， false) '
// 它是' c.Request.URL.Query().Get(key) ' GET /?name=Manu&lastname= ("Manu"， true) == c.GetQuery("name") (""， false) == c.GetQuery("id") (""， true) == c.GetQuery("lastname")
# <翻译结束>


<原文开始>
// QueryArray returns a slice of strings for a given query key.
// The length of the slice depends on the number of params with the given key.
<原文结束>

# <翻译开始>
// QueryArray返回给定查询键的字符串切片
// 切片的长度取决于具有给定键的参数的数量
# <翻译结束>


<原文开始>
// GetQueryArray returns a slice of strings for a given query key, plus
// a boolean value whether at least one value exists for the given key.
<原文结束>

# <翻译开始>
// GetQueryArray返回给定查询键的字符串切片，以及一个布尔值，用于判断给定键是否至少存在一个值
# <翻译结束>


<原文开始>
// QueryMap returns a map for a given query key.
<原文结束>

# <翻译开始>
// QueryMap返回给定查询键的映射
# <翻译结束>


<原文开始>
// GetQueryMap returns a map for a given query key, plus a boolean value
// whether at least one value exists for the given key.
<原文结束>

# <翻译开始>
// GetQueryMap返回给定查询键的映射，加上一个布尔值，用于判断给定键是否至少存在一个值
# <翻译结束>


<原文开始>
// PostForm returns the specified key from a POST urlencoded form or multipart form
// when it exists, otherwise it returns an empty string `("")`.
<原文结束>

# <翻译开始>
// PostForm从存在的POST url编码表单或多部分表单返回指定的键，否则返回空字符串' ("") '
# <翻译结束>


<原文开始>
// DefaultPostForm returns the specified key from a POST urlencoded form or multipart form
// when it exists, otherwise it returns the specified defaultValue string.
// See: PostForm() and GetPostForm() for further information.
<原文结束>

# <翻译开始>
// DefaultPostForm从存在的POST url编码表单或多部分表单返回指定的键，否则返回指定的defaultValue字符串
// 参见:PostForm()和GetPostForm()了解更多信息
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
// GetPostForm类似于PostForm(key)
// 如果存在' (value, true) '(即使值是空字符串)，则从POST url编码形式或多部分形式返回指定的键，否则返回(""， false)
// 例如，在PATCH请求更新用户的电子邮件时:email=mail@example.com——>("mail@example.com"， true):= GetPostForm("email")设置email为"mail@example.com"电子邮件 =                  --& gt;(""， true):= GetPostForm("email")设置email为"——比;(""， false):= GetPostForm(&q
# <翻译结束>


<原文开始>
// PostFormArray returns a slice of strings for a given form key.
// The length of the slice depends on the number of params with the given key.
<原文结束>

# <翻译开始>
// PostFormArray返回给定表单键的字符串切片
// 切片的长度取决于具有给定键的参数的数量
# <翻译结束>


<原文开始>
// GetPostFormArray returns a slice of strings for a given form key, plus
// a boolean value whether at least one value exists for the given key.
<原文结束>

# <翻译开始>
// GetPostFormArray返回给定表单键的字符串切片，以及是否至少存在一个给定键的布尔值
# <翻译结束>


<原文开始>
// PostFormMap returns a map for a given form key.
<原文结束>

# <翻译开始>
// PostFormMap返回给定表单键的映射
# <翻译结束>


<原文开始>
// GetPostFormMap returns a map for a given form key, plus a boolean value
// whether at least one value exists for the given key.
<原文结束>

# <翻译开始>
// GetPostFormMap返回给定表单键的映射，以及一个布尔值，用于判断给定键是否至少存在一个值
# <翻译结束>


<原文开始>
// get is an internal method and returns a map which satisfies conditions.
<原文结束>

# <翻译开始>
// Get是一个内部方法，返回一个满足条件的映射
# <翻译结束>


<原文开始>
// FormFile returns the first file for the provided form key.
<原文结束>

# <翻译开始>
// FormFile返回所提供表单键的第一个文件
# <翻译结束>


<原文开始>
// MultipartForm is the parsed multipart form, including file uploads.
<原文结束>

# <翻译开始>
// MultipartForm是解析后的多部分表单，包括文件上传
# <翻译结束>


<原文开始>
// SaveUploadedFile uploads the form file to specific dst.
<原文结束>

# <翻译开始>
// SaveUploadedFile上传表单文件到指定的dst
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
// Bind检查方法和内容类型以自动选择绑定引擎，具体取决于“内容类型”
// 头文件使用了不同的绑定，例如:"application/json"——比;JSON绑定"application/xml"——比;如果Content-Type == "application/ JSON "使用JSON或XML作为JSON输入
// 它将json有效负载解码为指定为指针的结构
// 它会写一个400的错误，并设置Content-Type header "text/plain"在响应中，如果输入无效
# <翻译结束>


<原文开始>
// BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
<原文结束>

# <翻译开始>
// BindJSON是c.MustBindWith(obj, binding.JSON)的快捷方式
# <翻译结束>


<原文开始>
// BindXML is a shortcut for c.MustBindWith(obj, binding.BindXML).
<原文结束>

# <翻译开始>
// BindXML是c.MustBindWith(obj, binding.BindXML)的快捷方式
# <翻译结束>


<原文开始>
// BindQuery is a shortcut for c.MustBindWith(obj, binding.Query).
<原文结束>

# <翻译开始>
// BindQuery是c.MustBindWith(obj, binding.Query)的快捷方式
# <翻译结束>


<原文开始>
// BindYAML is a shortcut for c.MustBindWith(obj, binding.YAML).
<原文结束>

# <翻译开始>
// BindYAML是c.MustBindWith(obj, binding.YAML)的快捷方式
# <翻译结束>


<原文开始>
// BindTOML is a shortcut for c.MustBindWith(obj, binding.TOML).
<原文结束>

# <翻译开始>
// BindTOML是c.MustBindWith(obj, binding.TOML)的快捷方式
# <翻译结束>


<原文开始>
// BindHeader is a shortcut for c.MustBindWith(obj, binding.Header).
<原文结束>

# <翻译开始>
// BindHeader是c.MustBindWith(obj, binding.Header)的快捷方式
# <翻译结束>


<原文开始>
// BindUri binds the passed struct pointer using binding.Uri.
// It will abort the request with HTTP 400 if any error occurs.
<原文结束>

# <翻译开始>
// BindUri使用binding.Uri绑定传递的结构指针
// 如果发生任何错误，它将使用HTTP 400中止请求
# <翻译结束>


<原文开始>
//nolint: errcheck
<原文结束>

# <翻译开始>
// nolint: errcheck
// 翻译：// 不进行errcheck检查
# <翻译结束>


<原文开始>
// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
<原文结束>

# <翻译开始>
// MustBindWith使用指定的绑定引擎绑定传递的结构指针
// 如果发生任何错误，它将使用HTTP 400中止请求
// 参见绑定包
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
// shoulbind检查方法和内容类型，根据“内容类型”自动选择绑定引擎
// 头文件使用了不同的绑定，例如:"application/json"——比;JSON绑定"application/xml"——比;如果Content-Type == "application/ JSON "使用JSON或XML作为JSON输入
// 它将json有效负载解码为指定为指针的结构
// 与c.Bind()类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
# <翻译结束>


<原文开始>
// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
<原文结束>

# <翻译开始>
// ShouldBindJSON是c.ShouldBindWith(obj, binding.JSON)的快捷方式
# <翻译结束>


<原文开始>
// ShouldBindXML is a shortcut for c.ShouldBindWith(obj, binding.XML).
<原文结束>

# <翻译开始>
// ShouldBindXML是c.ShouldBindWith(obj, binding.XML)的快捷方式
# <翻译结束>


<原文开始>
// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
<原文结束>

# <翻译开始>
// ShouldBindQuery是c.ShouldBindWith(obj, binding.Query)的快捷方式
# <翻译结束>


<原文开始>
// ShouldBindYAML is a shortcut for c.ShouldBindWith(obj, binding.YAML).
<原文结束>

# <翻译开始>
// ShouldBindYAML是c.ShouldBindWith(obj, binding.YAML)的快捷方式
# <翻译结束>


<原文开始>
// ShouldBindTOML is a shortcut for c.ShouldBindWith(obj, binding.TOML).
<原文结束>

# <翻译开始>
// ShouldBindTOML是c.ShouldBindWith(obj, binding.TOML)的快捷方式
# <翻译结束>


<原文开始>
// ShouldBindHeader is a shortcut for c.ShouldBindWith(obj, binding.Header).
<原文结束>

# <翻译开始>
// ShouldBindHeader是c.ShouldBindWith(obj, binding.Header)的快捷方式
# <翻译结束>


<原文开始>
// ShouldBindUri binds the passed struct pointer using the specified binding engine.
<原文结束>

# <翻译开始>
// ShouldBindUri使用指定的绑定引擎绑定传递的结构指针
# <翻译结束>


<原文开始>
// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
<原文结束>

# <翻译开始>
// ShouldBindWith使用指定的绑定引擎绑定传递的结构指针
// 参见绑定包
# <翻译结束>


<原文开始>
// ShouldBindBodyWith is similar with ShouldBindWith, but it stores the request
// body into the context, and reuse when it is called again.
//
// NOTE: This method reads the body before binding. So you should use
// ShouldBindWith for better performance if you need to call only once.
<原文结束>

# <翻译开始>
// ShouldBindBodyWith与ShouldBindWith类似，但它将请求体存储到上下文中，并在再次调用时重用
// 注意:此方法在绑定前读取主体
// 因此，如果只需要调用一次，应该使用ShouldBindWith以获得更好的性能
# <翻译结束>


<原文开始>
// ClientIP implements one best effort algorithm to return the real client IP.
// It calls c.RemoteIP() under the hood, to check if the remote IP is a trusted proxy or not.
// If it is it will then try to parse the headers defined in Engine.RemoteIPHeaders (defaulting to [X-Forwarded-For, X-Real-Ip]).
// If the headers are not syntactically valid OR the remote IP does not correspond to a trusted proxy,
// the remote IP (coming from Request.RemoteAddr) is returned.
<原文结束>

# <翻译开始>
// ClientIP实现了一个最佳努力算法来返回真实的客户端IP
// 它在底层调用c.RemoteIP()来检查远程IP是否是可信代理
// 如果是，它将尝试解析Engine中定义的标头
// RemoteIPHeaders(缺省为[X-Forwarded-For, X-Real-Ip])
// 如果报头在语法上无效或远程IP不对应于可信代理，则返回远程IP(来自Request.RemoteAddr)
# <翻译结束>


<原文开始>
	// Check if we're running on a trusted platform, continue running backwards if error
<原文结束>

# <翻译开始>
// 检查我们是否运行在一个可信的平台上，如果错误继续运行
# <翻译结束>


<原文开始>
		// Developers can define their own header of Trusted Platform or use predefined constants
<原文结束>

# <翻译开始>
// 开发人员可以定义自己的可信平台头或使用预定义的常量
# <翻译结束>


<原文开始>
	// Legacy "AppEngine" flag
<原文结束>

# <翻译开始>
// 遗留“AppEngine"国旗
# <翻译结束>


<原文开始>
	// It also checks if the remoteIP is a trusted proxy or not.
	// In order to perform this validation, it will see if the IP is contained within at least one of the CIDR blocks
	// defined by Engine.SetTrustedProxies()
<原文结束>

# <翻译开始>
// 它还检查remoteIP是否是受信任的代理
// 为了执行此验证，它将查看IP是否包含在engine定义的至少一个CIDR块中
# <翻译结束>


<原文开始>
// RemoteIP parses the IP from Request.RemoteAddr, normalizes and returns the IP (without the port).
<原文结束>

# <翻译开始>
// RemoteIP解析来自Request的IP
// RemoteAddr，规范化并返回IP(不带端口)
# <翻译结束>


<原文开始>
// ContentType returns the Content-Type header of the request.
<原文结束>

# <翻译开始>
// ContentType返回请求的Content-Type报头
# <翻译结束>


<原文开始>
// IsWebsocket returns true if the request headers indicate that a websocket
// handshake is being initiated by the client.
<原文结束>

# <翻译开始>
// 如果请求头表明客户端正在发起websocket握手，IsWebsocket返回true
# <翻译结束>


<原文开始>
// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
<原文结束>

# <翻译开始>
// bodyAllowedForStatus是http的一个副本
// bodyAllowedForStatus非导出函数
# <翻译结束>


<原文开始>
// Status sets the HTTP response code.
<原文结束>

# <翻译开始>
// 状态设置HTTP响应码
# <翻译结束>


<原文开始>
// Header is an intelligent shortcut for c.Writer.Header().Set(key, value).
// It writes a header in the response.
// If value == "", this method removes the header `c.Writer.Header().Del(key)`
<原文结束>

# <翻译开始>
// Header是c.Writer.Header()的智能快捷方式
// 集(关键字,值)
// 它在响应中写入一个标头
// 如果value == ""，此方法将删除头' c.Writer.Header().Del(key) '
# <翻译结束>


<原文开始>
// GetHeader returns value from request headers.
<原文结束>

# <翻译开始>
// GetHeader从请求头返回值
# <翻译结束>


<原文开始>
// GetRawData returns stream data.
<原文结束>

# <翻译开始>
// GetRawData返回流数据
# <翻译结束>


<原文开始>
// SetSameSite with cookie
<原文结束>

# <翻译开始>
// SetSameSite 用于设置 cookie 的 SameSite 属性
# <翻译结束>


<原文开始>
// SetCookie adds a Set-Cookie header to the ResponseWriter's headers.
// The provided cookie must have a valid Name. Invalid cookies may be
// silently dropped.
<原文结束>

# <翻译开始>
// SetCookie在ResponseWriter的报头中添加一个Set-Cookie报头
// 提供的cookie必须有一个有效的Name
// 无效的cookie可能会被静默删除
# <翻译结束>


<原文开始>
// Cookie returns the named cookie provided in the request or
// ErrNoCookie if not found. And return the named cookie is unescaped.
// If multiple cookies match the given name, only one cookie will
// be returned.
<原文结束>

# <翻译开始>
// Cookie返回请求中提供的命名Cookie，如果没有找到，则返回ErrNoCookie
// 并返回未转义的命名cookie
// 如果多个cookie与给定的名称匹配，则只返回一个cookie
# <翻译结束>


<原文开始>
// Render writes the response headers and calls render.Render to render data.
<原文结束>

# <翻译开始>
// Render写入响应头并调用Render
// 渲染到渲染数据
# <翻译结束>


<原文开始>
		// Pushing error to c.Errors
<原文结束>

# <翻译开始>
// 将error推入c.Errors
# <翻译结束>


<原文开始>
// HTML renders the HTTP template specified by its file name.
// It also updates the HTTP code and sets the Content-Type as "text/html".
// See http://golang.org/doc/articles/wiki/
<原文结束>

# <翻译开始>
// HTML呈现由其文件名指定的HTTP模板
// 它还更新HTTP代码并将Content-Type设置为"text/html"
// 参见http://golang.org/doc/articles/wiki/
# <翻译结束>


<原文开始>
// IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.
// It also sets the Content-Type as "application/json".
// WARNING: we recommend using this only for development purposes since printing pretty JSON is
// more CPU and bandwidth consuming. Use Context.JSON() instead.
<原文结束>

# <翻译开始>
// indetedjson将给定的结构序列化为漂亮的JSON(缩进+ endlines)到响应体中
// 它还将Content-Type设置为“application/json”
// 警告:我们建议仅用于开发目的，因为打印漂亮的JSON会消耗更多的CPU和带宽
// 使用Context.JSON()代替
# <翻译结束>


<原文开始>
// SecureJSON serializes the given struct as Secure JSON into the response body.
// Default prepends "while(1)," to response body if the given struct is array values.
// It also sets the Content-Type as "application/json".
<原文结束>

# <翻译开始>
// SecureJSON将给定的结构作为安全JSON序列化到响应体中
// Default前面加上"while(1)，"如果给定的结构体是数组值，则返回响应体
// 它还将Content-Type设置为“application/json”
# <翻译结束>


<原文开始>
// JSONP serializes the given struct as JSON into the response body.
// It adds padding to response body to request data from a server residing in a different domain than the client.
// It also sets the Content-Type as "application/javascript".
<原文结束>

# <翻译开始>
// JSONP将给定的结构作为JSON序列化到响应体中
// 它向响应体添加填充，以便从位于与客户端不同域的服务器请求数据
// 它还将Content-Type设置为"application/javascript"
# <翻译结束>


<原文开始>
// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
<原文结束>

# <翻译开始>
// JSON将给定的结构作为JSON序列化到响应体中
// 它还将Content-Type设置为“application/json”
# <翻译结束>


<原文开始>
// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
<原文结束>

# <翻译开始>
// AsciiJSON将给定的结构作为JSON序列化到响应体中，并使用unicode到ASCII字符串
// 它还将Content-Type设置为“application/json”
# <翻译结束>


<原文开始>
// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
<原文结束>

# <翻译开始>
// PureJSON将给定的结构作为JSON序列化到响应体中
// 与JSON不同的是，PureJSON不会用它们的unicode实体替换特殊的html字符
# <翻译结束>


<原文开始>
// XML serializes the given struct as XML into the response body.
// It also sets the Content-Type as "application/xml".
<原文结束>

# <翻译开始>
// XML将给定的结构作为XML序列化到响应体中
// 它还将Content-Type设置为“application/xml”
# <翻译结束>


<原文开始>
// YAML serializes the given struct as YAML into the response body.
<原文结束>

# <翻译开始>
// YAML将给定的结构作为YAML序列化到响应体中
# <翻译结束>


<原文开始>
// TOML serializes the given struct as TOML into the response body.
<原文结束>

# <翻译开始>
// TOML将给定的结构作为TOML序列化到响应体中
# <翻译结束>


<原文开始>
// ProtoBuf serializes the given struct as ProtoBuf into the response body.
<原文结束>

# <翻译开始>
// ProtoBuf将给定的结构体作为ProtoBuf序列化到响应体中
# <翻译结束>


<原文开始>
// String writes the given string into the response body.
<原文结束>

# <翻译开始>
// String将给定的字符串写入响应体
# <翻译结束>


<原文开始>
// Redirect returns an HTTP redirect to the specific location.
<原文结束>

# <翻译开始>
// Redirect返回到特定位置的HTTP重定向
# <翻译结束>


<原文开始>
// Data writes some data into the body stream and updates the HTTP code.
<原文结束>

# <翻译开始>
// Data将一些数据写入主体流并更新HTTP代码
# <翻译结束>


<原文开始>
// DataFromReader writes the specified reader into the body stream and updates the HTTP code.
<原文结束>

# <翻译开始>
// DataFromReader将指定的阅读器写入正文流并更新HTTP代码
# <翻译结束>


<原文开始>
// File writes the specified file into the body stream in an efficient way.
<原文结束>

# <翻译开始>
// File以一种有效的方式将指定的文件写入体流
# <翻译结束>


<原文开始>
// FileFromFS writes the specified file from http.FileSystem into the body stream in an efficient way.
<原文结束>

# <翻译开始>
// FileFromFS从http写入指定的文件
// 文件系统以一种有效的方式进入主体流
# <翻译结束>


<原文开始>
// FileAttachment writes the specified file into the body stream in an efficient way
// On the client side, the file will typically be downloaded with the given filename
<原文结束>

# <翻译开始>
// FileAttachment以一种有效的方式将指定的文件写入正文流
// 在客户端，文件通常会以给定的文件名下载
# <翻译结束>


<原文开始>
// SSEvent writes a Server-Sent Event into the body stream.
<原文结束>

# <翻译开始>
// SSEvent将服务器发送的事件写入主体流
# <翻译结束>


<原文开始>
// Stream sends a streaming response and returns a boolean
// indicates "Is client disconnected in middle of stream"
<原文结束>

# <翻译开始>
// 流发送一个流响应并返回一个布尔值，表示“客户端在流的中间断开了连接”;
# <翻译结束>


<原文开始>
// Negotiate contains all negotiations data.
<原文结束>

# <翻译开始>
// Negotiate包含所有谈判数据
# <翻译结束>


<原文开始>
// Negotiate calls different Render according to acceptable Accept format.
<原文结束>

# <翻译开始>
// 根据可接受的Accept格式协商调用不同的Render
# <翻译结束>


<原文开始>
// NegotiateFormat returns an acceptable Accept format.
<原文结束>

# <翻译开始>
// NegotiateFormat返回一个可接受的Accept格式
# <翻译结束>


<原文开始>
			// According to RFC 2616 and RFC 2396, non-ASCII characters are not allowed in headers,
			// therefore we can just iterate over the string without casting it into []rune
<原文结束>

# <翻译开始>
// 根据RFC 2616和RFC 2396，头中不允许使用非ascii字符，因此我们可以迭代字符串，而不将其转换为[]rune
# <翻译结束>


<原文开始>
// SetAccepted sets Accept header data.
<原文结束>

# <翻译开始>
// SetAccepted设置接受报头数据
# <翻译结束>


<原文开始>
// hasRequestContext returns whether c.Request has Context and fallback.
<原文结束>

# <翻译开始>
// hasRequestContext返回c.Request是否有Context和fallback
# <翻译结束>


<原文开始>
// Deadline returns that there is no deadline (ok==false) when c.Request has no Context.
<原文结束>

# <翻译开始>
// 当c.Request没有Context时，Deadline返回没有Deadline (ok==false)
# <翻译结束>


<原文开始>
// Done returns nil (chan which will wait forever) when c.Request has no Context.
<原文结束>

# <翻译开始>
// 当c.Request没有上下文时，Done返回nil (chan将永远等待)
# <翻译结束>


<原文开始>
// Err returns nil when c.Request has no Context.
<原文结束>

# <翻译开始>
// 当c.Request没有Context时，Err返回nil
# <翻译结束>


<原文开始>
// Value returns the value associated with this context for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
<原文结束>

# <翻译开始>
// Value为key返回与此上下文关联的值，如果没有值与key关联，则返回nil
// 连续调用具有相同键的Value返回相同的结果
# <翻译结束>

