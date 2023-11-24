package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var GlobalConfig Config

type Config struct {
	Debug       bool   `yaml:"debug"`
	WebAddr     string `yaml:"web_addr"`
	DatabaseURL string `yaml:"database_url"`
}

func InitConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %v", filename, err)
	}

	err = yaml.Unmarshal(data, &GlobalConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file %s: %v", filename, err)
	}

	return nil
}
