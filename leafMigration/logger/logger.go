package logger

import (
	"github.com/labstack/gommon/log"
	"github.com/paulusrobin/leaf-utilities/leafMigration/config"
	leafLogrus "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"sync"
)

func GetLogger() leafLogger.Logger {
	var once sync.Once
	var logger leafLogger.Logger
	var err error
	once.Do(func() {
		configuration := config.GetConfig()
		formatter, err := leafLogrus.GetLoggerFormatter(configuration.LogFormatter)
		if err != nil {
			return
		}

		logger, err = leafLogrus.New(
			leafLogrus.WithFormatter(formatter),
			leafLogrus.WithLogFilePath(configuration.LogFilePath),
			leafLogrus.WithLevel(log.INFO))
	})
	if err != nil {
		panic(err)
	}
	return logger
}
