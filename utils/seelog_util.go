package utils

import (
	log "github.com/cihub/seelog"
)

var Logger log.LoggerInterface

func init() {
	defer log.Flush()
	//加载配置文件
	Logger, _ = log.LoggerFromConfigAsFile("config/seelog.xml")

	//替换记录器
	log.ReplaceLogger(Logger)
}
