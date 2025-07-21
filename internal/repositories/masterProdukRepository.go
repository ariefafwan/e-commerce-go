package repositories

import (
	"e-commerce-go/internal/models"

	"gorm.io/gorm"
)

type MasterProdukRepository interface {
	GetAll() ([]models.MasterProduk, error)
	GetByID(id string) (*models.MasterProduk, error)
	Create(produk *models.MasterProduk) error
	Update(produk *models.MasterProduk) error
	Delete(id string) error
}

type masterProdukRepo struct {
	db *gorm.DB
}

func NewMasterProdukRepository(db *gorm.DB) MasterProdukRepository {
	return &masterProdukRepo{db}
}

func (m *masterProdukRepo) GetAll() ([]models.MasterProduk, error) {
	var produk []models.MasterProduk
	err := m.db.Find(&produk).Error
	return produk, err
}

func (m *masterProdukRepo) GetByID(id string) (*models.MasterProduk, error) {
	var produk models.MasterProduk
	err := m.db.First(&produk, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &produk, nil
}

func (m *masterProdukRepo) Create(produk *models.MasterProduk) error {
	return m.db.Create(produk).Error
}

func (m *masterProdukRepo) Update(produk *models.MasterProduk) error {
	return m.db.Save(produk).Error
}

func (m *masterProdukRepo) Delete(id string) error {
	return m.db.Delete(&models.MasterProduk{}, "id = ?", id).Error
}