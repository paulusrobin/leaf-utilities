package leafGooglePubsub

import (
	"context"
	"fmt"
	leafSentryTracer "github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry"
	leafSentryDestination "github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/messageDestinationType"
	leafSentrySpanType "github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry/spanType"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func startMessagingProducerSpan(ctx context.Context, topic string) leafTracer.Span {
	span, found := tracer.SpanFromContext(ctx)
	if !found {
		return tracer.NoopSpan()
	}

	tracerSpan, err := tracer.StartSpanFromContext(ctx, fmt.Sprintf("[GooglePubSub-Producer] %s", topic),
		tracer.ChildOf(span.Context()),
		//taniNewRelicTracer.WithSpanType(taniNewRelicSpanType.MessageProducer),
		//taniNewRelicTracer.WithMessageProducer(taniNewRelicTracer.MessageProducerOption{
		//	Library:              "GooglePubSub",
		//	DestinationType:      taniNewRelicDestination.MessageTopic,
		//	DestinationName:      topic,
		//	DestinationTemporary: false,
		//}),
		leafSentryTracer.WithSpanType(leafSentrySpanType.MessageProducer),
		leafSentryTracer.WithMessageProducer(leafSentryTracer.MessageProducerOption{
			Library:              "GooglePubSub",
			DestinationType:      leafSentryDestination.MessageTopic,
			DestinationName:      topic,
			DestinationTemporary: false,
		}),
	)

	if err != nil {
		return tracer.NoopSpan()
	}

	return tracerSpan
}
