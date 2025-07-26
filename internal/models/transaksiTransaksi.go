package models

import (
	"database/sql/driver"
	"time"

	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusTransaksi string

const (
    Pending    	StatusTransaksi = "Pending"
	Cancelled 	StatusTransaksi = "Cancelled"
	Expired 	StatusTransaksi = "Expired"
    Paid 		StatusTransaksi = "Paid"
	Complete 	StatusTransaksi = "Complete"
)


// scan dan value ini untuk type enum, dokumentasi : https://stackoverflow.com/questions/68637265/how-can-i-add-enum-in-gorm
func (ct *StatusTransaksi) Scan(value interface{}) error {
    s, ok := value.(string)
    if !ok {
        b, ok := value.([]byte)
        if !ok {
            return fmt.Errorf("gagal scan status: status tidak dikenal")
        }
        s = string(b)
    }
    switch StatusTransaksi(s) {
		case Pending, Paid, Complete:
			*ct = StatusTransaksi(s)
			return nil
		default:
			return fmt.Errorf("nilai status tidak valid: %s", s)
    }
}

func (ct StatusTransaksi) Value() (driver.Value, error) {
    return string(ct), nil
}

type Transaksi struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	NoInvoice          string	 `gorm:"type:varchar(255);not null;uniqueIndex"`
	IDPelanggan        uuid.UUID `gorm:"type:uuid;not null;"`
	IDAlamatPelanggan  uuid.UUID `gorm:"type:uuid;not null;"`
	TotalHarga         float64	 `gorm:"type:float;not null;"`
	TotalOngkir        float64	 `gorm:"type:float;not null;"`
	JumlahItem		   int16	 `gorm:"type:int;not null;"`
	BeratTotal 		   float64	 `gorm:"type:float;not null;"`
	Pajak              float64	 `gorm:"type:float;not null;"`
	GrandTotal         float64	 `gorm:"type:float;not null;"`
	Notes              *string	 `gorm:"type:text;"`
	Status             StatusTransaksi 	`gorm:"type:varchar(255);not null;default:Pending"`
	ExpiredAt          *time.Time
	CancelledAt        *time.Time
	PaidAt     		   *time.Time
	CompleteAt         *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time

	DataPelanggan      MasterPelanggan 			`gorm:"foreignKey:IDPelanggan;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DataAlamat         MasterAlamatPelanggan 	`gorm:"foreignKey:IDAlamatPelanggan;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DataItems          []TransaksiItem			`gorm:"foreignKey:IDTransaksi"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}

func (t *Transaksi) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return nil
}