package config

import "os"

type Config struct {
	APIKey string
}

func LoadConfig() *Config {
	return &Config{
		APIKey: loadEnv("API_KEY", ""),
	}
}

func loadEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}
