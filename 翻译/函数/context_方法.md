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
#     如待替换: type authPair struct { //zz:^type *authPair处理函数

[func (c *Context) Copy() *Context {]
ff=取副本

[func (c *Context) HandlerName() string {]
ff=取主处理程序名称

[func (c *Context) HandlerNames() #左中括号##右中括号#string {]
ff=取处理程序数组

[func (c *Context) Handler() HandlerFunc {]
ff=取主处理程序

[func (c *Context) Next() {]
ff=中间件继续

[func (c *Context) Set(key string, value any) {]
ff=设置值
value=值
key=名称

[func (c *Context) Get(key string) (value any, exists bool) {]
ff=取值
exists=是否存在
value=返回值
key=名称

[func (c *Context) MustGet(key string) any {]
ff=取值PANI
key=名称

[func (c *Context) GetString(key string) (s string) {]
ff=取文本值
s=返回值
key=名称

[func (c *Context) GetBool(key string) (b bool) {]
ff=取布尔值
b=返回值
key=名称

[func (c *Context) GetInt(key string) (i int) {]
ff=取整数值
i=返回值
key=名称

[func (c *Context) GetInt64(key string) (i64 int64) {]
ff=取整数64位值
i64=返回值
key=名称

[func (c *Context) GetUint(key string) (ui uint) {]
ff=取正整数值
ui=返回值
key=名称

[func (c *Context) GetUint64(key string) (ui64 uint64) {]
ff=取正整数64位值
ui64=返回值
key=名称

[func (c *Context) GetFloat64(key string) (f64 float64) {]
ff=取小数64位值
f64=返回值
key=名称

[func (c *Context) GetTime(key string) (t time.Time) {]
ff=取时间值
t=返回值
key=名称

[func (c *Context) GetDuration(key string) (d time.Duration) {]
ff=取时长值
d=返回时长
key=名称

[func (c *Context) GetStringSlice(key string) (ss #左中括号##右中括号#string) {]
ff=取数组值
ss=返回数组
key=名称

[func (c *Context) GetStringMap(key string) (sm map#左中括号#string#右中括号#any) {]
ff=取Map值
sm=返回Map
key=名称

[func (c *Context) GetStringMapString(key string) (sms map#左中括号#string#右中括号#string) {]
ff=取文本Map值
sms=返回Map
key=名称

[func (c *Context) GetStringMapStringSlice(key string) (smss map#左中括号#string#右中括号##左中括号##右中括号#string) {]
ff=取数组Map值
smss=返回数组Map
key=名称

[func (c *Context) Param(key string) string {]
ff=取API参数值
key=名称

[func (c *Context) AddParam(key, value string) {]
ff=设置API参数值
value=值
key=名称

[func (c *Context) Query(key string) (value string) {]
ff=取URL参数值
value=返回值
key=名称

[func (c *Context) DefaultQuery(key, defaultValue string) string {]
ff=取URL参数值并带默认
defaultValue=默认值
key=名称

[func (c *Context) GetQuery(key string) (string, bool) {]
ff=取URL参数值2
key=名称

[func (c *Context) QueryArray(key string) (values #左中括号##右中括号#string) {]
ff=取URL参数数组值
values=返回数组
key=名称

[func (c *Context) GetQueryArray(key string) (values #左中括号##右中括号#string, ok bool) {]
ff=取URL参数数组值2
ok=是否存在
values=返回数组
key=名称

[func (c *Context) QueryMap(key string) (dicts map#左中括号#string#右中括号#string) {]
ff=取URL参数Map值
dicts=返回Map
key=名称

[func (c *Context) GetQueryMap(key string) (map#左中括号#string#右中括号#string, bool) {]
ff=取URL参数Map值2
key=名称

[func (c *Context) PostForm(key string) (value string) {]
ff=取表单参数值
value=返回值
key=名称

[func (c *Context) DefaultPostForm(key, defaultValue string) string {]
ff=取表单参数值并带默认
defaultValue=默认值
key=名称

[func (c *Context) GetPostForm(key string) (string, bool) {]
ff=取表单参数值2
key=名称

[func (c *Context) PostFormArray(key string) (values #左中括号##右中括号#string) {]
ff=取表单参数数组值
values=返回数组
key=名称

[func (c *Context) GetPostFormArray(key string) (values #左中括号##右中括号#string, ok bool) {]
ff=取参数数组值
ok=是否存在
values=返回数组
key=名称

[func (c *Context) PostFormMap(key string) (dicts map#左中括号#string#右中括号#string) {]
ff=取表单参数Map值
dicts=返回Map
key=名称

[func (c *Context) GetPostFormMap(key string) (map#左中括号#string#右中括号#string, bool) {]
ff=取参数Map值
key=名称

[func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {]
ff=取表单上传文件
name=名称

[func (c *Context) MultipartForm() (*multipart.Form, error) {]
ff=取表单multipart对象

[func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {]
ff=保存上传文件
dst=文件路径
file=文件对象

[func (c *Context) BindJSON(obj any) error {]
ff=取JSON参数到指针PANI
obj=结构指针

[func (c *Context) ClientIP() string {]
ff=取客户端ip

[func (c *Context) RemoteIP() string {]
ff=取协议头ip

[func (c *Context) ContentType() string {]
ff=取协议头ContentType

[func (c *Context) IsWebsocket() bool {]
ff=是否为Websocket请求

[func (c *Context) SetSameSite(samesite http.SameSite) {]
ff=设置cookie跨站

[func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {]
ff=设置cookie值
httpOnly=禁止js访问
secure=仅https生效
domain=域名
path=路径
maxAge=生效时间
value=值
name=名称

[func (c *Context) Cookie(name string) (string, error) {]
ff=取cookie值
name=名称

[func (c *Context) HTML(code int, name string, obj any) {]
ff=输出html模板
obj=结构
name=模板文件名
code=状态码

[func (c *Context) IndentedJSON(code int, obj any) {]
ff=输出JSON并美化
obj=结构
code=状态码

[func (c *Context) JSONP(code int, obj any) {]
ff=输出JSONP
obj=结构
code=状态码

[func (c *Context) JSON(code int, obj any) {]
ff=输出JSON
obj=结构
code=状态码

[func (c *Context) AsciiJSON(code int, obj any) {]
ff=输出JSON并按ASCII
obj=结构
code=状态码

[func (c *Context) PureJSON(code int, obj any) {]
ff=输出JSON并按原文
obj=结构
code=状态码

[func (c *Context) XML(code int, obj any) {]
ff=输出XML
obj=结构
code=状态码

[func (c *Context) YAML(code int, obj any) {]
ff=输出YAML
obj=结构
code=状态码

[func (c *Context) TOML(code int, obj any) {]
ff=输出TOML
obj=结构
code=状态码

[func (c *Context) String(code int, format string, values ...any) {]
ff=输出文本
values=文本s
format=格式
code=状态码

[func (c *Context) ShouldBind(obj any) error {]
ff=取参数到指针
obj=变量结构指针

[func (c *Context) ShouldBindJSON(obj any) error {]
ff=取JSON参数到指针
obj=JSON结构指针

[func (c *Context) ShouldBindXML(obj any) error {]
ff=取XML参数到指针
obj=XML结构指针

[func (c *Context) ShouldBindQuery(obj any) error {]
ff=取URL参数到指针
obj=表单结构指针

[func (c *Context) ShouldBindYAML(obj any) error {]
ff=取YAML参数到指针
obj=YAML结构指针

[func (c *Context) ShouldBindTOML(obj any) error {]
ff=取TOML参数到指针
obj=TOML结构指针

[func (c *Context) ShouldBindHeader(obj any) error {]
ff=取Header参数到指针
obj=Header结构指针

[func (c *Context) ShouldBindUri(obj any) error {]
ff=取Uri参数到指针
obj=Uri结构指针

[func (c *Context) ShouldBindWith(obj any, b binding.Binding) error {]
ff=取参数到指针并按类型
b=类型
obj=结构指针

[func (c *Context) Status(code int) {]
ff=设置状态码
code=状态码

[func (c *Context) Header(key, value string) {]
ff=设置响应协议头值
value=值
key=名称

[func (c *Context) GetHeader(key string) string {]
ff=取请求协议头值
key=名称

[func (c *Context) GetRawData() (#左中括号##右中括号#byte, error) {]
ff=取流数据

[func (c *Context) IsAborted() bool {]
ff=是否已停止

[func (c *Context) Bind(obj any) error {]
ff=取参数到指针PANI
obj=结构指针

[func (c *Context) BindXML(obj any) error {]
ff=取XML参数到指针PANI
obj=结构指针

[func (c *Context) BindQuery(obj any) error {]
ff=取URL参数到指针PANI
obj=结构指针

[func (c *Context) BindYAML(obj any) error {]
ff=取YAML参数到指针PANI
obj=结构指针

[func (c *Context) BindTOML(obj any) error {]
ff=取TOML参数到指针PANI
obj=结构指针

[func (c *Context) BindHeader(obj any) error {]
ff=取Header参数到指针PANI
obj=结构指针

[func (c *Context) BindUri(obj any) error {]
ff=取Uri参数到指针PANI
obj=结构指针

[func (c *Context) MustBindWith(obj any, b binding.Binding) error {]
ff=取参数到指针并按类型PANI
b=类型
obj=结构指针

[func (c *Context) SecureJSON(code int, obj any) {]
ff=输出JSON并防劫持
obj=结构
code=状态码

[func (c *Context) ProtoBuf(code int, obj any) {]
ff=输出ProtoBuf
obj=结构
code=状态码

[func (c *Context) ShouldBindBodyWith(obj any, bb binding.BindingBody) (err error) {]
err=错误
obj=结构指针
ff=取参数到指针并按类型且缓存

[func (c *Context) Redirect(code int, location string) {]
ff=重定向
location=重定向地址
code=状态码

[func (c *Context) FullPath() string {]
ff=取路由路径

[func (c *Context) Abort() {]
ff=停止

[func (c *Context) AbortWithStatus(code int) {]
ff=停止并带状态码
code=状态码

[func (c *Context) AbortWithStatusJSON(code int, jsonObj any) {]
ff=停止并带状态码且返回JSON
jsonObj=JSON结构
code=状态码

[func (c *Context) AbortWithError(code int, err error) *Error {]
ff=停止并带状态码与错误
err=错误
code=状态码

[func (c *Context) Error(err error) *Error {]
ff=错误
err=错误

[func (c *Context) Render(code int, r render.Render) {]
ff=Render底层方法
code=状态码

[func (c *Context) Data(code int, contentType string, data #左中括号##右中括号#byte) {]
ff=输出字节集
data=字节集
contentType=HTTP响应类型
code=状态码

[func (c *Context) DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map#左中括号#string#右中括号#string) {]
ff=输出字节集并按IO
extraHeaders=协议头Map
reader=写出IO数据
contentType=HTTP响应类型
contentLength=数据长度
code=状态码

[func (c *Context) File(filepath string) {]
ff=下载文件
filepath=文件路径

[func (c *Context) FileFromFS(filepath string, fs http.FileSystem) {]
ff=下载文件FS
filepath=文件路径

[func (c *Context) FileAttachment(filepath, filename string) {]
ff=下载文件并带文件名
filename=文件名
filepath=文件路径

[func (c *Context) Negotiate(code int, config Negotiate) {]
ff=Negotiate底层方法

[func (c *Context) NegotiateFormat(offered ...string) string {]
ff=NegotiateFormat底层方法

[func (c *Context) SetAccepted(formats ...string) {]
ff=SetAccepted底层方法
formats=类型名称s

 