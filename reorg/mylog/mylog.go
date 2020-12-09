package mylog

import (
	"log"
	"os"
	"time"
)

func MkLogDir() {
	// 断是否有logs路径,没有则进行创建
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		err = os.Mkdir("logs", 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NewLogger(name string) (*os.File, *log.Logger) {
	MkLogDir()
	time := time.Now().Format("20060102_150405")
	logFileName := "logs/" + name + "_" + time + ".log"
	logFile, _ := os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	return logFile, log.New(logFile, "", log.Ldate | log.Ltime)
}
