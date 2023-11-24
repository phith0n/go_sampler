package web

import (
	"github.com/urfave/cli/v2"
	"go_sampler/config"
)

var WebCommand = &cli.Command{
	Name:  "webserver",
	Usage: "start the webserver",
	Action: func(c *cli.Context) error {
		listen := c.String("listen")
		if listen != "" {
			config.GlobalConfig.WebAddr = listen
		}

		return StartGin(config.GlobalConfig.WebAddr)
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
