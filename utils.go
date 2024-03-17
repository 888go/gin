// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"encoding/xml"
	"net/http"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

// BindKey默认绑定键
const BindKey = "_gin-gonic/gin/bindkey"

// Bind是给定接口对象的辅助函数，并返回一个Gin中间件

// ff:
// val:
func Bind(val any) HandlerFunc {
	value := reflect.ValueOf(val)
	if value.Kind() == reflect.Ptr {
		panic(`Bind struct can not be a pointer. Example:
	Use: gin.Bind(Struct{}) instead of gin.Bind(&Struct{})
`)
	}
	typ := value.Type()

	return func(c *Context) {
		obj := reflect.New(typ).Interface()
		if c.Bind(obj) == nil {
			c.Set(BindKey, obj)
		}
	}
}

// WrapF是用于包装http的辅助函数
// HandlerFunc并返回一个Gin中间件

// ff:
// f:
func WrapF(f http.HandlerFunc) HandlerFunc {
	return func(c *Context) {
		f(c.Writer, c.Request)
	}
}

// WrapH是包装http的辅助函数
// 处理程序并返回一个Gin中间件

// ff:
// h:
func WrapH(h http.Handler) HandlerFunc {
	return func(c *Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// H是map[string]any的快捷方式
type H map[string]any

// MarshalXML允许类型H与xml.Marshal一起使用

// ff:
// start:
// e:
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func assert1(guard bool, text string) {
	if !guard {
		panic(text)
	}
}

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func chooseData(custom, wildcard any) any {
	if custom != nil {
		return custom
	}
	if wildcard != nil {
		return wildcard
	}
	panic("negotiation config is invalid")
}

func parseAccept(acceptHeader string) []string {
	parts := strings.Split(acceptHeader, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		if i := strings.IndexByte(part, ';'); i > 0 {
			part = part[:i]
		}
		if part = strings.TrimSpace(part); part != "" {
			out = append(out, part)
		}
	}
	return out
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func nameOfFunction(f any) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			debugPrint("Environment variable PORT=\"%s\"", port)
			return ":" + port
		}
		debugPrint("Environment variable PORT is undefined. Using port :8080 by default")
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

// 该注释引用自Stack Overflow上的一个问题：
// https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
// 这段Go代码是用于检查一个字符串是否只包含ASCII字符。
// 以下是更详细的中文注释（假设后面有具体的代码）：
// ```go
// 上述链接指向Stack Overflow上的一个问题，讨论如何检查一个字符串是否仅包含ASCII字符。
// 下面的Go代码实现了一个函数，用于检测给定的字符串是否全部由ASCII字符组成。
// ASCII字符集包含128个字符，包括英文字母、数字、标点符号和一些控制字符。
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
