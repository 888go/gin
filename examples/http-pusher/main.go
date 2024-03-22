package main

import (
	"html/template"
	"log"
	
	"github.com/888go/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin类.X创建默认对象()
	r.X绑定静态文件目录("/assets", "./assets")
	r.X设置Template模板(html)

	r.X绑定GET("/", func(c *gin类.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用pusher.Push()进行服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.X输出html模板(200, "https", gin类.H{
			"status": "success",
		})
	})

	// 在 https://127.0.0.1:8080 上监听并服务
	r.X监听TLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
