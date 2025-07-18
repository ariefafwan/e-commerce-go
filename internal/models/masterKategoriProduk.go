package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterKategoriProduk struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ParentID  *uuid.UUID
	Nama      string
	Slug      string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Children  []MasterKategoriProduk `gorm:"foreignKey:ParentID"`
}

func (MasterKategoriProduk) TableName() string {
	return "master_kategori_produk"
}

func (m *MasterKategoriProduk) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}