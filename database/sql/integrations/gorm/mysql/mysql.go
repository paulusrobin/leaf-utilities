package leafGormMySql

import (
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	leafGorm "github.com/paulusrobin/leaf-utilities/database/sql/integrations/gorm"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	leafLogrus "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func New(dbConnection DbConnection, options ...leafSql.ConnectionOption) (leafSql.ORM, error) {
	option := leafSql.ConnectionOptions{
		MaxIdleConnection: 10,
		MaxOpenConnection: 200,
		ConnMaxLifetime:   time.Hour,
		LogMode:           false,
	}
	leafSql.WithLogger(leafLogrus.DefaultLog()).Apply(&option)
	for _, opt := range options {
		opt.Apply(&option)
	}

	logLevel := []logger.LogLevel{
		logger.Silent,
		logger.Info,
		logger.Info,
		logger.Warn,
		logger.Error,
		logger.Silent,
		logger.Silent,
		logger.Silent,
	}[option.Logger().StandardLogger().Level()]

	if !option.LogMode {
		logLevel = logger.Silent
	}

	db, err := gorm.Open(mysql.Open(dbConnection.URI()), &gorm.Config{
		Logger: logger.New(option.Logger().StandardLogger(), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logLevel,
			Colorful:      true,
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql connection: %+v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql connection: %+v", err)
	}

	sqlDb.SetMaxIdleConns(option.MaxIdleConnection)
	sqlDb.SetMaxOpenConns(option.MaxOpenConnection)
	sqlDb.SetConnMaxLifetime(option.ConnMaxLifetime)
	return &leafGorm.Impl{
		GormDB:           db,
		Log:              option.Logger(),
		DatabaseName:     dbConnection.DbName,
		DataStoreProduct: newrelic.DatastoreMySQL,
	}, nil
}
