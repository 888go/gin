
<原文开始>
	// Disable Console Color
	// gin.DisableConsoleColor()
<原文结束>

# <翻译开始>
	// 禁用控制台颜色
	// gin.DisableConsoleColor()
# <翻译结束>


<原文开始>
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
<原文结束>

# <翻译开始>
	// 授权分组（使用gin.BasicAuth()中间件）
	// 等同于：
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar", 	// 用户名：密码
	//	  "manu": "123",
	// }))
	// 
	// 这段Go注释翻译成中文后的大致意思是：
	// 
	// 此处定义一个授权访问的分组，该分组将采用gin.BasicAuth()中间件进行身份验证。
	// 这与以下代码功能相同：
	// 首先创建一个名为authorized的路由分组，并将其根路径设置为"/"。
	// 然后在该分组中使用gin.BasicAuth()中间件进行基本认证，其中包含如下用户名和密码凭据：
	// 用户名 "foo" 对应的密码是 "bar"
	// 用户名 "manu" 对应的密码是 "123"
# <翻译结束>


<原文开始>
// Listen and Server in 0.0.0.0:8080
<原文结束>

# <翻译开始>
// 在0.0.0.0:8080监听并服务
# <翻译结束>

