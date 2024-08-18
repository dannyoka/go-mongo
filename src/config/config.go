package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

func InitConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	return New()
}
