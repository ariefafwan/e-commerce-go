package seeders

import (
	"e-commerce-go/internal/helpers"
	"e-commerce-go/internal/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	fmt.Println("Seeding User & Pelanggan...")
	user := models.User{
		ID: uuid.New(),
		Nama: "Pelanggan",
		Role: models.Pelanggan,
		Email: "Pelanggan@gmail.com",
		Password: helpers.Hash("123"),
	}

	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Error seeding User Pelanggan:", )
	}

	err := db.Create(&models.MasterPelanggan{
		NamaLengkap: "Pelanggan",
		IDUser: user.ID,
		NamaPanggilan: "Pelanggan",
		Phone: "08123456789",
	}).Error

	if err != nil {
		fmt.Println("Error seeding Master Pelanggan:", err)
	}
}