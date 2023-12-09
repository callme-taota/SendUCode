package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"senducode/conf"
	"senducode/tolog"
)

// Server represents the main Gin engine.
var Server *gin.Engine

// InitServer initializes the main Gin server with CORS configuration.
func InitServer() {
	// Create a new default Gin server instance.
	ginServer := gin.Default()

	// Configure CORS settings.
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Referer", "User-Agent", "Content-Type", "Authorization"}
	ginServer.Use(cors.New(config))

	// Set the global Server variable to the configured Gin server.
	Server = ginServer

	// Log server initialization information.
	tolog.Log().Info("Gin Main Server Start").PrintAndWriteSafe()
	port := conf.Server.Port
	LinkAPI()
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
