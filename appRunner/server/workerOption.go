package leafServer

import (
	leafWorker "github.com/enricodg/leaf-utilities/appRunner/worker"
	leafLogrus "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

type (
	WorkerHook    func(runners *leafWorker.Runners, logger leafLogger.Logger)
	workerOptions struct {
		// Server
		enable bool
		logger leafLogger.Logger

		// Hooks
		register   WorkerHook
		beforeRun  WorkerHook
		afterRun   WorkerHook
		beforeExit WorkerHook
		afterExit  WorkerHook
	}
	WorkerOption interface {
		Apply(o *workerOptions)
	}
)

func defaultWorkerOption() workerOptions {
	return workerOptions{
		logger:     leafLogrus.DefaultLog(),
		enable:     false,
		register:   nil,
		beforeRun:  nil,
		afterRun:   nil,
		beforeExit: nil,
		afterExit:  nil,
	}
}

type withWorkerEnable bool

func (w withWorkerEnable) Apply(o *workerOptions) {
	o.enable = bool(w)
}

func WithWorkerEnable(enable bool) WorkerOption {
	return withWorkerEnable(enable)
}

type withWorkerLogger struct{ leafLogger.Logger }

func (w withWorkerLogger) Apply(o *workerOptions) {
	o.logger = w.Logger
}

func WithWorkerLogger(logger leafLogger.Logger) WorkerOption {
	return withWorkerLogger{logger}
}

type withWorkerRegister WorkerHook

func (w withWorkerRegister) Apply(o *workerOptions) {
	o.register = WorkerHook(w)
}

func WithWorkerRegister(hook WorkerHook) WorkerOption {
	return withWorkerRegister(hook)
}

type withWorkerBeforeRun WorkerHook

func (w withWorkerBeforeRun) Apply(o *workerOptions) {
	o.beforeRun = WorkerHook(w)
}

func WithWorkerBeforeRun(hook WorkerHook) WorkerOption {
	return withWorkerBeforeRun(hook)
}

type withWorkerAfterRun WorkerHook

func (w withWorkerAfterRun) Apply(o *workerOptions) {
	o.afterRun = WorkerHook(w)
}

func WithWorkerAfterRun(hook WorkerHook) WorkerOption {
	return withWorkerAfterRun(hook)
}

type withWorkerBeforeExit WorkerHook

func (w withWorkerBeforeExit) Apply(o *workerOptions) {
	o.beforeExit = WorkerHook(w)
}

func WithWorkerBeforeExit(hook WorkerHook) WorkerOption {
	return withWorkerBeforeExit(hook)
}

type withWorkerAfterExit WorkerHook

func (w withWorkerAfterExit) Apply(o *workerOptions) {
	o.afterExit = WorkerHook(w)
}

func WithWorkerAfterExit(hook WorkerHook) WorkerOption {
	return withWorkerAfterExit(hook)
}
