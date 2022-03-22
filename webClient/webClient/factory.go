package leafWebClient

import "time"

type Factory interface {
	Create(timeout time.Duration) WebClient
	CreateWithRetry(timeout time.Duration, retryCount int) WebClient
}
