package pkg

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=public",
        GetEnv("DB_HOST", "localhost"),
        GetEnv("DB_USER", "postgres"),
        GetEnv("DB_PASSWORD", ""),
        GetEnv("DB_NAME", "sanbercode"),
        GetEnv("DB_PORT", "5432"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to DB:", err)
    }

    DB = db
    // DB.AutoMigrate(&models.Bioskop{})
    if err != nil {
    log.Fatalf("Gagal AutoMigrate: %v", err)
    } else {
        log.Println("AutoMigrate berhasil dijalankan.")
    }
}
