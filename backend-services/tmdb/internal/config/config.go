package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppName                  string
	Port                     string
	APIKey                   string
	APITimeout               int
	KafkaBrokers             string
	KafkaProducerMaxMessages int
}

func LoadConfig() *Config {
	return &Config{
		AppName:                  "tmdb",
		Port:                     loadEnv("PORT", "8080"),
		APIKey:                   loadEnv("API_KEY", ""),
		APITimeout:               loadIntEnv("API_TIMEOUT", 5),
		KafkaBrokers:             loadEnv("KAFKA_BROKERS", "localhost:9092"),
		KafkaProducerMaxMessages: loadIntEnv("KAFKA_PRODUCER_MAX_MESSAGES", 100_000),
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
