package leafHystrix

import (
	"context"
	"fmt"
	"github.com/gojek/heimdall/v7/hystrix"
	leafFunctions "github.com/paulusrobin/leaf-utilities/common/functions"
	"io"
	"net/http"
	netUrl "net/url"
)

type WebClient struct {
	Doer hystrix.Client
}

func (c *WebClient) Get(ctx context.Context, url string, headers http.Header, queryString map[string]string) (*http.Response, error) {
	if len(queryString) > 0 {
		var count = 0
		for key, value := range queryString {
			format := "?%s=%s"
			if count > 0 {
				format = "&%s=%s"
			}
			url += fmt.Sprintf(format, key, netUrl.QueryEscape(value))
			count++
		}
	}

	headers = leafFunctions.AppendMandatoryHeader(ctx, headers)
	return c.Doer.Get(url, headers)
}

func (c *WebClient) Post(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	headers = leafFunctions.AppendMandatoryHeader(ctx, headers)
	return c.Doer.Post(url, body, headers)
}

func (c *WebClient) Put(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	headers = leafFunctions.AppendMandatoryHeader(ctx, headers)
	return c.Doer.Put(url, body, headers)
}

func (c *WebClient) Patch(ctx context.Context, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	headers = leafFunctions.AppendMandatoryHeader(ctx, headers)
	return c.Doer.Patch(url, body, headers)
}

func (c *WebClient) Delete(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	headers = leafFunctions.AppendMandatoryHeader(ctx, headers)
	return c.Doer.Delete(url, headers)
}

func (c *WebClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header = leafFunctions.AppendMandatoryHeader(ctx, req.Header)
	return c.Doer.Do(req)
}
