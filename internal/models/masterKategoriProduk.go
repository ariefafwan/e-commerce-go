package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterKategoriProduk struct {
	ID        uuid.UUID 	`gorm:"type:char(36);primaryKey"`
	ParentID  *uuid.UUID	`gorm:"type:char(36);"`
	Nama      string		`gorm:"type:varchar(50);not null;uniqueIndex"`
	Slug      string 		`gorm:"uniqueIndex"`
	Urutan 	  uint8			`gorm:"not null;"`
	CreatedAt time.Time		
	UpdatedAt time.Time

	DataParent  *MasterKategoriProduk 	`gorm:"foreignKey:ParentID"`
	DataProduk  []MasterProduk 		   	`gorm:"many2many:master_produk_kategori_produk;joinForeignKey:IDKategori;JoinReferences:IDProduk"`
}

func (MasterKategoriProduk) TableName() string {
	return "master_kategori_produk"
}

func (m *MasterKategoriProduk) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}