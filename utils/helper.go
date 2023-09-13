package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(env string) {
	err := godotenv.Load(env)
	if err != nil {
		log.Fatalf("Error loading %s file", env)
	}
}

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
