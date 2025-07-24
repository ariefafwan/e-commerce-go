package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterToko struct {
	ID         uuid.UUID 	`gorm:"type:uuid;primaryKey"`
	Nama       string		`gorm:"type:varchar(255);not null;uniqueIndex"`
	Alamat     string		`gorm:"type:text;not null;"`
	IDProvinsi   string		`gorm:"type:uuid;not null;"`
	IDKota       string		`gorm:"type:uuid;not null;"`
	IDKecamatan  string		`gorm:"type:uuid;not null;"`
	Gambar     string		`gorm:"type:varchar(255);not null;"`
	NoTelp     string		`gorm:"type:varchar(255);not null;"`
	AturanPajak float64		`gorm:"type:float;not null;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	DataProvinsi  MasterProvinsi `gorm:"foreignKey:IDProvinsi;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataKota      MasterKota `gorm:"foreignKey:IDKota;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DataKecamatan MasterKecamatan `gorm:"foreignKey:IDKecamatan;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MasterToko) TableName() string {
	return "master_toko"
}

func (m *MasterToko) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}