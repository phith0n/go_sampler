package web

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"go_sampler/providers/config"
	"go_sampler/providers/mysql"
	"net/http"
)

var WebCommand = &cli.Command{
	Name:  "webserver",
	Usage: "start the webserver",
	Action: func(c *cli.Context) error {
		listen := c.String("listen")
		debug := c.Bool("debug")
		configFilename := c.String("config")

		fx.New(
			fx.Provide(fx.Annotated{
				Name: "listen",
				Target: func() string {
					return listen
				},
			}),
			fx.Provide(fx.Annotated{
				Name: "debug",
				Target: func() bool {
					return debug
				},
			}),
			fx.Provide(fx.Annotated{
				Name: "config",
				Target: func() string {
					return configFilename
				},
			}),
			fx.Provide(config.NewConfig, mysql.NewMysql, NewWebServer, NewHandler),
			fx.Invoke(func(*http.Server) {

			})).Run()
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
