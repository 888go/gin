
<原文开始>
// set ins property from opts
<原文结束>

# <翻译开始>
// 从opts设置ins属性
# <翻译结束>


<原文开始>
// 	if bundle isn't constructed then assign it from default
<原文结束>

# <翻译开始>
// 如果bundle未构建，则从默认值中分配它
# <翻译结束>


<原文开始>
// if getLngHandler isn't constructed then assign it from default
<原文结束>

# <翻译开始>
// 如果getLngHandler尚未构建，则从默认值中赋值
# <翻译结束>


<原文开始>
// GetMessage get the i18n message with error handling
// param is one of these type: messageID, *i18n.LocalizeConfig
// Example:
// GetMessage(context, "hello") // messageID is hello
//
//	GetMessage(context, &i18n.LocalizeConfig{
//	  MessageID: "welcomeWithName",
//	  TemplateData: map[string]string{
//	    "name": context.Param("name"),
//	  },
//	})
<原文结束>

# <翻译开始>
// GetMessage 获取带有错误处理功能的国际化消息
// 参数可以是以下类型之一：messageID, *i18n.LocalizeConfig
// 示例：
// GetMessage(context, "hello") // messageID 为 "hello"
//
// GetMessage(context, &i18n.LocalizeConfig{
//   MessageID: "welcomeWithName", // 消息ID为 "welcomeWithName"
//   TemplateData: map[string]string{ // 模板数据
//     "name": context.Param("name"), // 将 "name" 参数值注入模板
//   },
// })
# <翻译结束>


<原文开始>
// MustGetMessage get the i18n message without error handling
// param is one of these type: messageID, *i18n.LocalizeConfig
// Example:
// MustGetMessage(context, "hello") // messageID is hello
//
//	MustGetMessage(context, &i18n.LocalizeConfig{
//	  MessageID: "welcomeWithName",
//	  TemplateData: map[string]string{
//	    "name": context.Param("name"),
//	  },
//	})
<原文结束>

# <翻译开始>
// MustGetMessage 获取i18n消息，不进行错误处理
// 参数可以是以下类型之一：messageID, *i18n.LocalizeConfig
// 示例：
// MustGetMessage(context, "hello") // messageID 为 "hello"
//
// MustGetMessage(context, &i18n.LocalizeConfig{
//   MessageID: "welcomeWithName",
//   TemplateData: map[string]string{
//     "name": context.Param("name"),
//   },
// })
# <翻译结束>

