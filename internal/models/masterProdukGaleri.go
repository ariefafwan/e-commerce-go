package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterProdukGaleri struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDProduk  uuid.UUID
	Gambar    string
	Urutan    uint8
	CreatedAt time.Time
	UpdatedAt time.Time

	Produk    MasterProduk `gorm:"foreignKey:IDProduk"`
}

func (MasterProdukGaleri) TableName() string {
	return "master_produk_galeri"
}

func (m *MasterProdukGaleri) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}