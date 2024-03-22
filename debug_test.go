// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

// TODO：待办事项（需要实现或改进的功能）
// 
// func debugRoute(httpMethod, absolutePath string, handlers HandlersChain) {
//   // 函数功能：调试路由，接收HTTP方法、绝对路径和处理器链作为参数
// 
// func debugPrint(format string, values ...any) {
//   // 函数功能：调试打印，接收一个格式字符串和任意数量的参数，用于输出调试信息

func TestIsDebugging(t *testing.T) {
	X设置运行模式(X常量_运行模式_调试)
	assert.True(t, X是否为调试模式())
	X设置运行模式(X常量_运行模式_发布)
	assert.False(t, X是否为调试模式())
	X设置运行模式(X常量_运行模式_测试)
	assert.False(t, X是否为调试模式())
}

func TestDebugPrint(t *testing.T) {
	re := captureOutput(t, func() {
		X设置运行模式(X常量_运行模式_调试)
		X设置运行模式(X常量_运行模式_发布)
		debugPrint("DEBUG this!")
		X设置运行模式(X常量_运行模式_测试)
		debugPrint("DEBUG this!")
		X设置运行模式(X常量_运行模式_调试)
		debugPrint("these are %d %s", 2, "error messages")
		X设置运行模式(X常量_运行模式_测试)
	})
	assert.Equal(t, "[GIN-debug] these are 2 error messages\n", re)
}

func TestDebugPrintError(t *testing.T) {
	re := captureOutput(t, func() {
		X设置运行模式(X常量_运行模式_调试)
		debugPrintError(nil)
		debugPrintError(errors.New("this is an error"))
		X设置运行模式(X常量_运行模式_测试)
	})
	assert.Equal(t, "[GIN-debug] [ERROR] this is an error\n", re)
}

func TestDebugPrintRoutes(t *testing.T) {
	re := captureOutput(t, func() {
		X设置运行模式(X常量_运行模式_调试)
		debugPrintRoute("GET", "/path/to/route/:param", HandlersChain{func(c *Context) {}, handlerNameTest})
		X设置运行模式(X常量_运行模式_测试)
	})
	assert.Regexp(t, `^\[GIN-debug\] GET    /path/to/route/:param     --> (.*/vendor/)?github.com/888go/gin.handlerNameTest \(2 handlers\)\n$`, re)
}

func TestDebugPrintRouteFunc(t *testing.T) {
	DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		fmt.Fprintf(DefaultWriter, "[GIN-debug] %-6s %-40s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	re := captureOutput(t, func() {
		X设置运行模式(X常量_运行模式_调试)
		debugPrintRoute("GET", "/path/to/route/:param1/:param2", HandlersChain{func(c *Context) {}, handlerNameTest})
		X设置运行模式(X常量_运行模式_测试)
	})
	assert.Regexp(t, `^\[GIN-debug\] GET    /path/to/route/:param1/:param2           --> (.*/vendor/)?github.com/888go/gin.handlerNameTest \(2 handlers\)\n$`, re)
}

func TestDebugPrintLoadTemplate(t *testing.T) {
	re := captureOutput(t, func() {
		X设置运行模式(X常量_运行模式_调试)
		templ := template.Must(template.New("").Delims("{[{", "}]}").ParseGlob("./testdata/template/hello.tmpl"))
		debugPrintLoadTemplate(templ)
		X设置运行模式(X常量_运行模式_测试)
	})
	assert.Regexp(t, `^\[GIN-debug\] Loaded HTML Templates \(2\): \n(\t- \n|\t- hello\.tmpl\n){2}\n`, re)
}

func TestDebugPrintWARNINGSetHTMLTemplate(t *testing.T) {
	re := captureOutput(t, func() {
		X设置运行模式(X常量_运行模式_调试)
		debugPrintWARNINGSetHTMLTemplate()
		X设置运行模式(X常量_运行模式_测试)
	})
	assert.Equal(t, "[GIN-debug] [WARNING] Since SetHTMLTemplate() is NOT thread-safe. It should only be called\nat initialization. ie. before any route is registered or the router is listening in a socket:\n\n\trouter := gin.Default()\n\trouter.SetHTMLTemplate(template) // << good place\n\n", re)
}

func TestDebugPrintWARNINGDefault(t *testing.T) {
	re := captureOutput(t, func() {
		X设置运行模式(X常量_运行模式_调试)
		debugPrintWARNINGDefault()
		X设置运行模式(X常量_运行模式_测试)
	})
	m, e := getMinVer(runtime.Version())
	if e == nil && m < ginSupportMinGoVer {
		assert.Equal(t, "[GIN-debug] [WARNING] Now Gin requires Go 1.18+.\n\n[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.\n\n", re)
	} else {
		assert.Equal(t, "[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.\n\n", re)
	}
}

func TestDebugPrintWARNINGNew(t *testing.T) {
	re := captureOutput(t, func() {
		X设置运行模式(X常量_运行模式_调试)
		debugPrintWARNINGNew()
		X设置运行模式(X常量_运行模式_测试)
	})
	assert.Equal(t, "[GIN-debug] [WARNING] Running in \"debug\" mode. Switch to \"release\" mode in production.\n - using env:\texport GIN_MODE=release\n - using code:\tgin.SetMode(gin.ReleaseMode)\n\n", re)
}

func captureOutput(t *testing.T, f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	defaultWriter := DefaultWriter
	defaultErrorWriter := DefaultErrorWriter
	defer func() {
		DefaultWriter = defaultWriter
		DefaultErrorWriter = defaultErrorWriter
		log.SetOutput(os.Stderr)
	}()
	DefaultWriter = writer
	DefaultErrorWriter = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf strings.Builder
		wg.Done()
		_, err := io.Copy(&buf, reader)
		assert.NoError(t, err)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestGetMinVer(t *testing.T) {
	var m uint64
	var e error
	_, e = getMinVer("go1")
	assert.NotNil(t, e)
	m, e = getMinVer("go1.1")
	assert.Equal(t, uint64(1), m)
	assert.Nil(t, e)
	m, e = getMinVer("go1.1.1")
	assert.Nil(t, e)
	assert.Equal(t, uint64(1), m)
	_, e = getMinVer("go1.1.1.1")
	assert.NotNil(t, e)
}
