package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksiKeranjangItem struct {
	ID              uuid.UUID 	`gorm:"type:uuid;primaryKey"`
	IDKeranjang     uuid.UUID	`gorm:"type:uuid;not null;"`
	IDProduk        uuid.UUID	`gorm:"type:uuid;not null;"`
	IDVariantProduk uuid.UUID	`gorm:"type:uuid;not null;"`
	Quantity        int			`gorm:"type:int;not null;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	DataKeranjang   TransaksiKeranjang 	`gorm:"foreignKey:IDKeranjang;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataProduk      MasterProduk       	`gorm:"foreignKey:IDProduk;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	DataVariant     MasterProdukVariant `gorm:"foreignKey:IDVariantProduk;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

func (TransaksiKeranjangItem) TableName() string {
	return "transaksi_keranjang_item"
}

func (tki *TransaksiKeranjangItem) BeforeCreate(tx *gorm.DB) (err error) {
	tki.ID = uuid.New()
	return nil
}