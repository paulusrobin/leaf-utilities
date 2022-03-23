# Web Client

## Factory Interface
```go
type Factory interface {
	Create(timeout time.Duration) WebClient
	CreateWithRetry(timeout time.Duration, retryCount int) WebClient
}
```

## Factory with Circuit Breaker Interface
```go
type CircuitBreakerFactory interface {
	Create(options ...Option) WebClient
}
```

## Circuit Breaker Options
```go
type ClientOption struct {
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
```

## Circuit Breaker Retry Backoff Options
```go
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
```

## Web Client Interface
```go
type WebClient interface {
	Get(ctx context.Context, url string, headers http.Header, queryString map[string]string) (*http.Response, error)
	Post(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error)
	Put(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error)
	Patch(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error)
	Delete(ctx context.Context, url string, headers http.Header) (*http.Response, error)
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}
```

## Usage
Import the package
```go
import (
    ...
    "github.com/paulusrobin/leaf-utilities/webClient"
    ...
)
```