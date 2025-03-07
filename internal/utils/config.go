package util

import (
	"os"
	"strconv"
)

type Config struct {
	Host string
	Port int
}

func LoadConfig() *Config {
	config := &Config{
		Host: "localhost",
		Port: 8080,
	}

	if host := os.Getenv("MCP_HOST"); host != "" {
		config.Host = host
	}

	if portStr := os.Getenv("MCP_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			config.Port = port
		}
	}

	return config
}
