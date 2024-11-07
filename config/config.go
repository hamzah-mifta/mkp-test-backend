package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	LogLevel     string
	JWTSecretKey string
}

func LoadConfig() (config *Config) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	logLevel := os.Getenv("LOG_LEVEL")
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	return &Config{
		DatabaseURL:  databaseURL,
		LogLevel:     logLevel,
		JWTSecretKey: jwtSecretKey,
	}
}
