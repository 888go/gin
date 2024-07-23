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

[func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {]
ff=中间件
middleware=处理函数

[func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {]
ff=创建分组路由
handlers=处理函数
relativePath=路由规则

[func (group *RouterGroup) BasePath() string {]
ff=取路由基础路径

[func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {]
ff=绑定
handlers=处理函数
relativePath=路由规则
httpMethod=HTTP方法

[func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {]
ff=绑定POST
handlers=处理函数
relativePath=路由规则

[func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {]
ff=绑定GET
handlers=处理函数
relativePath=路由规则

[func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {]
ff=绑定DELETE
handlers=处理函数
relativePath=路由规则

[func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {]
ff=绑定PATCH
handlers=处理函数
relativePath=路由规则

[func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {]
ff=绑定PUT
handlers=处理函数
relativePath=路由规则

[func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes {]
ff=绑定OPTIONS
handlers=处理函数
relativePath=路由规则

[func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes {]
ff=绑定HEAD
handlers=处理函数
relativePath=路由规则

[func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes {]
ff=绑定Any
handlers=处理函数
relativePath=路由规则

[func (group *RouterGroup) StaticFile(relativePath, filepath string) IRoutes {]
ff=绑定静态单文件
filepath=文件路径
relativePath=URL路径

[func (group *RouterGroup) StaticFileFS(relativePath, filepath string, fs http.FileSystem) IRoutes {]
ff=绑定静态单文件FS
filepath=文件路径
relativePath=URL路径

[func (group *RouterGroup) Static(relativePath, root string) IRoutes {]
ff=绑定静态文件目录
root=绑定目录
relativePath=URL路径前缀

[func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes {]
ff=绑定静态文件目录FS
relativePath=URL路径前缀
