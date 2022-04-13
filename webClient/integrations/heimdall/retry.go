package leafHeimdall

import (
	leafCircuitBreaker "github.com/enricodg/leaf-utilities/webClient/webClient/circuitBreaker"
	"github.com/gojek/heimdall/v7"
)

func convertToHeimdallRetryBackoff(retryBackoff leafCircuitBreaker.RetryBackoffOption) heimdall.Retriable {
	var backoff heimdall.Backoff
	switch retryBackoff.GetType() {
	case 1:
		backoff = heimdall.NewConstantBackoff(retryBackoff.GetConstantConfig())
		break
	case 2:
		backoff = heimdall.NewExponentialBackoff(retryBackoff.GetExponentialConfig())
		break
	default:
		return heimdall.NewNoRetrier()
	}
	return heimdall.NewRetrier(backoff)
}
