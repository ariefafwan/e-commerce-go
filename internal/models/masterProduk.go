package models

import (
	"time"

	"github.com/google/uuid"
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
			*ct = StatusProduk(s) // Jika valid, tetapkan nilainya
			return nil
		default:
			// Jika tidak valid, kembalikan error
			return fmt.Errorf("nilai status tidak valid: %s", s)
    }
}

func (ct StatusProduk) Value() (driver.Value, error) {
    return string(ct), nil
}

type MasterProduk struct {
	ID         uuid.UUID 	`gorm:"type:char(36);primaryKey"`
	Nama       string		`gorm:"type:varchar(50);not null;"`
	Thumbnail  string		`gorm:"type:varchar(255);not null;"`
	Slug       string 		`gorm:"uniqueIndex"`
	Status     StatusProduk `gorm:"type:varchar(50);not null;default:Aktif"`
	MinHarga   float64		`gorm:"tyoe:float;not null;"`
    MaxHarga   float64		`gorm:"type:float;not null;"`
	Deskripsi  string		`gorm:"type:text;not null;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	DataKategori   []MasterKategoriProduk 	`gorm:"many2many:master_produk_kategori_produk;joinForeignKey:IDProduk;JoinReferences:IDKategori"`
	DataGaleri     []MasterProdukGaleri		`gorm:"foreignKey:IDProduk"`
	DataVariant    []MasterProdukVariant	`gorm:"foreignKey:IDProduk"`
}

func (MasterProduk) TableName() string {
	return "master_produk"
}

func (m *MasterProduk) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}