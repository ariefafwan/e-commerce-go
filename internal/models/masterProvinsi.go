package models

import (
	"time"
)

type MasterProvinsi struct {
	ID        string 		`gorm:"type:varchar(255);primaryKey"`
	Nama      string 		`gorm:"type:varchar(255);"`
	CreatedAt time.Time
	UpdatedAt time.Time

	DataKota []MasterKota `gorm:"foreignKey:IDProvinsi"`
}

func (MasterProvinsi) TableName() string {
	return "master_provinsi"
}