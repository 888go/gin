// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("路")
	dot       = []byte(".")
	slash     = []byte("/")
)

// recoveryfunction定义了CustomRecovery可传递的函数
type RecoveryFunc func(c *Context, err any)

// Recovery返回一个中间件，它可以从任何Panic中恢复，如果有Panic，则写入500

// ff:
func Recovery() HandlerFunc {
	return RecoveryWithWriter(DefaultErrorWriter)
}

// CustomRecovery返回一个中间件，它可以从任何Panic中恢复，并调用提供的handle函数来处理它

// ff:
// handle:
func CustomRecovery(handle RecoveryFunc) HandlerFunc {
	return RecoveryWithWriter(DefaultErrorWriter, handle)
}

// RecoveryWithWriter为给定的写入器返回一个中间件，该写入器可以从任何Panic中恢复并写入500(如果有的话)

// ff:
// recovery:
// out:
func RecoveryWithWriter(out io.Writer, recovery ...RecoveryFunc) HandlerFunc {
	if len(recovery) > 0 {
		return CustomRecoveryWithWriter(out, recovery[0])
	}
	return CustomRecoveryWithWriter(out, defaultHandleRecovery)
}

// CustomRecoveryWithWriter为给定的编写器返回一个中间件，该编写器可以从任何Panic中恢复，并调用提供的handle函数来处理它

// ff:
// handle:
// out:
func CustomRecoveryWithWriter(out io.Writer, handle RecoveryFunc) HandlerFunc {
	var logger *log.Logger
	if out != nil {
		logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
	}
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
// 检查是否有断开的连接，因为它实际上并不是需要进行紧急堆栈跟踪的条件
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if logger != nil {
					stack := stack(3)
					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					headers := strings.Split(string(httpRequest), "\r\n")
					for idx, header := range headers {
						current := strings.Split(header, ":")
						if current[0] == "Authorization" {
							headers[idx] = current[0] + ": *"
						}
					}
					headersToStr := strings.Join(headers, "\r\n")
					if brokenPipe {
						logger.Printf("%s\n%s%s", err, headersToStr, reset)
					} else if IsDebugging() {
						logger.Printf("[Recovery] %s panic recovered:\n%s\n%s\n%s%s",
							timeFormat(time.Now()), headersToStr, err, stack, reset)
					} else {
						logger.Printf("[Recovery] %s panic recovered:\n%s\n%s%s",
							timeFormat(time.Now()), err, stack, reset)
					}
				}
				if brokenPipe {
// 如果连接已死，我们就不能向它写入状态
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
				} else {
					handle(c, err)
				}
			}
		}()
		c.Next()
	}
}

func defaultHandleRecovery(c *Context, _ any) {
	c.AbortWithStatus(http.StatusInternalServerError)
}

// Stack返回一个格式良好的堆栈帧，跳过跳过帧
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // 返回的数据
// 在循环时，打开文件并读取它们
// 这些变量记录当前加载的文件
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // 跳过预期的帧数
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
// 至少打印这么多
// 如果我们找不到源头，就不会显示出来
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := os.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// Source返回第n行经过空格处理的切片
func source(lines [][]byte, n int) []byte {
	n-- // 在堆栈跟踪中，行是1索引的，但是数组是0索引的
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// 如果可能的话，函数返回包含PC的函数名
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T路ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contain dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.ReplaceAll(name, centerDot, dot)
	return name
}

// timeFormat返回记录器的自定义时间字符串
func timeFormat(t time.Time) string {
	return t.Format("2006/01/02 - 15:04:05")
}
