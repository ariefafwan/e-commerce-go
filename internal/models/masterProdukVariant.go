package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterProdukVariant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDProduk  uuid.UUID	`gorm:"type:uuid;not null;"`
	NamaVariant string	`gorm:"type:varchar(255);not null;"`
	Harga     float64	`gorm:"type:float;not null;"`
	Stok      int		`gorm:"type:int;not null;default:1"`
	SKU       string	`gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time

	DataProduk    MasterProduk `gorm:"foreignKey:IDProduk;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MasterProdukVariant) TableName() string {
	return "master_produk_variant"
}

const skuCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomSKU(length int) string {
	rand.Seed(time.Now().UnixNano())
	sku := make([]byte, length)
	for i := range sku {
		sku[i] = skuCharset[rand.Intn(len(skuCharset))]
	}
	return string(sku)
}

func (m *MasterProdukVariant) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	for {
		sku := fmt.Sprintf("VARIANT-%s", generateRandomSKU(10))
		var count int64
		err := tx.Model(&MasterProdukVariant{}).Where("sku = ?", sku).Count(&count).Error
		if err != nil {
			return err
		}
		if count == 0 {
			m.SKU = sku
			break
		}
	}
	return nil
}