package repositories

import (
	"e-commerce-go/internal/models"

	"gorm.io/gorm"
)

type TransaksiKeranjangItemRepository interface {
	GetAll() ([]models.TransaksiKeranjangItem, error)
	Create(item *models.TransaksiKeranjangItem) error
	Update(item *models.TransaksiKeranjangItem) error
	Delete(id string) error
	UpdateQuantity(id string, quantity int) error
}

type transaksiKeranjangItemRepo struct {
	db *gorm.DB
}

func NewTransaksiKeranjangItemRepository(db *gorm.DB) TransaksiKeranjangItemRepository {
	return &transaksiKeranjangItemRepo{db}
}

func (t *transaksiKeranjangItemRepo) GetAll() ([]models.TransaksiKeranjangItem, error) {
	var items []models.TransaksiKeranjangItem
	err := t.db.Find(&items).Preload("DataProduk").Preload("DataVariant").Preload("DataKeranjang.DataPelanggan").Error
	return items, err
}

func (t *transaksiKeranjangItemRepo) Create(item *models.TransaksiKeranjangItem) error {
	return t.db.Create(item).Error
}

func (t *transaksiKeranjangItemRepo) Update(item *models.TransaksiKeranjangItem) error {
	return t.db.Save(item).Error
}

func (t *transaksiKeranjangItemRepo) Delete(id string) error {
	return t.db.Where("id = ?", id).Delete(&models.TransaksiKeranjangItem{}).Error
}


func (t *transaksiKeranjangItemRepo) UpdateQuantity(id string, quantity int) error {
	return t.db.Model(&models.TransaksiKeranjangItem{}).Where("id = ?", id).Update("quantity", quantity).Error
}