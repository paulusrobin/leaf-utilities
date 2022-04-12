package handler

import (
	"context"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper/version"
	"github.com/paulusrobin/leaf-utilities/migrationCli/migrator"
	"github.com/pkg/errors"
	"strings"
)

func (h handler) Rollback(m migrator.Migrator, version version.Version, specific bool, verbose bool, migrationTypes ...string) error {
	if err := h.initializeConnection(m, migrationTypes); err != nil {
		return err
	}

	var errorMessage = make([]string, 0)
	for _, connection := range h.connections {
		if err := connection.Check(verbose); err != nil {
			log.Error(leafLogger.BuildMessage(context.Background(), "[%s] error on connection check: %+v",
				leafLogger.WithAttr("name", connection.Name()),
				leafLogger.WithAttr("error", err)))
			errorMessage = append(errorMessage, err.Error())
			continue
		}

		if err := connection.Rollback(version, specific); err != nil {
			log.Error(leafLogger.BuildMessage(context.Background(), "[%s] error on process rollback: %+v",
				leafLogger.WithAttr("name", connection.Name()),
				leafLogger.WithAttr("error", err)))
			errorMessage = append(errorMessage, err.Error())
		}
	}

	if len(errorMessage) == 0 {
		return nil
	}
	return errors.New(strings.Join(errorMessage, "\n"))
}
