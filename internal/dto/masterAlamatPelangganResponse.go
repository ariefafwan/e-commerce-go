package dto

import (
	"time"

	"github.com/google/uuid"
)

type MasterAlamatPelangganResponse struct {
	ID               uuid.UUID `json:"id"`
	IDPelanggan 	 uuid.UUID `json:"id_pelanggan"`
	AlamatLengkap    string    `json:"alamat_lengkap"`
	KodePos          string    `json:"kode_pos"`
	Kota             string    `json:"kota"`
	Negara           string    `json:"negara"`
	NomorPenerima    string    `json:"nomor_penerima"`
	NamaPenerima     string    `json:"nama_penerima"`
	IsDefault        bool      `json:"is_default"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}