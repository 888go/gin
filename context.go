// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

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

// Content-Type MIME of the most common data formats.
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

// BodyBytesKey indicates a default body bytes key.
const BodyBytesKey = "_gin-gonic/gin/bodybyteskey"

// ContextKey is the key that a Context returns itself for.
const ContextKey = "_gin-gonic/gin/contextkey"

// abortIndex represents a typical value used in abort functions.
const abortIndex int8 = math.MaxInt8 >> 1

// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
type Context struct {
	writermem responseWriter
	X请求   *http.Request
	Writer    ResponseWriter

	X参数   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine       *Engine
	params       *Params
	skippedNodes *[]skippedNode

	// This mutex protects Keys map.
	mu sync.RWMutex

	// Keys is a key/value pair exclusively for the context of each request.
	X上下文设置值Map map[string]any

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
	X错误s errorMsgs

	// Accepted defines a list of manually accepted formats for content negotiation.
	Accepted []string

	// queryCache caches the query result from c.Request.URL.Query().
	queryCache url.Values

	// formCache caches c.Request.PostForm, which contains the parsed form data from POST, PATCH,
	// or PUT body parameters.
	formCache url.Values

	// SameSite allows a server to define a cookie attribute making it impossible for
	// the browser to send this cookie along with cross-site requests.
	sameSite http.SameSite
}

/************************************/
/********** CONTEXT CREATION ********/
/************************************/

func (c *Context) reset() {
	c.Writer = &c.writermem
	c.X参数 = c.X参数[:0]
	c.handlers = nil
	c.index = -1

	c.fullPath = ""
	c.X上下文设置值Map = nil
	c.X错误s = c.X错误s[:0]
	c.Accepted = nil
	c.queryCache = nil
	c.formCache = nil
	c.sameSite = 0
	*c.params = (*c.params)[:0]
	*c.skippedNodes = (*c.skippedNodes)[:0]
}

// Copy returns a copy of the current context that can be safely used outside the request's scope.
// This has to be used when the context has to be passed to a goroutine.
func (c *Context) X取副本() *Context {
	cp := Context{
		writermem: c.writermem,
		X请求:   c.X请求,
		X参数:    c.X参数,
		engine:    c.engine,
	}
	cp.writermem.ResponseWriter = nil
	cp.Writer = &cp.writermem
	cp.index = abortIndex
	cp.handlers = nil
	cp.X上下文设置值Map = map[string]any{}
	for k, v := range c.X上下文设置值Map {
		cp.X上下文设置值Map[k] = v
	}
	paramCopy := make([]Param, len(cp.X参数))
	copy(paramCopy, cp.X参数)
	cp.X参数 = paramCopy
	return &cp
}

// HandlerName returns the main handler's name. For example if the handler is "handleGetUsers()",
// this function will return "main.handleGetUsers".
func (c *Context) X取主处理程序名称() string {
	return nameOfFunction(c.handlers.X取最后一个处理函数())
}

// HandlerNames returns a list of all registered handlers for this context in descending order,
// following the semantics of HandlerName()
func (c *Context) X取处理程序数组() []string {
	hn := make([]string, 0, len(c.handlers))
	for _, val := range c.handlers {
		hn = append(hn, nameOfFunction(val))
	}
	return hn
}

// Handler returns the main handler.
func (c *Context) X取主处理程序() HandlerFunc {
	return c.handlers.X取最后一个处理函数()
}

// FullPath returns a matched route full path. For not found routes
// returns an empty string.
//
//	router.GET("/user/:id", func(c *gin.Context) {
//	    c.FullPath() == "/user/:id" // true
//	})
func (c *Context) X取路由路径() string {
	return c.fullPath
}

/************************************/
/*********** FLOW CONTROL ***********/
/************************************/

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in GitHub.
func (c *Context) X中间件继续() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// IsAborted returns true if the current context was aborted.
func (c *Context) X是否已停止() bool {
	return c.index >= abortIndex
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
func (c *Context) X停止() {
	c.index = abortIndex
}

// AbortWithStatus calls `Abort()` and writes the headers with the specified status code.
// For example, a failed attempt to authenticate a request could use: context.AbortWithStatus(401).
func (c *Context) X停止并带状态码(状态码 int) {
	c.X设置状态码(状态码)
	c.Writer.WriteHeaderNow()
	c.X停止()
}

// AbortWithStatusJSON calls `Abort()` and then `JSON` internally.
// This method stops the chain, writes the status code and return a JSON body.
// It also sets the Content-Type as "application/json".
func (c *Context) X停止并带状态码且返回JSON(状态码 int, JSON结构 any) {
	c.X停止()
	c.X输出JSON(状态码, JSON结构)
}

// AbortWithError calls `AbortWithStatus()` and `Error()` internally.
// This method stops the chain, writes the status code and pushes the specified error to `c.Errors`.
// See Context.Error() for more details.
func (c *Context) X停止并带状态码与错误(状态码 int, 错误 error) *Error {
	c.X停止并带状态码(状态码)
	return c.X错误(错误)
}

/************************************/
/********* ERROR MANAGEMENT *********/
/************************************/

// Error attaches an error to the current context. The error is pushed to a list of errors.
// It's a good idea to call Error for each error that occurred during the resolution of a request.
// A middleware can be used to collect all the errors and push them to a database together,
// print a log, or append it in the HTTP response.
// Error will panic if err is nil.
func (c *Context) X错误(错误 error) *Error {
	if 错误 == nil {
		panic("err is nil")
	}

	var parsedError *Error
	ok := errors.As(错误, &parsedError)
	if !ok {
		parsedError = &Error{
			Err:  错误,
			Type: ErrorTypePrivate,
		}
	}

	c.X错误s = append(c.X错误s, parsedError)
	return parsedError
}

/************************************/
/******** METADATA MANAGEMENT********/
/************************************/

// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
func (c *Context) X设置值(名称 string, 值 any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.X上下文设置值Map == nil {
		c.X上下文设置值Map = make(map[string]any)
	}

	c.X上下文设置值Map[名称] = 值
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
func (c *Context) X取值(名称 string) (返回值 any, 是否存在 bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	返回值, 是否存在 = c.X上下文设置值Map[名称]
	return
}

// MustGet returns the value for the given key if it exists, otherwise it panics.
func (c *Context) X取值PANI(名称 string) any {
	if value, exists := c.X取值(名称); exists {
		return value
	}
	panic("Key \"" + 名称 + "\" does not exist")
}

// GetString returns the value associated with the key as a string.
func (c *Context) X取文本值(名称 string) (返回值 string) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(string)
	}
	return
}

// GetBool returns the value associated with the key as a boolean.
func (c *Context) X取布尔值(名称 string) (返回值 bool) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(bool)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) X取整数值(名称 string) (返回值 int) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(int)
	}
	return
}

// GetInt64 returns the value associated with the key as an integer.
func (c *Context) X取整数64位值(名称 string) (返回值 int64) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(int64)
	}
	return
}

// GetUint returns the value associated with the key as an unsigned integer.
func (c *Context) X取正整数值(名称 string) (返回值 uint) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(uint)
	}
	return
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func (c *Context) X取正整数64位值(名称 string) (返回值 uint64) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(uint64)
	}
	return
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Context) X取小数64位值(名称 string) (返回值 float64) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(float64)
	}
	return
}

// GetTime returns the value associated with the key as time.
func (c *Context) X取时间值(名称 string) (返回值 time.Time) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(time.Time)
	}
	return
}

// GetDuration returns the value associated with the key as a duration.
func (c *Context) X取时长值(名称 string) (返回时长 time.Duration) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回时长, _ = val.(time.Duration)
	}
	return
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (c *Context) X取数组值(名称 string) (返回数组 []string) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回数组, _ = val.([]string)
	}
	return
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (c *Context) X取Map值(名称 string) (返回Map map[string]any) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回Map, _ = val.(map[string]any)
	}
	return
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (c *Context) X取文本Map值(名称 string) (返回Map map[string]string) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回Map, _ = val.(map[string]string)
	}
	return
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (c *Context) X取数组Map值(名称 string) (返回数组Map map[string][]string) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回数组Map, _ = val.(map[string][]string)
	}
	return
}

/************************************/
/************ INPUT DATA ************/
/************************************/

// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
//
//	router.GET("/user/:id", func(c *gin.Context) {
//	    // a GET request to /user/john
//	    id := c.Param("id") // id == "/john"
//	    // a GET request to /user/john/
//	    id := c.Param("id") // id == "/john/"
//	})
func (c *Context) X取API参数值(名称 string) string {
	return c.X参数.ByName(名称)
}

// AddParam adds param to context and
// replaces path param key with given value for e2e testing purposes
// Example Route: "/user/:id"
// AddParam("id", 1)
// Result: "/user/1"
func (c *Context) X设置API参数值(名称, 值 string) {
	c.X参数 = append(c.X参数, Param{Key: 名称, Value: 值})
}

// Query returns the keyed url query value if it exists,
// otherwise it returns an empty string `("")`.
// It is shortcut for `c.Request.URL.Query().Get(key)`
//
//	    GET /path?id=1234&name=Manu&value=
//		   c.Query("id") == "1234"
//		   c.Query("name") == "Manu"
//		   c.Query("value") == ""
//		   c.Query("wtf") == ""
func (c *Context) X取URL参数值(名称 string) (返回值 string) {
	返回值, _ = c.X取URL参数值2(名称)
	return
}

// DefaultQuery returns the keyed url query value if it exists,
// otherwise it returns the specified defaultValue string.
// See: Query() and GetQuery() for further information.
//
//	GET /?name=Manu&lastname=
//	c.DefaultQuery("name", "unknown") == "Manu"
//	c.DefaultQuery("id", "none") == "none"
//	c.DefaultQuery("lastname", "none") == ""
func (c *Context) X取URL参数值并带默认(名称, 默认值 string) string {
	if value, ok := c.X取URL参数值2(名称); ok {
		return value
	}
	return 默认值
}

// GetQuery is like Query(), it returns the keyed url query value
// if it exists `(value, true)` (even when the value is an empty string),
// otherwise it returns `("", false)`.
// It is shortcut for `c.Request.URL.Query().Get(key)`
//
//	GET /?name=Manu&lastname=
//	("Manu", true) == c.GetQuery("name")
//	("", false) == c.GetQuery("id")
//	("", true) == c.GetQuery("lastname")
func (c *Context) X取URL参数值2(名称 string) (string, bool) {
	if values, ok := c.X取URL参数数组值2(名称); ok {
		return values[0], ok
	}
	return "", false
}

// QueryArray returns a slice of strings for a given query key.
// The length of the slice depends on the number of params with the given key.
func (c *Context) X取URL参数数组值(名称 string) (返回数组 []string) {
	返回数组, _ = c.X取URL参数数组值2(名称)
	return
}

func (c *Context) initQueryCache() {
	if c.queryCache == nil {
		if c.X请求 != nil {
			c.queryCache = c.X请求.URL.Query()
		} else {
			c.queryCache = url.Values{}
		}
	}
}

// GetQueryArray returns a slice of strings for a given query key, plus
// a boolean value whether at least one value exists for the given key.
func (c *Context) X取URL参数数组值2(名称 string) (返回数组 []string, 是否存在 bool) {
	c.initQueryCache()
	返回数组, 是否存在 = c.queryCache[名称]
	return
}

// QueryMap returns a map for a given query key.
func (c *Context) X取URL参数Map值(名称 string) (返回Map map[string]string) {
	返回Map, _ = c.X取URL参数Map值2(名称)
	return
}

// GetQueryMap returns a map for a given query key, plus a boolean value
// whether at least one value exists for the given key.
func (c *Context) X取URL参数Map值2(名称 string) (map[string]string, bool) {
	c.initQueryCache()
	return c.get(c.queryCache, 名称)
}

// PostForm returns the specified key from a POST urlencoded form or multipart form
// when it exists, otherwise it returns an empty string `("")`.
func (c *Context) X取表单参数值(名称 string) (返回值 string) {
	返回值, _ = c.X取表单参数值2(名称)
	return
}

// DefaultPostForm returns the specified key from a POST urlencoded form or multipart form
// when it exists, otherwise it returns the specified defaultValue string.
// See: PostForm() and GetPostForm() for further information.
func (c *Context) X取表单参数值并带默认(名称, 默认值 string) string {
	if value, ok := c.X取表单参数值2(名称); ok {
		return value
	}
	return 默认值
}

// GetPostForm is like PostForm(key). It returns the specified key from a POST urlencoded
// form or multipart form when it exists `(value, true)` (even when the value is an empty string),
// otherwise it returns ("", false).
// For example, during a PATCH request to update the user's email:
//
//	    email=mail@example.com  -->  ("mail@example.com", true) := GetPostForm("email") // set email to "mail@example.com"
//		   email=                  -->  ("", true) := GetPostForm("email") // set email to ""
//	                            -->  ("", false) := GetPostForm("email") // do nothing with email
func (c *Context) X取表单参数值2(名称 string) (string, bool) {
	if values, ok := c.X取参数数组值(名称); ok {
		return values[0], ok
	}
	return "", false
}

// PostFormArray returns a slice of strings for a given form key.
// The length of the slice depends on the number of params with the given key.
func (c *Context) X取表单参数数组值(名称 string) (返回数组 []string) {
	返回数组, _ = c.X取参数数组值(名称)
	return
}

func (c *Context) initFormCache() {
	if c.formCache == nil {
		c.formCache = make(url.Values)
		req := c.X请求
		if err := req.ParseMultipartForm(c.engine.X最大Multipart内存); err != nil {
			if !errors.Is(err, http.ErrNotMultipart) {
				debugPrint("error on parse multipart form array: %v", err)
			}
		}
		c.formCache = req.PostForm
	}
}

// GetPostFormArray returns a slice of strings for a given form key, plus
// a boolean value whether at least one value exists for the given key.
func (c *Context) X取参数数组值(名称 string) (返回数组 []string, 是否存在 bool) {
	c.initFormCache()
	返回数组, 是否存在 = c.formCache[名称]
	return
}

// PostFormMap returns a map for a given form key.
func (c *Context) X取表单参数Map值(名称 string) (返回Map map[string]string) {
	返回Map, _ = c.X取参数Map值(名称)
	return
}

// GetPostFormMap returns a map for a given form key, plus a boolean value
// whether at least one value exists for the given key.
func (c *Context) X取参数Map值(名称 string) (map[string]string, bool) {
	c.initFormCache()
	return c.get(c.formCache, 名称)
}

// get is an internal method and returns a map which satisfies conditions.
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

// FormFile returns the first file for the provided form key.
func (c *Context) X取表单上传文件(名称 string) (*multipart.FileHeader, error) {
	if c.X请求.MultipartForm == nil {
		if err := c.X请求.ParseMultipartForm(c.engine.X最大Multipart内存); err != nil {
			return nil, err
		}
	}
	f, fh, err := c.X请求.FormFile(名称)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}

// MultipartForm is the parsed multipart form, including file uploads.
func (c *Context) X取表单multipart对象() (*multipart.Form, error) {
	err := c.X请求.ParseMultipartForm(c.engine.X最大Multipart内存)
	return c.X请求.MultipartForm, err
}

// SaveUploadedFile uploads the form file to specific dst.
func (c *Context) X保存上传文件(文件对象 *multipart.FileHeader, 文件路径 string) error {
	src, err := 文件对象.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(文件路径), 0750); err != nil {
		return err
	}

	out, err := os.Create(文件路径)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// Bind checks the Method and Content-Type to select a binding engine automatically,
// Depending on the "Content-Type" header different bindings are used, for example:
//
//	"application/json" --> JSON binding
//	"application/xml"  --> XML binding
//
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// It writes a 400 error and sets Content-Type header "text/plain" in the response if input is not valid.
func (c *Context) X取参数到指针PANI(结构指针 any) error {
	b := binding.Default(c.X请求.Method, c.X取协议头ContentType())
	return c.X取参数到指针并按类型PANI(结构指针, b)
}

// BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
func (c *Context) X取JSON参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.JSON)
}

// BindXML is a shortcut for c.MustBindWith(obj, binding.BindXML).
func (c *Context) X取XML参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.XML)
}

// BindQuery is a shortcut for c.MustBindWith(obj, binding.Query).
func (c *Context) X取URL参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.Query)
}

// BindYAML is a shortcut for c.MustBindWith(obj, binding.YAML).
func (c *Context) X取YAML参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.YAML)
}

// BindTOML is a shortcut for c.MustBindWith(obj, binding.TOML).
func (c *Context) X取TOML参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.TOML)
}

// BindHeader is a shortcut for c.MustBindWith(obj, binding.Header).
func (c *Context) X取Header参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.Header)
}

// BindUri binds the passed struct pointer using binding.Uri.
// It will abort the request with HTTP 400 if any error occurs.
func (c *Context) X取Uri参数到指针PANI(结构指针 any) error {
	if err := c.X取Uri参数到指针(结构指针); err != nil {
		c.X停止并带状态码与错误(http.StatusBadRequest, err).SetType(ErrorTypeBind) //nolint: errcheck
		return err
	}
	return nil
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func (c *Context) X取参数到指针并按类型PANI(结构指针 any, 类型 binding.Binding) error {
	if err := c.X取参数到指针并按类型(结构指针, 类型); err != nil {
		c.X停止并带状态码与错误(http.StatusBadRequest, err).SetType(ErrorTypeBind) //nolint: errcheck
		return err
	}
	return nil
}

// ShouldBind checks the Method and Content-Type to select a binding engine automatically,
// Depending on the "Content-Type" header different bindings are used, for example:
//
//	"application/json" --> JSON binding
//	"application/xml"  --> XML binding
//
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// Like c.Bind() but this method does not set the response status code to 400 or abort if input is not valid.
func (c *Context) X取参数到指针(变量结构指针 any) error {
	b := binding.Default(c.X请求.Method, c.X取协议头ContentType())
	return c.X取参数到指针并按类型(变量结构指针, b)
}

// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
func (c *Context) X取JSON参数到指针(JSON结构指针 any) error {
	return c.X取参数到指针并按类型(JSON结构指针, binding.JSON)
}

// ShouldBindXML is a shortcut for c.ShouldBindWith(obj, binding.XML).
func (c *Context) X取XML参数到指针(XML结构指针 any) error {
	return c.X取参数到指针并按类型(XML结构指针, binding.XML)
}

// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
func (c *Context) X取URL参数到指针(表单结构指针 any) error {
	return c.X取参数到指针并按类型(表单结构指针, binding.Query)
}

// ShouldBindYAML is a shortcut for c.ShouldBindWith(obj, binding.YAML).
func (c *Context) X取YAML参数到指针(YAML结构指针 any) error {
	return c.X取参数到指针并按类型(YAML结构指针, binding.YAML)
}

// ShouldBindTOML is a shortcut for c.ShouldBindWith(obj, binding.TOML).
func (c *Context) X取TOML参数到指针(TOML结构指针 any) error {
	return c.X取参数到指针并按类型(TOML结构指针, binding.TOML)
}

// ShouldBindHeader is a shortcut for c.ShouldBindWith(obj, binding.Header).
func (c *Context) X取Header参数到指针(Header结构指针 any) error {
	return c.X取参数到指针并按类型(Header结构指针, binding.Header)
}

// ShouldBindUri binds the passed struct pointer using the specified binding engine.
func (c *Context) X取Uri参数到指针(Uri结构指针 any) error {
	m := make(map[string][]string)
	for _, v := range c.X参数 {
		m[v.Key] = []string{v.Value}
	}
	return binding.Uri.BindUri(m, Uri结构指针)
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func (c *Context) X取参数到指针并按类型(结构指针 any, 类型 binding.Binding) error {
	return 类型.Bind(c.X请求, 结构指针)
}

// ShouldBindBodyWith is similar with ShouldBindWith, but it stores the request
// body into the context, and reuse when it is called again.
//
// NOTE: This method reads the body before binding. So you should use
// ShouldBindWith for better performance if you need to call only once.
func (c *Context) X取参数到指针并按类型且缓存(结构指针 any, bb binding.BindingBody) (错误 error) {
	var body []byte
	if cb, ok := c.X取值(BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if body == nil {
		body, 错误 = io.ReadAll(c.X请求.Body)
		if 错误 != nil {
			return 错误
		}
		c.X设置值(BodyBytesKey, body)
	}
	return bb.BindBody(body, 结构指针)
}

// ClientIP implements one best effort algorithm to return the real client IP.
// It calls c.RemoteIP() under the hood, to check if the remote IP is a trusted proxy or not.
// If it is it will then try to parse the headers defined in Engine.RemoteIPHeaders (defaulting to [X-Forwarded-For, X-Real-Ip]).
// If the headers are not syntactically valid OR the remote IP does not correspond to a trusted proxy,
// the remote IP (coming from Request.RemoteAddr) is returned.
func (c *Context) X取客户端ip() string {
	// Check if we're running on a trusted platform, continue running backwards if error
	if c.engine.TrustedPlatform != "" {
		// Developers can define their own header of Trusted Platform or use predefined constants
		if addr := c.requestHeader(c.engine.TrustedPlatform); addr != "" {
			return addr
		}
	}

	// Legacy "AppEngine" flag
	if c.engine.AppEngine弃用 {
		log.Println(`The AppEngine flag is going to be deprecated. Please check issues #2723 and #2739 and use 'TrustedPlatform: gin.PlatformGoogleAppEngine' instead.`)
		if addr := c.requestHeader("X-Appengine-Remote-Addr"); addr != "" {
			return addr
		}
	}

	// It also checks if the remoteIP is a trusted proxy or not.
	// In order to perform this validation, it will see if the IP is contained within at least one of the CIDR blocks
	// defined by Engine.SetTrustedProxies()
	remoteIP := net.ParseIP(c.X取协议头ip())
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

// RemoteIP parses the IP from Request.RemoteAddr, normalizes and returns the IP (without the port).
func (c *Context) X取协议头ip() string {
	ip, _, err := net.SplitHostPort(strings.TrimSpace(c.X请求.RemoteAddr))
	if err != nil {
		return ""
	}
	return ip
}

// ContentType returns the Content-Type header of the request.
func (c *Context) X取协议头ContentType() string {
	return filterFlags(c.requestHeader("Content-Type"))
}

// IsWebsocket returns true if the request headers indicate that a websocket
// handshake is being initiated by the client.
func (c *Context) X是否为Websocket请求() bool {
	if strings.Contains(strings.ToLower(c.requestHeader("Connection")), "upgrade") &&
		strings.EqualFold(c.requestHeader("Upgrade"), "websocket") {
		return true
	}
	return false
}

func (c *Context) requestHeader(key string) string {
	return c.X请求.Header.Get(key)
}

/************************************/
/******** RESPONSE RENDERING ********/
/************************************/

// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
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

// Status sets the HTTP response code.
func (c *Context) X设置状态码(状态码 int) {
	c.Writer.WriteHeader(状态码)
}

// Header is an intelligent shortcut for c.Writer.Header().Set(key, value).
// It writes a header in the response.
// If value == "", this method removes the header `c.Writer.Header().Del(key)`
func (c *Context) X设置响应协议头值(名称, 值 string) {
	if 值 == "" {
		c.Writer.Header().Del(名称)
		return
	}
	c.Writer.Header().Set(名称, 值)
}

// GetHeader returns value from request headers.
func (c *Context) X取请求协议头值(名称 string) string {
	return c.requestHeader(名称)
}

// GetRawData returns stream data.
func (c *Context) X取流数据() ([]byte, error) {
	return io.ReadAll(c.X请求.Body)
}

// SetSameSite with cookie
func (c *Context) X设置cookie跨站(samesite http.SameSite) {
	c.sameSite = samesite
}

// SetCookie adds a Set-Cookie header to the ResponseWriter's headers.
// The provided cookie must have a valid Name. Invalid cookies may be
// silently dropped.
func (c *Context) X设置cookie值(名称, 值 string, 生效时间 int, 路径, 域名 string, 仅https生效, 禁止js访问 bool) {
	if 路径 == "" {
		路径 = "/"
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     名称,
		Value:    url.QueryEscape(值),
		MaxAge:   生效时间,
		Path:     路径,
		Domain:   域名,
		SameSite: c.sameSite,
		Secure:   仅https生效,
		HttpOnly: 禁止js访问,
	})
}

// Cookie returns the named cookie provided in the request or
// ErrNoCookie if not found. And return the named cookie is unescaped.
// If multiple cookies match the given name, only one cookie will
// be returned.
func (c *Context) X取cookie值(名称 string) (string, error) {
	cookie, err := c.X请求.Cookie(名称)
	if err != nil {
		return "", err
	}
	val, _ := url.QueryUnescape(cookie.Value)
	return val, nil
}

// Render writes the response headers and calls render.Render to render data.
func (c *Context) Render底层方法(状态码 int, r render.Render) {
	c.X设置状态码(状态码)

	if !bodyAllowedForStatus(状态码) {
		r.WriteContentType(c.Writer)
		c.Writer.WriteHeaderNow()
		return
	}

	if err := r.Render(c.Writer); err != nil {
		// Pushing error to c.Errors
		_ = c.X错误(err)
		c.X停止()
	}
}

// HTML renders the HTTP template specified by its file name.
// It also updates the HTTP code and sets the Content-Type as "text/html".
// See http://golang.org/doc/articles/wiki/
func (c *Context) X输出html模板(状态码 int, 模板文件名 string, 结构 any) {
	instance := c.engine.HTMLRender.Instance(模板文件名, 结构)
	c.Render底层方法(状态码, instance)
}

// IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.
// It also sets the Content-Type as "application/json".
// WARNING: we recommend using this only for development purposes since printing pretty JSON is
// more CPU and bandwidth consuming. Use Context.JSON() instead.
func (c *Context) X输出JSON并美化(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.IndentedJSON{Data: 结构})
}

// SecureJSON serializes the given struct as Secure JSON into the response body.
// Default prepends "while(1)," to response body if the given struct is array values.
// It also sets the Content-Type as "application/json".
func (c *Context) X输出JSON并防劫持(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.SecureJSON{Prefix: c.engine.secureJSONPrefix, Data: 结构})
}

// JSONP serializes the given struct as JSON into the response body.
// It adds padding to response body to request data from a server residing in a different domain than the client.
// It also sets the Content-Type as "application/javascript".
func (c *Context) X输出JSONP(状态码 int, 结构 any) {
	callback := c.X取URL参数值并带默认("callback", "")
	if callback == "" {
		c.Render底层方法(状态码, render.JSON{Data: 结构})
		return
	}
	c.Render底层方法(状态码, render.JsonpJSON{Callback: callback, Data: 结构})
}

// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func (c *Context) X输出JSON(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.JSON{Data: 结构})
}

// AsciiJSON serializes the given struct as JSON into the response body with unicode to ASCII string.
// It also sets the Content-Type as "application/json".
func (c *Context) X输出JSON并按ASCII(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.AsciiJSON{Data: 结构})
}

// PureJSON serializes the given struct as JSON into the response body.
// PureJSON, unlike JSON, does not replace special html characters with their unicode entities.
func (c *Context) X输出JSON并按原文(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.PureJSON{Data: 结构})
}

// XML serializes the given struct as XML into the response body.
// It also sets the Content-Type as "application/xml".
func (c *Context) X输出XML(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.XML{Data: 结构})
}

// YAML serializes the given struct as YAML into the response body.
func (c *Context) X输出YAML(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.YAML{Data: 结构})
}

// TOML serializes the given struct as TOML into the response body.
func (c *Context) X输出TOML(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.TOML{Data: 结构})
}

// ProtoBuf serializes the given struct as ProtoBuf into the response body.
func (c *Context) X输出ProtoBuf(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.ProtoBuf{Data: 结构})
}

// String writes the given string into the response body.
func (c *Context) X输出文本(状态码 int, 格式 string, 文本s ...any) {
	c.Render底层方法(状态码, render.String{Format: 格式, Data: 文本s})
}

// Redirect returns an HTTP redirect to the specific location.
func (c *Context) X重定向(状态码 int, 重定向地址 string) {
	c.Render底层方法(-1, render.Redirect{
		Code:     状态码,
		Location: 重定向地址,
		Request:  c.X请求,
	})
}

// Data writes some data into the body stream and updates the HTTP code.
func (c *Context) X输出字节集(状态码 int, HTTP响应类型 string, 字节集 []byte) {
	c.Render底层方法(状态码, render.Data{
		ContentType: HTTP响应类型,
		Data:        字节集,
	})
}

// DataFromReader writes the specified reader into the body stream and updates the HTTP code.
func (c *Context) X输出字节集并按IO(状态码 int, 数据长度 int64, HTTP响应类型 string, 写出IO数据 io.Reader, 协议头Map map[string]string) {
	c.Render底层方法(状态码, render.Reader{
		Headers:       协议头Map,
		ContentType:   HTTP响应类型,
		ContentLength: 数据长度,
		Reader:        写出IO数据,
	})
}

// File writes the specified file into the body stream in an efficient way.
func (c *Context) X下载文件(文件路径 string) {
	http.ServeFile(c.Writer, c.X请求, 文件路径)
}

// FileFromFS writes the specified file from http.FileSystem into the body stream in an efficient way.
func (c *Context) X下载文件FS(文件路径 string, fs http.FileSystem) {
	defer func(old string) {
		c.X请求.URL.Path = old
	}(c.X请求.URL.Path)

	c.X请求.URL.Path = 文件路径

	http.FileServer(fs).ServeHTTP(c.Writer, c.X请求)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// FileAttachment writes the specified file into the body stream in an efficient way
// On the client side, the file will typically be downloaded with the given filename
func (c *Context) X下载文件并带文件名(文件路径, 文件名 string) {
	if isASCII(文件名) {
		c.Writer.Header().Set("Content-Disposition", `attachment; filename="`+escapeQuotes(文件名)+`"`)
	} else {
		c.Writer.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(文件名))
	}
	http.ServeFile(c.Writer, c.X请求, 文件路径)
}

// SSEvent writes a Server-Sent Event into the body stream.
func (c *Context) SSEvent(name string, message any) {
	c.Render底层方法(-1, sse.Event{
		Event: name,
		Data:  message,
	})
}

// Stream sends a streaming response and returns a boolean
// indicates "Is client disconnected in middle of stream"
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

// Negotiate contains all negotiations data.
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

// Negotiate calls different Render according to acceptable Accept format.
func (c *Context) Negotiate底层方法(code int, config Negotiate) {
	switch c.NegotiateFormat底层方法(config.Offered...) {
	case binding.MIMEJSON:
		data := chooseData(config.JSONData, config.Data)
		c.X输出JSON(code, data)

	case binding.MIMEHTML:
		data := chooseData(config.HTMLData, config.Data)
		c.X输出html模板(code, config.HTMLName, data)

	case binding.MIMEXML:
		data := chooseData(config.XMLData, config.Data)
		c.X输出XML(code, data)

	case binding.MIMEYAML:
		data := chooseData(config.YAMLData, config.Data)
		c.X输出YAML(code, data)

	case binding.MIMETOML:
		data := chooseData(config.TOMLData, config.Data)
		c.X输出TOML(code, data)

	default:
		c.X停止并带状态码与错误(http.StatusNotAcceptable, errors.New("the accepted formats are not offered by the server")) //nolint: errcheck
	}
}

// NegotiateFormat returns an acceptable Accept format.
func (c *Context) NegotiateFormat底层方法(offered ...string) string {
	assert1(len(offered) > 0, "you must provide at least one offer")

	if c.Accepted == nil {
		c.Accepted = parseAccept(c.requestHeader("Accept"))
	}
	if len(c.Accepted) == 0 {
		return offered[0]
	}
	for _, accepted := range c.Accepted {
		for _, offer := range offered {
			// According to RFC 2616 and RFC 2396, non-ASCII characters are not allowed in headers,
			// therefore we can just iterate over the string without casting it into []rune
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

// SetAccepted sets Accept header data.
func (c *Context) SetAccepted底层方法(类型名称s ...string) {
	c.Accepted = 类型名称s
}

/************************************/
/***** GOLANG.ORG/X/NET/CONTEXT *****/
/************************************/

// hasRequestContext returns whether c.Request has Context and fallback.
func (c *Context) hasRequestContext() bool {
	hasFallback := c.engine != nil && c.engine.ContextWithFallback
	hasRequestContext := c.X请求 != nil && c.X请求.Context() != nil
	return hasFallback && hasRequestContext
}

// Deadline returns that there is no deadline (ok==false) when c.Request has no Context.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	if !c.hasRequestContext() {
		return
	}
	return c.X请求.Context().Deadline()
}

// Done returns nil (chan which will wait forever) when c.Request has no Context.
func (c *Context) Done() <-chan struct{} {
	if !c.hasRequestContext() {
		return nil
	}
	return c.X请求.Context().Done()
}

// Err returns nil when c.Request has no Context.
func (c *Context) Err() error {
	if !c.hasRequestContext() {
		return nil
	}
	return c.X请求.Context().Err()
}

// Value returns the value associated with this context for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
func (c *Context) Value(key any) any {
	if key == 0 {
		return c.X请求
	}
	if key == ContextKey {
		return c
	}
	if keyAsString, ok := key.(string); ok {
		if val, exists := c.X取值(keyAsString); exists {
			return val
		}
	}
	if !c.hasRequestContext() {
		return nil
	}
	return c.X请求.Context().Value(key)
}
