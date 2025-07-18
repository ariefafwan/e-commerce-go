package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterPelanggan struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDUser        uuid.UUID
	NamaLengkap   string
	NamaPanggilan string
	CreatedAt     time.Time
	UpdatedAt     time.Time

	User          User             `gorm:"foreignKey:IDUser"`
	Alamat        []MasterAlamatPelanggan
	Keranjang     []TransaksiKeranjang
	Transaksi     []Transaksi
}

func (MasterPelanggan) TableName() string {
	return "master_pelanggan"
}

func (m *MasterPelanggan) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}