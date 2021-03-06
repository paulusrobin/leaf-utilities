package migrator

import (
	leafNoSql "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

type (
	Migrator struct {
		mongo   func(conn leafNoSql.Mongo, log leafLogger.Logger) []migration.Migration
		mySql   func(conn leafSql.ORM, log leafLogger.Logger) []migration.Migration
		postgre func(conn leafSql.ORM, log leafLogger.Logger) []migration.Migration
	}
)

func New() *Migrator {
	return &Migrator{}
}

func (m Migrator) MySql() func(conn leafSql.ORM, log leafLogger.Logger) []migration.Migration {
	return m.mySql
}

func (m *Migrator) WithMySql(f func(conn leafSql.ORM, log leafLogger.Logger) []migration.Migration) *Migrator {
	m.mySql = f
	return m
}

func (m Migrator) Postgre() func(conn leafSql.ORM, log leafLogger.Logger) []migration.Migration {
	return m.postgre
}

func (m *Migrator) WithPostgre(f func(conn leafSql.ORM, log leafLogger.Logger) []migration.Migration) *Migrator {
	m.postgre = f
	return m
}

func (m Migrator) Mongo() func(conn leafNoSql.Mongo, log leafLogger.Logger) []migration.Migration {
	return m.mongo
}

func (m *Migrator) WithMongo(f func(conn leafNoSql.Mongo, log leafLogger.Logger) []migration.Migration) *Migrator {
	m.mongo = f
	return m
}
