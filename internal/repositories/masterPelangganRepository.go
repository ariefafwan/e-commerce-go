package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MasterPelangganRepository interface {
	GetAll() ([]dto.MasterPelangganResponse, error)
	GetByID(id string) (*dto.MasterPelangganResponse, error)
	Create(pelanggan *models.MasterPelanggan) error
	Update(pelanggan *models.MasterPelanggan) error
	Delete(id string) error
}

type masterPelangganRepo struct {
	db *gorm.DB
}

func NewMasterPelangganRepository(db *gorm.DB) MasterPelangganRepository {
	return &masterPelangganRepo{db}
}

func (r *masterPelangganRepo) GetAll() ([]dto.MasterPelangganResponse, error) {
	var pelanggan []models.MasterPelanggan
	err := r.db.Find(&pelanggan).Preload("DataUser").Error
	if err != nil {
		return nil, err
	}

	var pelangganResponse []dto.MasterPelangganResponse
	if err := copier.Copy(&pelangganResponse, &pelanggan); err != nil {
		return nil, err
	}
	return pelangganResponse, err
}

func (r *masterPelangganRepo) GetByID(id string) (*dto.MasterPelangganResponse, error) {
	var pelanggan models.MasterPelanggan
	err := r.db.First(&pelanggan, "id = ?", id).Preload("DataUser").Preload("DataAlamat").Error
	if err != nil {
		return nil, err
	}
	var pelangganResponse dto.MasterPelangganResponse
	if err := copier.Copy(&pelangganResponse, &pelanggan); err != nil {
		return nil, err
	}
	return &pelangganResponse, nil
}

func (r *masterPelangganRepo) Create(pelanggan *models.MasterPelanggan) error {
	return r.db.Create(pelanggan).Error
}

func (r *masterPelangganRepo) Update(pelanggan *models.MasterPelanggan) error {
	return r.db.Save(pelanggan).Error
}

func (r *masterPelangganRepo) Delete(id string) error {
	return r.db.Delete(&models.MasterPelanggan{}, "id = ?", id).Error
}