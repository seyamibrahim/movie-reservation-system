package configs

import (
	"log"
	"os"
)

type AppConfig struct {
	Port      string
	DBURL     string
	JWTSecret string
}

func LoadConfig() (*AppConfig, error) {

	return &AppConfig{
		Port:      getEnv("PORT"),
		DBURL:     getEnv("DB_URL"),
		JWTSecret: getEnv("JWT_SECRET"),
	}, nil
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("%s not found", key)
	}
	return value
}
