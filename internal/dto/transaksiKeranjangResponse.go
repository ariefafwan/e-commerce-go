package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransaksiKeranjangResponse struct {
	ID 			uuid.UUID `json:"id"`
	IDPelanggan uuid.UUID `json:"id_pelanggan"`
	BerlakuSampai time.Time `json:"berlaku_sampai"`
	DataPelanggan MasterPelangganPreload `json:"data_pelanggan"`
	DataItems     []TransaksiKeranjangItemResponse `json:"data_items"`

	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`

}

type TransaksiKeranjangItemResponse struct {
	ID          uuid.UUID `json:"id"`
	IDProduk    uuid.UUID `json:"id_produk"`
	IDKeranjang uuid.UUID `json:"id_keranjang"`
	IDVariantProduk uuid.UUID `json:"id_variant_produk"`
	Quantity      int       `json:"jumlah"`
	DataProduk    MasterProdukPreload `json:"data_produk"`
	DataProdukVariant MasterProdukVariantResponse `json:"data_produk_variant"`

	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}