package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"senducode/conf"
	"senducode/tolog"
)

var Server *gin.Engine

func InitServer() {
	ginServer := gin.Default()
	// 使用CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 这里设置允许所有来源，可以根据需要指定具体的域名
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Referer", "User-Agent", "Content-Type", "Authorization"} // 根据实际需要添加其他 header
	ginServer.Use(cors.New(config))
	Server = ginServer
	tolog.Log().Info("Gin主服务已开启").PrintAndWriteSafe()
	port := conf.Server.Port
	LinkAPI()
	tolog.Log().Infoln("Gin监听开启 端口:"+port, "本地地址 http://127.0.0.1:"+port).PrintAndWriteSafe()
	ginServer.Run(fmt.Sprintf(":%s", port))
}

func LinkAPI() {
	LinkUser()
	LinkMsg()
}
