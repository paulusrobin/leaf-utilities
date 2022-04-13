module github.com/enricodg/leaf-utilities/database/sql/integrations/gorm/mysql

go 1.18

require (
	github.com/enricodg/leaf-utilities/database/sql/integrations/gorm v0.0.0-20220331094736-b3cb24b8d0aa
	github.com/enricodg/leaf-utilities/database/sql/sql v0.0.0-20220331094736-b3cb24b8d0aa
	github.com/newrelic/go-agent/v3 v3.15.2
	github.com/paulusrobin/leaf-utilities/logger/integrations/logrus v0.0.0-20220323084925-3ece86cd22d6
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.3
)

replace github.com/enricodg/leaf-utilities/database/sql/integrations/gorm => github.com/paulusrobin/leaf-utilities/database/sql/integrations/gorm v0.0.0-20220407094130-ee2721541c6d
