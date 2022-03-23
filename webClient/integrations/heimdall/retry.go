package leafHeimdall

import (
	"github.com/gojek/heimdall/v7"
	leafWebClient "github.com/paulusrobin/leaf-utilities/webClient/webClient/circuitBreaker"
)

func convertToHeimdallRetryBackoff(retryBackoff leafWebClient.RetryBackoffOption) heimdall.Retriable {
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
