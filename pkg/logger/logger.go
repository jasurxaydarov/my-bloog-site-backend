package logger

import (
	"jasurxaydarov/my-bloog-site-backend/config"

	"github.com/saidamir98/udevs_pkg/logger"
)


type Log struct{
	logger.LoggerI
}

func(l *Log)CleanUp(){

	if err:=logger.Cleanup(l);err!= nil{
		l.Error("Faliled to cleanup logger ", logger.Error(err))
	}
}


func NewLogger(cfg config.GeneralConfig)Log{

	var loggerLevel string
	 switch cfg.Environment{

	 case config.DebugMode:
		loggerLevel=logger.LevelDebug

	 case config.TestMode:
		loggerLevel=logger.LevelDebug

	 default:
		loggerLevel=logger.LevelInfo

	 }

	logger:=logger.NewLogger(cfg.AppName,loggerLevel)
	return Log{logger}
}