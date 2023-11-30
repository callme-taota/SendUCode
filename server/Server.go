package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"senducode/conf"
	"senducode/tolog"
)

var Server *gin.Engine

func InitServer() {
	ginServer := gin.Default()
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
