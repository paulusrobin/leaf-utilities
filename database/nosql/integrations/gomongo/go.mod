module github.com/paulusrobin/leaf-utilities/database/nosql/integrations/gomongo

go 1.18

require (
	github.com/newrelic/go-agent/v3 v3.15.2
	github.com/paulusrobin/leaf-utilities/common v0.0.0-20220405011639-4753f34ca4a8
	github.com/paulusrobin/leaf-utilities/database/nosql/nosql v0.0.0-20220406024637-86c20de7d2ab
	github.com/paulusrobin/leaf-utilities/logger/integrations/logrus v0.0.0-20220406024637-86c20de7d2ab
	github.com/paulusrobin/leaf-utilities/logger/logger v0.0.0-20220406024637-86c20de7d2ab
	github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic v0.0.0-20220413061907-30a9a704dbc3
	github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry v0.0.0-20220407094130-ee2721541c6d
	github.com/paulusrobin/leaf-utilities/tracer/tracer v0.0.0-20220413055116-5143d33efda6
	go.mongodb.org/mongo-driver v1.9.0
)
