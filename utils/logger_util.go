package utils

import (
	"os"
	"time"
)

func LogFile() *os.File {
	today := time.Now().Format("2006-01-02")
	fileNmae :="D:/data/okay_logs/"+today + ".log"
	file,err := os.OpenFile(fileNmae,os.O_CREATE|os.O_WRONLY|os.O_APPEND, 102400)
	if err != nil {
		panic(err)
	}
	return file
}
