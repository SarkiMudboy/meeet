package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load("")
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err.Error())
	}
}

func GetEnv(variable string, fallback string) string {
	value, ok := os.LookupEnv(variable)
	if !ok {
		return fallback
	}

	return value
}
