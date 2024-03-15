package i18n

import (
	"github.com/888go/gin"
)

// newI18n ... 创建一个新的国际化（i18n）实例
func newI18n(opts ...Option) GinI18n {
// init 初始化
	ins := &ginI18nImpl{}

// 从opts设置ins属性
	for _, opt := range opts {
		opt(ins)
	}

// 如果bundle未构建，则从默认值中分配它
	if ins.bundle == nil {
		ins.setBundle(defaultBundleConfig)
	}

// 如果getLngHandler尚未构造，则从默认值中赋值
	if ins.getLngHandler == nil {
		ins.getLngHandler = defaultGetLngHandler
	}

	return ins
}

// Localize ... 本地化
func Localize(opts ...Option) gin.HandlerFunc {
	atI18n := newI18n(opts...)
	return func(context *gin.Context) {
		context.Set("i18n", atI18n)
	}
}

// GetMessage 获取具有错误处理功能的i18n消息
// 参数可以是以下类型之一：messageID, *i18n.LocalizeConfig
// 示例：
// GetMessage(context, "hello") // messageID 为 "hello"
//
//	GetMessage(context, &i18n.LocalizeConfig{
//	  MessageID: "welcomeWithName", // 消息ID为 "welcomeWithName"
//	  TemplateData: map[string]string{ // 模板数据，键值对形式
//	    "name": context.Param("name"), // 将 "name" 参数值注入模板中
//	  },
//	})
func GetMessage(context *gin.Context, param interface{}) (string, error) {
	atI18n := context.Value("i18n").(GinI18n)
	return atI18n.getMessage(context, param)
}

// MustGetMessage 必须获取i18n消息，无错误处理
// 参数可以是以下类型之一：messageID, *i18n.LocalizeConfig
// 示例：
// MustGetMessage(context, "hello") // messageID 为 "hello"
//
// MustGetMessage(context, &i18n.LocalizeConfig{
//   MessageID: "welcomeWithName", // 消息ID为 "welcomeWithName"
//   TemplateData: map[string]string{ // 模板数据为字符串到字符串的映射
//     "name": context.Param("name"), // 其中 "name" 键对应context中的参数值
//   },
// })
func MustGetMessage(context *gin.Context, param interface{}) string {
	atI18n := context.MustGet("i18n").(GinI18n)
	return atI18n.mustGetMessage(context, param)
}
