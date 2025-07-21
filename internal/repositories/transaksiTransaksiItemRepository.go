package repositories

import (
	"e-commerce-go/internal/models"

	"gorm.io/gorm"
)

type TransaksiItemRepository interface {
	GetAll() ([]models.TransaksiItem, error)
	Create(item *models.TransaksiItem) error
	Update(item *models.TransaksiItem) error
	Delete(id string) error
}

type transaksiItemRepo struct {
	db *gorm.DB
}

func NewTransaksiItemRepository(db *gorm.DB) TransaksiItemRepository {
	return &transaksiItemRepo{db}
}

func (m *transaksiItemRepo) GetAll() ([]models.TransaksiItem, error) {
	var item []models.TransaksiItem
	err := m.db.Find(&item).Preload("DataProduk").Preload("DataVariant").Preload("DataTransaksi.DataPelanggan").Preload("DataTransaksi.DataAlamat").Error
	return item, err
}

func (m *transaksiItemRepo) Create(item *models.TransaksiItem) error {
	return m.db.Create(item).Error
}

func (m *transaksiItemRepo) Update(item *models.TransaksiItem) error {
	return m.db.Save(item).Error
}

func (m *transaksiItemRepo) Delete(id string) error {
	return m.db.Delete(&models.TransaksiItem{}, "id = ?", id).Error
}