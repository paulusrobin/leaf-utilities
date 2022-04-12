package templates

const MigrationsTemplate = `package {{.MigrationType}}

import (
	"context"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	{{if .IsMongo}}nosqlConnection "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"{{else}}sqlConnection "github.com/enricodg/leaf-utilities/database/sql/sql"{{end}}
)

type migration_{{.Version}} struct {
	Log leafLogger.Logger
    {{if .IsMongo}}Conn nosqlConnection.Mongo{{else}}Conn sqlConnection.ORM{{end}}
}

// NOTE: DO NOT CHANGE MIGRATION Version
func (m *migration_{{.Version}}) Version() uint64 {
	return uint64({{.Version}})
}

// NOTE: DO NOT CHANGE MIGRATION Name
func (m *migration_{{.Version}}) Name() string {
	return "{{.Version}}_{{.MigrationName}}"
}

func (m *migration_{{.Version}}) Migrate() error {
	{{if .IsMongo}}
	panic("implement me"){{else}}
	script, err := file.ReadToString("./scripts/{{.MigrationType}}/{{.Version}}_{{.MigrationName}}_migrate.sql")
	if err != nil {
		return err
	}

	if err := m.Conn.Exec(context.Background(), script); err != nil {
		return err.Error()
	}

	return nil
	{{end}}
}

func (m *migration_{{.Version}}) Rollback() error {
	{{if .IsMongo}}panic("implement me"){{else}}script, err := file.ReadToString("./scripts/{{.MigrationType}}/{{.Version}}_{{.MigrationName}}_rollback.sql")
	if err != nil {
		return err
	}

	if err := m.Conn.Exec(context.Background(), script); err != nil {
		return err.Error()
	}

	return nil{{end}}
}
`
