package migration

import (
	"github.com/paulusrobin/leaf-utilities/migrationCli/helper/version"
	"os"
	"path/filepath"
	"strings"
)

type Migration interface {
	Name() string
	Version() uint64
	Migrate() error
	Rollback() error
}
type Tool interface {
	Name() string
	Migrations() []Migration
	Check(verbose bool) error
	CheckVersion(version version.Version) error
	Versions() []version.DataVersion
	Migrate(version version.Version, specific bool) error
	Rollback(version version.Version, specific bool) error
}

func LoadMigrations(path string) []string {
	var files = make([]string, 0)
	prefix := "migrations/" + path

	err := filepath.Walk(prefix, func(path string, info os.FileInfo, err error) error {
		stripPrefix := strings.TrimSpace(strings.ReplaceAll(path, prefix, ""))
		if len(stripPrefix) == 0 {
			return nil
		}

		stripSlash := strings.ReplaceAll(stripPrefix, string(filepath.Separator), "")

		extension := ".go"
		if strings.Index(stripSlash, extension) != len(stripSlash)-len(extension) {
			return nil
		}

		if stripSlash == "di.go" || stripSlash == "initialize.go" {
			return nil
		}

		stripExtension := strings.ReplaceAll(stripSlash, ".go", "")
		files = append(files, strings.Split(stripExtension, "_")[0])
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
