package command

import (
	"context"
	"github.com/paulusrobin/leaf-utilities/leafMigration/handler"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/urfave/cli/v2"
	"strings"
)

func Init() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "init --project <project URL>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "project",
				Aliases:  []string{"p"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			log := logger.GetLogger()
			project := strings.ToLower(c.String("project"))
			log.Info(leafLogger.BuildMessage(context.Background(), "initializing project..."))
			return handler.GetHandler().Init(project)
		},
	}
}
