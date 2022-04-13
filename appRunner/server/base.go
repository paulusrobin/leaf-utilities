package leafServer

import (
	"os"
)

type (
	serveParam struct {
		serve     func(sig chan os.Signal)
		register  func()
		beforeRun func()
		afterRun  func()
	}
	shutdownParam struct {
		shutdown   func()
		beforeExit func()
		afterExit  func()
	}
	serverTemplate struct{}
)

func (s serverTemplate) serve(sig chan os.Signal, param serveParam) {
	param.register()
	param.beforeRun()
	param.serve(sig)
	param.afterRun()
}

func (s serverTemplate) shutdown(param shutdownParam) {
	param.beforeExit()
	param.shutdown()
	param.afterExit()
}
