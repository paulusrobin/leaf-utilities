module github.com/paulusrobin/leaf-utilities/messageQueue/integrations/googlePubsub

go 1.18

require (
	cloud.google.com/go/pubsub v1.19.0
	github.com/paulusrobin/leaf-utilities/common v0.0.0-20220407084001-602cbec02989
	github.com/paulusrobin/leaf-utilities/logger/integrations/zap v0.0.0-20220407084001-602cbec02989
	github.com/paulusrobin/leaf-utilities/logger/logger v0.0.0-20220407084001-602cbec02989
	github.com/paulusrobin/leaf-utilities/mandatory v0.0.0-20220322085140-66e0cded624f
	github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue v0.0.0-20220407084001-602cbec02989
	github.com/paulusrobin/leaf-utilities/slack v0.0.0-20220407084001-602cbec02989
	github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry v0.0.0-20220407094130-ee2721541c6d
	github.com/paulusrobin/leaf-utilities/tracer/tracer v0.0.0-20220407094130-ee2721541c6d
	google.golang.org/api v0.70.0
	google.golang.org/grpc v1.44.0
)
