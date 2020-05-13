package config

import (
	"os"
)

type Config struct {
	Mysql_url string
	Mysql_db  string
	Http_port string
}

func Load() *Config {
	return &Config{
		Mysql_url: getEnv("MYSQL_URL", "127.0.0.1:3306"),
		Mysql_db:  getEnv("MYSQL_DB", "qok_url_shortener"),
		Http_port: getEnv("HTTP_PORT", "8787"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
