package leafWebClient

type CircuitBreakerFactory interface {
	Create(options ...Option) WebClient
}
