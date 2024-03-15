// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"errors"
	"io"
	"log"
	"math"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	
	"github.com/gin-contrib/sse"
	"github.com/888go/gin/binding"
	"github.com/888go/gin/render"
)

// 内容类型MIME最常用的数据格式
const (
	MIMEJSON              = binding.MIMEJSON
	MIMEHTML              = binding.MIMEHTML
	MIMEXML               = binding.MIMEXML
	MIMEXML2              = binding.MIMEXML2
	MIMEPlain             = binding.MIMEPlain
	MIMEPOSTForm          = binding.MIMEPOSTForm
	MIMEMultipartPOSTForm = binding.MIMEMultipartPOSTForm
	MIMEYAML              = binding.MIMEYAML
	MIMETOML              = binding.MIMETOML
)

// BodyBytesKey默认的体字节键
const BodyBytesKey = "_gin-gonic/gin/bodybyteskey"

// ContextKey是Context返回自身的键
const ContextKey = "_gin-gonic/gin/contextkey"

// abortIndex表示中止函数中使用的典型值
const abortIndex int8 = math.MaxInt8 >> 1

// 环境是杜松子酒最重要的部分
// 例如，它允许我们在中间件之间传递变量、管理流、验证请求的JSON并呈现JSON响应
type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine       *Engine
	params       *Params
	skippedNodes *[]skippedNode

// 这个互斥锁保护键映射
	mu sync.RWMutex

// Keys是每个请求上下文专用的键/值对
	Keys map[string]any

// Errors是附加到使用此上下文的所有处理程序/中间件的错误列表
	Errors errorMsgs

// Accepted定义了一个手动接受的格式列表，用于内容协商
	Accepted []string

// queryCache缓存c.Request.URL.Query()的查询结果
	queryCache url.Values

// c.Request
// PostForm，它包含来自POST、PATCH或PUT主体参数的解析表单数据
	formCache url.Values

// SameSite允许服务器定义cookie属性，使浏览器无法将此cookie与跨站点请求一起发送
	sameSite http.SameSite
}

/************************************/
/********** CONTEXT CREATION ********/
/************************************/

func (c *Context) reset() {
	c.Writer = &c.writermem
	c.Params = c.Params[:0]
	c.handlers = nil
	c.index = -1

	c.fullPath = ""
	c.Keys = nil
	c.Errors = c.Errors[:0]
	c.Accepted = nil
	c.queryCache = nil
	c.formCache = nil
	c.sameSite = 0
	*c.params = (*c.params)[:0]
	*c.skippedNodes = (*c.skippedNodes)[:0]
}

// Copy返回当前上下文的副本，该副本可在请求作用域之外安全地使用
// 当必须将上下文传递给程序时，必须使用此方法
func (c *Context) Copy() *Context {
	cp := Context{
		writermem: c.writermem,
		Request:   c.Request,
		Params:    c.Params,
		engine:    c.engine,
	}
	cp.writermem.ResponseWriter = nil
	cp.Writer = &cp.writermem
	cp.index = abortIndex
	cp.handlers = nil
	cp.Keys = map[string]any{}
	for k, v := range c.Keys {
		cp.Keys[k] = v
	}
	paramCopy := make([]Param, len(cp.Params))
	copy(paramCopy, cp.Params)
	cp.Params = paramCopy
	return &cp
}

// HandlerName返回主处理程序的名称
// 例如，如果处理程序为“handleGetUsers()”，则此函数将返回“main.handleGetUsers”
func (c *Context) HandlerName() string {
	return nameOfFunction(c.handlers.Last())
}

// HandlerNames按照HandlerName()的语义，按降序返回此上下文的所有已注册处理程序的列表
func (c *Context) HandlerNames() []string {
	hn := make([]string, 0, len(c.handlers))
	for _, val := range c.handlers {
		hn = append(hn, nameOfFunction(val))
	}
	return hn
}

// Handler返回主处理程序
func (c *Context) Handler() HandlerFunc {
	return c.handlers.Last()
}

// FullPath返回匹配的路由完整路径
// 对于未找到的路由返回一个空字符串
// router.GET("/user/:id"， func(c *gin.Context) {c. fullpath () == "/user/:id"真正})
func (c *Context) FullPath() string {
	return c.fullPath
}

/************************************/
/*********** FLOW CONTROL ***********/
/************************************/

// Next应该只在中间件内部使用
// 它执行调用处理程序内部链中的挂起处理程序
// 参见GitHub中的示例
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// 如果当前上下文被中止，IsAborted返回true
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

// Abort防止调用挂起的处理程序
// 注意，这不会停止当前处理程序
// 假设您有一个授权中间件，用于验证当前请求是否已授权
// 如果授权失败(例如:密码不匹配)，调用Abort以确保不调用此请求的其余处理程序
func (c *Context) Abort() {
	c.index = abortIndex
}

// AbortWithStatus调用`Abort()`并写入带有指定状态码的头文件
// 例如，验证请求失败时可以使用:context.AbortWithStatus(401)
func (c *Context) AbortWithStatus(code int) {
	c.Status(code)
	c.Writer.WriteHeaderNow()
	c.Abort()
}

// AbortWithStatusJSON调用' Abort() '，然后在内部调用' JSON '
// 此方法停止链，编写状态代码并返回JSON主体
// 它还将Content-Type设置为“application/json”
func (c *Context) AbortWithStatusJSON(code int, jsonObj any) {
	c.Abort()
	c.JSON(code, jsonObj)
}

// AbortWithError在内部调用`AbortWithStatus()`和`Error()`
// 此方法停止链，写入状态码并将指定的错误推入' c.Errors '
// 有关详细信息，请参阅Context.Error()
func (c *Context) AbortWithError(code int, err error) *Error {
	c.AbortWithStatus(code)
	return c.Error(err)
}

/************************************/
/********* ERROR MANAGEMENT *********/
/************************************/

// Error将错误附加到当前上下文
// 错误被推入错误列表
// 对解析请求期间发生的每个错误调用Error是一个好主意
// 中间件可用于收集所有错误并将它们一起推送到数据库、打印日志或将其附加到HTTP响应中
// 如果err为nil, Error将出现Panic
func (c *Context) Error(err error) *Error {
	if err == nil {
		panic("err is nil")
	}

	var parsedError *Error
	ok := errors.As(err, &parsedError)
	if !ok {
		parsedError = &Error{
			Err:  err,
			Type: ErrorTypePrivate,
		}
	}

	c.Errors = append(c.Errors, parsedError)
	return parsedError
}

/************************************/
/******** METADATA MANAGEMENT********/
/************************************/

// Set用于存储专门用于此上下文的新键/值对
// 如果以前没有使用c.Keys，它也会延迟初始化它
func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
}

// Get返回给定键的值，即:(value, true)
// 如果值不存在，则返回(nil, false)
func (c *Context) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists = c.Keys[key]
	return
}

// 如果给定的键存在，则必须返回该键的值，否则会产生Panic
func (c *Context) MustGet(key string) any {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

// GetString以字符串的形式返回与键相关的值
func (c *Context) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetBool返回与键相关联的值作为布尔值
func (c *Context) GetBool(key string) (b bool) {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

// GetInt以整数形式返回与键相关的值
func (c *Context) GetInt(key string) (i int) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

// GetInt64以整数形式返回与键关联的值
func (c *Context) GetInt64(key string) (i64 int64) {
	if val, ok := c.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

// GetUint以无符号整数的形式返回与键相关的值
func (c *Context) GetUint(key string) (ui uint) {
	if val, ok := c.Get(key); ok && val != nil {
		ui, _ = val.(uint)
	}
	return
}

// GetUint64以无符号整数的形式返回与键相关的值
func (c *Context) GetUint64(key string) (ui64 uint64) {
	if val, ok := c.Get(key); ok && val != nil {
		ui64, _ = val.(uint64)
	}
	return
}

// GetFloat64返回与该键相关的值作为float64
func (c *Context) GetFloat64(key string) (f64 float64) {
	if val, ok := c.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

// GetTime返回与键相关的值作为time
func (c *Context) GetTime(key string) (t time.Time) {
	if val, ok := c.Get(key); ok && val != nil {
		t, _ = val.(time.Time)
	}
	return
}

// GetDuration以持续时间的形式返回与键相关的值
func (c *Context) GetDuration(key string) (d time.Duration) {
	if val, ok := c.Get(key); ok && val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

// GetStringSlice以字符串切片的形式返回与键相关的值
func (c *Context) GetStringSlice(key string) (ss []string) {
	if val, ok := c.Get(key); ok && val != nil {
		ss, _ = val.([]string)
	}
	return
}

// GetStringMap以接口映射的形式返回与键相关的值
func (c *Context) GetStringMap(key string) (sm map[string]any) {
	if val, ok := c.Get(key); ok && val != nil {
		sm, _ = val.(map[string]any)
	}
	return
}

// GetStringMapString以字符串映射的形式返回与键相关的值
func (c *Context) GetStringMapString(key string) (sms map[string]string) {
	if val, ok := c.Get(key); ok && val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}

// GetStringMapStringSlice返回与键相关的值，作为到字符串切片的映射
func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string) {
	if val, ok := c.Get(key); ok && val != nil {
		smss, _ = val.(map[string][]string)
	}
	return
}

/************************************/
/************ INPUT DATA ************/
/************************************/

// 参数返回URL参数的值
// 它是c. param . byname (key) router.GET("/user/:id"， func(c *gin.Context) {GET请求/user/john id:= c. param ("id") id == "/john"一个GET请求到/user/john/ id:= c.参数("id") id == "/john/"}）
func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}

// AddParam将参数添加到上下文中，并用给定的值替换路径参数键，用于端到端测试
// 示例Route: "/user/:id"AddParam("id"， 1) Result: "/user/1"
func (c *Context) AddParam(key, value string) {
	c.Params = append(c.Params, Param{Key: key, Value: value})
}

// Query如果存在则返回键控url查询值，否则返回空字符串' ("") '
// 这是快捷方式的' c.Request.URL.Query().Get(key) ' GET /path?id=1234&name= manual &value= c.Query("id") == "1234"c.Query("name") == " manual "c.Query("value") == "c.查询("wtf") == ";
func (c *Context) Query(key string) (value string) {
	value, _ = c.GetQuery(key)
	return
}

// 如果存在，则返回键控url查询值，否则返回指定的defaultValue字符串
// 更多信息请参见:Query()和GetQuery()
// GET / ?name=姓名&lastname= c.DefaultQuery("name"， "unknown") ==姓名"c.DefaultQuery("id"， "none") == "none"c.DefaultQuery("lastname"， "none") == "
func (c *Context) DefaultQuery(key, defaultValue string) string {
	if value, ok := c.GetQuery(key); ok {
		return value
	}
	return defaultValue
}

// GetQuery类似于Query()，如果存在' (value, true) '(即使值是空字符串)，它返回键控url查询值，否则它返回' (""， false) '
// 它是' c.Request.URL.Query().Get(key) ' GET /?name=Manu&lastname= ("Manu"， true) == c.GetQuery("name") (""， false) == c.GetQuery("id") (""， true) == c.GetQuery("lastname")
func (c *Context) GetQuery(key string) (string, bool) {
	if values, ok := c.GetQueryArray(key); ok {
		return values[0], ok
	}
	return "", false
}

// QueryArray返回给定查询键的字符串切片
// 切片的长度取决于具有给定键的参数的数量
func (c *Context) QueryArray(key string) (values []string) {
	values, _ = c.GetQueryArray(key)
	return
}

func (c *Context) initQueryCache() {
	if c.queryCache == nil {
		if c.Request != nil {
			c.queryCache = c.Request.URL.Query()
		} else {
			c.queryCache = url.Values{}
		}
	}
}

// GetQueryArray返回给定查询键的字符串切片，以及一个布尔值，用于判断给定键是否至少存在一个值
func (c *Context) GetQueryArray(key string) (values []string, ok bool) {
	c.initQueryCache()
	values, ok = c.queryCache[key]
	return
}

// QueryMap返回给定查询键的映射
func (c *Context) QueryMap(key string) (dicts map[string]string) {
	dicts, _ = c.GetQueryMap(key)
	return
}

// GetQueryMap返回给定查询键的映射，加上一个布尔值，用于判断给定键是否至少存在一个值
func (c *Context) GetQueryMap(key string) (map[string]string, bool) {
	c.initQueryCache()
	return c.get(c.queryCache, key)
}

// PostForm从存在的POST url编码表单或多部分表单返回指定的键，否则返回空字符串' ("") '
func (c *Context) PostForm(key string) (value string) {
	value, _ = c.GetPostForm(key)
	return
}

// DefaultPostForm从存在的POST url编码表单或多部分表单返回指定的键，否则返回指定的defaultValue字符串
// 参见:PostForm()和GetPostForm()了解更多信息
func (c *Context) DefaultPostForm(key, defaultValue string) string {
	if value, ok := c.GetPostForm(key); ok {
		return value
	}
	return defaultValue
}

// GetPostForm类似于PostForm(key)
// 如果存在' (value, true) '(即使值是空字符串)，则从POST url编码形式或多部分形式返回指定的键，否则返回(""， false)
// 例如，在PATCH请求更新用户的电子邮件时:email=mail@example.com——>("mail@example.com"， true):= GetPostForm("email")设置email为"mail@example.com"电子邮件 =                  --& gt;(""， true):= GetPostForm("email")设置email为"——比;(""， false):= GetPostForm(&q
func (c *Context) GetPostForm(key string) (string, bool) {
	if values, ok := c.GetPostFormArray(key); ok {
		return values[0], ok
	}
	return "", false
}

// PostFormArray返回给定表单键的字符串切片
// 切片的长度取决于具有给定键的参数的数量
func (c *Context) PostFormArray(key string) (values []string) {
	values, _ = c.GetPostFormArray(key)
	return
}

func (c *Context) initFormCache() {
	if c.formCache == nil {
		c.formCache = make(url.Values)
		req := c.Request
		if err := req.ParseMultipartForm(c.engine.MaxMultipartMemory); err != nil {
			if !errors.Is(err, http.ErrNotMultipart) {
				debugPrint("error on parse multipart form array: %v", err)
			}
		}
		c.formCache = req.PostForm
	}
}

// GetPostFormArray返回给定表单键的字符串切片，以及是否至少存在一个给定键的布尔值
func (c *Context) GetPostFormArray(key string) (values []string, ok bool) {
	c.initFormCache()
	values, ok = c.formCache[key]
	return
}

// PostFormMap返回给定表单键的映射
func (c *Context) PostFormMap(key string) (dicts map[string]string) {
	dicts, _ = c.GetPostFormMap(key)
	return
}

// GetPostFormMap返回给定表单键的映射，以及一个布尔值，用于判断给定键是否至少存在一个值
func (c *Context) GetPostFormMap(key string) (map[string]string, bool) {
	c.initFormCache()
	return c.get(c.formCache, key)
}

// Get是一个内部方法，返回一个满足条件的映射
func (c *Context) get(m map[string][]string, key string) (map[string]string, bool) {
	dicts := make(map[string]string)
	exist := false
	for k, v := range m {
		if i := strings.IndexByte(k, '['); i >= 1 && k[0:i] == key {
			if j := strings.IndexByte(k[i+1:], ']'); j >= 1 {
				exist = true
				dicts[k[i+1:][:j]] = v[0]
			}
		}
	}
	return dicts, exist
}

// FormFile返回所提供表单键的第一个文件
func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	if c.Request.MultipartForm == nil {
		if err := c.Request.ParseMultipartForm(c.engine.MaxMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := c.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}

// MultipartForm是解析后的多部分表单，包括文件上传
func (c *Context) MultipartForm() (*multipart.Form, error) {
	err := c.Request.ParseMultipartForm(c.engine.MaxMultipartMemory)
	return c.Request.MultipartForm, err
}

// SaveUploadedFile上传表单文件到指定的dst
func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// Bind检查方法和内容类型以自动选择绑定引擎，具体取决于“内容类型”
// 头文件使用了不同的绑定，例如:"application/json"——比;JSON绑定"application/xml"——比;如果Content-Type == "application/ JSON "使用JSON或XML作为JSON输入
// 它将json有效负载解码为指定为指针的结构
// 它会写一个400的错误，并设置Content-Type header "text/plain"在响应中，如果输入无效
func (c *Context) Bind(obj any) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.MustBindWith(obj, b)
}

// BindJSON是c.MustBindWith(obj, binding.JSON)的快捷方式
func (c *Context) BindJSON(obj any) error {
	return c.MustBindWith(obj, binding.JSON)
}

// BindXML是c.MustBindWith(obj, binding.BindXML)的快捷方式
func (c *Context) BindXML(obj any) error {
	return c.MustBindWith(obj, binding.XML)
}

// BindQuery是c.MustBindWith(obj, binding.Query)的快捷方式
func (c *Context) BindQuery(obj any) error {
	return c.MustBindWith(obj, binding.Query)
}

// BindYAML是c.MustBindWith(obj, binding.YAML)的快捷方式
func (c *Context) BindYAML(obj any) error {
	return c.MustBindWith(obj, binding.YAML)
}

// BindTOML是c.MustBindWith(obj, binding.TOML)的快捷方式
func (c *Context) BindTOML(obj any) error {
	return c.MustBindWith(obj, binding.TOML)
}

// BindHeader是c.MustBindWith(obj, binding.Header)的快捷方式
func (c *Context) BindHeader(obj any) error {
	return c.MustBindWith(obj, binding.Header)
}

// BindUri使用binding.Uri绑定传递的结构指针
// 如果发生任何错误，它将使用HTTP 400中止请求
func (c *Context) BindUri(obj any) error {
	if err := c.ShouldBindUri(obj); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(ErrorTypeBind) //nolint: errcheck
		return err
	}
	return nil
}

// MustBindWith使用指定的绑定引擎绑定传递的结构指针
// 如果发生任何错误，它将使用HTTP 400中止请求
// 参见绑定包
func (c *Context) MustBindWith(obj any, b binding.Binding) error {
	if err := c.ShouldBindWith(obj, b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(ErrorTypeBind) //nolint: errcheck
		return err
	}
	return nil
}

// shoulbind检查方法和内容类型，根据“内容类型”自动选择绑定引擎
// 头文件使用了不同的绑定，例如:"application/json"——比;JSON绑定"application/xml"——比;如果Content-Type == "application/ JSON "使用JSON或XML作为JSON输入
// 它将json有效负载解码为指定为指针的结构
// 与c.Bind()类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
func (c *Context) ShouldBind(obj any) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

// ShouldBindJSON是c.ShouldBindWith(obj, binding.JSON)的快捷方式
func (c *Context) ShouldBindJSON(obj any) error {
	return c.ShouldBindWith(obj, binding.JSON)
}

// ShouldBindXML是c.ShouldBindWith(obj, binding.XML)的快捷方式
func (c *Context) ShouldBindXML(obj any) error {
	return c.ShouldBindWith(obj, binding.XML)
}

// ShouldBindQuery是c.ShouldBindWith(obj, binding.Query)的快捷方式
func (c *Context) ShouldBindQuery(obj any) error {
	return c.ShouldBindWith(obj, binding.Query)
}

// ShouldBindYAML是c.ShouldBindWith(obj, binding.YAML)的快捷方式
func (c *Context) ShouldBindYAML(obj any) error {
	return c.ShouldBindWith(obj, binding.YAML)
}

// ShouldBindTOML是c.ShouldBindWith(obj, binding.TOML)的快捷方式
func (c *Context) ShouldBindTOML(obj any) error {
	return c.ShouldBindWith(obj, binding.TOML)
}

// ShouldBindHeader是c.ShouldBindWith(obj, binding.Header)的快捷方式
func (c *Context) ShouldBindHeader(obj any) error {
	return c.ShouldBindWith(obj, binding.Header)
}

// ShouldBindUri使用指定的绑定引擎绑定传递的结构指针
func (c *Context) ShouldBindUri(obj any) error {
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}
	return binding.Uri.BindUri(m, obj)
}

// ShouldBindWith使用指定的绑定引擎绑定传递的结构指针
// 参见绑定包
func (c *Context) ShouldBindWith(obj any, b binding.Binding) error {
	return b.Bind(c.Request, obj)
}

// ShouldBindBodyWith与ShouldBindWith类似，但它将请求体存储到上下文中，并在再次调用时重用
// 注意:此方法在绑定前读取主体
// 因此，如果只需要调用一次，应该使用ShouldBindWith以获得更好的性能
func (c *Context) ShouldBindBodyWith(obj any, bb binding.BindingBody) (err error) {
	var body []byte
	if cb, ok := c.Get(BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if body == nil {
		body, err = io.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		c.Set(BodyBytesKey, body)
	}
	return bb.BindBody(body, obj)
}

// ClientIP实现了一个最佳努力算法来返回真实的客户端IP
// 它在底层调用c.RemoteIP()来检查远程IP是否是可信代理
// 如果是，它将尝试解析Engine中定义的标头
// RemoteIPHeaders(缺省为[X-Forwarded-For, X-Real-Ip])
// 如果报头在语法上无效或远程IP不对应于可信代理，则返回远程IP(来自Request.RemoteAddr)
func (c *Context) ClientIP() string {
// 检查我们是否运行在一个可信的平台上，如果错误继续运行
	if c.engine.TrustedPlatform != "" {
// 开发人员可以定义自己的可信平台头或使用预定义的常量
		if addr := c.requestHeader(c.engine.TrustedPlatform); addr != "" {
			return addr
		}
	}

// 遗留“AppEngine"国旗
	if c.engine.AppEngine {
		log.Println(`The AppEngine flag is going to be deprecated. Please check issues #2723 and #2739 and use 'TrustedPlatform: gin.PlatformGoogleAppEngine' instead.`)
		if addr := c.requestHeader("X-Appengine-Remote-Addr"); addr != "" {
			return addr
		}
	}

// 它还检查remoteIP是否是受信任的代理
// 为了执行此验证，它将查看IP是否包含在engine定义的至少一个CIDR块中
	remoteIP := net.ParseIP(c.RemoteIP())
	if remoteIP == nil {
		return ""
	}
	trusted := c.engine.isTrustedProxy(remoteIP)

	if trusted && c.engine.ForwardedByClientIP && c.engine.RemoteIPHeaders != nil {
		for _, headerName := range c.engine.RemoteIPHeaders {
			ip, valid := c.engine.validateHeader(c.requestHeader(headerName))
			if valid {
				return ip
			}
		}
	}
	return remoteIP.String()
}

// RemoteIP解析来自Request的IP
// RemoteAddr，规范化并返回IP(不带端口)
func (c *Context) RemoteIP() string {
	ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr))
	if err != nil {
		return ""
	}
	return ip
}

// ContentType返回请求的Content-Type报头
func (c *Context) ContentType() string {
	return filterFlags(c.requestHeader("Content-Type"))
}

// 如果请求头表明客户端正在发起websocket握手，IsWebsocket返回true
func (c *Context) IsWebsocket() bool {
	if strings.Contains(strings.ToLower(c.requestHeader("Connection")), "upgrade") &&
		strings.EqualFold(c.requestHeader("Upgrade"), "websocket") {
		return true
	}
	return false
}

func (c *Context) requestHeader(key string) string {
	return c.Request.Header.Get(key)
}

/************************************/
/******** RESPONSE RENDERING ********/
/************************************/

// bodyAllowedForStatus是http的一个副本
// bodyAllowedForStatus非导出函数
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}

// 状态设置HTTP响应码
func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

// Header是c.Writer.Header()的智能快捷方式
// 集(关键字,值)
// 它在响应中写入一个标头
// 如果value == ""，此方法将删除头' c.Writer.Header().Del(key) '
func (c *Context) Header(key, value string) {
	if value == "" {
		c.Writer.Header().Del(key)
		return
	}
	c.Writer.Header().Set(key, value)
}

// GetHeader从请求头返回值
func (c *Context) GetHeader(key string) string {
	return c.requestHeader(key)
}

// GetRawData返回流数据
func (c *Context) GetRawData() ([]byte, error) {
	return io.ReadAll(c.Request.Body)
}

// SetSameSite with cookie
func (c *Context) SetSameSite(samesite http.SameSite) {
	c.sameSite = samesite
}

// SetCookie在ResponseWriter的报头中添加一个Set-Cookie报头
// 提供的cookie必须有一个有效的Name
// 无效的cookie可能会被静默删除
func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	if path == "" {
		path = "/"
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: c.sameSite,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
}

// Cookie返回请求中提供的命名Cookie，如果没有找到，则返回ErrNoCookie
// 并返回未转义的命名cookie
// 如果多个cookie与给定的名称匹配，则只返回一个cookie
func (c *Context) Cookie(name string) (string, error) {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	val, _ := url.QueryUnescape(cookie.Value)
	return val, nil
}

// Render写入响应头并调用Render
// 渲染到渲染数据
func (c *Context) Render(code int, r render.Render) {
	c.Status(code)

	if !bodyAllowedForStatus(code) {
		r.WriteContentType(c.Writer)
		c.Writer.WriteHeaderNow()
		return
	}

	if err := r.Render(c.Writer); err != nil {
// 将error推入c.Errors
		_ = c.Error(err)
		c.Abort()
	}
}

// HTML呈现由其文件名指定的HTTP模板
// 它还更新HTTP代码并将Content-Type设置为"text/html"
// 参见http://golang.org/doc/articles/wiki/
func (c *Context) HTML(code int, name string, obj any) {
	instance := c.engine.HTMLRender.Instance(name, obj)
	c.Render(code, instance)
}

// indetedjson将给定的结构序列化为漂亮的JSON(缩进+ endlines)到响应体中
// 它还将Content-Type设置为“application/json”
// 警告:我们建议仅用于开发目的，因为打印漂亮的JSON会消耗更多的CPU和带宽
// 使用Context.JSON()代替
func (c *Context) IndentedJSON(code int, obj any) {
	c.Render(code, render.IndentedJSON{Data: obj})
}

// SecureJSON将给定的结构作为安全JSON序列化到响应体中
// Default前面加上"while(1)，"如果给定的结构体是数组值，则返回响应体
// 它还将Content-Type设置为“application/json”
func (c *Context) SecureJSON(code int, obj any) {
	c.Render(code, render.SecureJSON{Prefix: c.engine.secureJSONPrefix, Data: obj})
}

// JSONP将给定的结构作为JSON序列化到响应体中
// 它向响应体添加填充，以便从位于与客户端不同域的服务器请求数据
// 它还将Content-Type设置为"application/javascript"
func (c *Context) JSONP(code int, obj any) {
	callback := c.DefaultQuery("callback", "")
	if callback == "" {
		c.Render(code, render.JSON{Data: obj})
		return
	}
	c.Render(code, render.JsonpJSON{Callback: callback, Data: obj})
}

// JSON将给定的结构作为JSON序列化到响应体中
// 它还将Content-Type设置为“application/json”
func (c *Context) JSON(code int, obj any) {
	c.Render(code, render.JSON{Data: obj})
}

// AsciiJSON将给定的结构作为JSON序列化到响应体中，并使用unicode到ASCII字符串
// 它还将Content-Type设置为“application/json”
func (c *Context) AsciiJSON(code int, obj any) {
	c.Render(code, render.AsciiJSON{Data: obj})
}

// PureJSON将给定的结构作为JSON序列化到响应体中
// 与JSON不同的是，PureJSON不会用它们的unicode实体替换特殊的html字符
func (c *Context) PureJSON(code int, obj any) {
	c.Render(code, render.PureJSON{Data: obj})
}

// XML将给定的结构作为XML序列化到响应体中
// 它还将Content-Type设置为“application/xml”
func (c *Context) XML(code int, obj any) {
	c.Render(code, render.XML{Data: obj})
}

// YAML将给定的结构作为YAML序列化到响应体中
func (c *Context) YAML(code int, obj any) {
	c.Render(code, render.YAML{Data: obj})
}

// TOML将给定的结构作为TOML序列化到响应体中
func (c *Context) TOML(code int, obj any) {
	c.Render(code, render.TOML{Data: obj})
}

// ProtoBuf将给定的结构体作为ProtoBuf序列化到响应体中
func (c *Context) ProtoBuf(code int, obj any) {
	c.Render(code, render.ProtoBuf{Data: obj})
}

// String将给定的字符串写入响应体
func (c *Context) String(code int, format string, values ...any) {
	c.Render(code, render.String{Format: format, Data: values})
}

// Redirect返回到特定位置的HTTP重定向
func (c *Context) Redirect(code int, location string) {
	c.Render(-1, render.Redirect{
		Code:     code,
		Location: location,
		Request:  c.Request,
	})
}

// Data将一些数据写入主体流并更新HTTP代码
func (c *Context) Data(code int, contentType string, data []byte) {
	c.Render(code, render.Data{
		ContentType: contentType,
		Data:        data,
	})
}

// DataFromReader将指定的阅读器写入正文流并更新HTTP代码
func (c *Context) DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) {
	c.Render(code, render.Reader{
		Headers:       extraHeaders,
		ContentType:   contentType,
		ContentLength: contentLength,
		Reader:        reader,
	})
}

// File以一种有效的方式将指定的文件写入体流
func (c *Context) File(filepath string) {
	http.ServeFile(c.Writer, c.Request, filepath)
}

// FileFromFS从http写入指定的文件
// 文件系统以一种有效的方式进入主体流
func (c *Context) FileFromFS(filepath string, fs http.FileSystem) {
	defer func(old string) {
		c.Request.URL.Path = old
	}(c.Request.URL.Path)

	c.Request.URL.Path = filepath

	http.FileServer(fs).ServeHTTP(c.Writer, c.Request)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// FileAttachment以一种有效的方式将指定的文件写入正文流
// 在客户端，文件通常会以给定的文件名下载
func (c *Context) FileAttachment(filepath, filename string) {
	if isASCII(filename) {
		c.Writer.Header().Set("Content-Disposition", `attachment; filename="`+escapeQuotes(filename)+`"`)
	} else {
		c.Writer.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(filename))
	}
	http.ServeFile(c.Writer, c.Request, filepath)
}

// SSEvent将服务器发送的事件写入主体流
func (c *Context) SSEvent(name string, message any) {
	c.Render(-1, sse.Event{
		Event: name,
		Data:  message,
	})
}

// 流发送一个流响应并返回一个布尔值，表示“客户端在流的中间断开了连接”;
func (c *Context) Stream(step func(w io.Writer) bool) bool {
	w := c.Writer
	clientGone := w.CloseNotify()
	for {
		select {
		case <-clientGone:
			return true
		default:
			keepOpen := step(w)
			w.Flush()
			if !keepOpen {
				return false
			}
		}
	}
}

/************************************/
/******** CONTENT NEGOTIATION *******/
/************************************/

// Negotiate包含所有谈判数据
type Negotiate struct {
	Offered  []string
	HTMLName string
	HTMLData any
	JSONData any
	XMLData  any
	YAMLData any
	Data     any
	TOMLData any
}

// 根据可接受的Accept格式协商调用不同的Render
func (c *Context) Negotiate(code int, config Negotiate) {
	switch c.NegotiateFormat(config.Offered...) {
	case binding.MIMEJSON:
		data := chooseData(config.JSONData, config.Data)
		c.JSON(code, data)

	case binding.MIMEHTML:
		data := chooseData(config.HTMLData, config.Data)
		c.HTML(code, config.HTMLName, data)

	case binding.MIMEXML:
		data := chooseData(config.XMLData, config.Data)
		c.XML(code, data)

	case binding.MIMEYAML:
		data := chooseData(config.YAMLData, config.Data)
		c.YAML(code, data)

	case binding.MIMETOML:
		data := chooseData(config.TOMLData, config.Data)
		c.TOML(code, data)

	default:
		c.AbortWithError(http.StatusNotAcceptable, errors.New("the accepted formats are not offered by the server")) //nolint: errcheck
	}
}

// NegotiateFormat返回一个可接受的Accept格式
func (c *Context) NegotiateFormat(offered ...string) string {
	assert1(len(offered) > 0, "you must provide at least one offer")

	if c.Accepted == nil {
		c.Accepted = parseAccept(c.requestHeader("Accept"))
	}
	if len(c.Accepted) == 0 {
		return offered[0]
	}
	for _, accepted := range c.Accepted {
		for _, offer := range offered {
// 根据RFC 2616和RFC 2396，头中不允许使用非ascii字符，因此我们可以迭代字符串，而不将其转换为[]rune
			i := 0
			for ; i < len(accepted) && i < len(offer); i++ {
				if accepted[i] == '*' || offer[i] == '*' {
					return offer
				}
				if accepted[i] != offer[i] {
					break
				}
			}
			if i == len(accepted) {
				return offer
			}
		}
	}
	return ""
}

// SetAccepted设置接受报头数据
func (c *Context) SetAccepted(formats ...string) {
	c.Accepted = formats
}

/************************************/
/***** GOLANG.ORG/X/NET/CONTEXT *****/
/************************************/

// hasRequestContext返回c.Request是否有Context和fallback
func (c *Context) hasRequestContext() bool {
	hasFallback := c.engine != nil && c.engine.ContextWithFallback
	hasRequestContext := c.Request != nil && c.Request.Context() != nil
	return hasFallback && hasRequestContext
}

// 当c.Request没有Context时，Deadline返回没有Deadline (ok==false)
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	if !c.hasRequestContext() {
		return
	}
	return c.Request.Context().Deadline()
}

// 当c.Request没有上下文时，Done返回nil (chan将永远等待)
func (c *Context) Done() <-chan struct{} {
	if !c.hasRequestContext() {
		return nil
	}
	return c.Request.Context().Done()
}

// 当c.Request没有Context时，Err返回nil
func (c *Context) Err() error {
	if !c.hasRequestContext() {
		return nil
	}
	return c.Request.Context().Err()
}

// Value为key返回与此上下文关联的值，如果没有值与key关联，则返回nil
// 连续调用具有相同键的Value返回相同的结果
func (c *Context) Value(key any) any {
	if key == 0 {
		return c.Request
	}
	if key == ContextKey {
		return c
	}
	if keyAsString, ok := key.(string); ok {
		if val, exists := c.Get(keyAsString); exists {
			return val
		}
	}
	if !c.hasRequestContext() {
		return nil
	}
	return c.Request.Context().Value(key)
}
