package dto

import (
	"time"
)

type MasterProvinsiResponse struct {
	ID        string 	`json:"id"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}