package tolog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	ToLogStatusInfo    = "info"
	ToLogStatusWarning = "warning"
	ToLogStatusError   = "error"
	ToLogStatusDebug   = "debug"
	ToLogStatusNotice  = "notice"
	ToLogStatusUnknown = "unknown"
)

var logFile *os.File

type ToLog struct {
	logType    string
	logContext string
	logTime    string
	FullLog    string
}

type ToLogOptions func(l *ToLog)

func Log() *ToLog {
	tolog := &ToLog{}
	return tolog
}

func (l *ToLog) Context(ctx string) *ToLog {
	l.logContext = ctx
	return l
}

func (l *ToLog) Type(level string) *ToLog {
	level = strings.ToLower(level)
	if level != ToLogStatusInfo && level != ToLogStatusWarning && level != ToLogStatusError && level != ToLogStatusNotice && level != ToLogStatusDebug {
		level = ToLogStatusUnknown
	}
	l.logType = level
	return l
}

func (l *ToLog) PrintLog() *ToLog {
	CreateFullLog(l)
	fmt.Println(l.FullLog)
	return l
}

func CreateFullLog(l *ToLog) {
	fullLog := "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + "[" + l.logType + "] " + l.logContext
	l.FullLog = fullLog
}

func (l *ToLog) Write() {
	if logFile == nil {
		initLog()
	}
	logFile.WriteString(l.FullLog + "\n")
}

func initLog() {
	currentDay := time.Now().Format("2006-01-02")
	logFilePath := "./tolog/logs/painter-blog-tolog-" + currentDay + ".tolog"

	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("[error]", err)
	}
	logFile = file
}
func CloseLogFile() {
	logFile.Close()
}
