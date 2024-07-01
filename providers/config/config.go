package config

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug       bool   `yaml:"debug"`
	WebAddr     string `yaml:"web_addr"`
	DatabaseURL string `yaml:"database_url"`
}

func NewConfig(lc fx.Lifecycle, filename string) *Config {
	var config Config
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			data, err := os.ReadFile(filename)
			if err != nil {
				return fmt.Errorf("failed to read config file %s: %v", filename, err)
			}

			err = yaml.Unmarshal(data, &config)
			if err != nil {
				return fmt.Errorf("failed to unmarshal config file %s: %v", filename, err)
			}

			return nil
		},
	})

	return &config
}
