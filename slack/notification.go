package leafSlack

import (
	"context"
	leafWebClient "github.com/enricodg/leaf-utilities/webClient/webClient"
	leafHeimdall "github.com/paulusrobin/leaf-utilities/webClient/integrations/heimdall"
	"net/http"
	"strings"
	"time"
)

type (
	Integration interface {
		Push(ctx context.Context, message Message) error
	}
	integration struct {
		option    option
		webClient leafWebClient.WebClient
	}
)

func (i *integration) Push(ctx context.Context, message Message) error {
	header := http.Header{
		"Content-type": []string{"application/json"},
	}
	response, err := i.webClient.Post(ctx, i.option.hook, strings.NewReader(message.Json()), header)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func Notification(options ...Option) (Integration, error) {
	o := option{
		timeout: 5 * time.Second,
	}

	for _, opt := range options {
		opt.Apply(&o)
	}

	if err := o.validate(); err != nil {
		return nil, err
	}

	return &integration{
		option:    o,
		webClient: leafHeimdall.NewClientFactory().Create(leafWebClient.NewDefaultWebClientOption(o.timeout)),
	}, nil
}
