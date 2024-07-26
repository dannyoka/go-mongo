package config

import (
	"fmt"
	"os"
)

type Config struct{}

func New() *Config {
	return &Config{}
}

func (c *Config) Get(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("No value found for key %s", key)
	}
	return value, nil
}
