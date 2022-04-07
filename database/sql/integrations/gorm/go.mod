module github.com/enricodg/leaf-utilities/database/sql/integrations/gorm

go 1.16

require (
	github.com/enricodg/leaf-utilities/database/sql/sql v0.0.0-20220331075104-d5bc5037862e
	github.com/newrelic/go-agent/v3 v3.15.2
	github.com/paulusrobin/leaf-utilities/common v0.0.0-20220323084925-3ece86cd22d6
	github.com/paulusrobin/leaf-utilities/encoding/json v0.0.0-20220407091457-d5651ac33646
	github.com/paulusrobin/leaf-utilities/logger/logger v0.0.0-20220323084925-3ece86cd22d6
	github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry v0.0.0-20220407093531-dd042bda6151
	github.com/paulusrobin/leaf-utilities/tracer/tracer v0.0.0-20220407093244-f51f6d55a7b5
	github.com/thoas/go-funk v0.9.2
	gorm.io/gorm v1.23.3
)
