package leafCircuitBreaker

import (
	"time"
)

// StatsdCollectorConfig provides configuration that the Statsd client will need.
type StatsdCollectorConfig struct {
	// StatsdAddr is the tcp address of the Statsd server
	StatsdAddr string
	// Prefix is the prefix that will be prepended to all metrics sent from this collector.
	Prefix string
}

type fallbackFunc func(error) error

type Option func(*CircuitBreakerOption)

type CircuitBreakerOption struct {
	timeout                   time.Duration
	circuitBreakerTimeout     time.Duration
	circuitBreakerCommandName string
	maxConcurrentRequests     int
	requestVolumeThreshold    int
	sleepWindow               int
	errorPercentThreshold     int
	retryCount                int
	fallbackFunc              func(err error) error
	statsD                    *StatsdCollectorConfig
	retryBackoffOption        *RetryBackoffOption
}

func WithCommandName(name string) Option {
	return func(c *CircuitBreakerOption) {
		c.circuitBreakerCommandName = name
	}
}

func (c *CircuitBreakerOption) GetCommandName() string {
	return c.circuitBreakerCommandName
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *CircuitBreakerOption) {
		c.timeout = timeout
	}
}

func (c *CircuitBreakerOption) GetTimeout() time.Duration {
	return c.timeout
}

func WithCircuitBreakerTimeout(timeout time.Duration) Option {
	return func(c *CircuitBreakerOption) {
		c.circuitBreakerTimeout = timeout
	}
}

func (c *CircuitBreakerOption) GetCircuitBreakerTimeout() time.Duration {
	return c.circuitBreakerTimeout
}

func WithMaxConcurrentRequests(maxConcurrentRequests int) Option {
	return func(c *CircuitBreakerOption) {
		c.maxConcurrentRequests = maxConcurrentRequests
	}
}

func (c *CircuitBreakerOption) GetMaxConcurrentRequests() int {
	return c.maxConcurrentRequests
}

func WithRequestVolumeThreshold(requestVolumeThreshold int) Option {
	return func(c *CircuitBreakerOption) {
		c.requestVolumeThreshold = requestVolumeThreshold
	}
}

func (c *CircuitBreakerOption) GetRequestVolumeThreshold() int {
	return c.requestVolumeThreshold
}

func WithSleepWindow(sleepWindow int) Option {
	return func(c *CircuitBreakerOption) {
		c.sleepWindow = sleepWindow
	}
}

func (c *CircuitBreakerOption) GetSleepWindow() int {
	return c.sleepWindow
}

func WithErrorPercentThreshold(errorPercentThreshold int) Option {
	return func(c *CircuitBreakerOption) {
		c.errorPercentThreshold = errorPercentThreshold
	}
}

func (c *CircuitBreakerOption) GetErrorPercentThreshold() int {
	return c.errorPercentThreshold
}

func WithFallbackFunc(fn fallbackFunc) Option {
	return func(c *CircuitBreakerOption) {
		c.fallbackFunc = fn
	}
}

func (c *CircuitBreakerOption) GetFallbackFunc() func(err error) error {
	return c.fallbackFunc
}

func WithRetryCount(retryCount int) Option {
	return func(c *CircuitBreakerOption) {
		c.retryCount = retryCount
	}
}

func (c *CircuitBreakerOption) GetRetryCount() int {
	return c.retryCount
}

func WithStatsDCollector(addr, prefix string) Option {
	return func(c *CircuitBreakerOption) {
		c.statsD = &StatsdCollectorConfig{StatsdAddr: addr, Prefix: prefix}
	}
}

func (c *CircuitBreakerOption) GetStatsDCollector() *StatsdCollectorConfig {
	return c.statsD
}

func WithRetryBackoffOption(retryBackoffOption RetryBackoffOption) Option {
	return func(c *CircuitBreakerOption) {
		c.retryBackoffOption = &retryBackoffOption
	}
}

func (c *CircuitBreakerOption) GetRetryBackoffOption() *RetryBackoffOption {
	return c.retryBackoffOption
}
