package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterProduk struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Nama       string
	Thumbnail  string
	Slug       string `gorm:"uniqueIndex"`
	Status     string `gorm:"type:enum('Aktif','Non Aktif')"`
	Deskripsi  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	Kategori   []MasterKategoriProduk `gorm:"many2many:master_produk_kategori_produk;joinForeignKey:IDProduk;JoinReferences:IDKategori"`
	Galeri     []MasterProdukGaleri
	Variant    []MasterProdukVariant
}

func (MasterProduk) TableName() string {
	return "master_produk"
}

func (m *MasterProduk) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}