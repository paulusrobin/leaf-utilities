package mysql

import (
	leafGormMySql "github.com/paulusrobin/leaf-utilities/database/sql/integrations/gorm/mysql"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/config"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"strings"
	"sync"
)

var (
	mysqlConnection leafSql.ORM
	once            sync.Once
)

func GetMysql() (leafSql.ORM, error) {
	var err error
	once.Do(func() {
		configuration := config.GetConfig()
		mysqlConnection, err = leafGormMySql.New(leafGormMySql.DbConnection{
			Address:  strings.Split(configuration.MySQLAddress, ","),
			Username: configuration.MySQLUsername,
			Password: configuration.MySQLPassword,
			DbName:   configuration.MySQLDbName,
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
	return mysqlConnection, nil
}
