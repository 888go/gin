# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 
# **_package.md 文件备注:
# bm= 包名,更换新的包名称, 如: package gin //bm:gin类
#
# **_其他.md 文件备注:
# qm= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
#     但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"
# zz= 正则表达式,用于结构名称替换或者复杂替换
#     如待替换: type authPair struct { //zz:^type *authPair处理函数

[AppEngine bool]
qm=AppEngine弃用

[Method      string]
qm=方法

[Path        string]
qm=路径

[HandlerFunc HandlerFunc]
qm=处理函数

[RedirectTrailingSlash bool]
qm=重定向尾部斜杠

[RedirectFixedPath bool]
qm=重定向固定路径

[UseRawPath bool]
qm=使用原始路径

[RemoveExtraSlash bool]
qm=删除多余斜杠

[MaxMultipartMemory int64]
qm=最大Multipart内存

[UseH2C bool]
qm=启用h2c支持
