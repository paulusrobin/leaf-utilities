package postgre

import (
	"context"
	"fmt"
	leafSql "github.com/enricodg/leaf-utilities/database/sql/sql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/db/postgre"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/connection"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/version"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"github.com/paulusrobin/leaf-utilities/leafMigration/migrator"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type (
	PostgreSQL struct {
		log             leafLogger.Logger
		migrations      []migration.Migration
		migrationFiles  map[uint64]migration.Migration
		sql             leafSql.ORM
		versions        []version.DataVersion
		executedVersion map[uint64]version.DataVersion
	}
)

var log = logger.GetLogger()

func New(m migrator.Migrator) (*PostgreSQL, error) {
	orm, err := postgre.GetPostgre()
	if err != nil {
		return nil, errors.New("cannot established connection to postgres")
	}

	migrations := m.Postgre()(orm, log)
	if len(migrations) < 1 {
		return nil, errors.New("no sql migrations file")
	}

	var migrationFiles = make(map[uint64]migration.Migration)
	for _, m := range migrations {
		migrationFiles[m.Version()] = m
	}

	return &PostgreSQL{
		log:            log,
		sql:            orm,
		migrations:     migrations,
		migrationFiles: migrationFiles,
	}, nil
}

func (p *PostgreSQL) Name() string {
	return connection.Postgre
}

func (p *PostgreSQL) Migrations() []migration.Migration {
	return p.migrations
}

func (p *PostgreSQL) Check(verbose bool) error {
	ctx := context.Background()
	if !p.sql.Migrator().HasTable(version.MigrationTable) {
		if err := p.sql.Migrator().CreateTable(&version.DataVersion{}); err != nil {
			return err
		}
		p.versions = make([]version.DataVersion, 0)
		p.executedVersion = make(map[uint64]version.DataVersion)
		return nil
	}

	if err := p.sql.Model(&version.DataVersion{}).
		Order("version asc").
		Find(ctx, &p.versions).Error(); err != nil {
		return err
	}

	p.executedVersion = make(map[uint64]version.DataVersion)
	for _, v := range p.versions {
		p.executedVersion[v.Version] = v
	}

	if verbose {
		for _, m := range p.migrations {
			if _, ok := p.executedVersion[m.Version()]; ok {
				log.Info(leafLogger.BuildMessage(ctx, "%d: UP", leafLogger.WithAttr("version", m.Version())))
			} else {
				log.Info(leafLogger.BuildMessage(ctx, "%d: DOWN", leafLogger.WithAttr("version", m.Version())))
			}
		}
	}
	return nil
}

func (p *PostgreSQL) CheckVersion(version version.Version) error {
	if _, ok := p.executedVersion[uint64(version)]; ok {
		return nil
	}
	return errors.New("record not found")
}

func (p *PostgreSQL) Versions() []version.DataVersion {
	return p.versions
}

func (p *PostgreSQL) logMigrated() {
	for _, v := range p.Versions() {
		if _, ok := p.migrationFiles[v.Version]; !ok {
			log.Warn(leafLogger.BuildMessage(context.Background(), "version %d - %s, already migrated but not available in current version",
				leafLogger.WithAttr("version", v.Version),
				leafLogger.WithAttr("name", v.Name)))
		}
	}
}

func (p *PostgreSQL) Migrate(ver version.Version, specific bool) error {
	ctx := context.Background()
	p.logMigrated()

	for _, m := range p.Migrations() {
		if specific {
			if m.Version() == uint64(ver) {
				if err := p.migrate(&ctx, m); err != nil {
					return err
				}
				return nil
			}
			continue
		}

		if uint64(ver) < m.Version() {
			return nil
		}

		if _, ok := p.executedVersion[m.Version()]; ok {
			continue
		}

		if err := p.migrate(&ctx, m); err != nil {
			return err
		}

		p.executedVersion[m.Version()] = version.DataVersion{Version: m.Version()}
	}
	return nil
}

func (p *PostgreSQL) migrate(ctx *context.Context, m migration.Migration) error {
	p.log.Info(leafLogger.BuildMessage(*ctx, "[%s] execute rollback version %d",
		leafLogger.WithAttr("name", p.Name()),
		leafLogger.WithAttr("version", m.Version())))
	if err := m.Migrate(); err != nil {
		p.log.Error(leafLogger.BuildMessage(*ctx, "[%s] error execute migration version %d: %+v",
			leafLogger.WithAttr("name", p.Name()),
			leafLogger.WithAttr("version", m.Version()),
			leafLogger.WithAttr("error", err)))
		return err
	}

	newVersion := version.DataVersion{
		ID:          fmt.Sprintf("%d_%s", m.Version(), strings.ReplaceAll(m.Name(), " ", "_")),
		Version:     m.Version(),
		Name:        m.Name(),
		ExecuteTime: time.Now().Format(time.RFC3339),
	}
	if err := p.sql.Create(*ctx, &newVersion).Error(); err != nil {
		p.log.Error(leafLogger.BuildMessage(*ctx, "[%s] error execute migration version %d: %+v",
			leafLogger.WithAttr("name", p.Name()),
			leafLogger.WithAttr("version", m.Version()),
			leafLogger.WithAttr("error", err)))
		return err
	} else {
		p.versions = append(p.versions, newVersion)
		p.executedVersion[newVersion.Version] = newVersion
	}
	p.log.Info(leafLogger.BuildMessage(*ctx, "[%s] finish execute migration version %d",
		leafLogger.WithAttr("name", p.Name()),
		leafLogger.WithAttr("version", m.Version())))
	return nil
}

func (p *PostgreSQL) Rollback(ver version.Version, specific bool) error {
	ctx := context.Background()
	p.logMigrated()

	latest := p.versions[len(p.versions)-1]
	if !specific && latest.Version < uint64(ver) {
		return nil
	}

	for i := len(p.migrations) - 1; i >= 0; i-- {
		if specific {
			if p.migrations[i].Version() == uint64(ver) {
				if err := p.rollback(&ctx, p.migrations[i]); err != nil {
					return err
				}
				return nil
			}
			continue
		}

		if latest.Version < p.migrations[i].Version() {
			continue
		}

		if p.migrations[i].Version() <= uint64(ver) {
			return nil
		}

		if _, ok := p.executedVersion[p.migrations[i].Version()]; ok {
			if err := p.rollback(&ctx, p.migrations[i]); err != nil {
				return err
			}
			delete(p.executedVersion, p.migrations[i].Version())
		}
	}
	return nil
}

func (p *PostgreSQL) rollback(ctx *context.Context, m migration.Migration) error {
	p.log.Info(leafLogger.BuildMessage(*ctx, "[%s] execute rollback version %d",
		leafLogger.WithAttr("name", p.Name()),
		leafLogger.WithAttr("version", m.Version())))
	if err := m.Rollback(); err != nil {
		p.log.Error(leafLogger.BuildMessage(*ctx, "[%s] error execute rollback version %d: %+v",
			leafLogger.WithAttr("name", p.Name()),
			leafLogger.WithAttr("version", m.Version()),
			leafLogger.WithAttr("error", err)))
		return err
	}

	newVersion := version.DataVersion{}
	if err := p.sql.Table(version.MigrationTable).
		Where("version = ?", m.Version()).
		First(*ctx, &newVersion).Error(); err != nil {
		return err
	}

	if err := p.sql.Delete(*ctx, &newVersion).Error(); err != nil {
		p.log.Error(leafLogger.BuildMessage(*ctx, "[%s] error execute rollback version %d: %+v",
			leafLogger.WithAttr("name", p.Name()),
			leafLogger.WithAttr("version", m.Version()),
			leafLogger.WithAttr("error", err)))
		return err
	}
	p.log.Info(leafLogger.BuildMessage(*ctx, "[%s] finish execute rollback version %d",
		leafLogger.WithAttr("name", p.Name()),
		leafLogger.WithAttr("version", m.Version())))
	return nil
}
