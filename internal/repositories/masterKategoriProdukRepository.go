package repositories

import (
	"e-commerce-go/internal/models"

	"gorm.io/gorm"
)

type MasterKategoriProdukRepository interface {
	GetAll() ([]models.MasterKategoriProduk, error)
	GetByID(id string) (*models.MasterKategoriProduk, error)
	Create(kategori *models.MasterKategoriProduk) error
	Update(kategori *models.MasterKategoriProduk) error
	Delete(id string) error
}

type masterKategoriProdukRepo struct {
	db *gorm.DB
}

func NewMasterKategoriProdukRepository(db *gorm.DB) MasterKategoriProdukRepository {
	return &masterKategoriProdukRepo{db}
}

func (r *masterKategoriProdukRepo) GetAll() ([]models.MasterKategoriProduk, error) {
	var kategori []models.MasterKategoriProduk
	err := r.db.Find(&kategori).Error
	return kategori, err
}

func (r *masterKategoriProdukRepo) GetByID(id string) (*models.MasterKategoriProduk, error) {
	var kategori models.MasterKategoriProduk
	err := r.db.First(&kategori, "id = ?", id).Preload("DataParent").Error
	if err != nil {
		return nil, err
	}
	return &kategori, nil
}

func (r *masterKategoriProdukRepo) Create(kategori *models.MasterKategoriProduk) error {
	return r.db.Create(kategori).Error
}

func (r *masterKategoriProdukRepo) Update(kategori *models.MasterKategoriProduk) error {
	return r.db.Save(kategori).Error
}

func (r *masterKategoriProdukRepo) Delete(id string) error {
	return r.db.Delete(&models.MasterKategoriProduk{}, "id = ?", id).Error
}



