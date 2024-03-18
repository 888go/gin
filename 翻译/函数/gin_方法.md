提示:
ff= 方法,重命名方法名称
hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
bm= 包名,更换新的包名称
th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
    但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"


[func New() *Engine {]
ff=创建

[func Default() *Engine {]
ff=创建默认对象

[func (engine *Engine) Delims(left, right string) *Engine {]
ff=设置模板分隔符
right=右边
left=左边

[func (engine *Engine) LoadHTMLGlob(pattern string) {]
ff=加载HTML模板目录
pattern=模板目录

[func (engine *Engine) LoadHTMLFiles(files ...string) {]
ff=加载HTML模板文件
files=模板文件s

[func (engine *Engine) Routes() (routes RoutesInfo) {]
ff=取路由数组
routes=路由s

[func (engine *Engine) Run(addr ...string) (err error) {]
ff=监听
err=错误
addr=地址与端口

[func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error) {]
ff=监听TLS
err=错误
keyFile=key文件
certFile=cert文件
addr=地址与端口

[func (engine *Engine) SetHTMLTemplate(templ *template.Template) {]
ff=设置Template模板
templ=Template模板

[func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {]
ff=设置Template模板函数
funcMap=函数Map

[func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {]
ff=中间件
middleware=中间件函数
