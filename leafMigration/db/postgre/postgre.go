package postgre

import (
	leafGormPostgreSql "github.com/enricodg/leaf-utilities/database/sql/integrations/gorm/postgresql"
	leafSql "github.com/enricodg/leaf-utilities/database/sql/sql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/config"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"strings"
	"sync"
)

var (
	postgreConnection leafSql.ORM
	once              sync.Once
)

func GetPostgre() (leafSql.ORM, error) {
	var err error
	once.Do(func() {
		configuration := config.GetConfig()
		postgreConnection, err = leafGormPostgreSql.New(leafGormPostgreSql.DbConnection{
			Address:  strings.Split(configuration.PostgreSQLAddress, ","),
			Username: configuration.PostgreSQLUsername,
			Password: configuration.PostgreSQLPassword,
			DbName:   configuration.PostgreSQLDbName,
		},
			leafSql.WithMaxIdleConnection(configuration.MySQLMaxIdleConnection),
			leafSql.WithMaxOpenConnection(configuration.MySQLMaxOpenConnection),
			leafSql.WithConnMaxLifetime(configuration.MySQLMaxLifetimeConnection),
			leafSql.WithLogMode(configuration.MySQLLogMode),
			leafSql.WithLogger(logger.GetLogger()),
		)
	})
	if err != nil {
		return nil, err
	}
	return postgreConnection, nil
}
