
<原文开始>
	// if Allow DirectoryIndex
	// r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
	// set prefix
	// r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))
<原文结束>

# <翻译开始>
// 如果允许目录索引
// r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
// 设置前缀
// r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))
// 翻译成中文：
// 如果允许目录索引功能
// 使用 Serve 函数，将根路径 "/" 映射到本地文件系统中 "/tmp" 目录，并开启目录索引
// r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
// 设置资源前缀为 "/static"
// 使用 Serve 函数，将 "/static" 路径映射到本地文件系统中 "/tmp" 目录，并开启目录索引
// r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))
# <翻译结束>


<原文开始>
	// Listen and Server in 0.0.0.0:8080
<原文结束>

# <翻译开始>
// 在0.0.0.0:8080监听并服务
# <翻译结束>

