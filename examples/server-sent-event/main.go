package main

import (
	"fmt"
	"io"
	"log"
	"time"
	
	"github.com/888go/gin"
)

// 它保留当前附加的客户机列表，并向这些客户机广播事件
type Event struct {
// 事件由主事件收集例程推送到此通道
	Message chan string

// 新的客户端连接
	NewClients chan chan string

// 关闭的客户端连接
	ClosedClients chan chan string

// 客户端总连接数
	TotalClients map[chan string]bool
}

// 将新事件消息广播到所有已注册的客户端连接通道
type ClientChan chan string

func main() {
	router := gin.Default()

// 初始化新的流媒体服务器
	stream := NewServer()

// 我们以10秒的间隔将当前时间流式传输给客户端
	go func() {
		for {
			time.Sleep(time.Second * 10)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

// 发送当前时间到客户端消息通道
			stream.Message <- currentTime
		}
	}()

// 基本身份验证
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin123", // 用户名:admin，密码:admin123
	}))

// 授权客户端可以流式传输事件
	authorized.GET("/stream", HeadersMiddleware(), stream.serveHTTP(), func(c *gin.Context) {
		v, ok := c.Get("clientChan")
		if !ok {
			return
		}
		clientChan, ok := v.(ClientChan)
		if !ok {
			return
		}
		c.Stream(func(w io.Writer) bool {
// 从消息通道流消息到客户端
			if msg, ok := <-clientChan; ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		})
	})

// 解析静态文件
	router.StaticFile("/", "./public/index.html")

	router.Run(":8085")
}

// 初始化事件并开始处理请求

// ff:
// event:

// ff:
// event:

// ff:
// event:

// ff:
// event:

// ff:
// event:

// ff:
// event:

// ff:
// event:
func NewServer() (event *Event) {
	event = &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}

	go event.listen()

	return
}

// 它监听来自客户端的所有传入请求
// 处理添加和删除客户端以及向客户端广播消息
func (stream *Event) listen() {
	for {
		select {
// 添加新的可用客户端
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

// 删除关闭的客户端
		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

// 向客户端广播消息
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (stream *Event) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
// 初始化客户端通道
		clientChan := make(ClientChan)

// 向事件服务器发送新连接
		stream.NewClients <- clientChan

		defer func() {
// 向事件服务器发送关闭的连接
			stream.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)

		c.Next()
	}
}


// ff:

// ff:

// ff:

// ff:

// ff:

// ff:

// ff:
func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}
