package scraping

import (
	"e-commerce-go/external/raja_ongkir"
	"e-commerce-go/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func ScrapingProvinsi(db *gorm.DB) {
	fmt.Println("Scraping Master Provinsi...")

	provinceResponse, err := raja_ongkir.GetProvince()
	if err != nil {
		fmt.Println("Gagal ambil data province:", err)
		return
	}

	for _, prov := range provinceResponse.Data {
		err := db.Create(&models.MasterProvinsi{
			ID:   fmt.Sprintf("%d", prov.ID),
			Nama: prov.Name,
		}).Error
		if err != nil {
			fmt.Println("Gagal simpan provinsi:", prov.Name, err)
		}
	}
}
