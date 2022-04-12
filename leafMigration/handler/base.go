package handler

import (
	"github.com/paulusrobin/leaf-utilities/leafMigration/config"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/connection"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration/mongo"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration/mysql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration/postgre"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/version"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"github.com/paulusrobin/leaf-utilities/leafMigration/migrator"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"sync"

	"github.com/pkg/errors"
)

var (
	instance handler
	once     sync.Once
	log      = logger.GetLogger()
)

type (
	NewRequestDTO struct {
		Version       version.Version
		MigrationType string
		MigrationName string
	}
	handler struct {
		connections []migration.Tool
		config      config.EnvConfig
		log         leafLogger.Logger
	}
)

func GetHandler() handler {
	once.Do(func() {
		instance = handler{
			config: config.GetConfig(),
			log:    logger.GetLogger(),
		}
	})
	return instance
}

var mappingConnection = map[string]func(m migrator.Migrator) (migration.Tool, error){
	connection.MySQL: func(m migrator.Migrator) (migration.Tool, error) {
		return mysql.New(m)
	},
	connection.Postgre: func(m migrator.Migrator) (migration.Tool, error) {
		return postgre.New(m)
	},
	connection.Mongo: func(m migrator.Migrator) (migration.Tool, error) {
		return mongo.New(m)
	},
}

func (h *handler) initializeConnection(m migrator.Migrator, types []string) error {
	for _, t := range types {
		f, ok := mappingConnection[t]
		if !ok {
			continue
		}
		h.addConnections(m, f)
	}

	if len(h.connections) < 1 {
		return errors.New("no connection created")
	}
	return nil
}

func (h *handler) addConnections(m migrator.Migrator, f func(m migrator.Migrator) (migration.Tool, error)) {
	conn, err := f(m)
	if err != nil {
		log.StandardLogger().Warnf("failed add connection: %s", err.Error())
	} else {
		h.connections = append(h.connections, conn)
	}
}
