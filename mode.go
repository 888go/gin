// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin
import (
	"flag"
	"io"
	"os"
	
	"e.coding.net/gogit/go/gin/binding"
	)
// EnvGinMode为gin模式的环境名
const EnvGinMode = "GIN_MODE"

const (
// DebugMode表示gin模式为debug
	DebugMode = "debug"
// ReleaseMode表示gin模式为release
	ReleaseMode = "release"
// TestMode表示gin模式为test
	TestMode = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

// defaultwwriter是默认的io
// Gin用于调试输出和中间件输出的写入器，如Logger()或Recovery()
// 请注意，Logger和Recovery都提供了自定义的方式来配置它们的输出
// 要在Windows中支持着色，请使用:import "github.com/mattn/go-colorable"杜松子酒
// defaultwwriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter是默认io
// Gin用来调试错误的写入器
var DefaultErrorWriter io.Writer = os.Stderr

var (
	ginMode  = debugCode
	modeName = DebugMode
)

func init() {
	mode := os.Getenv(EnvGinMode)
	SetMode(mode)
}

// SetMode根据输入的字符串设置gin模式
func SetMode(value string) {
	if value == "" {
		if flag.Lookup("test.v") != nil {
			value = TestMode
		} else {
			value = DebugMode
		}
	}

	switch value {
	case DebugMode:
		ginMode = debugCode
	case ReleaseMode:
		ginMode = releaseCode
	case TestMode:
		ginMode = testCode
	default:
		panic("gin mode unknown: " + value + " (available mode: debug release test)")
	}

	modeName = value
}

// DisableBindValidation关闭默认验证器
func DisableBindValidation() {
	binding.Validator = nil
}

// EnableJsonDecoderUseNumber为绑定设置为true
// EnableDecoderUseNumber以调用JSON Decoder实例上的UseNumber方法
func EnableJsonDecoderUseNumber() {
	binding.EnableDecoderUseNumber = true
}

// EnableJsonDecoderDisallowUnknownFields为绑定设置为true
// EnableDecoderDisallowUnknownFields调用JSON Decoder实例上的DisallowUnknownFields方法
func EnableJsonDecoderDisallowUnknownFields() {
	binding.EnableDecoderDisallowUnknownFields = true
}

// Mode返回当前gin模式
func Mode() string {
	return modeName
}
