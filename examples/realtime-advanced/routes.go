package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"
	
	"github.com/888go/gin"
)

func rateLimit(c *gin类.Context) {
	ip := c.X取客户端ip()
	value := int(ips.Add(ip, 1))
	if value%50 == 0 {
		fmt.Printf("ip: %s, count: %d\n", ip, value)
	}
	if value >= 200 {
		if value%200 == 0 {
			fmt.Println("ip blocked")
		}
		c.X停止()
		c.X输出文本(http.StatusServiceUnavailable, "you were automatically banned :)")
	}
}

func index(c *gin类.Context) {
	c.X重定向(http.StatusMovedPermanently, "/room/hn")
}

func roomGET(c *gin类.Context) {
	roomid := c.X取API参数值("roomid")
	nick := c.X取URL参数值("nick")
	if len(nick) < 2 {
		nick = ""
	}
	if len(nick) > 13 {
		nick = nick[0:12] + "..."
	}
	c.X输出html模板(http.StatusOK, "room_login.templ.html", gin类.H{
		"roomid":    roomid,
		"nick":      nick,
		"timestamp": time.Now().Unix(),
	})
}

func roomPOST(c *gin类.Context) {
	roomid := c.X取API参数值("roomid")
	nick := c.X取URL参数值("nick")
	message := c.X取表单参数值("message")
	message = strings.TrimSpace(message)

	validMessage := len(message) > 1 && len(message) < 200
	validNick := len(nick) > 1 && len(nick) < 14
	if !validMessage || !validNick {
		c.X输出JSON(http.StatusBadRequest, gin类.H{
			"status": "failed",
			"error":  "the message or nickname is too long",
		})
		return
	}

	post := gin类.H{
		"nick":    html.EscapeString(nick),
		"message": html.EscapeString(message),
	}
	messages.Add("inbound", 1)
	room(roomid).Submit(post)
	c.X输出JSON(http.StatusOK, post)
}

func streamRoom(c *gin类.Context) {
	roomid := c.X取API参数值("roomid")
	listener := openListener(roomid)
	ticker := time.NewTicker(1 * time.Second)
	users.Add("connected", 1)
	defer func() {
		closeListener(roomid, listener)
		ticker.Stop()
		users.Add("disconnected", 1)
	}()

	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-listener:
			messages.Add("outbound", 1)
			c.SSEvent("message", msg)
		case <-ticker.C:
			c.SSEvent("stats", Stats())
		}
		return true
	})
}
