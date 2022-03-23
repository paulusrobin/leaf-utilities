package leafWebClient

import (
	"github.com/enricodg/leaf-utilities/webClient/webClient/circuitBreaker"
	"time"
)

type Option func(*WebClientOption)

type WebClientOption struct {
	timeout        time.Duration
	retryCount     int
	circuitBreaker func() []leafCircuitBreaker.Option
}

func NewDefaultWebClientOption(timeout time.Duration) Option {
	return WithTimeout(timeout)
}

func NewWebClientOptionWithRetry(timeout time.Duration, retryCount int) []Option {
	return []Option{WithTimeout(timeout), WithRetryCount(retryCount)}
}

func NewWebClientOptionWithCircuitBreaker(opts ...leafCircuitBreaker.Option) Option {
	return WithCircuitBreaker(opts...)
}

func WithTimeout(timeout time.Duration) Option {
	return func(wc *WebClientOption) {
		wc.timeout = timeout
	}
}

func (wc *WebClientOption) GetTimeout() time.Duration {
	return wc.timeout
}

func WithRetryCount(retryCount int) Option {
	return func(wc *WebClientOption) {
		wc.retryCount = retryCount
	}
}

func (wc *WebClientOption) GetRetryCount() int {
	return wc.retryCount
}

func WithCircuitBreaker(opts ...leafCircuitBreaker.Option) Option {
	return func(wc *WebClientOption) {
		wc.circuitBreaker = func() []leafCircuitBreaker.Option {
			return opts
		}
	}
}

func (wc *WebClientOption) GetCircuitBreaker() func() []leafCircuitBreaker.Option {
	return wc.circuitBreaker
}
