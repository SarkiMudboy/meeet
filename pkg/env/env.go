package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err.Error())
	}
}

func GetString(variable string, fallback string) string {
	value, ok := os.LookupEnv(variable)
	if !ok {
		return fallback
	}

	return value
}

func GetInt(variable string, fallback int) int {
	value, ok := os.LookupEnv(variable)
	if !ok {
		return fallback
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intVal
}
