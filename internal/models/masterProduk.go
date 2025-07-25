package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gorm.io/gorm"

	"database/sql/driver"
	"fmt"
)

type StatusProduk string

const (
    Aktif    	StatusProduk = "Aktif"
    NonAktif 	StatusProduk = "Non Aktif"
)

func (ct *StatusProduk) Scan(value any) error {
    s, ok := value.(string)
    if !ok {
        b, ok := value.([]byte)
        if !ok {
            return fmt.Errorf("gagal scan status: status tidak dikenal")
        }
        s = string(b)
    }
    switch StatusProduk(s) {
		case Aktif, NonAktif:
			*ct = StatusProduk(s)
			return nil
		default:
			return fmt.Errorf("nilai status tidak valid: %s", s)
    }
}

func (ct StatusProduk) Value() (driver.Value, error) {
    return string(ct), nil
}

type MasterProduk struct {
	ID         uuid.UUID 	`gorm:"type:uuid;primaryKey"`
	Nama       string		`gorm:"type:varchar(255);not null;"`
	Thumbnail  string		`gorm:"type:varchar(255);not null;"`
	Slug       string 		`gorm:"uniqueIndex"`
	Status     StatusProduk `gorm:"type:varchar(50);not null;default:Non Aktif"`
	MinHarga   float64		`gorm:"type:float;not null;"`
    MaxHarga   float64		`gorm:"type:float;not null;"`
	Berat      float64		`gorm:"type:float;not null"`
	Deskripsi  string		`gorm:"type:text;not null;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	DataKategori   []MasterKategoriProduk 	`gorm:"many2many:master_produk_kategori_produk;foreignKey:ID;joinForeignKey:IDProduk;References:ID;joinReferences:IDKategori"`
	DataGaleri     []MasterProdukGaleri		`gorm:"foreignKey:IDProduk"`
	DataVariant    []MasterProdukVariant	`gorm:"foreignKey:IDProduk"`
}

func (MasterProduk) TableName() string {
	return "master_produk"
}

func (m *MasterProduk) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	baseSlug := slug.Make(m.Nama)
	slug := baseSlug
	counter := 1

	for {
		var count int64
		err := tx.Model(&MasterProduk{}).Where("slug = ?", slug).Count(&count).Error
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