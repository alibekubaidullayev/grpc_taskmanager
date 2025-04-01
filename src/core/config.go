package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbName string
	DbUser string
	DbPasw string
	DbHost string
	DbPort string
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env. Using default values instead")
	}

	dbName := getEnv("DB_NAME", "task_manager")
	dbUser := getEnv("DB_USER", "postgres")
	dbPasw := getEnv("DB_PASW", "123456")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")

	return Config{
		DbName: dbName,
		DbUser: dbUser,
		DbPasw: dbPasw,
		DbHost: dbHost,
		DbPort: dbPort,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
