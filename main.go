package main

import (
	"senducode/cache"
	"senducode/conf"
	"senducode/server"
)

func main() {
	conf.InitConf()
	cache.InitCache()
	server.InitServer()
}
