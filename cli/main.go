package main

import (
	"log"
	"os"

	"go_sampler/web"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:  "go_sampler",
		Usage: "",
		Commands: []*cli.Command{
			web.WebCommand,
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "enable debug mode",
				Value:   false,
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "config key or config file name",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
