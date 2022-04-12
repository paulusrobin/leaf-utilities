package config

import (
	leafConfig "github.com/paulusrobin/leaf-utilities/config"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	configuration EnvConfig
	once          sync.Once
)

type (
	EnvConfig struct {
		// - Logger config
		LogFilePath  string `envconfig:"LOG_FILE_NAME"`
		LogFormatter string `envconfig:"LOG_FORMATTER" default:"TEXT"`

		// - MySQL connection
		MySQLAddress               string        `envconfig:"MY_SQL_ADDRESS"`
		MySQLUsername              string        `envconfig:"MY_SQL_USERNAME"`
		MySQLPassword              string        `envconfig:"MY_SQL_PASSWORD"`
		MySQLDbName                string        `envconfig:"MY_SQL_DB_NAME"`
		MySQLMaxIdleConnection     int           `envconfig:"MY_SQL_MAX_IDLE_CONNECTION"`
		MySQLMaxOpenConnection     int           `envconfig:"MY_SQL_MAX_OPEN_CONNECTION"`
		MySQLMaxLifetimeConnection time.Duration `envconfig:"MY_SQL_MAX_LIFETIME_CONNECTION"`
		MySQLLogMode               bool          `envconfig:"MY_SQL_LOG_MODE" default:"false"`

		// - Postgres connection
		PostgreSQLAddress               string        `envconfig:"POSTGRE_SQL_ADDRESS"`
		PostgreSQLUsername              string        `envconfig:"POSTGRE_SQL_USERNAME"`
		PostgreSQLPassword              string        `envconfig:"POSTGRE_SQL_PASSWORD"`
		PostgreSQLDbName                string        `envconfig:"POSTGRE_SQL_DB_NAME"`
		PostgreSQLMaxIdleConnection     int           `envconfig:"POSTGRE_SQL_MAX_IDLE_CONNECTION"`
		PostgreSQLMaxOpenConnection     int           `envconfig:"POSTGRE_SQL_MAX_OPEN_CONNECTION"`
		PostgreSQLMaxLifetimeConnection time.Duration `envconfig:"POSTGRE_SQL_MAX_LIFETIME_CONNECTION"`
		PostgreSQLLogMode               bool          `envconfig:"POSTGRE_SQL_LOG_MODE" default:"false"`

		// - Mongo connection
		MongoUri      string `envconfig:"MONGO_DB_URI"`
		MongoDatabase string `envconfig:"MONGO_DB_DATABASE"`
	}
)

func (e EnvConfig) validate() error {
	if e.MySQLAddress != "" {
		if e.MySQLUsername == "" || e.MySQLDbName == "" {
			return errors.New("MY_SQL_USERNAME, MY_SQL_PASSWORD, MY_SQL_DB_NAME is required")
		}
		if e.MySQLMaxOpenConnection == 0 || e.MySQLMaxIdleConnection == 0 || e.MySQLMaxLifetimeConnection == 0 {
			return errors.New("MY_SQL_MAX_OPEN_CONNECTION, MY_SQL_MAX_IDLE_CONNECTION, MY_SQL_MAX_LIFETIME_CONNECTION is required")
		}
	}

	if e.PostgreSQLAddress != "" {
		if e.PostgreSQLUsername == "" || e.PostgreSQLDbName == "" {
			return errors.New("POSTGRE_SQL_USERNAME, POSTGRE_SQL_PASSWORD, POSTGRE_SQL_DB_NAME is required")
		}
		if e.PostgreSQLMaxOpenConnection == 0 || e.PostgreSQLMaxIdleConnection == 0 || e.PostgreSQLMaxLifetimeConnection == 0 {
			return errors.New("POSTGRE_SQL_MAX_OPEN_CONNECTION, POSTGRE_SQL_MAX_IDLE_CONNECTION, POSTGRE_SQL_MAX_LIFETIME_CONNECTION is required")
		}
	}

	if e.MongoUri != "" {
		if e.MongoDatabase == "" {
			return errors.New("MongoDatabase is required")
		}
	}

	return nil
}

func GetConfig() EnvConfig {
	once.Do(func() {
		if err := leafConfig.NewFromEnv(&configuration); err != nil {
			log.Fatal(err)
		}

		if err := configuration.validate(); err != nil {
			log.Fatal(err)
		}
	})
	return configuration
}
