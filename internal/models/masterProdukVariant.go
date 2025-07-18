package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterProdukVariant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDProduk  uuid.UUID
	Harga     float64
	Stok      int
	SKU       string
	CreatedAt time.Time
	UpdatedAt time.Time

	Produk    MasterProduk `gorm:"foreignKey:IDProduk"`
}

func (MasterProdukVariant) TableName() string {
	return "master_produk_variant"
}

func (m *MasterProdukVariant) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return nil
}