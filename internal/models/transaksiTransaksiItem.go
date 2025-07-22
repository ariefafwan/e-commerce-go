package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksiItem struct {
	ID              uuid.UUID `gorm:"type:char(36);primaryKey"`
	IDTransaksi     uuid.UUID `gorm:"type:char(36);not null;"`
	IDProduk        uuid.UUID `gorm:"type:char(36);not null;"`
	IDVariantProduk uuid.UUID `gorm:"type:char(36);not null;"`
	Harga           float64		`gorm:"type:float;not null;"`
	Quantity        int			`gorm:"type:int;not null;"`
	Subtotal        float64		`gorm:"type:float;not null;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	DataTransaksi   Transaksi 			`gorm:"foreignKey:IDTransaksi;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	DataProduk      MasterProduk       	`gorm:"foreignKey:IDProduk;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DataVariant     MasterProdukVariant `gorm:"foreignKey:IDVariantProduk;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (TransaksiItem) TableName() string {
	return "transaksi_item"
}

func (ti *TransaksiItem) BeforeCreate(tx *gorm.DB) (err error) {
	ti.ID = uuid.New()
	return nil
}