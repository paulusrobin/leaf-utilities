package leafCircuitBreaker

import "time"

type RetryBackoffType uint

const (
	Constant RetryBackoffType = iota + 1
	Exponential
)

type RetryBackoffOption struct {
	backoffType           RetryBackoffType
	backoffInterval       time.Duration
	maximumJitterInterval time.Duration
	initialTimeout        time.Duration
	maxTimeout            time.Duration
	exponentFactor        float64
}

func NewConstantRetryBackoff(backoffInterval, maximumJitterInterval time.Duration) *RetryBackoffOption {
	return &RetryBackoffOption{
		backoffType:           Constant,
		backoffInterval:       backoffInterval,
		maximumJitterInterval: maximumJitterInterval,
	}
}

func NewExponentialRetryBackoff(initialTimeout, maxTimeout time.Duration, exponentFactor float64, maximumJitterInterval time.Duration) *RetryBackoffOption {
	return &RetryBackoffOption{
		backoffType:           Exponential,
		maximumJitterInterval: maximumJitterInterval,
		exponentFactor:        exponentFactor,
		initialTimeout:        initialTimeout,
		maxTimeout:            maxTimeout,
	}
}

func (r *RetryBackoffOption) GetType() RetryBackoffType {
	return r.backoffType
}

func (r *RetryBackoffOption) GetConstantConfig() (time.Duration, time.Duration) {
	return r.backoffInterval, r.maximumJitterInterval
}

func (r *RetryBackoffOption) GetExponentialConfig() (time.Duration, time.Duration, float64, time.Duration) {
	return r.initialTimeout, r.maxTimeout, r.exponentFactor, r.maximumJitterInterval
}
