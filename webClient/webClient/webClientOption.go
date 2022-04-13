package leafWebClient

import (
	"github.com/paulusrobin/leaf-utilities/webClient/webClient/circuitBreaker"
	"time"
)

type Option func(*WebClientOption)

type WebClientOption struct {
	timeout        time.Duration
	retryCount     int
	circuitBreaker func() []leafCircuitBreaker.Option
}

func NewDefaultWebClientOption(timeout time.Duration) Option {
	return withTimeout(timeout)
}

func NewWebClientOptionWithRetry(timeout time.Duration, retryCount int) []Option {
	return []Option{withTimeout(timeout), withRetryCount(retryCount)}
}

func NewWebClientOptionWithCircuitBreaker(opts ...leafCircuitBreaker.Option) Option {
	return withCircuitBreaker(opts...)
}

func withTimeout(timeout time.Duration) Option {
	return func(wc *WebClientOption) {
		wc.timeout = timeout
	}
}

func (wc *WebClientOption) GetTimeout() time.Duration {
	return wc.timeout
}

func withRetryCount(retryCount int) Option {
	return func(wc *WebClientOption) {
		wc.retryCount = retryCount
	}
}

func (wc *WebClientOption) GetRetryCount() int {
	return wc.retryCount
}

func withCircuitBreaker(opts ...leafCircuitBreaker.Option) Option {
	return func(wc *WebClientOption) {
		wc.circuitBreaker = func() []leafCircuitBreaker.Option {
			return opts
		}
	}
}

func (wc *WebClientOption) GetCircuitBreaker() func() []leafCircuitBreaker.Option {
	return wc.circuitBreaker
}
