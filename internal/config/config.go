package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
		Name    string `yaml:"name"`
		User    string `yaml:"user"`
		PassEnv string `yaml:"passenv"`
	} `yaml:"database"`
}

func Load() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	cfgPath := filepath.Join("config", fmt.Sprintf("%s.yaml", env))
	raw, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", cfgPath, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML data in %s: %w", cfgPath, err)
	}

	return &cfg, err
}
