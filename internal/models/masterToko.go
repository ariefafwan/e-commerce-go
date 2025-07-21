package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterToko struct {
	ID         uuid.UUID 	`gorm:"type:char(36);primaryKey"`
	Nama       string		`gorm:"type:varchar(50);not null;uniqueIndex"`
	Alamat     string		`gorm:"type:text;not null;"`
	Gambar     string		`gorm:"type:varchar(50);not null;"`
	NomorToko  string		`gorm:"type:varchar(20);not null;"`
	AturanPajak float64		`gorm:"type:float;not null;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (MasterToko) TableName() string {
	return "master_toko"
}

func (m *MasterToko) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}