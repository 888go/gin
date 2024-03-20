// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin

import (
	"flag"
	"io"
	"os"
	
	"github.com/888go/gin/binding"
)

// EnvGinMode 指示 Gin 模式的环境名称。
const EnvGinMode = "GIN_MODE"

const (
	// DebugMode 指示 gin 模式为调试模式。
	DebugMode = "debug"
	// ReleaseMode 表示 gin 模式为发布模式。
	ReleaseMode = "release"
	// TestMode 表示 gin 模式为测试模式。
	TestMode = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter 是 Gin 默认使用的 io.Writer，用于调试输出以及中间件输出，如 Logger() 和 Recovery()。
// 注意，Logger 和 Recovery 都提供了自定义配置其输出 io.Writer 的方法。
// 若要在 Windows 系统中支持彩色输出，请使用：
//
//	导入 "github.com/mattn/go-colorable"
//	gin.DefaultWriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter 是 Gin 默认使用的 io.Writer，用于调试错误
var DefaultErrorWriter io.Writer = os.Stderr

var (
	ginMode  = debugCode
	modeName = DebugMode
)

func init() {
	mode := os.Getenv(EnvGinMode)
	SetMode(mode)
}

// SetMode 根据输入的字符串设置 gin 模式。
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

// DisableBindValidation 关闭默认的验证器。
func DisableBindValidation() {
	binding.Validator = nil
}

// EnableJsonDecoderUseNumber 将参数设置为 true 以启用 binding.EnableDecoderUseNumber，
// 这样就会在 JSON 解码器实例上调用 UseNumber 方法。
func EnableJsonDecoderUseNumber() {
	binding.EnableDecoderUseNumber = true
}

// EnableJsonDecoderDisallowUnknownFields 将 binding.EnableDecoderDisallowUnknownFields 设为 true，
// 以便在 JSON 解码器实例上调用 DisallowUnknownFields 方法。
func EnableJsonDecoderDisallowUnknownFields() {
	binding.EnableDecoderDisallowUnknownFields = true
}

// Mode 返回当前 gin 模式。
func Mode() string {
	return modeName
}
