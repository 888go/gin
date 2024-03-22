package expvar

import (
	"expvar"
	"fmt"
	
	"github.com/888go/gin"
)

// Handler for gin framework
func Handler() gin类.HandlerFunc {
	return func(c *gin类.Context) {
		w := c.Writer
		c.X设置响应协议头值("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte("{\n"))
		first := true
		expvar.Do(func(kv expvar.KeyValue) {
			if !first {
				_, _ = w.Write([]byte(",\n"))
			}
			first = false
			fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
		})
		_, _ = w.Write([]byte("\n}\n"))
		c.X停止并带状态码(200)
	}
}
