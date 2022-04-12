package mysql

import (
	"context"
	"fmt"
	leafSql "github.com/enricodg/leaf-utilities/database/sql/sql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/migrationCli/db/mysql"
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper/connection"
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper/migration"
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper/version"
	"github.com/paulusrobin/leaf-utilities/migrationCli/logger"
	"github.com/paulusrobin/leaf-utilities/migrationCli/migrator"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type (
	MySQL struct {
		log             leafLogger.Logger
		migrations      []migration.Migration
		migrationFiles  map[uint64]migration.Migration
		sql             leafSql.ORM
		versions        []version.DataVersion
		executedVersion map[uint64]version.DataVersion
	}
)

var log = logger.GetLogger()

func New(m migrator.Migrator) (*MySQL, error) {
	orm, err := mysql.GetMysql()
	if err != nil {
		return nil, errors.New("cannot established connection to mysql")
	}

	migrations := m.MySql()(orm, log)
	if len(migrations) < 1 {
		return nil, errors.New("no sql migrations file")
	}

	var migrationFiles = make(map[uint64]migration.Migration)
	for _, m := range migrations {
		migrationFiles[m.Version()] = m
	}

	return &MySQL{
		log:            log,
		sql:            orm,
		migrations:     migrations,
		migrationFiles: migrationFiles,
	}, nil
}

func (s *MySQL) Name() string {
	return connection.MySQL
}

func (s *MySQL) Migrations() []migration.Migration {
	return s.migrations
}

func (s *MySQL) Check(verbose bool) error {
	ctx := context.Background()
	if !s.sql.Migrator().HasTable(version.MigrationTable) {
		if err := s.sql.Migrator().CreateTable(&version.DataVersion{}); err != nil {
			return err
		}
		s.versions = make([]version.DataVersion, 0)
		s.executedVersion = make(map[uint64]version.DataVersion)
		return nil
	}

	if err := s.sql.Model(&version.DataVersion{}).
		Order("version asc").
		Find(ctx, &s.versions); err != nil {
		return err.Error()
	}

	s.executedVersion = make(map[uint64]version.DataVersion)
	for _, v := range s.versions {
		s.executedVersion[v.Version] = v
	}

	if verbose {
		for _, m := range s.migrations {
			if _, ok := s.executedVersion[m.Version()]; ok {
				log.Info(leafLogger.BuildMessage(ctx, "%d: UP", leafLogger.WithAttr("version", m.Version())))
			} else {
				log.Info(leafLogger.BuildMessage(ctx, "%d: DOWN", leafLogger.WithAttr("version", m.Version())))
			}
		}
	}
	return nil
}

func (s *MySQL) CheckVersion(version version.Version) error {
	if _, ok := s.executedVersion[uint64(version)]; ok {
		return nil
	}
	return errors.New("record not found")
}

func (s *MySQL) Versions() []version.DataVersion {
	return s.versions
}

func (s *MySQL) logMigrated() {
	for _, v := range s.Versions() {
		if _, ok := s.migrationFiles[v.Version]; !ok {
			log.Warn(leafLogger.BuildMessage(nil, "version %d - %s, already migrated but not available in current version",
				leafLogger.WithAttr("version", v.Version),
				leafLogger.WithAttr("name", v.Name)))
		}
	}
}

func (s *MySQL) Migrate(ver version.Version, specific bool) error {
	ctx := context.Background()
	s.logMigrated()

	for _, m := range s.Migrations() {
		if specific {
			if m.Version() == uint64(ver) {
				if err := s.migrate(&ctx, m); err != nil {
					return err
				}
				return nil
			}
			continue
		}

		if uint64(ver) < m.Version() {
			return nil
		}

		if _, ok := s.executedVersion[m.Version()]; ok {
			continue
		}

		if err := s.migrate(&ctx, m); err != nil {
			return err
		}
	}
	return nil
}

func (s *MySQL) migrate(ctx *context.Context, m migration.Migration) error {
	s.log.Info(leafLogger.BuildMessage(*ctx, "[%s] execute migration version %d: %+v",
		leafLogger.WithAttr("name", s.Name()),
		leafLogger.WithAttr("version", m.Version())))
	if err := m.Migrate(); err != nil {
		s.log.Error(leafLogger.BuildMessage(*ctx, "[%s] error execute migration version %d: %+v",
			leafLogger.WithAttr("name", s.Name()),
			leafLogger.WithAttr("version", m.Version()),
			leafLogger.WithAttr("error", err.Error())))
		return err
	}

	newVersion := version.DataVersion{
		ID:          fmt.Sprintf("%d_%s", m.Version(), strings.ReplaceAll(m.Name(), " ", "_")),
		Version:     m.Version(),
		Name:        m.Name(),
		ExecuteTime: time.Now().Format(time.RFC3339),
	}
	if err := s.sql.Create(*ctx, newVersion); err != nil {
		s.log.Error(leafLogger.BuildMessage(*ctx, "[%s] error execute migration version %d: %+v",
			leafLogger.WithAttr("name", s.Name()),
			leafLogger.WithAttr("version", m.Version()),
			leafLogger.WithAttr("error", err.Error())))
		return err.Error()
	} else {
		s.versions = append(s.versions, newVersion)
		s.executedVersion[newVersion.Version] = newVersion
	}
	s.log.Info(leafLogger.BuildMessage(*ctx, "[%s] finish execute migration version %d",
		leafLogger.WithAttr("name", s.Name()),
		leafLogger.WithAttr("version", m.Version())))
	return nil
}

func (s *MySQL) Rollback(ver version.Version, specific bool) error {
	ctx := context.Background()
	s.logMigrated()

	latest := s.versions[len(s.versions)-1]
	if !specific && latest.Version < uint64(ver) {
		return nil
	}

	for i := len(s.migrations) - 1; i >= 0; i-- {
		if specific {
			if s.migrations[i].Version() == uint64(ver) {
				if err := s.rollback(&ctx, s.migrations[i]); err != nil {
					return err
				}
				return nil
			}
			continue
		}

		if latest.Version < s.migrations[i].Version() {
			continue
		}

		if s.migrations[i].Version() <= uint64(ver) {
			return nil
		}

		if _, ok := s.executedVersion[s.migrations[i].Version()]; ok {
			if err := s.rollback(&ctx, s.migrations[i]); err != nil {
				return err
			}
			delete(s.executedVersion, s.migrations[i].Version())
		}
	}
	return nil
}

func (s *MySQL) rollback(ctx *context.Context, m migration.Migration) error {
	s.log.Info(leafLogger.BuildMessage(*ctx, "[%s] execute rollback version %d",
		leafLogger.WithAttr("name", s.Name()),
		leafLogger.WithAttr("version", m.Version())))
	if err := m.Rollback(); err != nil {
		s.log.Error(leafLogger.BuildMessage(*ctx, "[%s] error execute rollback version %d: %+v",
			leafLogger.WithAttr("name", s.Name()),
			leafLogger.WithAttr("version", m.Version()),
			leafLogger.WithAttr("error", err.Error())))
		return err
	}

	newVersion := version.DataVersion{}
	if err := s.sql.Table(version.MigrationTable).
		Where("version = ?", m.Version()).
		First(*ctx, &newVersion); err != nil {
		return err.Error()
	}

	if err := s.sql.Delete(*ctx, &newVersion); err != nil {
		s.log.Error(leafLogger.BuildMessage(*ctx, "[%s] error execute rollback version %d: %+v",
			leafLogger.WithAttr("name", s.Name()),
			leafLogger.WithAttr("version", m.Version()),
			leafLogger.WithAttr("error", err.Error())))
		return err.Error()
	}
	s.log.Info(leafLogger.BuildMessage(*ctx, "[%s] finish execute rollback version %d",
		leafLogger.WithAttr("name", s.Name()),
		leafLogger.WithAttr("version", m.Version())))
	return nil
}
