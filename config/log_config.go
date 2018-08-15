package config

import (
	"github.com/Sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init()  {
	Log = log_config()
}

func log_config() * logrus.Logger {
	logger := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			ForceColors:true,
			FullTimestamp:true,
			TimestampFormat:"2006-01-02 15:04:05",
		},
		Level:logrus.InfoLevel,
	}
	return logger
}
