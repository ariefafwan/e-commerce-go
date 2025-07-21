package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterProdukGaleri struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	IDProduk  uuid.UUID	`gorm:"type:char(36);not null;"`
	Gambar    string	`gorm:"type:varchar(255);not null;"`
	Urutan    uint8		`gorm:"not null;"`
	CreatedAt time.Time	
	UpdatedAt time.Time

	DataProduk    MasterProduk `gorm:"foreignKey:IDProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MasterProdukGaleri) TableName() string {
	return "master_produk_galeri"
}

func (m *MasterProdukGaleri) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}