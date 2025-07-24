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
		Gambar:     "1753166749_data3.jpg",
		IDProvinsi: "25",
		IDKota:     "433",
		IDKecamatan: "4240",
		NoTelp:  	"123456789",
		AturanPajak: 11.0,
	}).Error

	if err != nil {
		fmt.Println("Error seeding Master Toko:", err)
	}
}