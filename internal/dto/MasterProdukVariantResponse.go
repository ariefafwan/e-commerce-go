package dto

import (
	"time"

	"github.com/google/uuid"
)

type MasterProdukVariantResponse struct {
	ID        uuid.UUID `json:"id"`
	IDProduk  uuid.UUID	`json:"id_produk"`
	NamaVariant string	`json:"nama_variant"`
	Harga     float64	`json:"harga"`
	Stok      int		`json:"stok"`
	SKU       string	`json:"sku"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}