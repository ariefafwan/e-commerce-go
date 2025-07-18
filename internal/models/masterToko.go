package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterToko struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Nama       string
	Alamat     string
	Gambar     string
	NomorToko  string
	AturanPajak float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (MasterToko) TableName() string {
	return "master_toko"
}

func (m *MasterToko) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}