package repositories

import (
	"e-commerce-go/internal/models"

	"gorm.io/gorm"
)

type MasterPelangganRepository interface {
	GetAll() ([]models.MasterPelanggan, error)
	GetByID(id string) (*models.MasterPelanggan, error)
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

func (r *masterPelangganRepo) GetAll() ([]models.MasterPelanggan, error) {
	var pelanggan []models.MasterPelanggan
	err := r.db.Find(&pelanggan).Preload("DataUser").Error
	return pelanggan, err
}

func (r *masterPelangganRepo) GetByID(id string) (*models.MasterPelanggan, error) {
	var pelanggan models.MasterPelanggan
	err := r.db.First(&pelanggan, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &pelanggan, nil
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