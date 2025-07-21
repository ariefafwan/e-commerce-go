package dto

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/google/uuid"

	"e-commerce-go/internal/models"
)

type MasterTokoResponse struct {
	ID         uuid.UUID `json:"id"`
	Nama       string	 `json:"nama"`
	Alamat     string	 `json:"alamat"`
	Gambar     string	 `json:"gambar"`
	NomorToko  string	 `json:"nomor_toko"`
	AturanPajak float64	 `json:"aturan_pajak"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *MasterTokoResponse) FileUrl() string {
	return fmt.Sprintf("https://example.com/uploads/%s", m.Gambar)
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

func FromMasterToko(model models.MasterToko) MasterTokoResponse {
	return MasterTokoResponse{
		ID:          model.ID,
		Nama:        model.Nama,
		Alamat:      model.Alamat,
		Gambar:      model.Gambar,
		NomorToko:   model.NomorToko,
		AturanPajak: model.AturanPajak,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func FromMasterTokoList(models []models.MasterToko) []MasterTokoResponse {
	var list []MasterTokoResponse
	for _, m := range models {
		list = append(list, FromMasterToko(m))
	}
	return list
}