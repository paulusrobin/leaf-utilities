package handler

import (
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/version"
	"github.com/paulusrobin/leaf-utilities/leafMigration/migrator"
	"github.com/pkg/errors"
	"strings"
)

func (h handler) Migrate(m migrator.Migrator, version version.Version, specific bool, verbose bool, migrationTypes ...string) error {
	if err := h.initializeConnection(m, migrationTypes); err != nil {
		return err
	}

	var errorMessage = make([]string, 0)
	for _, connection := range h.connections {
		if err := connection.Check(verbose); err != nil {
			log.StandardLogger().Errorf("[%s] error on connection check: %+v", connection.Name(), err.Error())
			errorMessage = append(errorMessage, err.Error())
			continue
		}

		if err := connection.Migrate(version, specific); err != nil {
			log.StandardLogger().Errorf("[%s] error on process migration: %+v", connection.Name(), err.Error())
			errorMessage = append(errorMessage, err.Error())
		}
	}

	if len(errorMessage) == 0 {
		return nil
	}
	return errors.New(strings.Join(errorMessage, "\n"))
}
