package leafServer

import (
	leafLogrus "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
)

type (
	ConsumerHook    func(consumer leafMQ.Consumer, logger leafLogger.Logger)
	consumerOptions struct {
		// Server
		enable bool
		logger leafLogger.Logger

		// Hooks
		register   ConsumerHook
		beforeRun  ConsumerHook
		afterRun   ConsumerHook
		beforeExit ConsumerHook
		afterExit  ConsumerHook
	}
	ConsumerOption interface {
		Apply(o *consumerOptions)
	}
)

func defaultMessagingOption() consumerOptions {
	return consumerOptions{
		enable:     false,
		logger:     leafLogrus.DefaultLog(),
		register:   nil,
		beforeRun:  nil,
		afterRun:   nil,
		beforeExit: nil,
		afterExit:  nil,
	}
}

type withConsumerEnable bool

func (w withConsumerEnable) Apply(o *consumerOptions) {
	o.enable = bool(w)
}

func WithConsumerEnable(enable bool) ConsumerOption {
	return withConsumerEnable(enable)
}

type withConsumerLogger struct{ leafLogger.Logger }

func (w withConsumerLogger) Apply(o *consumerOptions) {
	o.logger = w.Logger
}

func WithConsumerLogger(logger leafLogger.Logger) ConsumerOption {
	return withConsumerLogger{logger}
}

type withConsumerRegister ConsumerHook

func (w withConsumerRegister) Apply(o *consumerOptions) {
	o.register = ConsumerHook(w)
}

func WithConsumerRegister(hook ConsumerHook) ConsumerOption {
	return withConsumerRegister(hook)
}

type withConsumerBeforeRun ConsumerHook

func (w withConsumerBeforeRun) Apply(o *consumerOptions) {
	o.beforeRun = ConsumerHook(w)
}

func WithConsumerBeforeRun(hook ConsumerHook) ConsumerOption {
	return withConsumerBeforeRun(hook)
}

type withConsumerAfterRun ConsumerHook

func (w withConsumerAfterRun) Apply(o *consumerOptions) {
	o.afterRun = ConsumerHook(w)
}

func WithConsumerAfterRun(hook ConsumerHook) ConsumerOption {
	return withConsumerAfterRun(hook)
}

type withConsumerBeforeExit ConsumerHook

func (w withConsumerBeforeExit) Apply(o *consumerOptions) {
	o.beforeExit = ConsumerHook(w)
}

func WithConsumerBeforeExit(hook ConsumerHook) ConsumerOption {
	return withConsumerBeforeExit(hook)
}

type withConsumerAfterExit ConsumerHook

func (w withConsumerAfterExit) Apply(o *consumerOptions) {
	o.afterExit = ConsumerHook(w)
}

func WithConsumerAfterExit(hook ConsumerHook) ConsumerOption {
	return withConsumerAfterExit(hook)
}
