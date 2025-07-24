package dto

import (
	"time"

	"github.com/google/uuid"
)

type MasterAlamatPelangganResponse struct {
	ID               uuid.UUID `json:"id"`
	Label            string    `json:"label"`
	IDPelanggan 	 uuid.UUID `json:"id_pelanggan"`
	AlamatLengkap    string    `json:"alamat_lengkap"`
	KodePos          string    `json:"kode_pos"`
	IDProvinsi       string    `json:"id_provinsi"`
	IDKota           string    `json:"id_kota"`
	IDKecamatan      string    `json:"id_kecamatan"`
	NomorPenerima    string    `json:"nomor_penerima"`
	NamaPenerima     string    `json:"nama_penerima"`
	DataPelanggan	 MasterPelangganPreload `json:"data_pelanggan"`
	DataKecamatan    MasterKecamatanResponse `json:"data_kecamatan"`
	IsDefault        bool      `json:"is_default"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type MasterAlamatPelangganPreload struct {
	ID               uuid.UUID `json:"id"`
	Label            string    `json:"label"`
	IDPelanggan 	 uuid.UUID `json:"id_pelanggan"`
	AlamatLengkap    string    `json:"alamat_lengkap"`
	KodePos          string    `json:"kode_pos"`
	IDProvinsi       string    `json:"id_provinsi"`
	IDKota           string    `json:"id_kota"`
	IDKecamatan      string    `json:"id_kecamatan"`
	Kelurahan        string    `json:"kelurahan"`
	NomorPenerima    string    `json:"nomor_penerima"`
	NamaPenerima     string    `json:"nama_penerima"`
	IsDefault        bool      `json:"is_default"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}