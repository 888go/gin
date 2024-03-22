// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

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
	X常量_运行模式_调试 = "debug"
	// ReleaseMode 表示 gin 模式为发布模式。
	X常量_运行模式_发布 = "release"
	// TestMode 表示 gin 模式为测试模式。
	X常量_运行模式_测试 = "test"
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
	modeName = X常量_运行模式_调试
)

func init() {
	mode := os.Getenv(EnvGinMode)
	X设置运行模式(mode)
}

// SetMode 根据输入的字符串设置 gin 模式。
func X设置运行模式(常量_运行模式 string) {
	if 常量_运行模式 == "" {
		if flag.Lookup("test.v") != nil {
			常量_运行模式 = X常量_运行模式_测试
		} else {
			常量_运行模式 = X常量_运行模式_调试
		}
	}

	switch 常量_运行模式 {
	case X常量_运行模式_调试:
		ginMode = debugCode
	case X常量_运行模式_发布:
		ginMode = releaseCode
	case X常量_运行模式_测试:
		ginMode = testCode
	default:
		panic("gin mode unknown: " + 常量_运行模式 + " (available mode: debug release test)")
	}

	modeName = 常量_运行模式
}

// DisableBindValidation 关闭默认的验证器。
func X关闭Validator验证器() {
	binding.Validator = nil
}

// EnableJsonDecoderUseNumber 将参数设置为 true 以启用 binding.EnableDecoderUseNumber，
// EnableDecoderUseNumber 用于调用 JSON 解码器实例上的 UseNumber 方法。启用 UseNumber 后，解码器将在反序列化数字时将其解析为 Number 类型而不是 float64 类型并存储到 any 类型变量中。
func X启用Json解码器使用Number() {
	binding.EnableDecoderUseNumber = true
}

// EnableJsonDecoderDisallowUnknownFields 将 binding.EnableDecoderDisallowUnknownFields 设为 true，
// EnableDecoderDisallowUnknownFields 用于调用 JSON 解码器实例上的 DisallowUnknownFields 方法。该方法启用后，当目标是一个结构体且输入包含与目标中非忽略的导出字段不匹配的对象键时，解码器会返回一个错误。
func X启用json解码器禁止未知字段() {
	binding.EnableDecoderDisallowUnknownFields = true
}

// Mode 返回当前 gin 模式。
func X取运行模式() string {
	return modeName
}
