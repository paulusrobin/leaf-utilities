package leafSql

import (
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"time"
)

type (
	ConnectionOption interface {
		Apply(o *ConnectionOptions)
	}
	ConnectionOptions struct {
		MaxIdleConnection, MaxOpenConnection int
		ConnMaxLifetime                      time.Duration
		LogMode                              bool
		logger                               leafLogger.Logger
	}
)

func (o ConnectionOptions) Logger() leafLogger.Logger {
	return o.logger
}

type withMaxIdleConnection int

func (w withMaxIdleConnection) Apply(o *ConnectionOptions) {
	o.MaxIdleConnection = int(w)
}

func WithMaxIdleConnection(maxIdleConnection int) ConnectionOption {
	return withMaxIdleConnection(maxIdleConnection)
}

type withMaxOpenConnection int

func (w withMaxOpenConnection) Apply(o *ConnectionOptions) {
	o.MaxOpenConnection = int(w)
}

func WithMaxOpenConnection(maxOpenConnection int) ConnectionOption {
	return withMaxOpenConnection(maxOpenConnection)
}

type withLogMode bool

func (w withLogMode) Apply(o *ConnectionOptions) {
	o.LogMode = bool(w)
}

func WithLogMode(logMode bool) ConnectionOption {
	return withLogMode(logMode)
}

type withConnMaxLifetime time.Duration

func (w withConnMaxLifetime) Apply(o *ConnectionOptions) {
	o.ConnMaxLifetime = time.Duration(w)
}

func WithConnMaxLifetime(connMaxLifetime time.Duration) ConnectionOption {
	return withConnMaxLifetime(connMaxLifetime)
}

type withLogger struct{ leafLogger.Logger }

func (w withLogger) Apply(o *ConnectionOptions) {
	o.logger = w.Logger
}

func WithLogger(logger leafLogger.Logger) ConnectionOption {
	return withLogger{logger}
}
