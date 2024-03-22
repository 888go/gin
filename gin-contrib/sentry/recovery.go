package sentry

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	
	"github.com/getsentry/raven-go"
	"github.com/888go/gin"
)

// Recovery中间件用于Sentry崩溃报告
func Recovery(client *raven.Client, onlyCrashes bool) gin类.HandlerFunc {

	return func(c *gin类.Context) {
		defer func() {
			flags := map[string]string{
				"endpoint": c.X请求.RequestURI,
			}
			if rval := recover(); rval != nil {
				debug.PrintStack()
				rvalStr := fmt.Sprint(rval)
				client.CaptureMessage(rvalStr, flags, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)),
					raven.NewHttp(c.X请求))
				c.X停止并带状态码(http.StatusInternalServerError)
			}
			if !onlyCrashes {
				for _, item := range c.X错误s {
					client.CaptureMessage(item.Error(), flags, &raven.Message{
						Message: item.Error(),
						Params:  []interface{}{item.Meta},
					},
						raven.NewHttp(c.X请求))
				}
			}
		}()

		c.X中间件继续()
	}
}
