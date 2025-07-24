package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type MasterTokoRepository interface {
	GetToko() (dto.MasterTokoResponse, error)
	GetByID(id string) (dto.MasterTokoResponse, error)
	// Create(toko *models.MasterToko) error
	Update(id string, toko models.MasterToko) error
	// Delete(id string) error
}

type masterTokoRepo struct {
	db *gorm.DB
}

func NewMasterTokoRepository(db *gorm.DB) MasterTokoRepository {
	return &masterTokoRepo{db}
}

func (r *masterTokoRepo) GetToko() (dto.MasterTokoResponse, error) {
	var toko models.MasterToko
	err := r.db.Take(&toko).Preload("DataKecamatan.DataKota.DataProvinsi").Error

	var tokoResponse dto.MasterTokoResponse
	if err := copier.Copy(&tokoResponse, &toko); err != nil {
		return dto.MasterTokoResponse{}, err
	}
	
	return tokoResponse, err
}

func (r *masterTokoRepo) GetByID(id string) (dto.MasterTokoResponse, error) {
	var toko models.MasterToko
	err := r.db.First(&toko, "id = ?", id).Error
	
	if err != nil {
		return dto.MasterTokoResponse{}, err
	}
	
	var tokoResponse dto.MasterTokoResponse
	if err := copier.Copy(&tokoResponse, &toko); err != nil {
		return dto.MasterTokoResponse{}, err
	}
	return tokoResponse, nil
}

// func (r *masterTokoRepo) Create(toko *models.MasterToko) error {
// 	return r.db.Create(toko).Error
// }

func (r *masterTokoRepo) Update(id string, toko models.MasterToko) error {
	return r.db.Model(&models.MasterToko{}).Where("id = ?", id).Updates(&toko).Error
}

// func (r *masterTokoRepo) Delete(id string) error {
// 	return r.db.Delete(&models.MasterToko{}, "id = ?", id).Error
// }

