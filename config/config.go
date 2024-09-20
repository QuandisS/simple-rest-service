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

	c := &Config{}
	port, found := os.LookupEnv("PORT")
	if !found {
		c.Port = "8080"
	} else {
		c.Port = port
	}

	c.URL = os.Getenv("URL")

	return c, nil
}
