package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterPelanggan struct {
	ID            uuid.UUID `gorm:"type:char(36);primaryKey"`
	IDUser        uuid.UUID	`gorm:"type:char(36);not null;"`
	NamaLengkap   string	`gorm:"type:varchar(50);not null;"`
	NamaPanggilan string	`gorm:"type:varchar(50);not null;"`
	Phone         string	`gorm:"type:varchar(20);not null;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DataAlamatPelanggan  	[]MasterAlamatPelanggan	`gorm:"foreignKey:IDPelanggan"`
	DataKeranjang     		[]TransaksiKeranjang	`gorm:"foreignKey:IDPelanggan"`
	DataTransaksi     		[]Transaksi				`gorm:"foreignKey:IDPelanggan"`
	DataUser   				User 					`gorm:"foreignKey:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (MasterPelanggan) TableName() string {
	return "master_pelanggan"
}

func (m *MasterPelanggan) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}