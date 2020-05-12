package config

import (
	"os"
)

type Config struct {
	Mysql_url string
	Mysql_db  string
}

func Load() *Config {
	return &Config{
		Mysql_url: getEnv("MYSQL_URL", "127.0.0.1:3306"),
		Mysql_db:  getEnv("MYSQL_DB", "qok_url_shortener"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
