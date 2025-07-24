package models

import (
	"time"
)

type MasterKota struct {
	ID         string 		`gorm:"type:varchar(255);primaryKey"`
	IDProvinsi string 		`gorm:"type:varchar(255);not null;"`
	Nama       string 		`gorm:"type:varchar(255);"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	DataProvinsi MasterProvinsi `gorm:"foreignKey:IDProvinsi;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MasterKota) TableName() string {
	return "master_kota"
}