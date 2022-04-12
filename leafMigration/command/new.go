package command

import (
	"context"
	"fmt"
	"github.com/paulusrobin/leaf-utilities/leafMigration/handler"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/connection"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/version"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
	"time"
)

func New() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "new --types <type> --name <name>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "types",
				Aliases:  []string{"t"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			log := logger.GetLogger()
			migrationType := strings.TrimSpace(strings.ToLower(c.String("types")))
			migrationName := strings.TrimSpace(strings.ToLower(c.String("name")))

			if len(migrationName) < 1 {
				return errors.New("migration name is required")
			}

			if strings.Index(migrationName, "test") == len(migrationName)-4 {
				return errors.New("migration name must not ended with \"test\"")
			}

			if !connection.IsValid(migrationType) {
				return errors.New("invalid migration type [mysql | mongo | postgre]")
			}

			log.Info(leafLogger.BuildMessage(context.Background(), "start creating new %s migrations file",
				leafLogger.WithAttr("migrationType", migrationType)))

			now := time.Now()
			year, month, date := now.Date()
			hour := now.Hour()
			minute := now.Minute()
			second := now.Second()

			var stringVersion = fmt.Sprintf("%04d%02d%02d%02d%02d%02d",
				year, month, date, hour, minute, second)
			ver, err := strconv.ParseUint(stringVersion, 10, 64)
			if err != nil {
				return err
			}

			return handler.GetHandler().New(version.Version(ver), migrationType, migrationName)
		},
	}
}
