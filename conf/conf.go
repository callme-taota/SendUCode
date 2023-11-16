package conf

import (
	"senducode/tolog"
	"senducode/utils"
	"fmt"
)

func InitConf() (error) {
	confjson , err := utils.JSONReader("./conf/conf.json")
	if err != nil {
		tolog.WriteLog(fmt.Sprintln("jsonReader",err),tolog.ToLogStatusWarning)
		return err
	}
	var server map[string]string
	server = confjson["server"].(map[string]interface{})
	name := server["name"]
	version := server["version"]
	port := server["port"]
	model := server["model"]
	tolog.WriteLog(fmt.Sprintln("SendUCode-Server Conf Start"),tolog.ToLogStatusInfo)
	tolog.WriteLog(fmt.Sprintln("ServerName:",name),tolog.ToLogStatusInfo)
	tolog.WriteLog(fmt.Sprintln("ServerVersion:",version),tolog.ToLogStatusInfo)
	tolog.WriteLog(fmt.Sprintln("ServerPort:",port),tolog.ToLogStatusInfo)
	tolog.WriteLog(fmt.Sprintln("Running on model:",model),tolog.ToLogStatusInfo)
	return nil
}