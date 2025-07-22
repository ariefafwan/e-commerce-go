package models

import (
	"e-commerce-go/internal/helpers"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterKategoriProduk struct {
	ID        uuid.UUID 	`gorm:"type:char(36);primaryKey"`
	IDParent  *uuid.UUID	`gorm:"type:char(36);"`
	Nama      string		`gorm:"type:varchar(50);not null;uniqueIndex"`
	Slug      string 		`gorm:"uniqueIndex"`
	Urutan 	  uint8			`gorm:"not null;"`
	CreatedAt time.Time		
	UpdatedAt time.Time

	DataParent  *MasterKategoriProduk 	`gorm:"foreignKey:IDParent;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DataProduk  []MasterProduk 		   	`gorm:"many2many:master_produk_kategori_produk;joinForeignKey:IDKategori;JoinReferences:IDProduk"`
}

func (MasterKategoriProduk) TableName() string {
	return "master_kategori_produk"
}

func (m *MasterKategoriProduk) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	
	baseSlug := helpers.GenerateSlug(m.Nama)
	slug := baseSlug
	counter := 1

	var maxUrutan uint8
	tx.Model(&MasterKategoriProduk{}).
		Select("COALESCE(MAX(urutan), 0)").
		Scan(&maxUrutan)
	m.Urutan = maxUrutan + 1

	for {
		var count int64
		err := tx.Model(&MasterKategoriProduk{}).Where("slug = ?", slug).Count(&count).Error
		if err != nil {
			return err
		}
		if count == 0 {
			break
		}
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}
	m.Slug = slug
	return nil
}

func (m *MasterKategoriProduk) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Nama") {
		baseSlug := helpers.GenerateSlug(m.Nama)
		slug := baseSlug
		counter := 1

		for {
			var count int64
			err := tx.Model(&MasterKategoriProduk{}).Where("slug = ?", slug).Count(&count).Error
			if err != nil {
				return err
			}
			if count == 0 {
				break
			}
			slug = fmt.Sprintf("%s-%d", baseSlug, counter)
			counter++
		}
		m.Slug = slug
	}
	return nil
}