package leafLogrus

import (
	"github.com/labstack/gommon/log"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

const (
	JSONFormatter Formatter = "JSON"
	TextFormatter Formatter = "TEXT"
)

type (
	Formatter string
	options   struct {
		level       log.Lvl
		prefix      string
		formatter   Formatter
		logFilePath string
		masking     leafLogger.MaskedEncoder
	}
	Option interface {
		Apply(o *options)
	}
)

func defaultOption() options {
	return options{
		level:       log.INFO,
		prefix:      "",
		formatter:   TextFormatter,
		logFilePath: "",
		masking:     make(map[string]leafLogger.Masked),
	}
}

type withLevel struct{ log.Lvl }

func (w withLevel) Apply(o *options) {
	o.level = w.Lvl
}

func WithLevel(lvl log.Lvl) Option {
	return &withLevel{lvl}
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

type withLogFilePath string

func (w withLogFilePath) Apply(o *options) {
	o.logFilePath = string(w)
}

func WithLogFilePath(path string) Option {
	return withLogFilePath(path)
}

type withFormatter string

func (w withFormatter) Apply(o *options) {
	o.formatter = Formatter(w)
}

func WithFormatter(formatter Formatter) Option {
	return withFormatter(formatter)
}
