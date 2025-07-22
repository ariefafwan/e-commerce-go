package seeders

import (
	"e-commerce-go/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func SeedMasterToko(db *gorm.DB) {
	fmt.Println("Seeding Master Toko...")
	err := db.Create(&models.MasterToko{
		Nama:       "Toko Sanbercode",
		Alamat:     "Jl. Sanbercode No. 1",
		Gambar:     "https://source.unsplash.com/random/400x400/?toko",
		NomorToko:  "123456789",
		AturanPajak: 10.0,
	}).Error

	if err != nil {
		fmt.Println("Error seeding Master Toko:", err)
	}
}