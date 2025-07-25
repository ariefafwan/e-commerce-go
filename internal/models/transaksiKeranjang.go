package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksiKeranjang struct {
	ID             uuid.UUID 	`gorm:"type:uuid;primaryKey"`
	IDPelanggan    uuid.UUID	`gorm:"type:uuid;not null;"`
	BerlakuSampai  time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time

	DataPelanggan  MasterPelanggan 			`gorm:"foreignKey:IDPelanggan;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataItems      []TransaksiKeranjangItem `gorm:"foreignKey:IDKeranjang"`
}

func (TransaksiKeranjang) TableName() string {
	return "transaksi_keranjang"
}

func (t *TransaksiKeranjang) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return nil
}