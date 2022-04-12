package helper

import (
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/templates"
	"os"
	"text/template"
)

type (
	InitializeProjectRequestDTO struct {
		ProjectName string
	}
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

func Initialize(data InitializeProjectRequestDTO) error {
	main, err := template.New("main.go").Parse(templates.MainTemplate)
	if err != nil {
		return err
	}

	goMod, err := template.New("go.mod").Parse(templates.GoModTemplate)
	if err != nil {
		return err
	}

	mainFile, err := os.Create("main.go")
	if err != nil {
		return err
	}
	defer mainFile.Close()

	err = main.Execute(mainFile, data)
	if err != nil {
		return err
	}

	goModFile, err := os.Create("go.mod")
	if err != nil {
		return err
	}
	defer goModFile.Close()

	err = goMod.Execute(goModFile, nil)
	if err != nil {
		return err
	}
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
