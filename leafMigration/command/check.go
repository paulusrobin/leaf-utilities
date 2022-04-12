package command

import (
	"context"
	"github.com/paulusrobin/leaf-utilities/leafMigration/handler"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/version"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"github.com/paulusrobin/leaf-utilities/leafMigration/migrator"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/urfave/cli/v2"
	"strings"
)

func Check(m migrator.Migrator) *cli.Command {
	return &cli.Command{
		Name:  "check",
		Usage: "check [--types <types>] [--version <version>]",
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:     "version",
				Aliases:  []string{"v"},
				Required: false,
			},
			&cli.StringFlag{
				Name:     "types",
				Aliases:  []string{"t"},
				Value:    "mysql,mongo,postgre",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			log := logger.GetLogger()
			log.Info(leafLogger.BuildMessage(context.Background(), "checking migrations..."))
			ver := c.Uint64("version")
			types := strings.ToLower(c.String("types"))
			migrationTypes := strings.Split(types, ",")
			return handler.GetHandler().Check(m, version.Version(ver), migrationTypes...)
		},
	}
}
