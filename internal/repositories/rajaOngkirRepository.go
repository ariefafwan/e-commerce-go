package repositories

import (
	"e-commerce-go/external/raja_ongkir"
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"
	"fmt"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type RajaOngkirRepository interface {
	GetProvince() ([]dto.MasterProvinsiResponse, error)
	GetCity(provinceID string) ([]dto.MasterKotaResponse, error)
	GetDistrict(cityID string) ([]dto.MasterKecamatanResponse, error)
}

type rajaOngkirRepository struct {
	DB *gorm.DB
}

func NewRajaOngkirRepository(db *gorm.DB) RajaOngkirRepository {
	return &rajaOngkirRepository{db}
}

func (r *rajaOngkirRepository) GetProvince() ([]dto.MasterProvinsiResponse, error) {
	var provinces []models.MasterProvinsi
	if err := r.DB.Find(&provinces).Error; err != nil {
		return nil, err
	}
	var provinceResponse []dto.MasterProvinsiResponse
	if err := copier.Copy(&provinceResponse, &provinces); err != nil {
		return nil, err
	}
	return provinceResponse, nil
}

func (r *rajaOngkirRepository) GetCity(provinceID string) ([]dto.MasterKotaResponse, error) {
	var cities []models.MasterKota
	if err := r.DB.Preload("DataProvinsi").Find(&cities, "id_provinsi = ?", provinceID).Error; err != nil {
		return nil, err
	}
	var cityResponse []dto.MasterKotaResponse
	if err := copier.Copy(&cityResponse, &cities); err != nil {
		return nil, err
	}
	return cityResponse, nil
}

func (r *rajaOngkirRepository) GetDistrict(cityID string) ([]dto.MasterKecamatanResponse, error) {
	var districts []models.MasterKecamatan
	var count int64
	
	err := r.DB.Model(&models.MasterKecamatan{}).Where("id_kota = ?", cityID).Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count == 0 {
		districtResponse, err := raja_ongkir.GetDistrict(cityID)
		if err != nil {
			fmt.Println("Gagal ambil data provinsi:", err)
			return nil, err
		}

		for _, dist := range districtResponse.Data {
			err := r.DB.Create(&models.MasterKecamatan{
				ID:   fmt.Sprintf("%d", dist.ID),
				IDKota: cityID,
				Nama: dist.Name,
			}).Error
			if err != nil {
				fmt.Println("Gagal simpan provinsi:", dist.Name, err)
			}
		}
	}

	if err := r.DB.Preload("DataKota.DataProvinsi").Find(&districts, "id_kota = ?", cityID).Error; err != nil {
		return nil, err
	}

	var districtResponse []dto.MasterKecamatanResponse
	if err := copier.Copy(&districtResponse, &districts); err != nil {
		return nil, err
	}
	return districtResponse, nil
}