package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"senducode/cache"
	"senducode/tolog"
	"strconv"
)

func LinkMsg() {
	tolog.Log().Info("Server-msg Create link").PrintAndWriteSafe()
	msgGroup := Server.Group("/msg")
	msgGroup.POST("/", newMsg)
	msgGroup.GET("/", getMsgs)
}

func getMsgs(c *gin.Context) {
	session := c.GetHeader("session")
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		tolog.Log().Warningf("limit need integer , err: ", err)
		c.JSON(http.StatusOK, gin.H{"msg": "limit need integer"})
		return
	}
	userid, err := cache.GetUserIDByUserSession(session)
	if err != nil {
		tolog.Log().Errorf("Invalid session for connection").PrintLog()
		return
	}

	tolog.Log().Infof("Get message for user: %s", userid).PrintLog()
	msgs, err := cache.GetMessagesFromSortedSetLimit(session, limit)
	c.JSON(http.StatusOK, msgs)
}

func newMsg(c *gin.Context) {
	session := c.GetHeader("session")
	userAgent := c.GetHeader("User-Agent")
	message := c.Query("message")
	userid, err := cache.GetUserIDByUserSession(session)
	if err != nil {
		tolog.Log().Errorf("Invalid session for connection").PrintLog()
		return
	}

	tolog.Log().Infof("New message from user: %s", userid).PrintLog()

	// 处理接收到的消息，你可以根据消息内容进行逻辑处理
	handleMessage(session, userAgent, message)
	msg, _ := cache.GetMessagesFromSortedSet(session)
	// 如果你要将消息广播给其他设备，可以调用 broadcastMessage 函数
	broadcastMessage(session, msg[0])
	c.JSON(http.StatusOK, gin.H{"msg": "send success"})
}

// 处理接收到的消息的函数
func handleMessage(session, userAgent, message string) {
	msg, err := cache.CreateMessage(message, userAgent)
	if err != nil {
		tolog.Log().Errorf("server-msg handleMessage %e", err)
		return
	}
	err = cache.AddMessageToList(session, msg)
	if err != nil {
		tolog.Log().Errorf("server-msg AddMessageToList %e", err)
		return
	}
	tolog.Log().Infof("Received message from user %s: %s", session, message)
}

// 广播消息给其他设备的函数
func broadcastMessage(session string, message cache.Message) {
	messageJSON, _ := json.Marshal(message)
	for id, client := range clients[session] {
		err := client.WriteMessage(websocket.TextMessage, messageJSON)
		if err != nil {
			tolog.Log().Errorf("Error sending message to user %s: %e", id, err)
		}
	}
}
