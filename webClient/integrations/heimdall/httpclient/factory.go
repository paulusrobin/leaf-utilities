package leafHeimdall

import (
	"github.com/gojek/heimdall/v7/httpclient"
	leafWebClient "github.com/paulusrobin/leaf-utilities/webClient/webClient"
	"time"
)

type webClientFactory struct{}

func (cf *webClientFactory) Create(timeout time.Duration) leafWebClient.WebClient {
	return &WebClient{
		Doer: *httpclient.NewClient(httpclient.WithHTTPTimeout(timeout)),
	}
}

func (cf *webClientFactory) CreateWithRetry(timeout time.Duration, retryCount int) leafWebClient.WebClient {
	return &WebClient{
		Doer: *httpclient.NewClient(
			httpclient.WithHTTPTimeout(timeout),
			httpclient.WithRetryCount(retryCount),
		),
	}
}

func NewClientFactory() leafWebClient.Factory {
	return &webClientFactory{}
}
