package dto

import (
	"time"

	"github.com/google/uuid"
)

type MasterPelangganResponse struct {
	ID        uuid.UUID 		`json:"id"`
	IDUser    uuid.UUID 		`json:"id_user"`
	NamaLengkap      string    	`json:"nama_lengkap"`
	NamaPanggilan    string    	`json:"nama_panggilan"`
	Phone            string    	`json:"phone"`
	DataUser  		 *UserResponse   		`json:"data_user"`
	DataAlamat		 *[]MasterAlamatPelangganPreload `json:"data_alamat"`
	CreatedAt time.Time 		`json:"created_at"`
	UpdatedAt time.Time 		`json:"updated_at"`
}

type MasterPelangganPreload struct {
	ID               uuid.UUID `json:"id"`
	IDUser    uuid.UUID 		`json:"id_user"`
	NamaLengkap      string    	`json:"nama_lengkap"`
	NamaPanggilan    string    	`json:"nama_panggilan"`
	Phone            string    	`json:"phone"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}