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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Referer", "User-Agent", "Content-Type", "Authorization"}
	ginServer.Use(cors.New(config))
	Server = ginServer
	tolog.Log().Info("Gin Main Server Start").PrintAndWriteSafe()
	port := conf.Server.Port
	LinkAPI()
	tolog.Log().Infoln("Gin listing on:"+port, "host: http://127.0.0.1:"+port).PrintAndWriteSafe()
	ginServer.Run(fmt.Sprintf(":%s", port))
}

func LinkAPI() {
	LinkUser()
	LinkMsg()
}
