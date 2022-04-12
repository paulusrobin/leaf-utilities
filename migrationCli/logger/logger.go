package logger

import (
	"github.com/labstack/gommon/log"
	leafLogrus "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/migrationCli/config"
	"sync"
)

func GetLogger() leafLogger.Logger {
	var once sync.Once
	var logger leafLogger.Logger
	var err error
	once.Do(func() {
		configuration := config.GetConfig()
		formatter, err := leafLogrus.GetLoggerFormatter("TEXT")
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
