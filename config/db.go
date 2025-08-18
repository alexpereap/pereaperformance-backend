// config/db.go
package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

func getenv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func LoadDB() DBConfig {
	port, _ := strconv.Atoi(getenv("DB_PORT", "5432"))

	return DBConfig{
		Host:     getenv("DB_HOST", "localhost"),
		Port:     port,
		User:     getenv("DB_USER", "alex"),
		Password: getenv("DB_PASSWORD", "alex"),
		Name:     getenv("DB_NAME", "pereaperformance"),
		SSLMode:  getenv("DB_SSLMODE", "disable"),
		TimeZone: getenv("DB_TIMEZONE", "America/Bogota"),
	}
}
