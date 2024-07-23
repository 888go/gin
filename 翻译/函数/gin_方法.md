# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func (c HandlersChain) Last() HandlerFunc {]
ff=取最后一个处理函数

[func New() *Engine {]
ff=创建

[func Default() *Engine {]
ff=创建默认对象

[func (engine *Engine) Handler() http.Handler {]
ff=取主处理程序

[func (engine *Engine) Delims(left, right string) *Engine {]
ff=设置模板分隔符
right=右边
left=左边

[func (engine *Engine) SecureJsonPrefix(prefix string) *Engine {]
ff=设置Json防劫持前缀
prefix=防劫持前缀

[func (engine *Engine) LoadHTMLGlob(pattern string) {]
ff=加载HTML模板目录
pattern=模板目录

[func (engine *Engine) LoadHTMLFiles(files ...string) {]
ff=加载HTML模板文件
files=模板文件s

[func (engine *Engine) SetHTMLTemplate(templ *template.Template) {]
ff=设置Template模板
templ=Template模板

[func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {]
ff=设置Template模板函数
funcMap=函数Map

[func (engine *Engine) NoRoute(handlers ...HandlerFunc) {]
ff=绑定404
handlers=处理函数s

[func (engine *Engine) NoMethod(handlers ...HandlerFunc) {]
ff=绑定405
handlers=处理函数s

[func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {]
ff=中间件
middleware=处理函数

[func (engine *Engine) Routes() (routes RoutesInfo) {]
ff=取路由切片
routes=路由s

[func (engine *Engine) Run(addr ...string) (err error) {]
ff=监听
err=错误
addr=地址与端口

[func (engine *Engine) SetTrustedProxies(trustedProxies #左中括号##右中括号#string) error {]
ff=设置受信任代理
trustedProxies=受信任代理

[func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error) {]
ff=监听TLS
err=错误
keyFile=key文件
certFile=cert文件
addr=地址与端口

[func (engine *Engine) RunUnix(file string) (err error) {]
ff=监听Unix
err=错误
file=文件路径

[func (engine *Engine) RunFd(fd int) (err error) {]
ff=监听Fd
err=错误

[func (engine *Engine) RunListener(listener net.Listener) (err error) {]
ff=监听Listener
err=错误

[func (engine *Engine) HandleContext(c *Context) {]
ff=HandleContext底层方法
