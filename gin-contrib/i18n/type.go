package i18n

import (
	"github.com/888go/gin"
)

type (
// GetLngHandler ... 
// （由于这段注释不完整，无法准确提供详细的中文翻译。但从命名推测，这可能是Go语言中的一个函数或处理程序的定义，该函数用于获取某种“Lng”相关的数据或执行与“Lng”相关的操作，"Lng"可能代表经度、某种标识符或其他含义，具体取决于上下文。）
// 获取Lng处理器 ...
	GetLngHandler = func(context *gin.Context, defaultLng string) string

// Option ... (待补充，该注释不完整，无法准确翻译)
// 在 Go 语言中，通常情况下 `// Option` 是用于表示一个可选项或者配置项的注释前缀。由于您提供的代码注释不完整，请提供更多的上下文信息以便进行准确翻译。例如：
// ```go
// Option 设置函数，用于设置某个结构体或对象的相关配置项
// type Option func(*SomeType)
// 或者
// Option 类型是定义配置选项的函数类型
// type Option func(config *Config) error
// 以上是对 `Option` 的两种可能含义的翻译，具体情况需要根据实际代码来判断。
	Option func(GinI18n)
)
