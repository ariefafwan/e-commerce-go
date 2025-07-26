package dto

import (
	"e-commerce-go/internal/models"
	"time"

	"github.com/google/uuid"
)

type TransaksiResponse struct {
	ID 			uuid.UUID `json:"id"`
	IDPelanggan uuid.UUID `json:"id_pelanggan"`
	IDAlamatPelanggan uuid.UUID `json:"id_alamat_pelanggan"`
	NoInvoice 	*string `json:"no_invoice"`
	TotalHarga 	float64 `json:"total_harga"`
	TotalOngkir float64 `json:"total_ongkir"`
	JumlahItem 	int16 `json:"jumlah_item"`
	BeratTotal 	float64 `json:"berat_total"`
	Pajak 		float64 `json:"pajak"`
	GrandTotal 	float64 `json:"grand_total"`
	Notes 		*string `json:"notes"`
	Status 		models.StatusTransaksi `json:"status"`
	PilihanOngkir *[]PilihanOngkirResponse `json:"pilihan_ongkir"`
	DataItems 	[]TransaksiItemResponse `json:"data_items"`
	DataPelanggan MasterPelangganPreload `json:"data_pelanggan"`
	DataAlamat   MasterAlamatPelangganResponse `json:"data_alamat"`
	PendingSampai *time.Time `json:"pending_sampai"`
	PaidAt  	*time.Time `json:"paid_at"`
	CompleteAt  *time.Time  `json:"complete_at"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

type TransaksiItemResponse struct {
	ID 			uuid.UUID `json:"id"`
	IDTransaksi uuid.UUID `json:"id_transaksi"`
	IDProduk    uuid.UUID `json:"id_produk"`
	IDVariantProduk uuid.UUID `json:"id_variant_produk"`
	Quantity 	int `json:"jumlah"`
	Subtotal 	float64 `json:"subtotal"`
	DataProduk 	MasterProdukPreload `json:"data_produk"`
	DataVariant MasterProdukVariantResponse `json:"data_produk_variant"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

type PilihanOngkirResponse struct {
	NamaLayanan string `json:"nama_layanan"`
	Estimasi    string `json:"estimasi"`
	Harga       int    `json:"harga"`
}