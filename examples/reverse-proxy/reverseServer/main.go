package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
	
	"github.com/888go/gin"
)

const (
// 这是我们的反向服务器IP地址
	ReverseServerAddr = "127.0.0.1:2002"
)

// 也许我们可以有很多真实的服务器地址，并做一些负载均衡策略
var RealAddr = []string{
	"http://127.0.0.1:2003",
}

// 一个伪函数，我们可以在这里做策略
func getLoadBalanceAddr() string {
	return RealAddr[0]
}

func main() {
	r := gin.Default()
	r.GET("/:path", func(c *gin.Context) {
// 步骤1:解析代理地址，更改请求中的方案和主机
		req := c.Request
		proxy, err := url.Parse(getLoadBalanceAddr())
		if err != nil {
			log.Printf("error in parse addr: %v", err)
			c.String(500, "error")
			return
		}
		req.URL.Scheme = proxy.Scheme
		req.URL.Host = proxy.Host

// 步骤2:使用http
// 将请求传输到实服务器
		transport := http.DefaultTransport
		resp, err := transport.RoundTrip(req)
		if err != nil {
			log.Printf("error in roundtrip: %v", err)
			c.String(500, "error")
			return
		}

// 步骤3:向上游返回实服务器响应
		for k, vv := range resp.Header {
			for _, v := range vv {
				c.Header(k, v)
			}
		}
		defer resp.Body.Close()
		bufio.NewReader(resp.Body).WriteTo(c.Writer)
		return
	})

	if err := r.Run(ReverseServerAddr); err != nil {
		log.Printf("Error: %v", err)
	}
}
