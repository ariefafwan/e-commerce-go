package dto

import (
	"time"
)

type MasterKecamatanResponse struct {
	ID        string `json:"id"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DataKota  MasterKotaResponse `json:"data_kota"` 
}