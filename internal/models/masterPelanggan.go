package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterPelanggan struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDUser        uuid.UUID	`gorm:"type:uuid;not null;"`
	NamaLengkap   string	`gorm:"type:varchar(255);not null;"`
	NamaPanggilan string	`gorm:"type:varchar(255);not null;"`
	Phone         string	`gorm:"type:varchar(255);not null;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DataAlamat  	[]MasterAlamatPelanggan	`gorm:"foreignKey:IDPelanggan"`
	DataKeranjang   []TransaksiKeranjang	`gorm:"foreignKey:IDPelanggan"`
	DataTransaksi   []Transaksi				`gorm:"foreignKey:IDPelanggan"`
	DataUser   		User 					`gorm:"foreignKey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (MasterPelanggan) TableName() string {
	return "master_pelanggan"
}

func (m *MasterPelanggan) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}