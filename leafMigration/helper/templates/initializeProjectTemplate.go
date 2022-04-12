package templates

const MainTemplate = `package main 

import (
	migration "github.com/paulusrobin/leaf-utilities/leafMigration"
)

func main() {
	migration.New().
		// WithMySql(mysql.InitializeMigrations).
		// WithMySql(postgre.InitializeMigrations).
		// WithMySql(mongo.InitializeMigrations).
		Run()
}`

const GoModTemplate = `module {{.ProjectName}}

go 1.18

require github.com/paulusrobin/leaf-utilities/leafMigration v0.0.0-20220412071838-c8969b42efe1 // indirect
`
