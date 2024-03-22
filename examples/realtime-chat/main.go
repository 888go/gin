package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	
	"github.com/888go/gin"
)

var roomManager *Manager

func main() {
	roomManager = NewRoomManager()
	router := gin类.X创建默认对象()
	router.X设置Template模板(html)

	router.X绑定GET("/room/:roomid", roomGET)
	router.X绑定POST("/room/:roomid", roomPOST)
	router.X绑定DELETE("/room/:roomid", roomDELETE)
	router.X绑定GET("/stream/:roomid", stream)

	router.X监听(":8080")
}

func stream(c *gin类.Context) {
	roomid := c.X取API参数值("roomid")
	listener := roomManager.OpenListener(roomid)
	defer roomManager.CloseListener(roomid, listener)

	clientGone := c.X请求.Context().Done()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-listener:
			c.SSEvent("message", message)
			return true
		}
	})
}

func roomGET(c *gin类.Context) {
	roomid := c.X取API参数值("roomid")
	userid := fmt.Sprint(rand.Int31())
	c.X输出html模板(http.StatusOK, "chat_room", gin类.H{
		"roomid": roomid,
		"userid": userid,
	})
}

func roomPOST(c *gin类.Context) {
	roomid := c.X取API参数值("roomid")
	userid := c.X取表单参数值("user")
	message := c.X取表单参数值("message")
	roomManager.Submit(userid, roomid, message)

	c.X输出JSON(http.StatusOK, gin类.H{
		"status":  "success",
		"message": message,
	})
}

func roomDELETE(c *gin类.Context) {
	roomid := c.X取API参数值("roomid")
	roomManager.DeleteBroadcast(roomid)
}
