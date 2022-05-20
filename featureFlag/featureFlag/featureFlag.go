package featureFlag

import (
	"fmt"
	"sync"
	"time"
)

type (
	FeatureFlag interface {
		Get(key string) interface{}
	}
	implementation struct {
		serviceName string
		option      option
		data        map[string]interface{}
		sync.Mutex
	}
)

func (f *implementation) init() {
	fmt.Println("init")
	
	f.Lock()
	defer f.Unlock()
	// call feature flag
	var data = map[string]interface{}{
		"listing": "v2",
	}
	f.data = data
}

func (f *implementation) Get(key string) interface{} {
	f.Lock()
	defer f.Unlock()
	return f.data[key]
}

func New(serviceName string, opts ...Option) (FeatureFlag, error) {
	impl := &implementation{serviceName: serviceName, option: defaultOption}

	impl.init()
	if impl.option.fnValidate != nil {
		if err := impl.option.fnValidate(impl.data); err != nil {
			return nil, err
		}
	}

	if impl.option.periodicallyUpdate.interval > 0 {
		go func() {
			for {
				select {
				case <-time.Tick(impl.option.periodicallyUpdate.interval):
					impl.init()
				}
			}
		}()
	}
	return impl, nil
}
