package typpostgres

import (
	"github.com/golang-migrate/migrate"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

func cmdMigrateDB(c *typbuildtool.Context) *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "Migrate Database",
		Action: func(cliCtx *cli.Context) (err error) {
			return migrateDB(c.BuildContext(cliCtx))
		},
	}
}

func migrateDB(c *typbuildtool.BuildContext) (err error) {
	var (
		migration *migrate.Migrate
		sourceURL = "file://" + DefaultMigrationSource
		cfg       *Config
	)

	if cfg, err = retrieveConfig(c); err != nil {
		return
	}

	c.Infof("Migrate database from source '%s'", sourceURL)
	if migration, err = migrate.New(sourceURL, dataSource(cfg)); err != nil {
		return err
	}
	defer migration.Close()
	return migration.Up()
}