package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterAlamatPelanggan struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDPelanggan   uuid.UUID
	AlamatLengkap string
	KodePos       string
	Kota          string
	Negara        string
	NomorPenerima string
	NamaPenerima  string
	IsDefault     bool
	CreatedAt     time.Time
	UpdatedAt     time.Time

	Pelanggan     MasterPelanggan `gorm:"foreignKey:IDPelanggan"`
}

func (MasterAlamatPelanggan) TableName() string {
	return "master_alamat_pelanggan"
}

func (m *MasterAlamatPelanggan) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}