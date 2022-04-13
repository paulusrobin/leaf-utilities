package leafServer

import (
	leafWorker "github.com/enricodg/leaf-utilities/appRunner/worker"
	"os"
	"time"
)

type (
	WorkerServer struct {
		runners leafWorker.Runners
		option  workerOptions
		serverTemplate
	}
)

/*
	====================
		WorkerServer
	====================
*/
func NewWorker(opts ...WorkerOption) *WorkerServer {
	o := defaultWorkerOption()
	for _, opt := range opts {
		opt.Apply(&o)
	}

	return &WorkerServer{
		runners: make([]leafWorker.IRunner, 0),
		option:  o,
	}
}

func (s *WorkerServer) Serve(sig chan os.Signal) {
	if !s.option.enable {
		s.option.logger.StandardLogger().Info("[WORKER-SERVER] server disabled")
		return
	}

	s.serve(sig, serveParam{
		serve: func(sig chan os.Signal) {
			s.option.logger.StandardLogger().Info("[WORKER-SERVER] starting server")
			for _, runner := range s.runners {
				runner.Serve(sig, s.option.logger)
			}
			time.Sleep(time.Second)
		},
		register: func() {
			if s.option.register != nil {
				s.option.logger.StandardLogger().Debug("[WORKER-SERVER] starting register hooks")
				s.option.register(&s.runners, s.option.logger)
			}
		},
		beforeRun: func() {
			if s.option.beforeRun != nil {
				s.option.logger.StandardLogger().Debug("[WORKER-SERVER] starting before run hooks")
				s.option.beforeRun(&s.runners, s.option.logger)
			}
		},
		afterRun: func() {
			if s.option.afterRun != nil {
				s.option.logger.StandardLogger().Debug("[WORKER-SERVER] starting after run hooks")
				s.option.afterRun(&s.runners, s.option.logger)
			}
		},
	})
}

func (s *WorkerServer) Shutdown() {
	if !s.option.enable {
		return
	}

	s.shutdown(shutdownParam{
		shutdown: func() {
			s.option.logger.StandardLogger().Info("[WORKER-SERVER] shutting down server")
		},
		beforeExit: func() {
			if s.option.beforeExit != nil {
				s.option.logger.StandardLogger().Debug("[WORKER-SERVER] starting before exit hooks")
				s.option.beforeExit(&s.runners, s.option.logger)
			}
		},
		afterExit: func() {
			if s.option.afterExit != nil {
				s.option.logger.StandardLogger().Debug("[WORKER-SERVER] starting after exit hooks")
				s.option.afterExit(&s.runners, s.option.logger)
			}
		},
	})
}
