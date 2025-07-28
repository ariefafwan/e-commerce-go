package pkg

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Peringatan: file .env tidak ditemukan.")
		} else {
			log.Fatalf("Error saat membaca file konfigurasi: %s", err)
		}
	}
}

func GetEnv(key string, fallback string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}
	return fallback
}