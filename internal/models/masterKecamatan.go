package models

import (
	"time"
)

type MasterKecamatan struct {
	ID        string 		`gorm:"type:varchar(255);primaryKey"`
	IDKota    string 		`gorm:"type:varchar(255);not null;"`
	Nama      string 		`gorm:"type:varchar(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
	DataKota  MasterKota `gorm:"foreignKey:IDKota;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MasterKecamatan) TableName() string {
	return "master_kecamatan"
}