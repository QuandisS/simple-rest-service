package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	URL  string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error when loading config: %w", err)
	}

	return &Config{
		Port: os.Getenv("PORT"),
		URL:  os.Getenv("URL"),
	}, nil
}
