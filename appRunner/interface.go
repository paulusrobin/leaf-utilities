package leafRunner

import "os"

type (
	IServable interface {
		Serve(sig chan os.Signal)
		Shutdown()
	}
	IRunner interface {
		With(runner IServable) IRunner
		Run()
	}
)
