package pkg

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"e-commerce-go/internal/models"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=public",
        GetEnv("DB_HOST", "localhost"),
        GetEnv("DB_USER", "postgres"),
        GetEnv("DB_PASSWORD", "postgres"),
        GetEnv("DB_NAME", "e-commerce"),
        GetEnv("DB_PORT", "5432"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    DB = db
    err = DB.AutoMigrate(
            &models.MasterToko{}, 
            &models.User{}, 
            &models.PersonalAccessToken{},
            &models.MasterPelanggan{},
            &models.MasterAlamatPelanggan{},
            &models.MasterKategoriProduk{},
        )
    
    if err != nil {
        panic(err)
    }

    var dbName string
    DB.Raw("SELECT current_database()").Scan(&dbName)
    
    log.Printf("Connected to DB: %s", dbName)
}
