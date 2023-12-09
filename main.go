package main

import (
	"senducode/cache"
	"senducode/conf"
	"senducode/server"
	"senducode/tolog"
)

func main() {
	err := conf.InitConf()
	if err != nil {
		tolog.Log().Errorf("Init config file error: %e", err)
		return
	}
	cache.InitCache()
	server.InitServer()
}
