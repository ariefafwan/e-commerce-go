package dto

import (
	"time"

	"github.com/google/uuid"
)

type MasterKategoriProdukResponse struct {
	ID        uuid.UUID `json:"id"`
	IDParent  *uuid.UUID `json:"id_parent"`
	Nama      string    `json:"nama"`
	Slug      string    `json:"slug"`
	Urutan    uint8     `json:"urutan"`
	DataParent  *[]MasterKategoriProdukResponse   `json:"data_parent"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MasterKategoriProdukPreload struct {
	ID        uuid.UUID `json:"id"`
	IDParent  *uuid.UUID `json:"id_parent"`
	Nama      string    `json:"nama"`
	Slug      string    `json:"slug"`
	Urutan    uint8     `json:"urutan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}