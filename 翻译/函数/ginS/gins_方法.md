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

[func LoadHTMLGlob(pattern string) {]
ff=加载HTML模板目录
pattern=模板目录

[func LoadHTMLFiles(files ...string) {]
ff=加载HTML模板文件
files=模板文件s

[func SetHTMLTemplate(templ *template.Template) {]
ff=设置Template模板
templ=Template模板

[func NoRoute(handlers ...gin.HandlerFunc) {]
ff=绑定404
handlers=处理函数s

[func NoMethod(handlers ...gin.HandlerFunc) {]
ff=绑定405
handlers=处理函数s

[func Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {]
ff=创建分组路由
handlers=处理函数s
relativePath=路由规则

[func Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定
handlers=处理函数s
relativePath=路由规则
httpMethod=HTTP方法

[func POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定POST
handlers=处理函数s
relativePath=路由规则

[func GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定GET
handlers=处理函数s
relativePath=路由规则

[func DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定DELETE
handlers=处理函数s
relativePath=路由规则

[func PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定PATCH
handlers=处理函数s
relativePath=路由规则

[func PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定PUT
handlers=处理函数s
relativePath=路由规则

[func OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定OPTIONS
handlers=处理函数s
relativePath=路由规则

[func HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定HEAD
handlers=处理函数s
relativePath=路由规则

[func Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定Any
handlers=处理函数s
relativePath=路由规则

[func StaticFile(relativePath, filepath string) gin.IRoutes {]
ff=绑定静态单文件
filepath=文件路径
relativePath=URL路径

[func Static(relativePath, root string) gin.IRoutes {]
ff=绑定静态文件目录
root=绑定目录
relativePath=URL路径前缀

[func StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {]
ff=绑定静态文件目录FS
relativePath=URL路径前缀

[func Use(middlewares ...gin.HandlerFunc) gin.IRoutes {]
ff=中间件

[func Routes() gin.RoutesInfo {]
ff=取路由切片

[func Run(addr ...string) (err error) {]
ff=监听
err=错误
addr=地址与端口

[func RunTLS(addr, certFile, keyFile string) (err error) {]
ff=监听TLS
err=错误
keyFile=key文件
certFile=cert文件
addr=地址与端口

[func RunUnix(file string) (err error) {]
ff=监听Unix
err=错误
file=文件路径

[func RunFd(fd int) (err error) {]
ff=监听Fd
err=错误
