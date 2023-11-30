package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"senducode/cache"
	"senducode/tolog"
)

var clients = make(map[string][]*websocket.Conn) // 保存连接的客户端

func LinkUser() {
	tolog.Log().Info("Server-user Create link").PrintAndWriteSafe()
	userGroup := Server.Group("/user")
	userGroup.POST("/", CreatUser)
	userGroup.POST("/check", CheckUsingSession)
	userGroup.GET("/ws", WebSocketHandler)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func CreatUser(c *gin.Context) {
	userid, ok := c.GetQuery("userid")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"msg": "err"})
		return
	}
	userSession, _ := cache.GetUserSessionByID(userid)
	if userSession != "" {
		c.JSON(http.StatusOK, gin.H{"msg": "created"})
		return
	}
	session, err := cache.AddUser(userid)
	if err != nil {
		tolog.Log().Errorf("CreateUser %e", err).PrintAndWriteSafe()
		c.JSON(http.StatusOK, gin.H{"msg": "err"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"session": session})
}

func CheckUsingSession(c *gin.Context) {
	session := c.Query("session")
	userid, err := cache.GetUserIDByUserSession(session)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User check successfully", "userid": userid})
}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		tolog.Log().Errorf("WebSocket upgrade error: %e", err).PrintLog()
		return
	}
	defer conn.Close()

	session := c.GetHeader("session")
	userAgent := c.GetHeader("User-Agent")
	userid, err := cache.GetUserIDByUserSession(session)
	if err != nil {
		tolog.Log().Errorf("Invalid session for WebSocket connection").PrintLog()
		return
	}

	tolog.Log().Infof("WebSocket connection established for user: %s", userid).PrintLog()

	clients[session] = append(clients[session], conn) // 将连接添加到客户端列表

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			tolog.Log().Errorf("Read message error: %e", err).PrintAndWriteSafe()
			delete(clients, session) // 从客户端列表中删除连接
			break
		}

		// 处理接收到的消息，你可以根据消息内容进行逻辑处理
		handleMessage(session, userAgent, string(p))
		msg, _ := cache.GetMessagesFromSortedSet(session)
		// 如果你要将消息广播给其他设备，可以调用 broadcastMessage 函数
		broadcastMessage(session, msg[0])
	}

}
