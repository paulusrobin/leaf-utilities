module github.com/paulusrobin/leaf-utilities/database/sql/integrations/gorm

go 1.18

require (
	github.com/paulusrobin/leaf-utilities/database/sql/sql v0.0.0-20220331075104-d5bc5037862e
	github.com/newrelic/go-agent/v3 v3.15.2
	github.com/paulusrobin/leaf-utilities/common v0.0.0-20220413101500-cfb038e5b795
	github.com/paulusrobin/leaf-utilities/encoding/json v0.0.0-20220413055116-5143d33efda6
	github.com/paulusrobin/leaf-utilities/logger/logger v0.0.0-20220323084925-3ece86cd22d6
	github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic v0.0.0-20220413061907-30a9a704dbc3
	github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry v0.0.0-20220407093531-dd042bda6151
	github.com/paulusrobin/leaf-utilities/tracer/tracer v0.0.0-20220413055116-5143d33efda6
	github.com/thoas/go-funk v0.9.2
	gorm.io/gorm v1.23.3
)
