package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gin-gonic/examples/grpc/example1/gen/helloworld/v1"
	
	"github.com/888go/gin"
	"google.golang.org/grpc"
)

func main() {
// 建立到服务器的连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := v1.NewGreeterClient(conn)

// 设置http服务器
	r := gin.Default()
	r.GET("/rest/n/:name", func(c *gin.Context) {
		name := c.Param("name")

// 联系服务器并打印出它的响应
		req := &v1.HelloRequest{Name: name}
		res, err := client.SayHello(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.Message),
		})
	})

// 运行http服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
