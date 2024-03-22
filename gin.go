// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

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

const defaultMultipartMemory = 32 << 20 // 32 MB

var (
	default404Body = []byte("404 page not found")
	default405Body = []byte("405 method not allowed")
)

var defaultPlatform string

var defaultTrustedCIDRs = []*net.IPNet{
	{ // 0.0.0.0/0 (IPv4)
		IP:   net.IP{0x0, 0x0, 0x0, 0x0},
		Mask: net.IPMask{0x0, 0x0, 0x0, 0x0},
	},
	{ // ::/0 (IPv6)
		IP:   net.IP{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Mask: net.IPMask{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
}

var regSafePrefix = regexp.MustCompile("[^a-zA-Z0-9/-]+")
var regRemoveRepeatedChar = regexp.MustCompile("/{2,}")

// HandlerFunc 定义了 Gin 中间件用作返回值的处理器。
type HandlerFunc func(*Context)

// HandlersChain 定义了一个 HandlerFunc 切片。
type HandlersChain []HandlerFunc

// Last 返回链中的最后一个处理器。即，最后一个处理器是主处理器。
func (c HandlersChain) X取最后一个处理函数() HandlerFunc {
	if length := len(c); length > 0 {
		return c[length-1]
	}
	return nil
}

// RouteInfo 表示请求路由的规范，其中包含方法、路径及其处理程序。
type RouteInfo struct {
	X方法      string
	X路径        string
	Handler     string
	X处理函数 HandlerFunc
}

// RoutesInfo 定义了一个 RouteInfo 切片。
type RoutesInfo []RouteInfo

// Trusted platforms
const (
// PlatformGoogleAppEngine：当在 Google App Engine 上运行时。信任 X-Appengine-Remote-Addr 头部来确定客户端的 IP 地址
	PlatformGoogleAppEngine = "X-Appengine-Remote-Addr"
// PlatformCloudflare 当使用Cloudflare的CDN时。信任CF-Connecting-IP来确定
// 客户端的IP地址
	PlatformCloudflare = "CF-Connecting-IP"
)

// Engine 是框架的实例，它包含了多路复用器（muxer）、中间件和配置设置。
// 通过使用 New() 或 Default() 创建 Engine 的一个实例。
type Engine struct {
	RouterGroup

// RedirectTrailingSlash 功能会自动重定向，当当前路由无法匹配，但存在一个与请求路径（有或无尾部斜杠）相匹配的处理程序时。
// 例如，如果请求了 /foo/，但仅存在 /foo 的路由，则客户端将被重定向到 /foo，并根据请求方法的不同返回不同的HTTP状态码：对于GET请求返回301，对于所有其他请求方法返回307。
	X重定向尾部斜杠 bool

// RedirectFixedPath：如果启用，当没有为当前请求路径注册处理程序时，路由器尝试修复该路径。
// 首先移除诸如 ../ 或 // 等多余的路径元素。
// 然后，路由器对清理后的路径进行不区分大小写的查找。
// 如果能找到与此路由匹配的处理程序，路由器将根据请求方法进行重定向：
// 对于 GET 请求，状态码为 301；对于所有其他请求方法，状态码为 307。
// 例如，/FOO 和 /..//Foo 可能会被重定向到 /foo。
// 该选项与 RedirectTrailingSlash 选项独立。
	X重定向固定路径 bool

// 如果启用HandleMethodNotAllowed，当当前请求无法被路由时，路由器会检查当前路由是否允许其他方法。
// 如果存在其他允许的方法，请求将得到响应'方法不允许'（Method Not Allowed）以及HTTP状态码405。
// 若没有其他方法被允许，则该请求会被转发至NotFound处理器进行处理。
	HandleMethodNotAllowed bool

// ForwardedByClientIP：如果启用，将会从请求头中解析客户端IP地址，这些请求头与存储在 `(*gin.Engine).RemoteIPHeaders` 中的相匹配。如果没有获取到IP地址，则会回退到通过 `(*gin.Context).Request.RemoteAddr` 获取的IP地址。
	ForwardedByClientIP bool

// AppEngine 已被弃用。
// 废弃: 请改用 `TrustedPlatform`，并设置其值为 `gin.PlatformGoogleAppEngine`
// #726 #755 如果启用，将会信任以 'X-AppEngine...' 开头的一些头部信息，
// 以便更好地与该 PaaS（平台即服务）进行集成。
	AppEngine弃用 bool

	// 如果启用UseRawPath，将使用url.RawPath来查找参数。
	X使用原始路径 bool

// UnescapePathValues 如果设为 true，路径值将被解码。
// 若 UseRawPath 为 false（默认情况），则 UnescapePathValues 实际上等同于 true，
// 因为此时会使用已经解码过的 url.Path。
	UnescapePathValues bool

// RemoveExtraSlash：即使存在额外的斜杠，参数也可以从URL中解析出来。
// 参见PR #1817和问题#1644
	X删除多余斜杠 bool

// RemoteIPHeaders 是一个头部列表，当 `(*gin.Engine).ForwardedByClientIP` 设置为 `true` 时，
// 如果 `(*gin.Context).Request.RemoteAddr` 与通过 `(*gin.Engine).SetTrustedProxies()` 方法定义的网络源列表中的至少一个匹配，
// 则会使用这些头部来获取客户端 IP 地址。
	RemoteIPHeaders []string

// TrustedPlatform 如果设置为gin.Platform*类型的常量值，表示信任该平台设置的头部信息，
// 例如用于确定客户端IP地址
	TrustedPlatform string

// MaxMultipartMemory 是提供给 http.Request 的 ParseMultipartForm 方法调用时的 'maxMemory' 参数的值。
	X最大Multipart内存 int64

	// UseH2C 启用 h2c 支持。
	X启用h2c支持 bool

	// ContextWithFallback 用于当 Context.Request.Context() 不为空时，启用备用的 Context.Deadline()，Context.Done()，Context.Err() 和 Context.Value() 方法。
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

// New 函数返回一个全新的、未附加任何中间件的空白 Engine 实例。
// 默认配置为：
// - RedirectTrailingSlash:  true （自动重定向末尾的斜杠）
// - RedirectFixedPath:      false （不进行固定路径重定向）
// - HandleMethodNotAllowed: false （不处理不允许的方法）
// - ForwardedByClientIP:    true （通过客户端 IP 转发请求头）
// - UseRawPath:             false （不使用原始路径，即不做 URL 解码）
// - UnescapePathValues:     true （对路径中的参数值进行解码）
func X创建() *Engine {
	debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		FuncMap:                template.FuncMap{},
		X重定向尾部斜杠:  true,
		X重定向固定路径:      false,
		HandleMethodNotAllowed: false,
		ForwardedByClientIP:    true,
		RemoteIPHeaders:        []string{"X-Forwarded-For", "X-Real-IP"},
		TrustedPlatform:        defaultPlatform,
		X使用原始路径:             false,
		X删除多余斜杠:       false,
		UnescapePathValues:     true,
		X最大Multipart内存:     defaultMultipartMemory,
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

// Default 返回一个已附加了 Logger 和 Recovery 中间件的 Engine 实例。
func X创建默认对象() *Engine {
	debugPrintWARNINGDefault()
	engine := X创建()
	engine.X中间件(Logger(), Recovery())
	return engine
}

func (engine *Engine) X取主处理程序() http.Handler {
	if !engine.X启用h2c支持 {
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

// Delims 设置模板左右分隔符，并返回一个 Engine 实例。
func (engine *Engine) X设置模板分隔符(左边, 右边 string) *Engine {
	engine.delims = render.Delims{Left: 左边, Right: 右边}
	return engine
}

// SecureJsonPrefix 设置在 Context.SecureJSON 中使用的 secureJSONPrefix。
func (engine *Engine) X设置Json防劫持前缀(防劫持前缀 string) *Engine {
	engine.secureJSONPrefix = 防劫持前缀
	return engine
}

// LoadHTMLGlob 通过 glob 模式加载 HTML 文件，并将结果与 HTML 渲染器关联。
func (engine *Engine) X加载HTML模板目录(模板目录 string) {
	left := engine.delims.Left
	right := engine.delims.Right
	templ := template.Must(template.New("").Delims(left, right).Funcs(engine.FuncMap).ParseGlob(模板目录))

	if X是否为调试模式() {
		debugPrintLoadTemplate(templ)
		engine.HTMLRender = render.HTMLDebug{Glob: 模板目录, FuncMap: engine.FuncMap, Delims: engine.delims}
		return
	}

	engine.X设置Template模板(templ)
}

// LoadHTMLFiles 加载一组 HTML 文件
// 并将结果与 HTML 渲染器关联。
func (engine *Engine) X加载HTML模板文件(模板文件s ...string) {
	if X是否为调试模式() {
		engine.HTMLRender = render.HTMLDebug{Files: 模板文件s, FuncMap: engine.FuncMap, Delims: engine.delims}
		return
	}

	templ := template.Must(template.New("").Delims(engine.delims.Left, engine.delims.Right).Funcs(engine.FuncMap).ParseFiles(模板文件s...))
	engine.X设置Template模板(templ)
}

// SetHTMLTemplate 将一个模板与HTML渲染器关联。
func (engine *Engine) X设置Template模板(Template模板 *template.Template) {
	if len(engine.trees) > 0 {
		debugPrintWARNINGSetHTMLTemplate()
	}

	engine.HTMLRender = render.HTMLProduction{Template: Template模板.Funcs(engine.FuncMap)}
}

// SetFuncMap 用于设置用于 template.FuncMap 的 FuncMap。
func (engine *Engine) X设置Template模板函数(函数Map template.FuncMap) {
	engine.FuncMap = 函数Map
}

// NoRoute 添加处理函数，用于未找到路由的情况（NoRoute）。默认情况下返回404状态码。
func (engine *Engine) X绑定404(处理函数s ...HandlerFunc) {
	engine.noRoute = 处理函数s
	engine.rebuild404Handlers()
}

// NoMethod 设置在 Engine.HandleMethodNotAllowed = true 时调用的处理器。
func (engine *Engine) X绑定405(处理函数s ...HandlerFunc) {
	engine.noMethod = 处理函数s
	engine.rebuild405Handlers()
}

// Use 方法将一个全局中间件附加到路由。也就是说，通过Use()方法附加的中间件将会
// 被包含在每一个请求的处理器链中。即便是404、405错误、静态文件等请求...
// 例如，这里适合放置日志记录器或错误管理中间件。
func (engine *Engine) X中间件(处理函数 ...HandlerFunc) IRoutes {
	engine.RouterGroup.X中间件(处理函数...)
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

// Routes 返回已注册路由的切片，其中包括一些有用的信息，比如：
// HTTP 方法、路径以及处理器名称。
func (engine *Engine) X取路由数组() (路由s RoutesInfo) {
	for _, tree := range engine.trees {
		路由s = iterate("", tree.method, 路由s, tree.root)
	}
	return 路由s
}

func iterate(path, method string, routes RoutesInfo, root *node) RoutesInfo {
	path += root.path
	if len(root.handlers) > 0 {
		handlerFunc := root.handlers.X取最后一个处理函数()
		routes = append(routes, RouteInfo{
			X方法:      method,
			X路径:        path,
			Handler:     nameOfFunction(handlerFunc),
			X处理函数: handlerFunc,
		})
	}
	for _, child := range root.children {
		routes = iterate(path, method, routes, child)
	}
	return routes
}

// Run 将路由器附加到 http.Server，并开始监听和处理 HTTP 请求。
// 这是 http.ListenAndServe(addr, router) 的一个快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的 goroutine。
func (engine *Engine) X监听(地址与端口 ...string) (错误 error) {
	defer func() { debugPrintError(错误) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.")
	}

	address := resolveAddress(地址与端口)
	debugPrint("Listening and serving HTTP on %s\n", address)
	错误 = http.ListenAndServe(address, engine.X取主处理程序())
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

// SetTrustedProxies 设置一个网络源列表（IPv4地址、IPv4 CIDR、IPv6地址或IPv6 CIDR），从这些源中信任请求头中包含的替代客户端IP。当`(*gin.Engine).ForwardedByClientIP`为`true`时生效。`TrustedProxies`特性默认启用，并且默认情况下信任所有代理。如果你想禁用此功能，使用Engine.SetTrustedProxies(nil)，那么Context.ClientIP()将直接返回远程地址。
func (engine *Engine) X设置受信任代理(受信任代理 []string) error {
	engine.trustedProxies = 受信任代理
	return engine.parseTrustedProxies()
}

// isUnsafeTrustedProxies 检查 Engine.trustedCIDRs 是否包含全部IP，如果包含（返回 true），则表示不安全
func (engine *Engine) isUnsafeTrustedProxies() bool {
	return engine.isTrustedProxy(net.ParseIP("0.0.0.0")) || engine.isTrustedProxy(net.ParseIP("::"))
}

// parseTrustedProxies 将 Engine.trustedProxies 解析为 Engine.trustedCIDRs
func (engine *Engine) parseTrustedProxies() error {
	trustedCIDRs, err := engine.prepareTrustedCIDRs()
	engine.trustedCIDRs = trustedCIDRs
	return err
}

// isTrustedProxy 将根据 Engine.trustedCIDRs 检查 IP 地址是否在信任列表中
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

// validateHeader 将解析 X-Forwarded-For 头部，并返回可信的客户端 IP 地址
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

// X-Forwarded-For 由代理服务器追加
// 按照逆序检查 IP 地址，并在找到不可信的代理时停止
		if (i == 0) || (!engine.isTrustedProxy(ip)) {
			return ipStr, true
		}
	}
	return "", false
}

// parseIP 将IP地址的字符串表示形式解析为 net.IP 类型，并返回一个字节表示形式最小的 IP，如果输入无效，则返回 nil。
func parseIP(ip string) net.IP {
	parsedIP := net.ParseIP(ip)

	if ipv4 := parsedIP.To4(); ipv4 != nil {
		// 返回一个4字节表示的IP地址
		return ipv4
	}

	// 返回一个16字节表示形式的IP地址，或返回nil
	return parsedIP
}

// RunTLS 将路由器附加到 http.Server，并开始监听和处理 HTTPS（安全）请求。
// 这是 http.ListenAndServeTLS(addr, certFile, keyFile, router) 的快捷方式。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的goroutine。
func (engine *Engine) X监听TLS(地址与端口, cert文件, key文件 string) (错误 error) {
	debugPrint("Listening and serving HTTPS on %s\n", 地址与端口)
	defer func() { debugPrintError(错误) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.")
	}

	错误 = http.ListenAndServeTLS(地址与端口, cert文件, key文件, engine.X取主处理程序())
	return
}

// RunUnix将路由器连接到http.Server，并开始通过指定的UNIX套接字（即文件）监听和处理HTTP请求。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用的goroutine。
func (engine *Engine) X监听Unix(文件路径 string) (错误 error) {
	debugPrint("Listening and serving HTTP on unix:/%s", 文件路径)
	defer func() { debugPrintError(错误) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.")
	}

	listener, 错误 := net.Listen("unix", 文件路径)
	if 错误 != nil {
		return
	}
	defer listener.Close()
	defer os.Remove(文件路径)

	错误 = http.Serve(listener, engine.X取主处理程序())
	return
}

// RunFd 将路由器连接到 http.Server，并开始通过指定的文件描述符监听和处理 HTTP 请求。
// 注意：除非发生错误，否则此方法将无限期地阻塞调用它的 goroutine。
func (engine *Engine) X监听Fd(fd int) (错误 error) {
	debugPrint("Listening and serving HTTP on fd@%d", fd)
	defer func() { debugPrintError(错误) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.")
	}

	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd@%d", fd))
	listener, 错误 := net.FileListener(f)
	if 错误 != nil {
		return
	}
	defer listener.Close()
	错误 = engine.X监听Listener(listener)
	return
}

// RunListener 将路由器附加到 http.Server，并开始通过指定的 net.Listener 监听和处理 HTTP 请求
func (engine *Engine) X监听Listener(listener net.Listener) (错误 error) {
	debugPrint("Listening and serving HTTP on listener what's bind with address@%s", listener.Addr())
	defer func() { debugPrintError(错误) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.")
	}

	错误 = http.Serve(listener, engine.X取主处理程序())
	return
}

// ServeHTTP 符合 http.Handler 接口。
//
// 注意!!! 此方法不能翻译, 因为是http包的接口实现
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	c.writermem.reset(w)
	c.X请求 = req
	c.reset()

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

// HandleContext 该方法会重新载入一个被重写的context(可以通过c.Request.URL.Path来实现).
//
// 注意:该方法可能造成context的循环使用(会绕死你,谨慎使用)
func (engine *Engine) HandleContext底层方法(c *Context) {
	oldIndexValue := c.index
	c.reset()
	engine.handleHTTPRequest(c)

	c.index = oldIndexValue
}

func (engine *Engine) handleHTTPRequest(c *Context) {
	httpMethod := c.X请求.Method
	rPath := c.X请求.URL.Path
	unescape := false
	if engine.X使用原始路径 && len(c.X请求.URL.RawPath) > 0 {
		rPath = c.X请求.URL.RawPath
		unescape = engine.UnescapePathValues
	}

	if engine.X删除多余斜杠 {
		rPath = cleanPath(rPath)
	}

	// 为给定的HTTP方法查找树的根节点
	t := engine.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		// Find route in tree
		value := root.getValue(rPath, c.params, c.skippedNodes, unescape)
		if value.params != nil {
			c.X参数 = *value.params
		}
		if value.handlers != nil {
			c.handlers = value.handlers
			c.fullPath = value.fullPath
			c.X中间件继续()
			c.writermem.WriteHeaderNow()
			return
		}
		if httpMethod != http.MethodConnect && rPath != "/" {
			if value.tsr && engine.X重定向尾部斜杠 {
				redirectTrailingSlash(c)
				return
			}
			if engine.X重定向固定路径 && redirectFixedPath(c, root, engine.X重定向固定路径) {
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
	c.X中间件继续()
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
	req := c.X请求
	p := req.URL.Path
	if prefix := path.Clean(c.X请求.Header.Get("X-Forwarded-Prefix")); prefix != "." {
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
	req := c.X请求
	rPath := req.URL.Path

	if fixedPath, ok := root.findCaseInsensitivePath(cleanPath(rPath), trailingSlash); ok {
		req.URL.Path = bytesconv.BytesToString(fixedPath)
		redirectRequest(c)
		return true
	}
	return false
}

func redirectRequest(c *Context) {
	req := c.X请求
	rPath := req.URL.Path
	rURL := req.URL.String()

	code := http.StatusMovedPermanently // 永久重定向，使用GET方法请求
	if req.Method != http.MethodGet {
		code = http.StatusTemporaryRedirect
	}
	debugPrint("redirecting request %d: %s --> %s", code, rPath, rURL)
	http.Redirect(c.Writer, req, rURL, code)
	c.writermem.WriteHeaderNow()
}
