package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    ServerPort string
    DBDriver   string
    DBDSN      string
}

var Cfg Config

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using default environment variables")
    }

    Cfg = Config{
        ServerPort: getEnv("SERVER_PORT", "8080"),
        DBDriver:   getEnv("DB_DRIVER", "sqlite3"),
        DBDSN:      getEnv("DB_DSN", "schedulogo.db"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
