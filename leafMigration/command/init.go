package command

import (
	"github.com/paulusrobin/leaf-utilities/leafMigration/handler"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
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
			log.StandardLogger().Infof("[%s] initializing project...", project)
			return handler.GetHandler().Init(project)
		},
	}
}
