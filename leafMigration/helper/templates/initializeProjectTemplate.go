package templates

const MainTemplate = `package main 

import (
	migration "github.com/paulusrobin/leaf-utilities/leafMigration"
)

func main() {
	migration.New().
		// WithMySql(mysql.InitializeMigrations).
		// WithPostgre(postgre.InitializeMigrations).
		// WithMongo(mongo.InitializeMigrations).
		Run()
}`

const GoModTemplate = `module {{.ProjectName}}

go 1.18

require github.com/paulusrobin/leaf-utilities/leafMigration v0.0.0-20220412083712-e942d511cfe8 // indirect
`
