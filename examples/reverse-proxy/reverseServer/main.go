package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
	
	"github.com/888go/gin"
)

const (
	// this is our reverse server ip address
	ReverseServerAddr = "127.0.0.1:2002"
)

// maybe we can have many real server addresses and do some load balanced strategy.
var RealAddr = []string{
	"http://127.0.0.1:2003",
}

// a fake function that we can do strategy here.
func getLoadBalanceAddr() string {
	return RealAddr[0]
}

func main() {
	r := gin类.X创建默认对象()
	r.X绑定GET("/:path", func(c *gin类.Context) {
		// step 1: resolve proxy address, change scheme and host in requets
		req := c.X请求
		proxy, err := url.Parse(getLoadBalanceAddr())
		if err != nil {
			log.Printf("error in parse addr: %v", err)
			c.X输出文本(500, "error")
			return
		}
		req.URL.Scheme = proxy.Scheme
		req.URL.Host = proxy.Host

		// step 2: use http.Transport to do request to real server.
		transport := http.DefaultTransport
		resp, err := transport.RoundTrip(req)
		if err != nil {
			log.Printf("error in roundtrip: %v", err)
			c.X输出文本(500, "error")
			return
		}

		// step 3: return real server response to upstream.
		for k, vv := range resp.Header {
			for _, v := range vv {
				c.X设置响应协议头值(k, v)
			}
		}
		defer resp.Body.Close()
		bufio.NewReader(resp.Body).WriteTo(c.Writer)
		return
	})

	if err := r.X监听(ReverseServerAddr); err != nil {
		log.Printf("Error: %v", err)
	}
}
