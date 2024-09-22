package web

import (
	"log/slog"
	"net/http"

	"go_sampler/providers/config"
	"go_sampler/providers/db"
	"go_sampler/providers/logging"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

var WebCommand = &cli.Command{
	Name:  "webserver",
	Usage: "start the webserver",
	Action: func(c *cli.Context) error {
		listen := c.String("listen")
		debug := c.Bool("debug")
		configFilename := c.String("config")

		app := fx.New(
			fx.NopLogger,
			fx.Provide(func() (*config.Config, error) {
				if cfg, err := config.NewConfig(configFilename); err != nil {
					return nil, err
				} else {
					if listen != "" {
						cfg.WebAddr = listen
					}
					if debug {
						cfg.Debug = true
					}
					return cfg, nil
				}
			}),
			fx.Provide(logging.New, db.NewPostgres),
			fx.Provide(NewAPI, NewHandler, NewWebServer),
			fx.Invoke(func(*slog.Logger, *http.Server) {}),
		)
		if err := app.Err(); err != nil {
			return cli.Exit(err, 1)
		}

		app.Run()
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "listen",
			Aliases: []string{"l"},
			Usage:   "listen address",
			Value:   ":8080",
		},
	},
}
