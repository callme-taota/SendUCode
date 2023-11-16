package tolog

import (
	"fmt"
	"os"
	"strings"
	"time"
)


const (
	ToLogStatusInfo = "info"
	ToLogStatusWarning = "warning"
	ToLogStatusError = "error"
	ToLogStatusDebug = "debug"
	ToLogStatusNotice = "notice"
)

// PrintLog in the terminal
// called before write the tolog
func PrintLog(str string) {
	fmt.Println(str)
}

// init the logFile to always open while server running
var logFile *os.File

// initLog init logfile function
func initLog() {
	//getDay
	currentDay := time.Now().Format("2006-01-02")
	logFilePath := "./tolog/logs/painter-blog-tolog-" + currentDay + ".tolog"

	//写入logFile
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("[error]", err)
	}
	logFile = file
}

// CloseLogFile should call when main function done or the server closed
func CloseLogFile() {
	logFile.Close()
}

// WriteLog main part of tolog to write tolog in the file
// tolog file sep by days each day will create a new tolog file
func WriteLog(content, level string) {

	if logFile == nil {
		initLog()
	}

	//check level
	level = strings.ToLower(level)
	if level != "info" && level != "warning" && level != "error" && level != "debug" && level != "notice" {
		level = "Unknow"
	}

	//create content
	content = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + "[" + level + "] " + content

	//print tolog
	PrintLog(content)

	// 写入内容到文件
	logFile.WriteString(content + "\n")

}
