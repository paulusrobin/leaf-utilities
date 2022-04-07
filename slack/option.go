package leafSlack

import (
	"fmt"
	"time"
)

type (
	Option interface {
		Apply(o *option)
	}
	option struct {
		hook    string
		timeout time.Duration
	}
)

func (o option) validate() error {
	if "" == o.hook {
		return fmt.Errorf("hook is required")
	}
	return nil
}

type withHook string

func (w withHook) Apply(o *option) {
	o.hook = string(w)
}

func WithHook(hook string) Option {
	return withHook(hook)
}

type withTimeout time.Duration

func (w withTimeout) Apply(o *option) {
	o.timeout = time.Duration(w)
}

func WithTimeout(timeout time.Duration) Option {
	return withTimeout(timeout)
}
