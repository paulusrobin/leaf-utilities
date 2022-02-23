package leafZap

import (
	"github.com/labstack/gommon/log"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	options struct {
		level       log.Lvl
		prefix      string
		masking     leafLogger.MaskedEncoder
		zapOptions  []zap.Option
		timeEncoder zapcore.TimeEncoder
	}
	Option interface {
		Apply(o *options)
	}
)

func defaultOption() options {
	return options{
		level:       log.INFO,
		prefix:      "",
		masking:     make(map[string]leafLogger.Masked),
		zapOptions:  make([]zap.Option, 0),
		timeEncoder: zapcore.RFC3339TimeEncoder,
	}
}

type withLevel struct{ log.Lvl }

func (w withLevel) Apply(o *options) {
	o.level = w.Lvl
}

func WithLevel(lvl log.Lvl) Option {
	return &withLevel{lvl}
}

type withTimeEncoder struct{ zapcore.TimeEncoder }

func (w withTimeEncoder) Apply(o *options) {
	o.timeEncoder = w.TimeEncoder
}

func WithTimeEncoder(timeEncoder zapcore.TimeEncoder) Option {
	return &withTimeEncoder{timeEncoder}
}

type withPrefix string

func (w withPrefix) Apply(o *options) {
	o.prefix = string(w)
}

func WithPrefix(prefix string) Option {
	return withPrefix(prefix)
}

type withMasking struct {
	key    string
	masked leafLogger.Masked
}

func (w withMasking) Apply(o *options) {
	if _, found := o.masking[w.key]; !found {
		o.masking[w.key] = w.masked
	}
}

func WithMasking(key string, masked leafLogger.Masked) Option {
	return &withMasking{key: key, masked: masked}
}

type withZapOption struct{ zap.Option }

func (w withZapOption) Apply(o *options) {
	o.zapOptions = append(o.zapOptions, w.Option)
}

func WithZapOption(opt zap.Option) Option {
	return &withZapOption{opt}
}

