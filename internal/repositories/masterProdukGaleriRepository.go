package repositories

import (
	"e-commerce-go/internal/models"

	"gorm.io/gorm"
)

type MasterProdukGaleriRepository interface {
	GetAll() ([]models.MasterProdukGaleri, error)
	Create(galeri *models.MasterProdukGaleri) error
	Update(galeri *models.MasterProdukGaleri) error
	Delete(id string) error
}

type masterProdukGaleriRepo struct {
	db *gorm.DB
}

func NewMasterProdukGaleriRepository(db *gorm.DB) MasterProdukGaleriRepository {
	return &masterProdukGaleriRepo{db}
}

func (m *masterProdukGaleriRepo) GetAll() ([]models.MasterProdukGaleri, error) {
	var galeri []models.MasterProdukGaleri
	err := m.db.Find(&galeri).Error
	return galeri, err
}

func (m *masterProdukGaleriRepo) Create(galeri *models.MasterProdukGaleri) error {
	return m.db.Create(galeri).Error
}

func (m *masterProdukGaleriRepo) Update(galeri *models.MasterProdukGaleri) error {
	return m.db.Save(galeri).Error
}

func (m *masterProdukGaleriRepo) Delete(id string) error {
	return m.db.Delete(&models.MasterProdukGaleri{}, "id = ?", id).Error
}