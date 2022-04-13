package leafServer

import (
	leafLogrus "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

type (
	profilerOptions struct {
		// Server
		enable    bool
		logger    leafLogger.Logger
		projectID string
	}
	ProfilerOption interface {
		Apply(o *profilerOptions)
	}
)

func defaultProfilerOption() profilerOptions {
	return profilerOptions{
		enable:    false,
		logger:    leafLogrus.DefaultLog(),
		projectID: "",
	}
}

type withProfilerEnable bool

func (w withProfilerEnable) Apply(o *profilerOptions) {
	o.enable = bool(w)
}

func WithProfilerEnable(enable bool) ProfilerOption {
	return withProfilerEnable(enable)
}

type withProfilerLogger struct{ leafLogger.Logger }

func (w withProfilerLogger) Apply(o *profilerOptions) {
	o.logger = w.Logger
}

func WithProfilerLogger(logger leafLogger.Logger) ProfilerOption {
	return withProfilerLogger{logger}
}

type withProfilerProjectID string

func (w withProfilerProjectID) Apply(o *profilerOptions) {
	o.projectID = string(w)
}

func WithProfilerProjectID(projectID string) ProfilerOption {
	return withProfilerProjectID(projectID)
}
