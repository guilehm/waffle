package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port       string
	APIKey     string
	APITimeout int
}

func LoadConfig() *Config {

	return &Config{
		Port:       loadEnv("PORT", "8080"),
		APIKey:     loadEnv("API_KEY", ""),
		APITimeout: loadIntEnv("API_TIMEOUT", 5),
	}
}

func loadEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}

func loadIntEnv(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if exists {
		i, _ := strconv.Atoi(value)
		return i
	}
	return fallback
}
