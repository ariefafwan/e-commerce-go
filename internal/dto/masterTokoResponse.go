package dto

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/google/uuid"
)

type MasterTokoResponse struct {
	ID         uuid.UUID `json:"id"`
	Nama       string	 `json:"nama"`
	Alamat     string	 `json:"alamat"`
	Gambar     string	 `json:"gambar"`
	NoTelp  string	 `json:"no_telp"`
	AturanPajak float64	 `json:"aturan_pajak"`
	DataKecamatan MasterKecamatanResponse `json:"data_kecamatan"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *MasterTokoResponse) FileUrl() string {
	return fmt.Sprintf("https://res.cloudinary.com/dnabtsqjy/image/upload/Toko/%s", m.Gambar)
}

func (m MasterTokoResponse) MarshalJSON() ([]byte, error) {
	type Alias MasterTokoResponse
	return json.Marshal(&struct {
		*Alias
		Gambar string `json:"gambar_url"`
	}{
		Alias:  (*Alias)(&m),
		Gambar: m.FileUrl(),
	})
}