package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        string
	JWTSecret   string
	JWTTTLHours int
	DBDriver    string
	SQLitePath  string
	MySQLDSN    string
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}

func Load() *Config {
	ttlStr := getenv("JWT_TTL_HOURS", "24")
	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		ttl = 24
	}

	return &Config{
		Port:        getenv("PORT", "8080"),
		JWTSecret:   getenv("JWT_SECRET", "dev_secret"),
		JWTTTLHours: ttl,
		DBDriver:    getenv("DB_DRIVER", "sqlite"),
		SQLitePath:  getenv("SQLITE_PATH", "blog.db"),
		MySQLDSN:    getenv("MYSQL_DSN", ""),
	}
}
