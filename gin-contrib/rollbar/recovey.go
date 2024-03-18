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

// ff:
// onlyCrashes:

// ff:
// onlyCrashes:

// ff:
// onlyCrashes:

// ff:
// onlyCrashes:

// ff:
// onlyCrashes:

// ff:
// onlyCrashes:
func Recovery(onlyCrashes bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rval := recover(); rval != nil {
				debug.PrintStack()

				rollbar.Critical(errors.New(fmt.Sprint(rval)), getCallers(3), map[string]string{
					"endpoint": c.Request.RequestURI,
				})

				c.AbortWithStatus(http.StatusInternalServerError)
			}

			if !onlyCrashes {
				for _, item := range c.Errors {
					rollbar.Error(item.Err, map[string]string{
						"meta":     fmt.Sprint(item.Meta),
						"endpoint": c.Request.RequestURI,
					})
				}
			}
		}()

		c.Next()
	}
}

func getCallers(skip int) (pc []uintptr) {
	pc = make([]uintptr, 1000)
	i := runtime.Callers(skip+1, pc)
	return pc[0:i]
}
