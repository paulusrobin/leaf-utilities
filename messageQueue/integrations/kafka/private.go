package leafKafka

import (
	"context"
	"fmt"
	leafNewRelicTracer "github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic"
	leafNewRelicDestination "github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/messageDestinationType"
	leafNewRelicSpanType "github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic/spanType"
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

	return tracer.StartSpan(fmt.Sprintf("[Kafka-Producer] %s", topic),
		tracer.ChildOf(span.Context()),
		leafNewRelicTracer.WithSpanType(leafNewRelicSpanType.MessageProducer),
		leafNewRelicTracer.WithMessageProducer(leafNewRelicTracer.MessageProducerOption{
			Library:              "Kafka",
			DestinationType:      leafNewRelicDestination.MessageTopic,
			DestinationName:      topic,
			DestinationTemporary: false,
		}),
		leafSentryTracer.WithSpanType(leafSentrySpanType.MessageProducer),
		leafSentryTracer.WithMessageProducer(leafSentryTracer.MessageProducerOption{
			Library:              "Kafka",
			DestinationType:      leafSentryDestination.MessageTopic,
			DestinationName:      topic,
			DestinationTemporary: false,
		}),
	)
}
