package rollbar

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
	
	"github.com/888go/gin"
	"github.com/rollbar/rollbar-go"
)

// Recovery中间件用于Rollbar错误监控
func Recovery(onlyCrashes bool) gin类.HandlerFunc {
	return func(c *gin类.Context) {
		defer func() {
			if rval := recover(); rval != nil {
				debug.PrintStack()

				rollbar.Critical(errors.New(fmt.Sprint(rval)), getCallers(3), map[string]string{
					"endpoint": c.X请求.RequestURI,
				})

				c.X停止并带状态码(http.StatusInternalServerError)
			}

			if !onlyCrashes {
				for _, item := range c.X错误s {
					rollbar.Error(item.Err, map[string]string{
						"meta":     fmt.Sprint(item.Meta),
						"endpoint": c.X请求.RequestURI,
					})
				}
			}
		}()

		c.X中间件继续()
	}
}

func getCallers(skip int) (pc []uintptr) {
	pc = make([]uintptr, 1000)
	i := runtime.Callers(skip+1, pc)
	return pc[0:i]
}
