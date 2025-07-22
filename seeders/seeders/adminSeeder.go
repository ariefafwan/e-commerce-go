package seeders

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	fmt.Println("Seeding User Admin...")
	err := db.Create(&models.User{
		Nama: "admin",
		Role: "Admin",
		Email: "admin@admin.com",
		Password: helpers.Hash("123"),
	}).Error

	if err != nil {
		fmt.Println("Error seeding User:", err)
	}
}