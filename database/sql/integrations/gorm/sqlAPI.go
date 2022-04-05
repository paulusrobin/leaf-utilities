package leafGorm

import (
	"context"
	"database/sql"
	leafSql "github.com/enricodg/leaf-utilities/database/sql/sql"
	"gorm.io/gorm"
)

// Session create new db session
func (i *Impl) Session(config *gorm.Session) leafSql.ORM {
	db := i.GormDB.Session(config)
	return i.newImpl(db)
}

// WithContext change current instance db's context to ctx
func (i *Impl) WithContext(ctx context.Context) leafSql.ORM {
	db := i.GormDB.WithContext(ctx)
	return i.newImpl(db)
}

// Debug start debug mode
func (i *Impl) Debug() leafSql.ORM {
	db := i.GormDB.Debug()
	return i.newImpl(db)
}

// Set store value with key into current db instance's context
func (i *Impl) Set(key string, value interface{}) leafSql.ORM {
	db := i.GormDB.Set(key, value)
	return i.newImpl(db)
}

// Get get value with key from current db instance's context
func (i *Impl) Get(key string) (interface{}, bool) {
	return i.GormDB.Get(key)
}

// InstanceSet store value with key into current db instance's context
func (i *Impl) InstanceSet(key string, value interface{}) leafSql.ORM {
	db := i.GormDB.InstanceSet(key, value)
	return i.newImpl(db)
}

// InstanceGet get value with key from current db instance's context
func (i *Impl) InstanceGet(key string) (interface{}, bool) {
	return i.GormDB.InstanceGet(key)
}

// AddError add error to db
func (i *Impl) AddError(err error) error {
	return i.GormDB.AddError(err)
}

// DB returns `*sql.DB`
func (i *Impl) DB() (*sql.DB, error) {
	db, err := i.GormDB.DB()
	return db, err
}

func (i *Impl) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return i.GormDB.SetupJoinTable(model, field, joinTable)
}

func (i *Impl) Use(plugin gorm.Plugin) error {
	return i.GormDB.Use(plugin)
}
