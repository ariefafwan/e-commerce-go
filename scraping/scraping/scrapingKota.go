package scraping

import (
	"e-commerce-go/external/raja_ongkir"
	"e-commerce-go/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func ScrapingKota(db *gorm.DB) {
	fmt.Println("Scraping Master Kota...")

	var provinces []models.MasterProvinsi
	if err := db.Find(&provinces).Error; err != nil {
		fmt.Println("Gagal ambil data provinsi:", err)
		return
	}

	for _, prov := range provinces {
		cityRes, err := raja_ongkir.GetCity(prov.ID)
		if err != nil {
			fmt.Println("Gagal ambil data provinsi:", err)
			return
		}

		for _, city := range cityRes.Data {
			err := db.Create(&models.MasterKota{
				ID:         fmt.Sprintf("%d", city.ID),
				Nama:       city.Name,
				IDProvinsi: prov.ID,
			}).Error
			if err != nil {
				fmt.Println("Gagal simpan kota:", city.Name, err)
			}
		}
	}
}
