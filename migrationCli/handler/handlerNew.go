package handler

import (
	"context"
	"fmt"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper"
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper/connection"
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper/migration"
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper/version"
	"github.com/pkg/errors"
	"path/filepath"
	"strings"
)

func (h handler) New(version version.Version, migrationType string, migrationName string) error {
	if !connection.IsValid(migrationType) {
		return errors.New(fmt.Sprintf("migration type %s is not supported", migrationType))
	}

	neutralizedName := strings.ReplaceAll(migrationName, " ", "_")
	migrationsPath := fmt.Sprintf("migrations/%s", migrationType)

	fileName := fmt.Sprintf("%d_%s.go", version, neutralizedName)

	// NOTE: create migrations file
	h.log.Info(leafLogger.BuildMessage(context.Background(), "[%s] creating migration file version: %d_%s.go",
		leafLogger.WithAttr("migrationType", strings.ToUpper(migrationType)),
		leafLogger.WithAttr("version", version),
		leafLogger.WithAttr("neutralizedName", neutralizedName)))
	if err := helper.CreateMigrations(filepath.Join(migrationsPath, fileName),
		helper.MigrationRequestDTO{
			Version:       uint64(version),
			MigrationType: migrationType,
			MigrationName: migrationName,
			IsMongo:       connection.IsMongo(migrationType),
		}); err != nil {
		h.log.Error(leafLogger.BuildMessage(context.Background(), "[%s] error creating migration file version: %d_%s.go: %s",
			leafLogger.WithAttr("migrationType", strings.ToUpper(migrationType)),
			leafLogger.WithAttr("version", version),
			leafLogger.WithAttr("neutralizedName", neutralizedName),
			leafLogger.WithAttr("error", err.Error())))
		return err
	}

	// NOTE: create migrations initialization
	h.log.Info(leafLogger.BuildMessage(context.Background(), "[%s] initialize migration file version: %d_%s.go",
		leafLogger.WithAttr("migrationType", strings.ToUpper(migrationType)),
		leafLogger.WithAttr("version", version),
		leafLogger.WithAttr("neutralizedName", neutralizedName)))
	files := migration.LoadMigrations(migrationType)
	if err := helper.CreateInitialization(filepath.Join(migrationsPath, "initialize.go"),
		helper.InitializeRequestDTO{
			MigrationType: migrationType,
			IsMongo:       connection.IsMongo(migrationType),
			Versions:      files,
		}); err != nil {
		h.log.Error(leafLogger.BuildMessage(context.Background(), "[%s] error initialize migration file version: %d_%s.go: %s",
			leafLogger.WithAttr("migrationType", strings.ToUpper(migrationType)),
			leafLogger.WithAttr("version", version),
			leafLogger.WithAttr("neutralizedName", neutralizedName),
			leafLogger.WithAttr("error", err.Error())))
		return err
	}

	h.log.Info(leafLogger.BuildMessage(context.Background(), "[%s] finish migration file version: %d_%s.go",
		leafLogger.WithAttr("migrationType", strings.ToUpper(migrationType)),
		leafLogger.WithAttr("version", version),
		leafLogger.WithAttr("neutralizedName", neutralizedName)))
	return nil
}
