package leafMigration

import (
	"context"
	leafSql "github.com/enricodg/leaf-utilities/database/sql/sql"
	leafNoSql "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/command"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration"
	"github.com/paulusrobin/leaf-utilities/leafMigration/logger"
	"github.com/paulusrobin/leaf-utilities/leafMigration/migrator"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/urfave/cli/v2"
	"os"
)

type (
	Cli struct {
		migrator *migrator.Migrator
	}
)

var log = logger.GetLogger()

func New() *Cli {
	return &Cli{migrator: migrator.New()}
}

func (c *Cli) WithMySql(f func(conn leafSql.ORM, log leafLogger.Logger) []migration.Migration) *Cli {
	c.migrator.WithMySql(f)
	return c
}

func (c *Cli) WithPostgre(f func(conn leafSql.ORM, log leafLogger.Logger) []migration.Migration) *Cli {
	c.migrator.WithPostgre(f)
	return c
}

func (c *Cli) WithMongo(f func(conn leafNoSql.Mongo, log leafLogger.Logger) []migration.Migration) *Cli {
	c.migrator.WithMongo(f)
	return c
}

func (c Cli) Run() *cli.App {
	app := cli.NewApp()
	app.Name = "Leaf Migration"
	app.Usage = "Command Line Tools for Databases Migration"
	app.UsageText = "command [command options] [arguments...]"
	app.HelpName = "leafMigration"
	app.Description = "CLI migration tools"
	app.Version = "v1.0.0"
	app.Commands = []*cli.Command{
		command.Init(),
		command.New(),
		command.Migrate(*c.migrator),
		command.Rollback(*c.migrator),
		command.Check(*c.migrator),
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(leafLogger.BuildMessage(context.Background(), "Run Error: %+v",
			leafLogger.WithAttr("error", err)))
	}

	return app
}
