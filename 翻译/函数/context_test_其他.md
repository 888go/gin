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



[assert.Regexp(t, "^(.*/vendor/)?github.com/gin-gonic/gin.handlerNameTest$", c.HandlerName())]
th=assert.Regexp(t, "^(.*/vendor/)?github.com/888go/gin.handlerNameTest$", c.HandlerName())
代码块=	assert.Regexp(t, "^(.*/vendor/)?github.com/gin-gonic/gin.handlerNameTest$", c.HandlerName()) //th:assert.Regexp(t, "^(.*/vendor/)?github.com/888go/gin.handlerNameTest$", c.HandlerName())               

[assert.Regexp(t, `^(.*/vendor/)?(github\.com/gin-gonic/gin\.){1}(TestContextHandlerNames\.func.*){0,1}(handlerNameTest.*){0,1}`, name)]
th=assert.Regexp(t, `^(.*/vendor/)?(github\.com/888go/gin\.){1}(TestContextHandlerNames\.func.*){0,1}(handlerNameTest.*){0,1}`, name)
代码块=		assert.Regexp(t, `^(.*/vendor/)?(github\.com/gin-gonic/gin\.){1}(TestContextHandlerNames\.func.*){0,1}(handlerNameTest.*){0,1}`, name) //th:assert.Regexp(t, `^(.*/vendor/)?(github\.com/888go/gin\.){1}(TestContextHandlerNames\.func.*){0,1}(handlerNameTest.*){0,1}`, name)               

[assert.Contains(t, w.Body.String(), "func New() *Engine {")]
th=assert.Contains(t, w.Body.String(), "func X创建() *Engine {")
代码块=	assert.Contains(t, w.Body.String(), "func New() *Engine {") //th:assert.Contains(t, w.Body.String(), "func X创建() *Engine {")
