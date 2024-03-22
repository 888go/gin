package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	
	"github.com/888go/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r := gin类.X创建默认对象()
	r.X设置Template模板(html)

	r.X绑定GET("/welcome", func(c *gin类.Context) {
		c.X输出html模板(http.StatusOK, "https", gin类.H{
			"status": "success",
		})
	})

	// 在 https://127.0.0.1:8080 上监听并服务
	r.X监听TLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
