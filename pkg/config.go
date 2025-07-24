package pkg

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnv(key string, fallback string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config: %s",err )
	}

	value := viper.GetString(key)
    if value == "" {
        return fallback
    }
    return value
}