package templates

const MainTemplate = `package main 

import (
	"github.com/paulusrobin/leaf-utilities/leafMigration"
)

func main() {
	migrationsCli.New().
		Run()
}`

const GoModTemplate = `module {{.ProjectName}}

go 1.18

require github.com/paulusrobin/leaf-utilities/leafMigration v0.0.0-20220412053718-f62a405d11f1 // indirect
`
