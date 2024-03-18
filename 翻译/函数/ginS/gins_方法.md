# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 
# **_package.md 文件备注:
# bm= 包名,更换新的包名称, 如: package gin //bm:gin类
#
# **_其他.md 文件备注:
# hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
#     但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"
# zz= 正则表达式,用于结构名称替换或者复杂替换
#     如待替换: type authPair struct { //zz:^type *authPair

[func Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
handlers=处理函数
relativePath=路由规则
httpMethod=HTTP方法
ff=绑定

[func POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定POST
handlers=处理函数
relativePath=路由规则

[func GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定GET
handlers=处理函数
relativePath=路由规则

[func DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定DELETE
handlers=处理函数
relativePath=路由规则

[func PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定PATCH
handlers=处理函数
relativePath=路由规则

[func PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定PUT
handlers=处理函数
relativePath=路由规则

[func OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定OPTIONS
handlers=处理函数
relativePath=路由规则

[func HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定HEAD
handlers=处理函数
relativePath=路由规则

[func Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {]
ff=绑定Any
handlers=处理函数
relativePath=路由规则

[func Use(middlewares ...gin.HandlerFunc) gin.IRoutes {]
ff=中间件
