package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type FireStoreConfig struct {
	ProjectID string `yaml:"project_id"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Server   ServerConfig    `yaml:"server"`
	Database FireStoreConfig `yaml:"database"`
}

func LoadConfigFromFile(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("no se pudo leer el archivo de configuración %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("no se pudo parsear YAML de configuración: %w", err)
	}

	return &cfg, nil
}

func LoadDefaultConfig() (*Config, error) {
	return LoadConfigFromFile("configs/config.yaml")
}
