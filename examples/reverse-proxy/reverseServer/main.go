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

// 可能我们可以拥有多个真实的服务器地址，并执行某种负载均衡策略。
var RealAddr = []string{
	"http://127.0.0.1:2003",
}

// 这是一个模拟函数，我们可以在其中实现策略。
func getLoadBalanceAddr() string {
	return RealAddr[0]
}

func main() {
	r := gin.Default()
	r.GET("/:path", func(c *gin.Context) {
		// 步骤1：解析代理地址，更改请求中的方案和主机
		req := c.Request
		proxy, err := url.Parse(getLoadBalanceAddr())
		if err != nil {
			log.Printf("error in parse addr: %v", err)
			c.String(500, "error")
			return
		}
		req.URL.Scheme = proxy.Scheme
		req.URL.Host = proxy.Host

		// 步骤 2：使用 http.Transport 向真实服务器发起请求。
		transport := http.DefaultTransport
		resp, err := transport.RoundTrip(req)
		if err != nil {
			log.Printf("error in roundtrip: %v", err)
			c.String(500, "error")
			return
		}

		// 步骤3：将真实服务器响应返回给上游。
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
