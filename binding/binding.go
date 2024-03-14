// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

//go:build !nomsgpack

package binding


import (
	"net/http"
	)
// 内容类型MIME最常用的数据格式
const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
	MIMEYAML              = "application/x-yaml"
	MIMETOML              = "application/toml"
)

// 绑定描述了需要实现的接口，用于绑定请求中的数据，如JSON请求体、查询参数或表单POST
type Binding interface {
	Name() string
	Bind(*http.Request, any) error
}

// BindingBody在Binding中添加BindBody方法
// BindBody与Bind类似，但它从提供的字节中读取主体，而不是从req.Body中读取主体
type BindingBody interface {
	Binding
	BindBody([]byte, any) error
}

// BindingUri将BindUri方法添加到Binding中
// BindUri与Bind类似，但它读取Params
type BindingUri interface {
	Name() string
	BindUri(map[string][]string, any) error
}

// StructValidator是需要实现的最小接口，以便将其用作确保请求正确性的验证器引擎
// Gin为此提供了一个默认实现，使用https://github.com/go-playground/validator/tree/v10.6.1
type StructValidator interface {
// ValidateStruct可以接收任何类型，即使配置不正确，它也不会panic
// 如果接收到的类型是slice数组，则应该对每个元素执行验证
// 如果接收到的类型不是struct或slice|array，则应该跳过任何验证，并且必须返回nil
// 如果接收到的类型是结构体或指向结构体的指针，则应该执行验证
// 如果结构无效或验证本身失败，则应返回描述性错误
// 否则必须返回nil
	ValidateStruct(any) error

// Engine返回为StructValidator实现提供动力的底层验证器引擎
	Engine() any
}

// Validator是默认的验证器，它实现了StructValidator接口
// 它在引擎盖下使用https://github.com/go-playground/validator/tree/v10.6.1
var Validator StructValidator = &defaultValidator{}

// 它们实现了Binding接口，可用于将请求中的数据绑定到struct实例
var (
	JSON          = jsonBinding{}
	XML           = xmlBinding{}
	Form          = formBinding{}
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	ProtoBuf      = protobufBinding{}
	MsgPack       = msgpackBinding{}
	YAML          = yamlBinding{}
	Uri           = uriBinding{}
	Header        = headerBinding{}
	TOML          = tomlBinding{}
)

// Default根据HTTP方法和内容类型返回适当的Binding实例
func Default(method, contentType string) Binding {
	if method == http.MethodGet {
		return Form
	}

	switch contentType {
	case MIMEJSON:
		return JSON
	case MIMEXML, MIMEXML2:
		return XML
	case MIMEPROTOBUF:
		return ProtoBuf
	case MIMEMSGPACK, MIMEMSGPACK2:
		return MsgPack
	case MIMEYAML:
		return YAML
	case MIMETOML:
		return TOML
	case MIMEMultipartPOSTForm:
		return FormMultipart
	default: // 案例MIMEPOSTForm:
		return Form
	}
}

func validate(obj any) error {
	if Validator == nil {
		return nil
	}
	return Validator.ValidateStruct(obj)
}
