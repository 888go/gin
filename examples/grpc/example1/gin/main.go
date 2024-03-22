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
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := v1.NewGreeterClient(conn)

	// Set up a http server.
	r := gin类.X创建默认对象()
	r.X绑定GET("/rest/n/:name", func(c *gin类.Context) {
		name := c.X取API参数值("name")

		// Contact the server and print out its response.
		req := &v1.HelloRequest{Name: name}
		res, err := client.SayHello(c, req)
		if err != nil {
			c.X输出JSON(http.StatusInternalServerError, gin类.H{
				"error": err.Error(),
			})
			return
		}

		c.X输出JSON(http.StatusOK, gin类.H{
			"result": fmt.Sprint(res.Message),
		})
	})

	// Run http server
	if err := r.X监听(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
