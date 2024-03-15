package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/888go/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := gin.Default()
	r.GET("/test_stream", func(c *gin.Context) {
		w := c.Writer
		header := w.Header()
		header.Set("Transfer-Encoding", "chunked")
		header.Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<html>
					<body>
		`))
		w.(http.Flusher).Flush()
		for i := 0; i < 10; i++ {
			w.Write([]byte(fmt.Sprintf(`
				<h1>%d</h1>
			`, i)))
			w.(http.Flusher).Flush()
			time.Sleep(time.Duration(1) * time.Second)
		}
		w.Write([]byte(`
			
					</body>
			</html>
		`))
		w.(http.Flusher).Flush()
	})

	r.Run("127.0.0.1:8080")
}

/*
browser test url:
http:// 这是一个Go语言的代码注释，其内容翻译为中文如下：
// 在127.0.0.1主机上的8080端口上监听 /test_stream 路径
// 然而，根据提供的信息，这可能并不是一个完整的Go语言代码片段，它看起来更像是对服务地址的一个描述。如果这是HTTP服务器的监听地址，则对应的Go代码可能是这样的：
// ```go
// 在本地IP 127.0.0.1 的 8080 端口上启动一个HTTP服务器，并处理/test_stream路径的请求
// http.HandleFunc("/test_stream", handleTestStream)
// log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
*/
