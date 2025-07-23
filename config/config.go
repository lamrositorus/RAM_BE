package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string // Full connection string
	DBHost      string
	DBPort      string
	DBUser       string
	DBPassword  string
	DBName      string
	UseURL      bool
}

func LoadConfig() Config {
	// Memuat variabel dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	useURL := false
	if dbURL != "" {
		useURL = true
	}

	cfg := Config{
		DatabaseURL: dbURL,
		DBHost:      getEnv("DB_HOST", "db.filgqowjdfvbnlqbcymu.supabase.co"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "admin"),
		DBName:      getEnv("DB_NAME", "postgres"),
		UseURL:      useURL,
	}

	if useURL {
		log.Println("Using DATABASE_URL for database connection")
	} else {
		log.Println("Using individual DB connection parameters")
	}

	return cfg
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
