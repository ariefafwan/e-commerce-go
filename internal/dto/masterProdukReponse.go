package dto

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MasterProdukResponse struct {
	ID                uuid.UUID `json:"id"`
	Nama              string    `json:"nama"`
	Thumbnail         string    `json:"thumbnail"`
	Slug              string    `json:"slug"`
	Status            string    `json:"status"`
	MinHarga          float64   `json:"min_harga"`
	MaxHarga          float64   `json:"max_harga"`
	Deskripsi         string    `json:"deskripsi"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DataKategori 	   *[]MasterProdukKategoriProduk `json:"data_kategori"`
	DataGaleri         *[]MasterProdukGaleriResponse `json:"data_galeri"`
	DataVariant        *[]MasterProdukVariantResponse `json:"data_variant"`
}

type MasterProdukPreload struct {
	ID                uuid.UUID `json:"id"`
	Nama              string    `json:"nama"`
	Thumbnail         string    `json:"thumbnail"`
	Slug              string    `json:"slug"`
	Status            string    `json:"status"`
	MinHarga          float64   `json:"min_harga"`
	MaxHarga          float64   `json:"max_harga"`
	Deskripsi         string    `json:"deskripsi"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type MasterProdukKategoriProduk struct {
	ID     uuid.UUID 	`json:"id_kategori"`
}

func (m *MasterProdukResponse) FileUrl() string {
	return fmt.Sprintf("https://res.cloudinary.com/dnabtsqjy/image/upload/Toko/%s", m.Thumbnail)
}

func (m MasterProdukResponse) MarshalJSON() ([]byte, error) {
	type Alias MasterProdukResponse
	return json.Marshal(&struct {
		*Alias
		Thumbnail string `json:"thumbnail_url"`
	}{
		Alias:  (*Alias)(&m),
		Thumbnail: m.FileUrl(),
	})
}

func (m *MasterProdukPreload) FileUrl() string {
	return fmt.Sprintf("https://res.cloudinary.com/dnabtsqjy/image/upload/Toko/%s", m.Thumbnail)
}

func (m MasterProdukPreload) MarshalJSON() ([]byte, error) {
	type Alias MasterProdukPreload
	return json.Marshal(&struct {
		*Alias
		Thumbnail string `json:"thumbnail_url"`
	}{
		Alias:  (*Alias)(&m),
		Thumbnail: m.FileUrl(),
	})
}

