package handler

import (
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/version"
	"github.com/paulusrobin/leaf-utilities/leafMigration/migrator"
	"github.com/pkg/errors"
	"strings"
)

func (h handler) Check(m migrator.Migrator, version version.Version, migrationTypes ...string) error {
	if err := h.initializeConnection(m, migrationTypes); err != nil {
		return err
	}

	verbose := true
	if version != 0 {
		verbose = false
	}

	var errorMessage = make([]string, 0)
	for _, connection := range h.connections {
		if err := connection.Check(verbose); err != nil {
			log.StandardLogger().Errorf("[%s] error on connection check: %+v", connection.Name(), err.Error())
			errorMessage = append(errorMessage, err.Error())
			continue
		}

		if version != 0 {
			if err := connection.CheckVersion(version); err != nil {
				return err
			}
			log.StandardLogger().Infof("%d is UP", version)
		}
	}

	if len(errorMessage) == 0 {
		return nil
	}
	return errors.New(strings.Join(errorMessage, "\n"))
}
