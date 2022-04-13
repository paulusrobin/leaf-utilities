package leafServer

import (
	"cloud.google.com/go/profiler"
	"os"
)

type (
	ProfilerServer struct {
		serviceName, serviceVersion string
		option                      profilerOptions
		serverTemplate
	}
)

func NewProfiler(serviceName, serviceVersion string, opts ...ProfilerOption) *ProfilerServer {
	o := defaultProfilerOption()
	for _, opt := range opts {
		opt.Apply(&o)
	}
	return &ProfilerServer{
		serviceName:    serviceName,
		serviceVersion: serviceVersion,
		option:         o,
	}
}

func (p *ProfilerServer) Serve(sig chan os.Signal) {
	if !p.option.enable {
		p.option.logger.StandardLogger().Info("[PROFILER-SERVER] server disabled")
		return
	}

	config := profiler.Config{
		Service:        p.serviceName,
		ServiceVersion: p.serviceVersion,
	}

	if p.option.projectID != "" {
		config.ProjectID = p.option.projectID
	}

	p.serve(sig, serveParam{
		serve: func(sig chan os.Signal) {
			p.option.logger.StandardLogger().Info("[PROFILER-SERVER] starting server")
			if err := profiler.Start(config); err != nil {
				p.option.logger.StandardLogger().Error("[PROFILER-SERVER] failed to start profiler")
			}
		},
		register:  func() {},
		beforeRun: func() {},
		afterRun:  func() {},
	})
}

func (p *ProfilerServer) Shutdown() {
	if !p.option.enable {
		return
	}

	p.shutdown(shutdownParam{
		shutdown: func() {
			p.option.logger.StandardLogger().Info("[PROFILER-SERVER] shutting down server")
		},
		beforeExit: func() {},
		afterExit:  func() {},
	})
}
