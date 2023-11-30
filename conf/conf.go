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
	Host     string
	Port     string
	Password string
	DB       string
}

var CacheConf cacheConf

func InitConf() error {
	confjson, err := utils.JSONReader("./conf/conf.json")
	if err != nil {
		tolog.Log().Error(fmt.Sprintf("jsonReader%e", err)).PrintAndWriteSafe()
		return err
	}

	serverMap := confjson["server"]
	server := utils.JSONConvertToMapString(serverMap)
	name, version, port, model := server["name"], server["version"], server["port"], server["model"]
	Server.Port, Server.Model = port, model

	tolog.Log().Info("SendUCode-Server Conf Start").PrintAndWriteSafe()
	tolog.Log().Infof("ServerName:%s", name).PrintAndWriteSafe()
	tolog.Log().Infof("ServerVersion:%s", version).PrintAndWriteSafe()
	tolog.Log().Infof("ServerPort:%s", port).PrintAndWriteSafe()
	tolog.Log().Infof("Running on model:%s", model).PrintAndWriteSafe()

	cacheMap := confjson["redis"]
	cache := utils.JSONConvertToMapString(cacheMap)
	cacheHost, cachePort, cachePassword, cacheDB := cache["host"], cache["port"], cache["password"], cache["db"]
	CacheConf.Host, CacheConf.Port, CacheConf.Password, CacheConf.DB = cacheHost, cachePort, cachePassword, cacheDB

	return nil
}
