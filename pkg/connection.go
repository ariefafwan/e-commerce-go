package pkg

import (
	"e-commerce-go/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
    DB_HOST     = "localhost"
    DB_USER     = "postgres"
    DB_PASSWORD = ""
    DB_NAME     = "sanbercode"
    DB_PORT     = "5432"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=public",
        DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    DB = db
    err = DB.AutoMigrate(&models.MasterToko{})
    if err != nil {
        panic(err)
    }

    var dbName string
    DB.Raw("SELECT current_database()").Scan(&dbName)
    
    log.Printf("Connected to DB: %s", dbName)
}
