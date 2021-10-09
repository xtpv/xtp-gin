package global

import (
	"github.com/xtp2217866847/xtp-gin/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var Log *logger.Logger

func initLogger() error {
	Log = logger.NewLogger(&lumberjack.Logger{
		Filename:  AppConfig.LogSavePath + "/" + AppConfig.LogFileName + AppConfig.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
