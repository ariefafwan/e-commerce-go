package pkg

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    viper.AutomaticEnv() // prioritize env from OS
}

func GetEnv(key string, fallback string) string {
    if value := viper.GetString(key); value != "" {
        return value
    }
    return fallback
}