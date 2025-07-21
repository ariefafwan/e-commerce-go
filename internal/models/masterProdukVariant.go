package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterProdukVariant struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	IDProduk  uuid.UUID	`gorm:"type:char(36);not null;"`
	NamaVariant string	`gorm:"type:varchar(50);not null;"`
	Harga     float64	`gorm:"type:float;not null;"`
	Stok      int		`gorm:"type:int;not null;default:1"`
	SKU       string	`gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time

	DataProduk    MasterProduk `gorm:"foreignKey:IDProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MasterProdukVariant) TableName() string {
	return "master_produk_variant"
}

func (m *MasterProdukVariant) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return nil
}