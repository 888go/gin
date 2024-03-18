// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	
	"github.com/888go/gin/internal/bytesconv"
	"github.com/888go/gin/render"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const defaultMultipartMemory = 32 << 20 // 32 MB （32兆字节）

var (
	default404Body = []byte("404 page not found")
	default405Body = []byte("405 method not allowed")
)

var defaultPlatform string

var defaultTrustedCIDRs = []*net.IPNet{
	{ // 0.0.0.0/0 (IPv4) （IPv4地址）：表示整个IPv4地址空间的通配符，相当于所有IPv4地址的集合。
		IP:   net.IP{0x0, 0x0, 0x0, 0x0},
		Mask: net.IPMask{0x0, 0x0, 0x0, 0x0},
	},
	{ // /:: 0 (IPv6) // （此注释内容较为简略，直译为“IPv6的/:: 0”）
// 这个注释可能是在表示一个IPv6地址的特殊表示形式，"/::" 表示IPv6地址中的零压缩写法，其中 "::" 可以替换连续的一串零。当IPv6地址中包含较长的连续零时，可以使用这种简写方式。例如 "/::" 可以代表一串全零的部分，而 "0" 可能是指特定的IPv6地址部分（可能是指IPv6地址的剩余部分为全零）。
// 但由于上下文不完整，这里的具体含义可能需要根据代码的实际应用场景来判断。
		IP:   net.IP{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Mask: net.IPMask{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
}

var regSafePrefix = regexp.MustCompile("[^a-zA-Z0-9/-]+")
var regRemoveRepeatedChar = regexp.MustCompile("/{2,}")

// HandlerFunc定义了gin中间件使用的处理程序作为返回值
type HandlerFunc func(*Context)

// HandlersChain定义了一个handlerfuncc片
type HandlersChain []HandlerFunc

// Last返回链中的最后一个处理程序
// 也就是说，最后一个处理器是主处理器

// ff:
func (c HandlersChain) Last() HandlerFunc {
	if length := len(c); length > 0 {
		return c[length-1]
	}
	return nil
}

// RouteInfo表示一个请求路由的规范，它包含方法、路径和它的处理器
type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}

// RoutesInfo定义了一个RouteInfo切片
type RoutesInfo []RouteInfo

// 信任的平台
const (
// 在Google应用引擎上运行时的平台googleappengine
// 信任X-Appengine-Remote-Addr来确定客户端的IP
	PlatformGoogleAppEngine = "X-Appengine-Remote-Addr"
// 使用Cloudflare的CDN时的平台Cloudflare
// Trust CF-Connecting-IP用于确定客户端的IP
	PlatformCloudflare = "CF-Connecting-IP"
)

// 引擎是框架的实例，它包含了复用器、中间件和配置设置
// 使用New()或Default()创建Engine实例
type Engine struct {
	RouterGroup

// RedirectTrailingSlash在当前路由不能匹配的情况下启用自动重定向，但是存在一个带有(不带有)尾斜杠的路径处理程序
// 例如，如果请求/foo/，但只存在/foo的路由，则客户端被重定向到/foo, GET请求的http状态码为301，所有其他请求方法的http状态码为307
	RedirectTrailingSlash bool

// RedirectFixedPath如果启用，如果没有为它注册句柄，路由器会尝试修复当前的请求路径
// 首先是多余的路径元素，比如…/或被移除
// 之后，路由器会对清理后的路径进行不区分大小写的查找
// 如果能找到该路由的句柄，路由器就会重定向到正确的路径，GET请求的状态码为301，其他所有请求方法的状态码为307
// 例如/FOO和/..Foo可以重定向到/ Foo
// RedirectTrailingSlash与此选项无关
	RedirectFixedPath bool

// handlemethodnotalallowed如果使能，如果当前请求不能被路由，则路由器检查当前路由是否允许另一个方法
// 如果是这种情况，请求将返回“方法不允许”和HTTP状态码405
// 如果不允许使用其他方法，则将请求委托给NotFound处理程序
	HandleMethodNotAllowed bool

// 如果启用了ForwardedByClientIP，客户端IP将从与存储在' (*gin.Engine). remoteipheaders '匹配的请求头中解析
// 如果没有获取到IP，则返回到从' (*gin.Context). request . remoteaddr '获取的IP
	ForwardedByClientIP bool

// AppEngine已弃用
// 已弃用:使用' TrustedPlatform ' WITH VALUE ' gin
// 如果启用，它将信任一些以“X-AppEngine…”开头的标头
// 以便与该PaaS更好地集成
	AppEngine bool //hs:AppEngine弃用     

// UseRawPath如果启用，则为url
// RawPath将用于查找参数
	UseRawPath bool

// UnescapePathValues如果为true，则不转义路径值
// 如果UseRawPath为false(默认情况下)，UnescapePathValues有效地为true，如url
// 路径将被使用，它已经是未转义的
	UnescapePathValues bool

// 即使使用额外的斜杠，也可以从URL解析RemoveExtraSlash参数
// 见PR #1817和issue #1644
	RemoveExtraSlash bool

// RemoteIPHeaders获取客户端IP时使用的报头列表(*gin.Engine)
// ForwardedByClientIP '是' true '和' (*gin.Context). request
// RemoteAddr '被' (*gin.Engine). settrustedproxies() '定义的列表的至少一个网络源匹配
	RemoteIPHeaders []string

// TrustedPlatform设置为一个值为gin的常量
// 例如，平台*信任由该平台设置的报头来确定客户端IP
	TrustedPlatform string

// 给http的“maxMemory”参数的MaxMultipartMemory值请求的parsemmultipartform方法调用
	MaxMultipartMemory int64

// 启用h2c支持
	UseH2C bool

// 当Context.Request.Context()不是nil时，启用回退Context.Deadline()、Context.Done()、Context.Err()和Context.Value()
	ContextWithFallback bool

	delims           render.Delims
	secureJSONPrefix string
	HTMLRender       render.HTMLRender
	FuncMap          template.FuncMap
	allNoRoute       HandlersChain
	allNoMethod      HandlersChain
	noRoute          HandlersChain
	noMethod         HandlersChain
	pool             sync.Pool
	trees            methodTrees
	maxParams        uint16
	maxSections      uint16
	trustedProxies   []string
	trustedCIDRs     []*net.IPNet
}

var _ IRouter = (*Engine)(nil)

// New返回一个新的空白Engine实例，没有附加任何中间件
// 默认配置为:—RedirectTrailingSlash: true—RedirectFixedPath: false—handlemethodnotalallowed: false—ForwardedByClientIP: true—UseRawPath: false—UnescapePathValues: true

// ff:创建
func New() *Engine {
	debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		FuncMap:                template.FuncMap{},
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      false,
		HandleMethodNotAllowed: false,
		ForwardedByClientIP:    true,
		RemoteIPHeaders:        []string{"X-Forwarded-For", "X-Real-IP"},
		TrustedPlatform:        defaultPlatform,
		UseRawPath:             false,
		RemoveExtraSlash:       false,
		UnescapePathValues:     true,
		MaxMultipartMemory:     defaultMultipartMemory,
		trees:                  make(methodTrees, 0, 9),
		delims:                 render.Delims{Left: "{{", Right: "}}"},
		secureJSONPrefix:       "while(1);",
		trustedProxies:         []string{"0.0.0.0/0", "::/0"},
		trustedCIDRs:           defaultTrustedCIDRs,
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() any {
		return engine.allocateContext(engine.maxParams)
	}
	return engine
}

// Default返回一个Engine实例，其中已经附加了Logger和Recovery中间件

// ff:创建默认对象
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}


// ff:
func (engine *Engine) Handler() http.Handler {
	if !engine.UseH2C {
		return engine
	}

	h2s := &http2.Server{}
	return h2c.NewHandler(engine, h2s)
}

func (engine *Engine) allocateContext(maxParams uint16) *Context {
	v := make(Params, 0, maxParams)
	skippedNodes := make([]skippedNode, 0, engine.maxSections)
	return &Context{engine: engine, params: &v, skippedNodes: &skippedNodes}
}

// Delims设置模板的左和右分隔符并返回Engine实例

// ff:设置模板分隔符
// right:右边
// left:左边
func (engine *Engine) Delims(left, right string) *Engine {
	engine.delims = render.Delims{Left: left, Right: right}
	return engine
}

// SecureJsonPrefix设置Context.SecureJSON中使用的SecureJsonPrefix

// ff:
// prefix:
func (engine *Engine) SecureJsonPrefix(prefix string) *Engine {
	engine.secureJSONPrefix = prefix
	return engine
}

// LoadHTMLGlob加载由glob模式标识的HTML文件，并将结果与HTML渲染器相关联

// ff:加载HTML模板目录
// pattern:模板目录
func (engine *Engine) LoadHTMLGlob(pattern string) {
	left := engine.delims.Left
	right := engine.delims.Right
	templ := template.Must(template.New("").Delims(left, right).Funcs(engine.FuncMap).ParseGlob(pattern))

	if IsDebugging() {
		debugPrintLoadTemplate(templ)
		engine.HTMLRender = render.HTMLDebug{Glob: pattern, FuncMap: engine.FuncMap, Delims: engine.delims}
		return
	}

	engine.SetHTMLTemplate(templ)
}

// LoadHTMLFiles加载一段HTML文件，并将结果与HTML渲染器相关联

// ff:加载HTML模板文件
// files:模板文件s
func (engine *Engine) LoadHTMLFiles(files ...string) {
	if IsDebugging() {
		engine.HTMLRender = render.HTMLDebug{Files: files, FuncMap: engine.FuncMap, Delims: engine.delims}
		return
	}

	templ := template.Must(template.New("").Delims(engine.delims.Left, engine.delims.Right).Funcs(engine.FuncMap).ParseFiles(files...))
	engine.SetHTMLTemplate(templ)
}

// SetHTMLTemplate将模板与HTML渲染器关联

// ff:
// templ:
func (engine *Engine) SetHTMLTemplate(templ *template.Template) {
	if len(engine.trees) > 0 {
		debugPrintWARNINGSetHTMLTemplate()
	}

	engine.HTMLRender = render.HTMLProduction{Template: templ.Funcs(engine.FuncMap)}
}

// SetFuncMap设置用于template.FuncMap的FuncMap

// ff:
// funcMap:
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.FuncMap = funcMap
}

// NoRoute为NoRoute添加处理程序
// 默认情况下，它返回404代码

// ff:
// handlers:
func (engine *Engine) NoRoute(handlers ...HandlerFunc) {
	engine.noRoute = handlers
	engine.rebuild404Handlers()
}

// NoMethod设置引擎时调用的处理程序
// handlemethodnotalallowed = true

// ff:
// handlers:
func (engine *Engine) NoMethod(handlers ...HandlerFunc) {
	engine.noMethod = handlers
	engine.rebuild405Handlers()
}

// Use将全局中间件附加到路由器上
// 也就是说，通过Use()附加的中间件将被包含在每个请求的处理程序链中
// 甚至404、405、静态文件……例如，这是日志记录器或错误管理中间件的正确位置

// ff:
// middleware:
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
	engine.RouterGroup.Use(middleware...)
	engine.rebuild404Handlers()
	engine.rebuild405Handlers()
	return engine
}

func (engine *Engine) rebuild404Handlers() {
	engine.allNoRoute = engine.combineHandlers(engine.noRoute)
}

func (engine *Engine) rebuild405Handlers() {
	engine.allNoMethod = engine.combineHandlers(engine.noMethod)
}

func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
	assert1(path[0] == '/', "path must begin with '/'")
	assert1(method != "", "HTTP method can not be empty")
	assert1(len(handlers) > 0, "there must be at least one handler")

	debugPrintRoute(method, path, handlers)

	root := engine.trees.get(method)
	if root == nil {
		root = new(node)
		root.fullPath = "/"
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)

	if paramsCount := countParams(path); paramsCount > engine.maxParams {
		engine.maxParams = paramsCount
	}

	if sectionsCount := countSections(path); sectionsCount > engine.maxSections {
		engine.maxSections = sectionsCount
	}
}

// Routes返回已注册路由的切片，其中包括一些有用的信息，例如:http方法、路径和处理程序名称

// ff:取路由数组
// routes:路由s
func (engine *Engine) Routes() (routes RoutesInfo) {
	for _, tree := range engine.trees {
		routes = iterate("", tree.method, routes, tree.root)
	}
	return routes
}

func iterate(path, method string, routes RoutesInfo, root *node) RoutesInfo {
	path += root.path
	if len(root.handlers) > 0 {
		handlerFunc := root.handlers.Last()
		routes = append(routes, RouteInfo{
			Method:      method,
			Path:        path,
			Handler:     nameOfFunction(handlerFunc),
			HandlerFunc: handlerFunc,
		})
	}
	for _, child := range root.children {
		routes = iterate(path, method, routes, child)
	}
	return routes
}

// Run将路由器附加到http上
// 服务器并开始监听和服务HTTP请求
// 它是http的快捷方式
// 注意:除非发生错误，否则此方法将无限期地阻塞调用例程

// ff:监听
// err:错误
// addr:地址与端口
func (engine *Engine) Run(addr ...string) (err error) {
	defer func() { debugPrintError(err) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.")
	}

	address := resolveAddress(addr)
	debugPrint("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine.Handler())
	return
}

func (engine *Engine) prepareTrustedCIDRs() ([]*net.IPNet, error) {
	if engine.trustedProxies == nil {
		return nil, nil
	}

	cidr := make([]*net.IPNet, 0, len(engine.trustedProxies))
	for _, trustedProxy := range engine.trustedProxies {
		if !strings.Contains(trustedProxy, "/") {
			ip := parseIP(trustedProxy)
			if ip == nil {
				return cidr, &net.ParseError{Type: "IP address", Text: trustedProxy}
			}

			switch len(ip) {
			case net.IPv4len:
				trustedProxy += "/32"
			case net.IPv6len:
				trustedProxy += "/128"
			}
		}
		_, cidrNet, err := net.ParseCIDR(trustedProxy)
		if err != nil {
			return cidr, err
		}
		cidr = append(cidr, cidrNet)
	}
	return cidr, nil
}

// SetTrustedProxies设置了一个网络起源列表(IPv4地址，IPv4 cidr, IPv6地址或IPv6 cidr)，从其中信任请求的头包含替代客户端IP时' (* gin.com engine)
// ForwardedByClientIP '为' true '
// ' TrustedProxies '功能是默认启用的，它也默认信任所有代理
// 如果您想禁用此功能，请使用Engine.SetTrustedProxies(nil)，然后Context.ClientIP()将直接返回远程地址

// ff:
// trustedProxies:
func (engine *Engine) SetTrustedProxies(trustedProxies []string) error {
	engine.trustedProxies = trustedProxies
	return engine.parseTrustedProxies()
}

// isUnsafeTrustedProxies检查引擎
// trustedCIDRs包含了所有的ip地址，如果有，则不安全(返回true)
func (engine *Engine) isUnsafeTrustedProxies() bool {
	return engine.isTrustedProxy(net.ParseIP("0.0.0.0")) || engine.isTrustedProxy(net.ParseIP("::"))
}

// parseTrustedProxies解析引擎
// trustedproxy to engine . trustedidrs
func (engine *Engine) parseTrustedProxies() error {
	trustedCIDRs, err := engine.prepareTrustedCIDRs()
	engine.trustedCIDRs = trustedCIDRs
	return err
}

// isTrustedProxy会根据Engine.trustedCIDRs检查IP地址是否在可信列表中
func (engine *Engine) isTrustedProxy(ip net.IP) bool {
	if engine.trustedCIDRs == nil {
		return false
	}
	for _, cidr := range engine.trustedCIDRs {
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}

// validateHeader将解析X-Forwarded-For报头并返回受信任的客户端IP地址
func (engine *Engine) validateHeader(header string) (clientIP string, valid bool) {
	if header == "" {
		return "", false
	}
	items := strings.Split(header, ",")
	for i := len(items) - 1; i >= 0; i-- {
		ipStr := strings.TrimSpace(items[i])
		ip := net.ParseIP(ipStr)
		if ip == nil {
			break
		}

// 以相反的顺序检查ip，当发现不受信任的代理时停止
		if (i == 0) || (!engine.isTrustedProxy(ip)) {
			return ipStr, true
		}
	}
	return "", false
}

// 解析IP的字符串表示形式并返回一个net
// 具有最小字节表示的IP，如果输入无效则为nil
func parseIP(ip string) net.IP {
	parsedIP := net.ParseIP(ip)

	if ipv4 := parsedIP.To4(); ipv4 != nil {
// 返回4字节表示的IP
		return ipv4
	}

// 返回16字节表示形式的IP或nil
	return parsedIP
}

// RunTLS将路由器附加到http
// 服务器并开始监听和服务HTTPS(安全)请求
// 它是http的快捷方式
// 注意:除非发生错误，否则此方法将无限期地阻塞调用例程

// ff:监听TLS
// err:错误
// keyFile:key文件
// certFile:cert文件
// addr:地址与端口
func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error) {
	debugPrint("Listening and serving HTTPS on %s\n", addr)
	defer func() { debugPrintError(err) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.")
	}

	err = http.ListenAndServeTLS(addr, certFile, keyFile, engine.Handler())
	return
}

// RunUnix将路由器附加到http
// 服务器并通过指定的unix套接字(即文件)开始侦听和服务HTTP请求
// 注意:除非发生错误，否则此方法将无限期地阻塞调用例程

// ff:
// err:
// file:
func (engine *Engine) RunUnix(file string) (err error) {
	debugPrint("Listening and serving HTTP on unix:/%s", file)
	defer func() { debugPrintError(err) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.")
	}

	listener, err := net.Listen("unix", file)
	if err != nil {
		return
	}
	defer listener.Close()
	defer os.Remove(file)

	err = http.Serve(listener, engine.Handler())
	return
}

// RunFd将路由器附加到http
// 服务器并通过指定的文件描述符开始侦听和服务HTTP请求
// 注意:除非发生错误，否则此方法将无限期地阻塞调用例程

// ff:
// err:
// fd:
func (engine *Engine) RunFd(fd int) (err error) {
	debugPrint("Listening and serving HTTP on fd@%d", fd)
	defer func() { debugPrintError(err) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.")
	}

	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd@%d", fd))
	listener, err := net.FileListener(f)
	if err != nil {
		return
	}
	defer listener.Close()
	err = engine.RunListener(listener)
	return
}

// RunListener将路由器附加到http
// 服务器并开始通过指定的网络侦听和服务HTTP请求
// 侦听器

// ff:
// err:
// listener:
func (engine *Engine) RunListener(listener net.Listener) (err error) {
	debugPrint("Listening and serving HTTP on listener what's bind with address@%s", listener.Addr())
	defer func() { debugPrintError(err) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.")
	}

	err = http.Serve(listener, engine.Handler())
	return
}

// ServeHTTP符合http
// 处理程序接口

// ff:
// req:
// w:
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

// HandleContext重新进入一个已经重写的上下文
// 这可以通过将c.Request.URL.Path设置为新目标来实现
// 免责声明:你可以循环自己来处理这个问题，明智地使用

// ff:
// c:
func (engine *Engine) HandleContext(c *Context) {
	oldIndexValue := c.index
	c.reset()
	engine.handleHTTPRequest(c)

	c.index = oldIndexValue
}

func (engine *Engine) handleHTTPRequest(c *Context) {
	httpMethod := c.Request.Method
	rPath := c.Request.URL.Path
	unescape := false
	if engine.UseRawPath && len(c.Request.URL.RawPath) > 0 {
		rPath = c.Request.URL.RawPath
		unescape = engine.UnescapePathValues
	}

	if engine.RemoveExtraSlash {
		rPath = cleanPath(rPath)
	}

// 查找给定HTTP方法的树的根
	t := engine.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
// 在树中查找路由
		value := root.getValue(rPath, c.params, c.skippedNodes, unescape)
		if value.params != nil {
			c.Params = *value.params
		}
		if value.handlers != nil {
			c.handlers = value.handlers
			c.fullPath = value.fullPath
			c.Next()
			c.writermem.WriteHeaderNow()
			return
		}
		if httpMethod != http.MethodConnect && rPath != "/" {
			if value.tsr && engine.RedirectTrailingSlash {
				redirectTrailingSlash(c)
				return
			}
			if engine.RedirectFixedPath && redirectFixedPath(c, root, engine.RedirectFixedPath) {
				return
			}
		}
		break
	}

	if engine.HandleMethodNotAllowed {
		for _, tree := range engine.trees {
			if tree.method == httpMethod {
				continue
			}
			if value := tree.root.getValue(rPath, nil, c.skippedNodes, unescape); value.handlers != nil {
				c.handlers = engine.allNoMethod
				serveError(c, http.StatusMethodNotAllowed, default405Body)
				return
			}
		}
	}
	c.handlers = engine.allNoRoute
	serveError(c, http.StatusNotFound, default404Body)
}

var mimePlain = []string{MIMEPlain}

func serveError(c *Context, code int, defaultMessage []byte) {
	c.writermem.status = code
	c.Next()
	if c.writermem.Written() {
		return
	}
	if c.writermem.Status() == code {
		c.writermem.Header()["Content-Type"] = mimePlain
		_, err := c.Writer.Write(defaultMessage)
		if err != nil {
			debugPrint("cannot write message to writer during serve error: %v", err)
		}
		return
	}
	c.writermem.WriteHeaderNow()
}

func redirectTrailingSlash(c *Context) {
	req := c.Request
	p := req.URL.Path
	if prefix := path.Clean(c.Request.Header.Get("X-Forwarded-Prefix")); prefix != "." {
		prefix = regSafePrefix.ReplaceAllString(prefix, "")
		prefix = regRemoveRepeatedChar.ReplaceAllString(prefix, "/")

		p = prefix + "/" + req.URL.Path
	}
	req.URL.Path = p + "/"
	if length := len(p); length > 1 && p[length-1] == '/' {
		req.URL.Path = p[:length-1]
	}
	redirectRequest(c)
}

func redirectFixedPath(c *Context, root *node, trailingSlash bool) bool {
	req := c.Request
	rPath := req.URL.Path

	if fixedPath, ok := root.findCaseInsensitivePath(cleanPath(rPath), trailingSlash); ok {
		req.URL.Path = bytesconv.BytesToString(fixedPath)
		redirectRequest(c)
		return true
	}
	return false
}

func redirectRequest(c *Context) {
	req := c.Request
	rPath := req.URL.Path
	rURL := req.URL.String()

	code := http.StatusMovedPermanently // 永久重定向，请求使用GET方法
	if req.Method != http.MethodGet {
		code = http.StatusTemporaryRedirect
	}
	debugPrint("redirecting request %d: %s --> %s", code, rPath, rURL)
	http.Redirect(c.Writer, req, rURL, code)
	c.writermem.WriteHeaderNow()
}
