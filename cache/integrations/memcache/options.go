package leafMemcache

import (
	leafCache "github.com/paulusrobin/leaf-utilities/cache/cache"
	leafLogrus "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

type (
	Option interface {
		Apply(o *option)
	}
	option struct {
		maxEntrySize       int
		maxEntriesKey      int
		maxEntriesInWindow int
		onRemove           func(key string, value interface{})
		onRemoveWithReason func(key string, reason string)
		logger             leafLogger.Logger
	}
)

func defaultOption() option {
	return option{
		maxEntrySize:       1024 * 1024,
		maxEntriesKey:      leafCache.Infinite,
		maxEntriesInWindow: 2 * 1024 * 1024 * 1024,
		onRemove:           nil,
		onRemoveWithReason: nil,
		logger:             leafLogrus.DefaultLog(),
	}
}

type withMaxEntrySize int

func (w withMaxEntrySize) Apply(o *option) {
	o.maxEntrySize = int(w)
}

func WithMaxEntrySize(maxEntrySize int) Option {
	return withMaxEntrySize(maxEntrySize)
}

type withMaxEntriesKey int

func (w withMaxEntriesKey) Apply(o *option) {
	o.maxEntriesKey = int(w)
}

func WithMaxEntriesKey(maxEntriesKey int) Option {
	return withMaxEntriesKey(maxEntriesKey)
}

type withMaxEntriesInWindow int

func (w withMaxEntriesInWindow) Apply(o *option) {
	o.maxEntriesInWindow = int(w)
}

func WithMaxEntriesInWindow(maxEntriesKey int) Option {
	return withMaxEntriesInWindow(maxEntriesKey)
}

type withOnRemove func(key string, value interface{})

func (w withOnRemove) Apply(o *option) {
	o.onRemove = w
}

func WithOnRemove(f func(key string, value interface{})) Option {
	return withOnRemove(f)
}

type withOnRemoveWithReason func(key string, reason string)

func (w withOnRemoveWithReason) Apply(o *option) {
	o.onRemoveWithReason = w
}

func WithOnRemoveWithReason(f func(key string, reason string)) Option {
	return withOnRemoveWithReason(f)
}

type withLogger struct{ leafLogger.Logger }

func (w withLogger) Apply(o *option) {
	o.logger = w.Logger
}

func WithLogger(log leafLogger.Logger) Option {
	return withLogger{log}
}
