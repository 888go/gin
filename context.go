// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

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

// Content-Type MIME 是最常见的数据格式的 MIME 类型。
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

// BodyBytesKey 表示一个默认的正文字节键。
const BodyBytesKey = "_gin-gonic/gin/bodybyteskey"

// ContextKey 是一个键，用于在 Context 中返回其自身。
const ContextKey = "_gin-gonic/gin/contextkey"

// abortIndex 表示在中止函数中常用的一个典型值。
const abortIndex int8 = math.MaxInt8 >> 1

// Context 是 gin 中最重要的部分。它允许我们在中间件之间传递变量，管理流程，验证请求的 JSON，并例如渲染 JSON 响应。
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

	// 这个互斥锁保护了Keys映射。
	mu sync.RWMutex

	// Keys 是一组键值对，它在每个请求的上下文中具有唯一性。
	X上下文设置值Map map[string]any

	// Errors 是一个错误列表，其中包含了所有使用了此上下文的处理器/中间件所附加的错误。
	X错误s errorMsgs

	// Accepted 定义了一个手动接受的内容协商格式列表。
	Accepted []string

	// queryCache 对从 c.Request.URL.Query() 获取的查询结果进行缓存。
	queryCache url.Values

// formCache 对 c.Request.PostForm 进行缓存，其中包含从 POST、PATCH 或 PUT 请求体参数解析得到的表单数据。
	formCache url.Values

// SameSite 允许服务器定义一个 cookie 属性，使得浏览器无法在跨站请求中携带此 cookie。
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

// Copy 返回当前上下文的副本，该副本可以在请求范围之外安全使用。
// 当需要将上下文传递给一个goroutine时，必须使用此方法。
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

// HandlerName 返回主处理程序的名称。例如，如果处理程序是 "handleGetUsers()"，
// 该函数将返回 "main.handleGetUsers"。
// 例如:
// 如果处理程序为“handleGetUsers()”，则此函数将返回“main.handleGetUsers”
// 包名为"github.com/888go/gin",返回如下:
// github.com/888go/gin.handleGetUsers
func (c *Context) X取主处理程序名称() string {
	return nameOfFunction(c.handlers.X取最后一个处理函数())
}

// HandlerNames 返回与此上下文关联的已注册处理程序的降序列表，遵循HandlerName()的语义
// 返回数组参考如下:
// 0 = {string} "github.com/888go/gin.TestContextHandlerNames.func1"
// 1 = {string} "github.com/888go/gin.handlerNameTest"
// 2 = {string} "github.com/888go/gin.TestContextHandlerNames.func2"
// 3 = {string} "github.com/888go/gin.handlerNameTest2"
func (c *Context) X取处理程序数组() []string {
	hn := make([]string, 0, len(c.handlers))
	for _, val := range c.handlers {
		hn = append(hn, nameOfFunction(val))
	}
	return hn
}

// Handler 返回主处理程序。
func (c *Context) X取主处理程序() HandlerFunc {
	return c.handlers.X取最后一个处理函数()
}

// FullPath 返回已匹配路由的完整路径。对于未找到的路由，返回一个空字符串。
//
// 示例：
//   router.GET("/user/:id", func(c *gin.Context) {
//       c.FullPath() == "/user/:id" // 将会返回 true
//   })
func (c *Context) X取路由路径() string {
	return c.fullPath
}

/************************************/
/*********** FLOW CONTROL ***********/
/************************************/

// Next 应仅在中间件内部使用。
// 它在调用处理程序内部执行链中待处理的后续处理程序。
// 参考 GitHub 上的示例。
func (c *Context) X中间件继续() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// IsAborted 返回当前上下文是否已中止。
func (c *Context) X是否已停止() bool {
	return c.index >= abortIndex
}

// Abort 阻止待处理的中间件被调用。请注意，这不会停止当前处理器。
// 假设你有一个授权中间件用于验证当前请求是否已授权。
// 如果授权失败（例如，密码不匹配），则调用 Abort 来确保该请求的剩余处理器不会被调用。
func (c *Context) X停止() {
	c.index = abortIndex
}

// AbortWithStatus 方法调用 `Abort()`，并使用指定的状态码写入头部信息。
// 例如，在尝试验证请求失败时，可以这样使用：context.AbortWithStatus(401)。
func (c *Context) X停止并带状态码(状态码 int) {
	c.X设置状态码(状态码)
	c.Writer.WriteHeaderNow()
	c.X停止()
}

// AbortWithStatusJSON 在内部调用`Abort()`和`JSON`方法。
// 该方法中断执行链，写入状态码并返回一个JSON格式的响应体。
// 同时将Content-Type设置为"application/json"。
func (c *Context) X停止并带状态码且返回JSON(状态码 int, JSON结构 any) {
	c.X停止()
	c.X输出JSON(状态码, JSON结构)
}

// AbortWithError 在内部调用 `AbortWithStatus()` 和 `Error()`。
// 该方法停止执行链，写入状态码并将指定错误推送到 `c.Errors`。
// 有关更多详细信息，请参阅 Context.Error()。
func (c *Context) X停止并带状态码与错误(状态码 int, 错误 error) *Error {
	c.X停止并带状态码(状态码)
	return c.X错误(错误)
}

/************************************/
/********* ERROR MANAGEMENT *********/
/************************************/

// Error 将错误附着到当前上下文中。该错误会被推送到错误列表中。
// 在请求解析过程中，对于发生的每个错误调用 Error 是一个好主意。
// 可以使用中间件来收集所有错误，并将它们一起推送到数据库、打印日志或将其添加到HTTP响应中。
// 如果err为nil，Error将会触发panic。
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

// Set 用于为当前上下文独占存储一个新的键值对。
// 如果之前未使用过，它还会初始化 c.Keys。
func (c *Context) X设置值(名称 string, 值 any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.X上下文设置值Map == nil {
		c.X上下文设置值Map = make(map[string]any)
	}

	c.X上下文设置值Map[名称] = 值
}

// Get 方法根据给定的键返回其对应的值，即：(value, true)。
// 若该值不存在，则返回 (nil, false)。
func (c *Context) X取值(名称 string) (返回值 any, 是否存在 bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	返回值, 是否存在 = c.X上下文设置值Map[名称]
	return
}

// MustGet 返回给定键对应的值，如果该键存在。否则，函数会触发panic异常。
func (c *Context) X取值PANI(名称 string) any {
	if value, exists := c.X取值(名称); exists {
		return value
	}
	panic("Key \"" + 名称 + "\" does not exist")
}

// GetString 方法返回与键关联的值，以字符串形式。
func (c *Context) X取文本值(名称 string) (返回值 string) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(string)
	}
	return
}

// GetBool返回与key关联的值，将其转化为布尔类型。
func (c *Context) X取布尔值(名称 string) (返回值 bool) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(bool)
	}
	return
}

// GetInt 通过键返回与其关联的整数值。
func (c *Context) X取整数值(名称 string) (返回值 int) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(int)
	}
	return
}

// GetInt64 以整数形式返回与键关联的值。
func (c *Context) X取整数64位值(名称 string) (返回值 int64) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(int64)
	}
	return
}

// GetUint 返回与键关联的值，以无符号整数形式。
func (c *Context) X取正整数值(名称 string) (返回值 uint) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(uint)
	}
	return
}

// GetUint64返回与key关联的值，将其转化为无符号整数。
func (c *Context) X取正整数64位值(名称 string) (返回值 uint64) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(uint64)
	}
	return
}

// GetFloat64 通过key返回关联的float64类型的值。
func (c *Context) X取小数64位值(名称 string) (返回值 float64) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(float64)
	}
	return
}

// GetTime 函数通过键返回其关联的时间值。
func (c *Context) X取时间值(名称 string) (返回值 time.Time) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回值, _ = val.(time.Time)
	}
	return
}

// GetDuration返回与键关联的值，其类型为持续时间。
func (c *Context) X取时长值(名称 string) (返回时长 time.Duration) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回时长, _ = val.(time.Duration)
	}
	return
}

// GetStringSlice 函数返回与键关联的值，该值为字符串切片。
func (c *Context) X取数组值(名称 string) (返回数组 []string) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回数组, _ = val.([]string)
	}
	return
}

// GetStringMap 返回与键关联的值，该值为接口映射（map）类型。
func (c *Context) X取Map值(名称 string) (返回Map map[string]any) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回Map, _ = val.(map[string]any)
	}
	return
}

// GetStringMapString返回与键关联的值，该值为字符串映射（map）类型。
func (c *Context) X取文本Map值(名称 string) (返回Map map[string]string) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回Map, _ = val.(map[string]string)
	}
	return
}

// GetStringMapStringSlice 返回与键关联的值，该值为字符串到字符串切片的映射。
func (c *Context) X取数组Map值(名称 string) (返回数组Map map[string][]string) {
	if val, ok := c.X取值(名称); ok && val != nil {
		返回数组Map, _ = val.(map[string][]string)
	}
	return
}

/************************************/
/************ INPUT DATA ************/
/************************************/

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
func (c *Context) X取API参数值(名称 string) string {
	return c.X参数.ByName(名称)
}

// AddParam 将参数添加到上下文，并为了端到端测试的目的，用给定的值替换路径参数键
// 示例路由："/user/:id"
// AddParam("id", 1)
// 结果："/user/1"
func (c *Context) X设置API参数值(名称, 值 string) {
	c.X参数 = append(c.X参数, Param{Key: 名称, Value: 值})
}

// Query方法返回键所对应的URL查询值，如果该值存在，则返回该值，否则返回一个空字符串 `("")`。
// 这是 `c.Request.URL.Query().Get(key)` 的快捷方式。
//
//    GET /path?id=1234&name=Manu&value=
//       c.Query("id") 返回 "1234"
//       c.Query("name") 返回 "Manu"
//       c.Query("value") 返回 ""
//       c.Query("wtf") 返回 ""
func (c *Context) X取URL参数值(名称 string) (返回值 string) {
	返回值, _ = c.X取URL参数值2(名称)
	return
}

// DefaultQuery 返回键值对形式的URL查询参数的值，如果该参数存在，则返回其值；否则返回指定的defaultValue字符串。
// 有关更多详细信息，请参阅：Query() 和 GetQuery()。
//
// 示例：
// 请求 GET /?name=Manu&lastname=
// c.DefaultQuery("name", "unknown") 将返回 "Manu"
// c.DefaultQuery("id", "none") 将返回 "none"
// c.DefaultQuery("lastname", "none") 将返回 ""
func (c *Context) X取URL参数值并带默认(名称, 默认值 string) string {
	if value, ok := c.X取URL参数值2(名称); ok {
		return value
	}
	return 默认值
}

// GetQuery 方法类似于 Query()，当给定键的 URL 查询值存在时，它返回该查询值及其对应的布尔值 `(value, true)`（即使该值是一个空字符串）；
// 否则，它返回 `("", false)`。此方法是 `c.Request.URL.Query().Get(key)` 的快捷方式。
//
// 示例：
// 请求 GET /?name=Manu&lastname=
// ("Manu", true) 等价于 c.GetQuery("name")
// ("", false) 等价于 c.GetQuery("id")
// ("", true) 等价于 c.GetQuery("lastname")
func (c *Context) X取URL参数值2(名称 string) (string, bool) {
	if values, ok := c.X取URL参数数组值2(名称); ok {
		return values[0], ok
	}
	return "", false
}

// QueryArray 函数针对给定的查询键返回一个字符串切片。
// 返回切片的长度取决于具有该键的参数的数量。
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

// GetQueryArray 返回给定查询键的字符串切片，以及
// 一个布尔值，表示该键是否存在至少一个值。
func (c *Context) X取URL参数数组值2(名称 string) (返回数组 []string, 是否存在 bool) {
	c.initQueryCache()
	返回数组, 是否存在 = c.queryCache[名称]
	return
}

// QueryMap 根据给定的查询键返回一个映射（map）。
func (c *Context) X取URL参数Map值(名称 string) (返回Map map[string]string) {
	返回Map, _ = c.X取URL参数Map值2(名称)
	return
}

// GetQueryMap 为给定的查询键返回一个映射（map），同时返回一个布尔值，
// 表示该键是否存在至少一个值。
func (c *Context) X取URL参数Map值2(名称 string) (map[string]string, bool) {
	c.initQueryCache()
	return c.get(c.queryCache, 名称)
}

// PostForm 返回从 POST 请求中 urlencoded 表单或 multipart 表单获取的指定键值，如果该键存在，则返回其对应的值；否则返回空字符串 `("")`。
func (c *Context) X取表单参数值(名称 string) (返回值 string) {
	返回值, _ = c.X取表单参数值2(名称)
	return
}

// DefaultPostForm 函数在 POST 请求的 urlencoded 表单或 multipart 表单中查找指定键的值，
// 如果该键存在，则返回对应的值，否则返回指定的 defaultValue 字符串。
// 有关更多信息，请参阅 PostForm() 和 GetPostForm() 函数。
func (c *Context) X取表单参数值并带默认(名称, 默认值 string) string {
	if value, ok := c.X取表单参数值2(名称); ok {
		return value
	}
	return 默认值
}

// 以下是将给定的Go注释翻译成中文：
// 
// GetPostForm 类似于 PostForm(key)。当存在时，它从POST urlencoded表单或multipart表单中返回指定键的值 `(value, true)`（即使该值为空字符串），
// 否则返回 ("", false)。
// 例如，在进行PATCH请求以更新用户邮箱时：
//
//	    email=mail@example.com  -->  ("mail@example.com", true) := GetPostForm("email") // 将邮箱设置为 "mail@example.com"
//		   email=                  -->  ("", true) := GetPostForm("email") // 将邮箱设置为空字符串
//	                            -->  ("", false) := GetPostForm("email") // 对邮箱不做任何处理
func (c *Context) X取表单参数值2(名称 string) (string, bool) {
	if values, ok := c.X取参数数组值(名称); ok {
		return values[0], ok
	}
	return "", false
}

// PostFormArray 为给定的表单键返回一个字符串切片。
// 切片的长度取决于具有该键的参数的数量。
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

// GetPostFormArray 针对给定表单键返回一个字符串切片，以及
// 一个布尔值，表示该键是否存在至少一个值。
func (c *Context) X取参数数组值(名称 string) (返回数组 []string, 是否存在 bool) {
	c.initFormCache()
	返回数组, 是否存在 = c.formCache[名称]
	return
}

// PostFormMap 为给定的表单键返回一个映射（map）。
func (c *Context) X取表单参数Map值(名称 string) (返回Map map[string]string) {
	返回Map, _ = c.X取参数Map值(名称)
	return
}

// GetPostFormMap 为给定的表单键返回一个映射，同时返回一个布尔值，
// 表示是否存在至少一个为此给定键的值。
func (c *Context) X取参数Map值(名称 string) (map[string]string, bool) {
	c.initFormCache()
	return c.get(c.formCache, 名称)
}

// get 是一个内部方法，它返回一个满足特定条件的地图（map）。
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

// FormFile返回提供的表单键所对应的第一个文件。
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

// MultipartForm 是已解析的多部分表单，包括文件上传。
func (c *Context) X取表单multipart对象() (*multipart.Form, error) {
	err := c.X请求.ParseMultipartForm(c.engine.X最大Multipart内存)
	return c.X请求.MultipartForm, err
}

// SaveUploadedFile 将表单文件上传到指定的dst。
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

// Bind 会根据 Method 和 Content-Type 自动选择绑定引擎，
// 根据 "Content-Type" 头部的不同，使用不同的绑定方式，例如：
//
//	"application/json" --> JSON 绑定
//	"application/xml"  --> XML 绑定
//
// 若 Content-Type 为 "application/json"，它将把请求体解析为 JSON，同时可将 XML 视为 JSON 输入进行处理。
// 它会将 json 数据解码到指定的结构体指针中。
// 如果输入无效，则在响应中写入 400 错误，并设置 Content-Type 头部为 "text/plain"。
func (c *Context) X取参数到指针PANI(结构指针 any) error {
	b := binding.Default(c.X请求.Method, c.X取协议头ContentType())
	return c.X取参数到指针并按类型PANI(结构指针, b)
}

// BindJSON 是一个快捷方式，等同于 c.MustBindWith(obj, binding.JSON)。
func (c *Context) X取JSON参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.JSON)
}

// BindXML 是一个快捷方式，用于 c.MustBindWith(obj, binding.BindXML)。 
// 
// 更详细的翻译：
// 
// BindXML 是一个便捷方法，它等同于调用 c.MustBindWith(obj, binding.BindXML)。
// 其中，c 通常代表上下文（Context），obj 代表要绑定的对象，binding.BindXML 表示使用 XML 绑定方式进行数据绑定。这个方法会确保 XML 数据成功绑定到对象上，如果绑定失败，则会触发 panic。
func (c *Context) X取XML参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.XML)
}

// BindQuery 是一个快捷方式，用于 c.MustBindWith(obj, binding.Query)。
// 如果解析错误,它将使用HTTP 400中止请求
// BindQuery 函数只绑定 url 查询参数而忽略 post 数据。参阅详细信息:
// https://gin-gonic.com/zh-cn/docs/examples/only-bind-query-string/
func (c *Context) X取URL参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.Query)
}

// BindYAML 是一个快捷方式，等同于 c.MustBindWith(obj, binding.YAML)。
// 注意: 如果解析错误,它将使用HTTP 400中止请求
func (c *Context) X取YAML参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.YAML)
}

// BindTOML 是一个快捷方式，用于 c.MustBindWith(obj, binding.TOML)。
// 注意: 如果解析错误,它将使用HTTP 400中止请求
func (c *Context) X取TOML参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.TOML)
}

// BindHeader 是一个快捷方式，等同于 c.MustBindWith(obj, binding.Header)。
// 注意: 如果解析错误,它将使用HTTP 400中止请求
func (c *Context) X取Header参数到指针PANI(结构指针 any) error {
	return c.X取参数到指针并按类型PANI(结构指针, binding.Header)
}

// BindUri通过binding.Uri将传递的结构体指针进行绑定。
// 如果发生任何错误，它将使用HTTP 400中止请求。
func (c *Context) X取Uri参数到指针PANI(结构指针 any) error {
	if err := c.X取Uri参数到指针(结构指针); err != nil {
		c.X停止并带状态码与错误(http.StatusBadRequest, err).SetType(ErrorTypeBind) //nolint: errcheck
		return err
	}
	return nil
}

// MustBindWith 使用指定的绑定引擎绑定传入的结构体指针。
// 如果在执行过程中出现任何错误，它将终止请求并返回HTTP状态码400。
// 请参阅binding包以获取更多信息。
func (c *Context) X取参数到指针并按类型PANI(结构指针 any, 类型 binding.Binding) error {
	if err := c.X取参数到指针并按类型(结构指针, 类型); err != nil {
		c.X停止并带状态码与错误(http.StatusBadRequest, err).SetType(ErrorTypeBind) //nolint: errcheck
		return err
	}
	return nil
}

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
func (c *Context) X取参数到指针(变量结构指针 any) error {
	b := binding.Default(c.X请求.Method, c.X取协议头ContentType())
	return c.X取参数到指针并按类型(变量结构指针, b)
}

// ShouldBindJSON 是 c.ShouldBindWith(obj, binding.JSON) 的快捷方式。
//
// 注意: c.ShouldBind***方法不能多次被调用, 如果绑定类型为" JSON, XML, MsgPack, ProtoBuf", 第一次绑定之后 c.Request.Body会设置成EOF, 如果需要多次绑定, 可以使用c.ShouldBindBodyWith
func (c *Context) X取JSON参数到指针(JSON结构指针 any) error {
	return c.X取参数到指针并按类型(JSON结构指针, binding.JSON)
}

// ShouldBindXML 是 c.ShouldBindWith(obj, binding.XML) 的快捷方式。
//
// 注意: c.ShouldBind***方法不能多次被调用, 如果绑定类型为" JSON, XML, MsgPack, ProtoBuf", 第一次绑定之后 c.Request.Body会设置成EOF, 如果需要多次绑定, 可以使用c.ShouldBindBodyWith
func (c *Context) X取XML参数到指针(XML结构指针 any) error {
	return c.X取参数到指针并按类型(XML结构指针, binding.XML)
}

// ShouldBindQuery 是一个快捷方式，用于 c.ShouldBindWith(obj, binding.Query)。
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
// ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据。参阅详细信息:
// https://gin-gonic.com/zh-cn/docs/examples/only-bind-query-string/
func (c *Context) X取URL参数到指针(表单结构指针 any) error {
	return c.X取参数到指针并按类型(表单结构指针, binding.Query)
}

// ShouldBindYAML 是 c.ShouldBindWith(obj, binding.YAML) 的快捷方式。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
func (c *Context) X取YAML参数到指针(YAML结构指针 any) error {
	return c.X取参数到指针并按类型(YAML结构指针, binding.YAML)
}

// ShouldBindTOML 是 c.ShouldBindWith(obj, binding.TOML) 的快捷方式。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
func (c *Context) X取TOML参数到指针(TOML结构指针 any) error {
	return c.X取参数到指针并按类型(TOML结构指针, binding.TOML)
}

// ShouldBindHeader 是一个快捷方式，用于 c.ShouldBindWith(obj, binding.Header)。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
func (c *Context) X取Header参数到指针(Header结构指针 any) error {
	return c.X取参数到指针并按类型(Header结构指针, binding.Header)
}

// ShouldBindUri 使用指定的绑定引擎，将传入的结构体指针进行绑定。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
func (c *Context) X取Uri参数到指针(Uri结构指针 any) error {
	m := make(map[string][]string)
	for _, v := range c.X参数 {
		m[v.Key] = []string{v.Value}
	}
	return binding.Uri.BindUri(m, Uri结构指针)
}

// ShouldBindWith 使用指定的绑定引擎绑定传入的结构体指针。
// 请参阅binding包。
//
// 与c.Bind***()方法类似，但此方法不会将响应状态码设置为400，也不会在输入无效时中止
func (c *Context) X取参数到指针并按类型(结构指针 any, 类型 binding.Binding) error {
	return 类型.Bind(c.X请求, 结构指针)
}

// ShouldBindBodyWith 与 ShouldBindWith 类似，但它会将请求体存储到上下文中，并在再次调用时重用。
//
// 注意：此方法在绑定前读取请求体。因此，如果你只需要调用一次，为了获得更好的性能，你应该使用 ShouldBindWith。
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

// ClientIP 实现了一种尽力而为的算法，用于返回真实的客户端 IP 地址。
// 在底层，它调用 c.RemoteIP() 来检查远程 IP 是否为可信代理。
// 如果是可信代理，则尝试解析 Engine.RemoteIPHeaders 中定义的头部（默认为 [X-Forwarded-For, X-Real-Ip]）。
// 如果这些头部格式不合法 或者 远程 IP 不对应于一个可信代理，
// 则返回来自 Request.RemoteAddr 的远程 IP 地址。
func (c *Context) X取客户端ip() string {
	// 检查我们是否在受信任的平台上运行，如果有错误则继续向后执行
	if c.engine.TrustedPlatform != "" {
		// 开发者可以定义自己的可信平台头文件，也可以使用预定义的常量
		if addr := c.requestHeader(c.engine.TrustedPlatform); addr != "" {
			return addr
		}
	}

	// "AppEngine"老版本标志
	if c.engine.AppEngine弃用 {
		log.Println(`The AppEngine flag is going to be deprecated. Please check issues #2723 and #2739 and use 'TrustedPlatform: gin.PlatformGoogleAppEngine' instead.`)
		if addr := c.requestHeader("X-Appengine-Remote-Addr"); addr != "" {
			return addr
		}
	}

// 它还会检查 remoteIP 是否为可信代理。
// 为了执行此验证，它会查看该 IP 是否至少包含在由 Engine.SetTrustedProxies() 方法定义的一个 CIDR 块中。
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

// RemoteIP 从 Request.RemoteAddr 解析 IP，进行规范化处理并返回不含端口号的 IP 地址。
//
// 注意: 协议头ip容易被伪造.应当使用ClientIP
func (c *Context) X取协议头ip() string {
	ip, _, err := net.SplitHostPort(strings.TrimSpace(c.X请求.RemoteAddr))
	if err != nil {
		return ""
	}
	return ip
}

// ContentType 返回请求的 Content-Type 头部信息。
func (c *Context) X取协议头ContentType() string {
	return filterFlags(c.requestHeader("Content-Type"))
}

// IsWebsocket 返回一个布尔值，如果请求头表明客户端正在进行websocket握手，则返回true。
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

// bodyAllowedForStatus 是 http 包中未导出函数 bodyAllowedForStatus 的复制版本。
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

// Status 设置 HTTP 响应代码。
func (c *Context) X设置状态码(状态码 int) {
	c.Writer.WriteHeader(状态码)
}

// Header 是一个智能快捷方式，用于 c.Writer.Header().Set(key, value)。
// 它在响应中写入一个头信息。
// 如果 value 等于 "", 则此方法会删除相应头信息：`c.Writer.Header().Del(key)`。
func (c *Context) X设置响应协议头值(名称, 值 string) {
	if 值 == "" {
		c.Writer.Header().Del(名称)
		return
	}
	c.Writer.Header().Set(名称, 值)
}

// GetHeader 从请求头中返回值。
func (c *Context) X取请求协议头值(名称 string) string {
	return c.requestHeader(名称)
}

// GetRawData 返回原始数据流。
func (c *Context) X取流数据() ([]byte, error) {
	return io.ReadAll(c.X请求.Body)
}

// SetSameSite 设置 cookie 的同站属性
func (c *Context) X设置cookie跨站(samesite http.SameSite) {
	c.sameSite = samesite
}

// SetCookie 向 ResponseWriter 的头信息中添加一个 Set-Cookie 头部。提供的 cookie 必须具有有效的名称。不合法的 cookie 可能会被悄悄丢弃。
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

// Cookie返回请求中提供的指定名称的cookie，如果未找到，则返回ErrNoCookie错误。同时返回的指定名称的cookie是经过解码的。
// 如果多个cookie与给定名称匹配，则只返回一个cookie。
func (c *Context) X取cookie值(名称 string) (string, error) {
	cookie, err := c.X请求.Cookie(名称)
	if err != nil {
		return "", err
	}
	val, _ := url.QueryUnescape(cookie.Value)
	return val, nil
}

// Render方法会写入响应头并调用render.Render来渲染数据。
func (c *Context) Render底层方法(状态码 int, r render.Render) {
	c.X设置状态码(状态码)

	if !bodyAllowedForStatus(状态码) {
		r.WriteContentType(c.Writer)
		c.Writer.WriteHeaderNow()
		return
	}

	if err := r.Render(c.Writer); err != nil {
		// 将错误推送到c.Errors
		_ = c.X错误(err)
		c.X停止()
	}
}

// HTML 根据其文件名渲染 HTTP 模板。
// 同时，它还会更新 HTTP 状态码，并将 Content-Type 设置为 "text/html"。
// 详情参见：http://golang.org/doc/articles/wiki/
func (c *Context) X输出html模板(状态码 int, 模板文件名 string, 结构 any) {
	instance := c.engine.HTMLRender.Instance(模板文件名, 结构)
	c.Render底层方法(状态码, instance)
}

// IndentedJSON 将给定的结构体序列化为美观的 JSON（缩进+换行符）并写入响应体中。
// 同时，它还会将 Content-Type 设置为 "application/json"。
// 警告：我们建议仅在开发目的下使用此方法，因为打印美观的 JSON 会消耗更多的 CPU 和带宽。请改用 Context.JSON()。
func (c *Context) X输出JSON并美化(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.IndentedJSON{Data: 结构})
}

// SecureJSON将给定的结构体作为安全的JSON序列化到响应体中。
// 默认情况下，如果给定的结构体是数组值，则会在响应体前缀添加 "while(1),"。
// 同时，它还会将Content-Type设置为"application/json"。
func (c *Context) X输出JSON并防劫持(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.SecureJSON{Prefix: c.engine.secureJSONPrefix, Data: 结构})
}

// JSONP将给定的结构体以JSON格式序列化到响应体中。
// 它在响应体中添加填充，以便从与客户端不同域的服务器请求数据。
// 同时，它还将Content-Type设置为"application/javascript"。
func (c *Context) X输出JSONP(状态码 int, 结构 any) {
	callback := c.X取URL参数值并带默认("callback", "")
	if callback == "" {
		c.Render底层方法(状态码, render.JSON{Data: 结构})
		return
	}
	c.Render底层方法(状态码, render.JsonpJSON{Callback: callback, Data: 结构})
}

// JSON将给定的结构体以JSON格式序列化到响应体中。
// 同时，它还将Content-Type设置为"application/json"。
func (c *Context) X输出JSON(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.JSON{Data: 结构})
}

// AsciiJSON 将给定的结构体按 JSON 格式序列化，并以 ASCII 字符串形式写入响应体中。
// 同时，它还会将 Content-Type 设置为 "application/json"。
func (c *Context) X输出JSON并按ASCII(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.AsciiJSON{Data: 结构})
}

// PureJSON 将给定的结构体作为 JSON 序列化到响应体中。
// 与 JSON 不同，PureJSON 不会将特殊 HTML 字符替换为它们的 Unicode 实体。
func (c *Context) X输出JSON并按原文(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.PureJSON{Data: 结构})
}

// XML将给定的结构体作为XML序列化到响应体中。
// 同时，它还会将Content-Type设置为"application/xml"。
func (c *Context) X输出XML(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.XML{Data: 结构})
}

// YAML 将给定的结构体以 YAML 格式序列化并写入响应体中。
func (c *Context) X输出YAML(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.YAML{Data: 结构})
}

// TOML将给定的结构体序列化为TOML格式，并写入响应体中。
func (c *Context) X输出TOML(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.TOML{Data: 结构})
}

// ProtoBuf将给定的结构体作为ProtoBuf序列化到响应体中。
func (c *Context) X输出ProtoBuf(状态码 int, 结构 any) {
	c.Render底层方法(状态码, render.ProtoBuf{Data: 结构})
}

// String将给定的字符串写入响应体中。
func (c *Context) X输出文本(状态码 int, 格式 string, 文本s ...any) {
	c.Render底层方法(状态码, render.String{Format: 格式, Data: 文本s})
}

// Redirect 返回一个HTTP重定向到特定位置。
func (c *Context) X重定向(状态码 int, 重定向地址 string) {
	c.Render底层方法(-1, render.Redirect{
		Code:     状态码,
		Location: 重定向地址,
		Request:  c.X请求,
	})
}

// Data 将一些数据写入主体流并更新HTTP状态码。
func (c *Context) X输出字节集(状态码 int, HTTP响应类型 string, 字节集 []byte) {
	c.Render底层方法(状态码, render.Data{
		ContentType: HTTP响应类型,
		Data:        字节集,
	})
}

// DataFromReader 将指定读取器的内容写入主体流，并更新HTTP状态码。
func (c *Context) X输出字节集并按IO(状态码 int, 数据长度 int64, HTTP响应类型 string, 写出IO数据 io.Reader, 协议头Map map[string]string) {
	c.Render底层方法(状态码, render.Reader{
		Headers:       协议头Map,
		ContentType:   HTTP响应类型,
		ContentLength: 数据长度,
		Reader:        写出IO数据,
	})
}

// File 以高效的方式将指定的文件写入正文流中。
func (c *Context) X下载文件(文件路径 string) {
	http.ServeFile(c.Writer, c.X请求, 文件路径)
}

// FileFromFS 以高效的方式将指定的文件从 http.FileSystem 写入到 body 流中。
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

// FileAttachment 以高效的方式将指定文件写入主体流
// 在客户端，该文件通常会以给定的文件名下载
func (c *Context) X下载文件并带文件名(文件路径, 文件名 string) {
	if isASCII(文件名) {
		c.Writer.Header().Set("Content-Disposition", `attachment; filename="`+escapeQuotes(文件名)+`"`)
	} else {
		c.Writer.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(文件名))
	}
	http.ServeFile(c.Writer, c.X请求, 文件路径)
}

// SSEvent 将一个服务器发送事件写入到主体数据流中。
func (c *Context) SSEvent(name string, message any) {
	c.Render底层方法(-1, sse.Event{
		Event: name,
		Data:  message,
	})
}

// Stream 发送一个流式响应，并返回一个布尔值
// 表示“在流传输过程中客户端是否已断开连接”
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

// Negotiate 包含所有协商数据。
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

// Negotiate 根据可接受的 Accept 格式调用不同的 Render 方法。
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

// NegotiateFormat 返回一个可接受的 Accept 格式。
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
// 根据RFC 2616和RFC 2396的规定，非ASCII字符在头部中是不允许出现的，
// 因此我们可以在不将其转换为[]rune的情况下直接遍历该字符串。
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

// SetAccepted 设置 Accept 头部数据。
func (c *Context) SetAccepted底层方法(类型名称s ...string) {
	c.Accepted = 类型名称s
}

/************************************/
/***** GOLANG.ORG/X/NET/CONTEXT *****/
/************************************/

// hasRequestContext 返回 c.Request 是否包含 Context 以及回退机制。
func (c *Context) hasRequestContext() bool {
	hasFallback := c.engine != nil && c.engine.ContextWithFallback
	hasRequestContext := c.X请求 != nil && c.X请求.Context() != nil
	return hasFallback && hasRequestContext
}

// Deadline 返回当 c.Request 没有 Context 时，表示没有截止时间（ok==false）。
//
// 注意!!! 此方法不能翻译, 因为是http包的接口实现
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	if !c.hasRequestContext() {
		return
	}
	return c.X请求.Context().Deadline()
}

// 当c.Request没有Context时，Done返回nil（表示一个将永远等待的通道）。
//
// 注意!!! 此方法不能翻译, 因为是http包的接口实现
func (c *Context) Done() <-chan struct{} {
	if !c.hasRequestContext() {
		return nil
	}
	return c.X请求.Context().Done()
}

// Err在c.Request没有Context时返回nil。
//
// 注意!!! 此方法不能翻译, 因为是http包的接口实现
func (c *Context) Err() error {
	if !c.hasRequestContext() {
		return nil
	}
	return c.X请求.Context().Err()
}

// Value 方法返回与该上下文关联的键key所对应的值，如果该键没有关联任何值，则返回nil。对同一键连续调用Value方法将返回相同的结果。
//
// 注意!!! 此方法不能翻译, 因为是http包的接口实现 
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
