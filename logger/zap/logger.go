package leafZap

import (
	"github.com/labstack/gommon/log"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"go.uber.org/zap"
)

type (
	logger struct {
		instance *zap.SugaredLogger
		option   options
		logger   standardLogger
	}
)

func (l logger) Info(message leafLogger.Message) {
	if l.option.level != log.OFF {
		l.logWithField(message).Info(message.String())
	}
}

func (l logger) Warn(message leafLogger.Message) {
	if l.option.level != log.OFF {
		l.logWithField(message).Warn(message.String())
	}
}

func (l logger) Error(message leafLogger.Message) {
	if l.option.level != log.OFF {
		l.logWithField(message).Error(message.String())
	}
}

func (l logger) Debug(message leafLogger.Message) {
	if l.option.level != log.OFF {
		l.logWithField(message).Debug(message.String())
	}
}

func (l logger) StandardLogger() leafLogger.StandardLogger {
	return &l.logger
}

func (l logger) logWithField(message leafLogger.Message) *zap.SugaredLogger {
	logging := l.instance
	for key, message := range message {
		if l.option.masking != nil {
			message = l.option.masking.Encode(key, message)
		}
		logging = logging.With(key, message)
	}
	return logging
}

func DefaultLog() leafLogger.Logger {
	logger, _ := New()
	return logger
}

func New(opts ...Option) (leafLogger.Logger, error) {
	o := defaultOption()
	for _, opt := range opts {
		opt.Apply(&o)
	}

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = o.timeEncoder

	instance, err := zapConfig.Build(o.zapOptions...)
	if err != nil {
		return nil, err
	}
	return &logger{
		instance: instance.Sugar(),
		option:   o,
		logger: standardLogger{
			instance: instance.Sugar(),
			prefix:   o.prefix,
			level:    o.level,
		},
	}, nil
}
