package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug       bool   `yaml:"debug"`
	WebAddr     string `yaml:"web_addr"`
	DatabaseURL string `yaml:"database_url"`
}

func NewConfig(filename string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %v", filename, err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file %s: %v", filename, err)
	}

	return &config, nil
}
