package repositories

import (
	"e-commerce-go/internal/models"

	"gorm.io/gorm"
)

type TransaksiKeranjangRepository interface {
	GetAll() ([]models.TransaksiKeranjang, error)
	GetByID(id string) (*models.TransaksiKeranjang, error)
	Create(keranjang *models.TransaksiKeranjang) error
	Update(keranjang *models.TransaksiKeranjang) error
	Delete(id string) error
	
}

type transaksiKeranjangRepo struct {
	db *gorm.DB
}

func NewTransaksiKeranjangRepository(db *gorm.DB) TransaksiKeranjangRepository {
	return &transaksiKeranjangRepo{db}
}

func (t *transaksiKeranjangRepo) GetAll() ([]models.TransaksiKeranjang, error) {
	var keranjangs []models.TransaksiKeranjang
	err := t.db.Find(&keranjangs).Preload("DataPelanggan").Preload("DataItems").Error
	return keranjangs, err
}

func (t *transaksiKeranjangRepo) GetByID(id string) (*models.TransaksiKeranjang, error) {
	var keranjang models.TransaksiKeranjang
	err := t.db.First(&keranjang, "id = ?", id).Error
	return &keranjang, err
}

func (t *transaksiKeranjangRepo) Create(keranjang *models.TransaksiKeranjang) error {
	return t.db.Create(keranjang).Error
}

func (t *transaksiKeranjangRepo) Update(keranjang *models.TransaksiKeranjang) error {
	return t.db.Save(keranjang).Error
}

func (t *transaksiKeranjangRepo) Delete(id string) error {
	return t.db.Delete(&models.TransaksiKeranjang{}, "id = ?", id).Error
}
