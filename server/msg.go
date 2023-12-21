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

// LinkMsg creates a new message-related route group.
func LinkMsg() {
	tolog.Log().Info("Server-msg Create link").PrintAndWriteSafe()
	msgGroup := Server.Group("/msg")
	msgGroup.POST("", newMsg)
	msgGroup.GET("/", getMsgs)
}

// getMsgs retrieves messages for a user based on their session.
func getMsgs(c *gin.Context) {
	session := c.GetHeader("session")
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		tolog.Log().Warningf("limit needs to be an integer, err: %e", err)
		c.JSON(http.StatusOK, gin.H{"msg": "limit needs to be an integer"})
		return
	}
	userid, err := cache.GetUserIDByUserSession(session)
	if err != nil {
		tolog.Log().Errorf("Invalid session for connection").PrintLog()
		c.JSON(http.StatusOK, gin.H{"msg": "Invalid session for connection", "ok": "false"})
		return
	}

	tolog.Log().Infof("Get messages for user: %s", userid).PrintLog()
	msgs, err := cache.GetMessagesFromSortedSetLimit(session, limit)
	c.JSON(http.StatusOK, msgs)
}

type newMsgJSON struct {
	Message string `json:"message"`
}

// newMsg handles the creation of a new message.
func newMsg(c *gin.Context) {
	session := c.GetHeader("session")
	userAgent := c.GetHeader("User-Agent")
	var nm newMsgJSON
	if err := c.ShouldBindJSON(&nm); err != nil {
		c.JSON(400, gin.H{"msg": "post data error", "ok": "false"})
		return
	}
	message := nm.Message
	userid, err := cache.GetUserIDByUserSession(session)
	if err != nil {
		tolog.Log().Errorf("Invalid session for connection").PrintLog()
		c.JSON(http.StatusOK, gin.H{"msg": "Invalid session for connection", "ok": "false"})
		return
	}
	tolog.Log().Infof("New message from user: %s", userid).PrintLog()
	handleMessage(session, userAgent, message)
	msg, _ := cache.GetMessagesFromSortedSet(session)
	broadcastMessage(session, msg[0])
	c.JSON(http.StatusOK, gin.H{"msg": "send success", "ok": "true"})
}

// handleMessage creates a new message and adds it to the user's message list in Redis.
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

// broadcastMessage sends a message to all connected clients for a given user session.
func broadcastMessage(session string, message cache.Message) {
	messageJSON, _ := json.Marshal(message)
	for id, client := range clients[session] {
		err := client.WriteMessage(websocket.TextMessage, messageJSON)
		if err != nil {
			tolog.Log().Errorf("Error sending message to user %d: %e", id, err)
		}
	}
}
