package leafWebClient

type Factory interface {
	Create(options ...Option) WebClient
}
