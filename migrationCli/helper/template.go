package helper

import (
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper/templates"
	"os"
	"text/template"
)

type (
	MigrationRequestDTO struct {
		Version       uint64
		MigrationType string
		MigrationName string
		IsMongo       bool
	}
	InitializeRequestDTO struct {
		IsMongo       bool
		MigrationType string
		Versions      []string
	}
)

func CreateEmptyFile(outputPath string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	_ = f.Close()
	return nil
}

func CreateMigrations(outputPath string, data MigrationRequestDTO) error {
	tmpl, err := template.New(outputPath).Parse(templates.MigrationsTemplate)
	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}

func CreateInitialization(outputPath string, data InitializeRequestDTO) error {
	tmpl, err := template.New(outputPath).Parse(templates.InitializeTemplate)
	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}
