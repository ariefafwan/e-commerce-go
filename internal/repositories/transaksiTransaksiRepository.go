package repositories

import (
	"e-commerce-go/internal/models"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	GetAll() ([]models.Transaksi, error)
	GetByID(id string) (*models.Transaksi, error)
	Create(transaksi *models.Transaksi) error
	Update(transaksi *models.Transaksi) error
	SetComplete(id string) error
}

type transaksiRepo struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepo{db}
}

func (m *transaksiRepo) GetAll() ([]models.Transaksi, error) {
	var data []models.Transaksi
	err := m.db.Find(&data).Preload("DataPelanggan").Preload("DataAlamat").Preload("DataItems").Error
	return data, err
}

func (m *transaksiRepo) GetByID(id string) (*models.Transaksi, error) {
	var data models.Transaksi
	err := m.db.First(&data, "id = ?", id).Error
	return &data, err
}

func (m *transaksiRepo) Create(transaksi *models.Transaksi) error {
	return m.db.Create(transaksi).Error
}

func (m *transaksiRepo) Update(transaksi *models.Transaksi) error {
	return m.db.Save(transaksi).Error
}

func (m *transaksiRepo) SetComplete(id string) error {
	return m.db.Model(&models.Transaksi{}).Where("id = ?", id).Update("status", "Complete").Error
}