package command

import (
	"github.com/paulusrobin/leaf-utilities/leafMigration/handler"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/connection"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/version"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"github.com/paulusrobin/leaf-utilities/leafMigration/migrator"
	"github.com/urfave/cli/v2"
	"strings"
)

func Rollback(m migrator.Migrator) *cli.Command {
	return &cli.Command{
		Name:  "rollback",
		Usage: "rollback --types <types> --version <version> [--verbose] [--specific]",
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:     "version",
				Aliases:  []string{"v"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "types",
				Aliases:  []string{"t"},
				Value:    "mysql,mongo,postgre",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "verbose",
				Required: false,
				Value:    false,
			},
			&cli.BoolFlag{
				Name:     "specific",
				Required: false,
				Value:    false,
			},
		},
		Action: func(c *cli.Context) error {
			log := logger.GetLogger()
			log.StandardLogger().Info("starting rollback migrations")
			ver := c.Uint64("version")
			verbose := c.Bool("verbose")
			specific := c.Bool("specific")
			types := strings.ToLower(c.String("types"))
			migrationsTypes := strings.Split(types, ",")
			if err := connection.CheckConnection(migrationsTypes); err != nil {
				return err
			}
			return handler.GetHandler().Rollback(m, version.Version(ver), specific, verbose, migrationsTypes...)
		},
	}
}
