package leafGorm

import (
	"context"
	"github.com/newrelic/go-agent/v3/newrelic"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"gorm.io/gorm"
)

type (
	Impl struct {
		GormDB           *gorm.DB
		GormDBDryRun     *gorm.DB
		Log              leafLogger.Logger
		DatabaseName     string
		DataStoreProduct newrelic.DatastoreProduct
	}
)

func (i *Impl) Ping(ctx context.Context) error {
	db, err := i.GormDB.WithContext(ctx).DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func (i *Impl) Gorm() *gorm.DB {
	return i.GormDB
}
func (i *Impl) Dialector() gorm.Dialector {
	return i.GormDB.Dialector
}
func (i *Impl) Migrator() gorm.Migrator {
	return i.GormDB.Migrator()
}

func (i *Impl) Error() error {
	return i.GormDB.Error
}
func (i *Impl) RowsAffected() int64 {
	return i.GormDB.RowsAffected
}

func (i *Impl) AutoMigrate(dst ...interface{}) error {
	return i.GormDB.AutoMigrate(dst...)
}
func (i *Impl) Association(column string) *gorm.Association {
	return i.GormDB.Association(column)
}
func (i *Impl) Statement() *gorm.Statement {
	return i.GormDBDryRun.Statement
}
