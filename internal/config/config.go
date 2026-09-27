package config

import "os"

type Config struct {
	Port string
	DSN  string
}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		DSN:  getEnv("DB_DSN", "postgres://postgres:postgres@localhost:5432/shop?sslmode=disable"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
