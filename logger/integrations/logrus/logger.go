package leafLogrus

import (
	"fmt"
	"github.com/labstack/gommon/log"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
)

type (
	logger struct {
		instance *logrus.Logger
		option   options
		logger   standardLogger
	}
)

func (l logger) Info(message leafLogger.Message) {
	if l.option.level != log.OFF {
		logging := l.instance
		logging.WithFields(l.logWithField(message)).
			Info(message.String())
	}
}

func (l logger) Warn(message leafLogger.Message) {
	if l.option.level != log.OFF {
		logging := l.instance
		logging.WithFields(l.logWithField(message)).
			Warn(message.String())
	}
}

func (l logger) Error(message leafLogger.Message) {
	if l.option.level != log.OFF {
		logging := l.instance
		logging.WithFields(l.logWithField(message)).
			Error(message.String())
	}
}

func (l logger) Debug(message leafLogger.Message) {
	if l.option.level != log.OFF {
		logging := l.instance
		logging.WithFields(l.logWithField(message)).
			Debug(message.String())
	}
}

func (l logger) StandardLogger() leafLogger.StandardLogger {
	return &l.logger
}

func (l logger) logWithField(message leafLogger.Message) map[string]interface{} {
	var messageData = make(map[string]interface{})
	for key, message := range message {
		if l.option.masking != nil {
			message = l.option.masking.Encode(key, message)
		}
		messageData[key] = message
	}
	return messageData
}

func DefaultLog() leafLogger.Logger {
	logger, _ := New()
	return logger
}

func GetLoggerFormatter(formatter string) (Formatter, error) {
	if string(JSONFormatter) == formatter {
		return JSONFormatter, nil
	}

	if string(TextFormatter) == formatter {
		return TextFormatter, nil
	}

	return "", fmt.Errorf("invalid log format")
}

func New(opts ...Option) (leafLogger.Logger, error) {
	o := defaultOption()
	for _, opt := range opts {
		opt.Apply(&o)
	}

	instance := logrus.New()

	switch o.level {
	case log.INFO:
		instance.Level = logrus.InfoLevel
		break
	case log.DEBUG:
		instance.Level = logrus.DebugLevel
		break
	case log.WARN:
		instance.Level = logrus.WarnLevel
		break
	case log.ERROR:
		instance.Level = logrus.ErrorLevel
		break
	default:
		instance.Level = logrus.ErrorLevel
		break
	}

	var formatter logrus.Formatter

	if o.formatter == JSONFormatter {
		formatter = &logrus.JSONFormatter{}
	} else {
		formatter = &logrus.TextFormatter{}
	}

	instance.Formatter = formatter

	// - check if log file path does exists
	if o.logFilePath != "" {
		if _, err := os.Stat(o.logFilePath); os.IsNotExist(err) {
			if _, err = os.Create(o.logFilePath); err != nil {
				return nil, fmt.Errorf("failed to create log file %s\nerror: %+v", o.logFilePath, err)
			}
		}
		maps := lfshook.PathMap{
			logrus.InfoLevel:  o.logFilePath,
			logrus.DebugLevel: o.logFilePath,
			logrus.ErrorLevel: o.logFilePath,
		}
		instance.Hooks.Add(lfshook.NewHook(maps, formatter))
	}

	return &logger{
		instance: instance,
		option:   o,
		logger: standardLogger{
			instance: instance,
			level:    o.level,
			prefix:   o.prefix,
		},
	}, nil
}

func getLevel(lvl log.Lvl) logrus.Level {
	switch lvl {
	case log.INFO:
		return logrus.InfoLevel
	case log.DEBUG:
		return logrus.DebugLevel
	case log.WARN:
		return logrus.WarnLevel
	case log.ERROR:
		return logrus.ErrorLevel
	default:
		return logrus.ErrorLevel
	}
}
