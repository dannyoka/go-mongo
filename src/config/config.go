package config

import (
	"fmt"
	"os"
)

type Config struct{}

func New() *Config {
	return &Config{}
}

func (c *Config) Get(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Error getting %s: %v", key, "value is empty"))
	}
	return value
}
