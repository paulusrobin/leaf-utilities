package mongo

import (
	"context"
	"fmt"
	leafNoSql "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/db/mongo"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/connection"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/version"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"github.com/paulusrobin/leaf-utilities/leafMigration/migrator"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Mongo struct {
		log             leafLogger.Logger
		nosql           leafNoSql.Mongo
		migrations      []migration.Migration
		migrationFiles  map[uint64]migration.Migration
		versions        []version.DataVersion
		executedVersion map[uint64]version.DataVersion
	}
)

var log = logger.GetLogger()

func New(m migrator.Migrator) (*Mongo, error) {
	nosql, err := mongo.GetMongoConnection()
	if err != nil {
		return nil, errors.New("cannot established connection to mongo")
	}

	migrations := m.Mongo()(nosql, log)
	if len(migrations) < 1 {
		return nil, errors.New("no nosql migrations file")
	}

	var migrationFiles = make(map[uint64]migration.Migration)
	for _, m := range migrations {
		migrationFiles[m.Version()] = m
	}

	return &Mongo{
		nosql:          nosql,
		log:            log,
		migrations:     migrations,
		migrationFiles: migrationFiles,
	}, nil
}

func (mgo *Mongo) Name() string {
	return connection.Mongo
}

func (mgo *Mongo) Migrations() []migration.Migration {
	return mgo.migrations
}

func (mgo *Mongo) Check(verbose bool) error {
	ctx := context.Background()
	if !mgo.nosql.DB().HasCollection(ctx, version.MigrationTable) {
		if err := mgo.nosql.DB().CreateCollection(ctx, version.MigrationTable); err != nil {
			return err
		}
		mgo.versions = make([]version.DataVersion, 0)
		mgo.executedVersion = make(map[uint64]version.DataVersion)
		return nil
	}

	var filter = bson.M{}
	err := mgo.nosql.
		FindAll(ctx, version.MigrationTable, filter, &mgo.versions, &options.FindOptions{
			Sort: bson.M{"version": 1},
		})
	if err != nil {
		return err
	}

	mgo.executedVersion = make(map[uint64]version.DataVersion)
	for _, v := range mgo.versions {
		mgo.executedVersion[v.Version] = v
	}

	if verbose {
		for _, m := range mgo.migrations {
			if _, ok := mgo.executedVersion[m.Version()]; ok {
				log.StandardLogger().Infof("%d: UP", m.Version())
			} else {
				log.StandardLogger().Infof("%d: DOWN", m.Version())
			}
		}
	}
	return nil
}

func (mgo *Mongo) CheckVersion(version version.Version) error {
	if _, ok := mgo.executedVersion[uint64(version)]; ok {
		return nil
	}
	return errors.New("record not found")
}

func (mgo *Mongo) Versions() []version.DataVersion {
	return mgo.versions
}

func (mgo *Mongo) logMigrated() {
	for _, v := range mgo.Versions() {
		if _, ok := mgo.migrationFiles[v.Version]; !ok {
			log.StandardLogger().Warnf("version %d - %s, already migrated but not available in current version",
				v.Version, v.Name)
		}
	}
}

func (mgo *Mongo) Migrate(ver version.Version, specific bool) error {
	ctx := context.Background()
	for _, m := range mgo.Migrations() {
		if specific {
			if m.Version() == uint64(ver) {
				if err := mgo.migrate(&ctx, m); err != nil {
					return err
				}
				return nil
			}
			continue
		}

		if uint64(ver) < m.Version() {
			return nil
		}

		if _, ok := mgo.executedVersion[m.Version()]; ok {
			continue
		}

		if err := mgo.migrate(&ctx, m); err != nil {
			return err
		}
	}
	return nil
}

func (mgo *Mongo) migrate(ctx *context.Context, m migration.Migration) error {
	mgo.log.StandardLogger().Infof("[%s] execute migration version %d", mgo.Name(), m.Version())
	if err := m.Migrate(); err != nil {
		mgo.log.StandardLogger().Errorf("[%s] error execute migration version %d: %+v",
			mgo.Name(), m.Version(), err.Error())
		return err
	}

	newVersion := version.DataVersion{
		ID:          fmt.Sprintf("%d_%s", m.Version(), m.Name()),
		Version:     m.Version(),
		Name:        m.Name(),
		ExecuteTime: time.Now().Format(time.RFC3339),
	}

	_, err := mgo.nosql.DB().Collection(version.MigrationTable).InsertOne(*ctx, newVersion.ToBson())
	if err != nil {
		return err
	} else {
		mgo.versions = append(mgo.versions, newVersion)
		mgo.executedVersion[newVersion.Version] = newVersion
	}
	mgo.log.StandardLogger().Infof("[%s] finish execute migration version %d", mgo.Name(), m.Version())
	return nil
}

func (mgo *Mongo) Rollback(ver version.Version, specific bool) error {
	ctx := context.Background()
	latest := mgo.versions[len(mgo.versions)-1]
	if !specific && latest.Version < uint64(ver) {
		return nil
	}

	for i := len(mgo.migrations) - 1; i >= 0; i-- {
		if specific {
			if mgo.migrations[i].Version() == uint64(ver) {
				if err := mgo.rollback(&ctx, mgo.migrations[i]); err != nil {
					return err
				}
				return nil
			}
			continue
		}

		if latest.Version < mgo.migrations[i].Version() {
			continue
		}

		if mgo.migrations[i].Version() <= uint64(ver) {
			return nil
		}

		if _, ok := mgo.executedVersion[mgo.migrations[i].Version()]; ok {
			if err := mgo.rollback(&ctx, mgo.migrations[i]); err != nil {
				return err
			}
			delete(mgo.executedVersion, mgo.migrations[i].Version())
		}
	}
	return nil
}

func (mgo *Mongo) rollback(ctx *context.Context, m migration.Migration) error {
	mgo.log.StandardLogger().Infof("[%s] execute rollback version %d", mgo.Name(), m.Version())
	if err := m.Rollback(); err != nil {
		mgo.log.StandardLogger().Errorf("[%s] error execute rollback version %d: %+v", mgo.Name(), m.Version(),
			err.Error())
		return err
	}

	filter := bson.M{"id": fmt.Sprintf("%d_%s", m.Version(), m.Name())}
	_, err := mgo.nosql.DB().Collection(version.MigrationTable).DeleteOne(*ctx, filter)
	if err != nil {
		mgo.log.StandardLogger().Errorf("[%s] error execute rollback version %d: %+v", mgo.Name(), m.Version(),
			err.Error())
		return err
	}
	mgo.log.StandardLogger().Infof("[%s] finish execute rollback version %d", mgo.Name(), m.Version())
	return nil
}
