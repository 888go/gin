package main

import (
	"fmt"
	"io"
	"log"
	"time"
	
	"github.com/888go/gin"
)

// 它维护一个当前已连接的客户端列表，并向这些客户端广播事件。
type Event struct {
	// 主要事件收集程序会将事件推送到此通道
	Message chan string

	// New client connections
	NewClients chan chan string

	// 已关闭的客户端连接
	ClosedClients chan chan string

	// 总客户端连接数
	TotalClients map[chan string]bool
}

// 新的事件消息将被广播到所有已注册的客户端连接通道
type ClientChan chan string

func main() {
	router := gin.Default()

	// 初始化新的流媒体服务器
	stream := NewServer()

	// 我们以10秒间隔向客户端流式传输当前时间
	go func() {
		for {
			time.Sleep(time.Second * 10)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			// 将当前时间发送到客户端消息通道
			stream.Message <- currentTime
		}
	}()

	// Basic Authentication
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "admin123", // 用户名：admin，密码：admin123
	}))

// 授权的客户端可以流式接收该事件
// 添加事件流传输所需的头部信息
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
			// 从消息通道向客户端流式传输消息
			if msg, ok := <-clientChan; ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		})
	})

	// Parse Static files
	router.StaticFile("/", "./public/index.html")

	router.Run(":8085")
}

// 初始化事件并开始处理请求
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

// 它监听来自客户端的所有入站请求。
// 处理客户端的添加和移除，并向客户端广播消息。
func (stream *Event) listen() {
	for {
		select {
		// 添加新的可用客户端
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		// Remove closed client
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

		// 将新的连接发送至事件服务器
		stream.NewClients <- clientChan

		defer func() {
			// 将关闭的连接发送到事件服务器
			stream.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)

		c.Next()
	}
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}
