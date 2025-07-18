package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksiKeranjangItem struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDKeranjang     uuid.UUID
	IDProduk        uuid.UUID
	IDVariantProduk uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Keranjang       TransaksiKeranjang `gorm:"foreignKey:IDKeranjang"`
	Produk          MasterProduk       `gorm:"foreignKey:IDProduk"`
	Variant         MasterProdukVariant `gorm:"foreignKey:IDVariantProduk"`
}

func (TransaksiKeranjangItem) TableName() string {
	return "transaksi_keranjang_item"
}

func (tki *TransaksiKeranjangItem) BeforeCreate(tx *gorm.DB) (err error) {
	tki.ID = uuid.New()
	return nil
}