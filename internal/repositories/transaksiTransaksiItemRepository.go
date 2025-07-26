package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type TransaksiItemRepository interface {
	GetAll(id_transaksi string) ([]dto.TransaksiItemResponse, error)
}

type transaksiItemRepo struct {
	db *gorm.DB
}

func NewTransaksiItemRepository(db *gorm.DB) TransaksiItemRepository {
	return &transaksiItemRepo{db}
}

func (m *transaksiItemRepo) GetAll(id_transaksi string) ([]dto.TransaksiItemResponse ,error) {
	var data models.TransaksiItem
	err := m.db.Preload("DataProduk").Preload("DataVariant").Preload("DataTransaksi.DataPelanggan").Find(&data, "id_transaksi = ?", id_transaksi).Error

	var response []dto.TransaksiItemResponse
	if err := copier.Copy(&response, &data); err != nil {
		return nil, err
	}
	return response, err
}