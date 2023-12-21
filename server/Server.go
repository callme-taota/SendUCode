package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"senducode/conf"
	"senducode/tolog"
	"time"
)

// Server represents the main Gin engine.
var Server *gin.Engine

// InitServer initializes the main Gin server with CORS configuration.
func InitServer() {
	// Create a new default Gin server instance.
	ginServer := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	// Configure CORS settings.
	ginServer.Use(cors.New(cors.Config{ //conf
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "session"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, //time
		ExposeHeaders:    []string{"Content-Length"},
		AllowOriginFunc: func(origin string) bool { //allow
			return true //all
		},
	}))

	// Set the global Server variable to the configured Gin server.
	Server = ginServer
	LinkAPI()

	// Log server initialization information.
	tolog.Log().Info("Gin Main Server Start").PrintAndWriteSafe()
	port := conf.Server.Port
	ginServer.POST("/test", TestPost)
	tolog.Log().Infoln("Gin listening on:"+port, "host: http://127.0.0.1:"+port).PrintAndWriteSafe()

	// Run the Gin server on the specified port.
	ginServer.Run(fmt.Sprintf(":%s", port))
}

// LinkAPI connects various API routes to the main Gin server.
func LinkAPI() {
	// Link User and Message APIs to the main server.
	LinkUser()
	LinkMsg()
}

type PostData struct {
	// 定义与前端POST数据结构对应的Go结构体
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestPost(c *gin.Context) {
	// 创建一个空的PostData对象
	var postData PostData

	// 使用ShouldBindJSON将前端发送的JSON数据绑定到PostData对象
	if err := c.ShouldBindJSON(&postData); err != nil {
		// 如果绑定失败，返回错误信息
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 在这里可以使用postData中的数据进行处理
	// 例如，打印用户名和密码
	println("Username:", postData.Username)
	println("Password:", postData.Password)

	// 返回成功的响应
	c.JSON(200, gin.H{"message": "Data received successfully"})
}
