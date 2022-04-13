package leafRunner

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type runner struct {
	runners []IServable
}

func (r *runner) With(runner IServable) IRunner {
	r.runners = append(r.runners, runner)
	return r
}

func (r *runner) Run() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	for _, servable := range r.runners {
		go func(runner IServable, sig chan os.Signal) {
			runner.Serve(sig)
		}(servable, sig)
	}
	<-sig

	var wg sync.WaitGroup
	wg.Add(len(r.runners))
	for _, servable := range r.runners {
		go func(runner IServable) {
			runner.Shutdown()
			wg.Done()
		}(servable)
	}
	wg.Wait()
}

func New() IRunner {
	return &runner{
		runners: make([]IServable, 0),
	}
}
