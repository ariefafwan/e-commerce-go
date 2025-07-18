package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksiItem struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDTransaksi     uuid.UUID
	IDProduk        uuid.UUID
	IDVariantProduk uuid.UUID
	Harga           float64
	Quantity        int
	Subtotal        float64
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Transaksi       Transaksi 			`gorm:"foreignKey:IDTransaksi"`
	Produk          MasterProduk       	`gorm:"foreignKey:IDProduk"`
	Variant         MasterProdukVariant `gorm:"foreignKey:IDVariantProduk"`
}

func (TransaksiItem) TableName() string {
	return "transaksi_transaksi_item"
}

func (ti *TransaksiItem) BeforeCreate(tx *gorm.DB) (err error) {
	ti.ID = uuid.New()
	return nil
}