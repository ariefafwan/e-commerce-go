package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterAlamatPelanggan struct {
	ID            uuid.UUID `gorm:"type:char(36);primaryKey"`
	IDPelanggan   uuid.UUID `gorm:"type:char(36);not null;"`
	AlamatLengkap string	`gorm:"type:text;not null;"`
	KodePos       string	`gorm:"type:char(10);not null;"`
	Kota          string	`gorm:"type:varchar(50);not null;"`
	Negara        string	`gorm:"type:varchar(50);not null;"`
	NomorPenerima string	`gorm:"type:varchar(20);not null;"`
	NamaPenerima  string	`gorm:"type:varchar(50);not null;"`
	IsDefault     bool		`gorm:"type:boolean;not null;default:false"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	DataPelanggan MasterPelanggan `gorm:"foreignKey:IDPelanggan;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MasterAlamatPelanggan) TableName() string {
	return "master_alamat_pelanggan"
}

func (m *MasterAlamatPelanggan) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}