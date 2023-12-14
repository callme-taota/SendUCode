package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"senducode/cache"
	"senducode/tolog"
	"sync"
)

// clients is a map to store WebSocket connections for each user session.
var clients = make(map[string][]*websocket.Conn)
var clientsLock sync.Mutex

// LinkUser creates routes related to user management.
func LinkUser() {
	tolog.Log().Info("Server-user Create link").PrintAndWriteSafe()
	userGroup := Server.Group("/user")
	userGroup.POST("/", CreatUser)
	userGroup.POST("/check", CheckUsingSession)
	userGroup.GET("/ws", WebSocketHandler)
	userGroup.DELETE("/", DeleteUser)
}

// upgrader is used to upgrade an HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// CreatUser handles the creation of a new user and returns a session.
func CreatUser(c *gin.Context) {
	userid := c.Query("userid")
	userSession, _ := cache.GetUserSessionByID(userid)
	if userSession != "" {
		c.JSON(http.StatusOK, gin.H{"msg": "created", "ok": "false", "session": ""})
		return
	}
	session, err := cache.AddUser(userid)
	if err != nil {
		tolog.Log().Errorf("CreateUser %e", err).PrintAndWriteSafe()
		c.JSON(http.StatusOK, gin.H{"msg": "err", "ok": "false", "session": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "ok": "true", "session": session})
}

// CheckUsingSession checks the validity of a session.
func CheckUsingSession(c *gin.Context) {
	session := c.Query("session")
	userid, err := cache.GetUserIDByUserSession(session)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "invalid session", "ok": "false", "userid": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "User check successfully", "ok": "true", "userid": userid})
}

// addClient adds a WebSocket connection to the list of clients for a specific session.
func addClient(session string, conn *websocket.Conn) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	clients[session] = append(clients[session], conn)
}

// removeClient removes a WebSocket connection from the list of clients for a specific session.
func removeClient(session string, conn *websocket.Conn) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	var updatedClients []*websocket.Conn
	for _, c := range clients[session] {
		if c != conn {
			updatedClients = append(updatedClients, c)
		}
	}
	clients[session] = updatedClients
}

// WebSocketHandler handles WebSocket connections for users.
func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		tolog.Log().Errorf("WebSocket upgrade error: %e", err).PrintLog()
		return
	}
	defer conn.Close()

	session := c.Query("session")
	userAgent := c.GetHeader("User-Agent")
	userid, err := cache.GetUserIDByUserSession(session)
	if err != nil {
		tolog.Log().Errorf("Invalid session for WebSocket connection").PrintLog()
		return
	}

	tolog.Log().Infof("WebSocket connection established for user: %s", userid).PrintLog()

	addClient(session, conn)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			tolog.Log().Errorf("Read message error: %e", err).PrintAndWriteSafe()
			removeClient(session, conn)
			break
		}

		handleMessage(session, userAgent, string(p))
		msg, _ := cache.GetMessagesFromSortedSet(session)
		broadcastMessage(session, msg[0])
	}
}

// DeleteUser handles the deletion of a user and their associated data.
func DeleteUser(c *gin.Context) {
	session, ok := c.GetQuery("session")
	if ok != true {
		tolog.Log().Warningf("DeleteUser : not found session")
		c.JSON(http.StatusOK, gin.H{"ok": "false", "msg": "session error"})
		return
	}
	ok, err := cache.DeleteUserBySession(session)
	if ok != true && err != nil {
		tolog.Log().Warningf("DeleteUser : %e", err)
		c.JSON(http.StatusOK, gin.H{"ok": "false", "msg": "delete fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": "true", "msg": "delete success"})
	return
}
