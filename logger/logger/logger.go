package leafLogger

import (
	"github.com/labstack/gommon/log"
	"io"
)

type (
	Logger interface {
		Info(message Message)
		Warn(message Message)
		Error(message Message)
		Debug(message Message)
		StandardLogger() StandardLogger
	}
	StandardLogger interface {
		Output() io.Writer
		SetOutput(w io.Writer)
		Prefix() string
		SetPrefix(prefix string)
		Level() log.Lvl
		SetLevel(v log.Lvl)
		SetHeader(header string)
		Print(i ...interface{})
		Println(i ...interface{})
		Printf(format string, i ...interface{})
		Printj(j log.JSON)
		Debug(i ...interface{})
		Debugf(format string, i ...interface{})
		Debugj(j log.JSON)
		Info(i ...interface{})
		Infof(format string, i ...interface{})
		Infoj(j log.JSON)
		Warn(i ...interface{})
		Warnf(format string, i ...interface{})
		Warnj(j log.JSON)
		Error(i ...interface{})
		Errorf(format string, i ...interface{})
		Errorj(j log.JSON)
		Fatal(i ...interface{})
		Fatalf(format string, i ...interface{})
		Fatalj(j log.JSON)
		Panic(i ...interface{})
		Panicf(format string, i ...interface{})
		Panicj(j log.JSON)
		Instance() interface{}
		Log(msg string)
	}
)

func GetLoggerLevel(level string) log.Lvl {
	switch level {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	default:
		return log.INFO
	}
}
