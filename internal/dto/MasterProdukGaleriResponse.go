package dto

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MasterProdukGaleriResponse struct {
	ID        uuid.UUID `json:"id"`
	IDProduk  uuid.UUID	`json:"id_produk"`
	Gambar    string	`json:"gambar"`
	Urutan    uint8		`json:"urutan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *MasterProdukGaleriResponse) FileUrl() string {
	return fmt.Sprintf("https://res.cloudinary.com/dnabtsqjy/image/upload/Toko/%s", m.Gambar)
}

func (m MasterProdukGaleriResponse) MarshalJSON() ([]byte, error) {
	type Alias MasterProdukGaleriResponse
	return json.Marshal(&struct {
		*Alias
		Gambar string `json:"gambar_url"`
	}{
		Alias:  (*Alias)(&m),
		Gambar: m.FileUrl(),
	})
}