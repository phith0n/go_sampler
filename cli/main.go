package main

import (
	"log"
	"os"

	"go_sampler/config"
	"go_sampler/db"
	"go_sampler/logging"
	"go_sampler/web"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var logger = logging.GetSugar()

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
		Before: func(c *cli.Context) error {
			configFile := c.String("config")
			err := config.InitConfig(configFile)
			if err != nil {
				return cli.Exit("failed to load config", 1)
			}

			debug := c.Bool("debug")
			if debug {
				config.GlobalConfig.Debug = true
			}

			err = logging.InitLogger(config.GlobalConfig.Debug)
			if err != nil {
				return err
			}
			logger.Infof("debug mode = %v", config.GlobalConfig.Debug)

			if config.GlobalConfig.Debug {
				gin.SetMode(gin.DebugMode)
			} else {
				gin.SetMode(gin.ReleaseMode)
			}

			err = db.InitMysql(config.GlobalConfig.DatabaseURL, debug)
			if err != nil {
				return cli.Exit("failed to initial MySQL database", 1)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
