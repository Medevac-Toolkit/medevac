package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Strategy struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Threshold int    `yaml:"threshold,omitempty"`
	Path      string `yaml:"path,omitempty"`
}

type Config struct {
	Strategies []Strategy `yaml:"strategies"`
}

func LoadConfig(configPath string) (Config, error) {
	var config Config

	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
