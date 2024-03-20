// 版权所有 ? 2020 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build nomsgpack

package binding

import (
	"net/http"
)

// Content-Type MIME 是最常见的数据格式的 MIME 类型。
const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEYAML              = "application/x-yaml"
	MIMETOML              = "application/toml"
)

// Binding描述了需要实现的接口，目的是为了将请求中携带的数据（如JSON请求体、查询参数或表单POST数据）进行绑定。
type Binding interface {
	Name() string
	Bind(*http.Request, any) error
}

// BindingBody 为 Binding 添加了 BindBody 方法。BindBody 与 Bind 类似，
// 但是它从提供的字节中读取请求体，而不是从 req.Body 中读取。
type BindingBody interface {
	Binding
	BindBody([]byte, any) error
}

// BindingUri 向 Binding 结构体添加 BindUri 方法。BindUri 与 Bind 类似，
// 但它读取的是 Params 参数。
type BindingUri interface {
	Name() string
	BindUri(map[string][]string, any) error
}

// StructValidator 是一个最小接口，为了能够用作验证请求正确性的验证引擎，需要实现这个接口。
// Gin 提供了一个默认实现，使用了 https://github.com/go-playground/validator/tree/v10.6.1。
type StructValidator interface {
// ValidateStruct 可以接收任何类型的值，并且即使配置不正确，也绝不应该引发 panic。如果接收到的类型不是结构体，则应跳过所有验证并返回 nil。
// 如果接收到的类型是结构体或指向结构体的指针，则应执行验证操作。
// 若结构体无效或验证过程本身失败，则应返回一个描述性错误信息。否则必须返回 nil。
	ValidateStruct(any) error

// Engine 方法返回底层驱动验证器引擎，该引擎为 StructValidator 实现提供支持。
	Engine() any
}

// Validator 是默认的验证器，实现了 StructValidator 接口。在底层，它使用了 https://github.com/go-playground/validator/tree/v10.6.1 。
var Validator StructValidator = &defaultValidator{}

// 这些实现了Binding接口，可用于将请求中呈现的数据绑定到结构体实例。
var (
	JSON          = jsonBinding{}
	XML           = xmlBinding{}
	Form          = formBinding{}
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	ProtoBuf      = protobufBinding{}
	YAML          = yamlBinding{}
	Uri           = uriBinding{}
	Header        = headerBinding{}
	TOML          = tomlBinding{}
)

// Default 根据 HTTP 方法和内容类型返回相应的 Binding 实例。

// ff:
// contentType:
// method:
func Default(method, contentType string) Binding {
	if method == "GET" {
		return Form
	}

	switch contentType {
	case MIMEJSON:
		return JSON
	case MIMEXML, MIMEXML2:
		return XML
	case MIMEPROTOBUF:
		return ProtoBuf
	case MIMEYAML:
		return YAML
	case MIMEMultipartPOSTForm:
		return FormMultipart
	case MIMETOML:
		return TOML
	default: // case MIMEPOSTForm:
		return Form
	}
}

func validate(obj any) error {
	if Validator == nil {
		return nil
	}
	return Validator.ValidateStruct(obj)
}
