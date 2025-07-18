package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaksi struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	Status             string    `gorm:"type:enum('Pending','Paid','Complete')"`
	IDPelanggan        uuid.UUID
	IDAlamatPelanggan  uuid.UUID
	TotalHarga         float64
	TotalOngkir        float64
	Pajak              float64
	MulaiTransaksi     time.Time
	SelesaiTransaksi   time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time

	Pelanggan          MasterPelanggan `gorm:"foreignKey:IDPelanggan"`
	Alamat             MasterAlamatPelanggan `gorm:"foreignKey:IDAlamatPelanggan"`
	Items              []TransaksiItem
}

func (Transaksi) TableName() string {
	return "transaksi_transaksi"
}

func (t *Transaksi) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return nil
}