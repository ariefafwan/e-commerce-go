package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterAlamatPelanggan struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDPelanggan   uuid.UUID `gorm:"type:uuid;not null;"`
	Label         string	`gorm:"type:varchar(255);not null;"`
	AlamatLengkap string	`gorm:"type:text;not null;"`
	KodePos       string	`gorm:"type:varchar(255);not null;"`
	IDProvinsi    string 	`gorm:"type:varchar(255);not null;"`
	IDKota        string 	`gorm:"type:varchar(255);not null;"`
	IDKecamatan   string 	`gorm:"type:varchar(255);not null;"`
	NomorPenerima string	`gorm:"type:varchar(255);not null;"`
	NamaPenerima  string	`gorm:"type:varchar(255);not null;"`
	IsDefault     bool		`gorm:"type:boolean;not null;default:false"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	DataProvinsi  MasterProvinsi 	`gorm:"foreignKey:IDProvinsi;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataKota      MasterKota 		`gorm:"foreignKey:IDKota;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataKecamatan MasterKecamatan 	`gorm:"foreignKey:IDKecamatan;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataPelanggan MasterPelanggan 	`gorm:"foreignKey:IDPelanggan;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MasterAlamatPelanggan) TableName() string {
	return "master_alamat_pelanggan"
}

func (m *MasterAlamatPelanggan) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}