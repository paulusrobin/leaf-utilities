# Web Client

## Factory Interface
```go
type Factory interface {
	Create(timeout time.Duration) WebClient
	CreateWithRetry(timeout time.Duration, retryCount int) WebClient
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