package conf

import (
	"fmt"
	"senducode/tolog"
	"senducode/utils"
)

type server struct {
	Port  string
	Model string
}

var Server server

type cacheConf struct {
	Host string
	Port string
}

var CacheConf cacheConf

func InitConf() error {
	confjson, err := utils.JSONReader("./conf/conf.json")
	if err != nil {
		tolog.Log().Context(fmt.Sprintln("jsonReader", err)).Type(tolog.ToLogStatusError).PrintLog().Write()
		return err
	}
	serverMap := confjson["server"]
	server := utils.JSONConvertToMapString(serverMap)
	name := server["name"]
	version := server["version"]
	port := server["port"]
	model := server["model"]
	Server.Port = port
	Server.Model = model
	tolog.Log().Context("SendUCode-Server Conf Start").Type(tolog.ToLogStatusInfo).PrintLog().Write()
	tolog.Log().Context(fmt.Sprintln("ServerName:", name)).Type(tolog.ToLogStatusInfo).PrintLog()
	tolog.Log().Context(fmt.Sprintln("ServerVersion:", version)).Type(tolog.ToLogStatusInfo).PrintLog()
	tolog.Log().Context(fmt.Sprintln("ServerPort:", port)).Type(tolog.ToLogStatusInfo).PrintLog()
	tolog.Log().Context(fmt.Sprintln("Running on model:", model)).Type(tolog.ToLogStatusInfo).PrintLog()
	cacheMap := confjson["redis"]
	cache := utils.JSONConvertToMapString(cacheMap)
	cacheHost := cache["host"]
	cachePort := cache["port"]
	CacheConf.Host = cacheHost
	CacheConf.Port = cachePort
	return nil
}
