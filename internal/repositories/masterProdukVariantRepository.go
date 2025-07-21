package repositories

import (
	"e-commerce-go/internal/models"

	"errors"

	"gorm.io/gorm"
)

type MasterProdukVariantRepository interface {
	GetAll() ([]models.MasterProdukVariant, error)
	Create(masterProdukVariant models.MasterProdukVariant) (models.MasterProdukVariant, error)
	Update(masterProdukVariant models.MasterProdukVariant) (models.MasterProdukVariant, error)
	Delete(id string) error
}

type masterProdukVariantRepo struct {
	db *gorm.DB
}

func NewMasterProdukVariantRepository(db *gorm.DB) MasterProdukVariantRepository {
	return &masterProdukVariantRepo{db}
}

func (m *masterProdukVariantRepo) GetAll() ([]models.MasterProdukVariant, error) {
	var masterProdukVariant []models.MasterProdukVariant
	err := m.db.Find(&masterProdukVariant).Error
	return masterProdukVariant, err
}

func (m *masterProdukVariantRepo) Create(masterProdukVariant models.MasterProdukVariant) (models.MasterProdukVariant, error) {
	err := m.db.Create(&masterProdukVariant).Error
	return masterProdukVariant, err
}

func (m *masterProdukVariantRepo) Update(masterProdukVariant models.MasterProdukVariant) (models.MasterProdukVariant, error) {
	err := m.db.Save(&masterProdukVariant).Error
	return masterProdukVariant, err
}

var ErrorProdukVariant = errors.New("produk harus memiliki setidaknya 1 variant, silahkan update jika ingin merubah data ini")

func (m *masterProdukVariantRepo) Delete(id string) error {
	var count int64
	var data models.MasterProdukVariant
	err := m.db.First(&data, "id = ?", id).Error
	if err != nil {
		return err
	}
	m.db.Model(&models.MasterProdukVariant{}).Where("id_produk = ?", data.IDProduk).Count(&count)
	if count < 2  {
		return ErrorProdukVariant
	}
	return m.db.Delete(&models.MasterProdukVariant{}, "id = ?", id).Error
}