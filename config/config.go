package config

import (
	"fireflyiiiapi/config/models"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	UBS models.UBS `yaml:"ubs,omitempty"`
}

func ReadConfig(file string) (*Config, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
