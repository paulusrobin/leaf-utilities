package leafKafka

//import (
//	"context"
//	"fmt"
//	taniTracer "gitlab.tanihub.com/rnd/utilities/tracer/tracer"
//	taniNewRelicTracer "gitlab.tanihub.com/rnd/utilities/tracer/tracer-new-relic"
//	taniNewRelicDestination "gitlab.tanihub.com/rnd/utilities/tracer/tracer-new-relic/messageDestinationType"
//	taniNewRelicSpanType "gitlab.tanihub.com/rnd/utilities/tracer/tracer-new-relic/spanType"
//	taniSentryTracer "gitlab.tanihub.com/rnd/utilities/tracer/tracer-sentrygo"
//	taniSentryDestination "gitlab.tanihub.com/rnd/utilities/tracer/tracer-sentrygo/messageDestinationType"
//	taniSentrySpanType "gitlab.tanihub.com/rnd/utilities/tracer/tracer-sentrygo/spanType"
//	"gitlab.tanihub.com/rnd/utilities/tracer/tracer/tracer"
//)
//
//func startMessagingProducerSpan(ctx context.Context, topic string) taniTracer.Span {
//	span, found := tracer.SpanFromContext(ctx)
//	if !found {
//		return tracer.NoopSpan()
//	}
//
//	return tracer.StartSpan(fmt.Sprintf("[Kafka-Producer] %s", topic),
//		tracer.ChildOf(span.Context()),
//		taniNewRelicTracer.WithSpanType(taniNewRelicSpanType.MessageProducer),
//		taniNewRelicTracer.WithMessageProducer(taniNewRelicTracer.MessageProducerOption{
//			Library:              "Kafka",
//			DestinationType:      taniNewRelicDestination.MessageTopic,
//			DestinationName:      topic,
//			DestinationTemporary: false,
//		}),
//		taniSentryTracer.WithSpanType(taniSentrySpanType.MessageProducer),
//		taniSentryTracer.WithMessageProducer(taniSentryTracer.MessageProducerOption{
//			Library:              "Kafka",
//			DestinationType:      taniSentryDestination.MessageTopic,
//			DestinationName:      topic,
//			DestinationTemporary: false,
//		}),
//	)
//}
