package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleUser string

const (
    Pelanggan    RoleUser = "Pelanggan"
    Admin 		 RoleUser = "Admin"
)

func (ct *RoleUser) Scan(value interface{}) error {
    s, ok := value.(string)
    if !ok {
        b, ok := value.([]byte)
        if !ok {
            return fmt.Errorf("gagal scan RoleUser: tipe data tidak dikenal")
        }
        s = string(b)
    }
    switch RoleUser(s) {
		case Pelanggan, Admin:
			*ct = RoleUser(s) // Jika valid, tetapkan nilainya
			return nil
		default:
			// Jika tidak valid, kembalikan error
			return fmt.Errorf("nilai RoleUser tidak valid: %s", s)
    }
}

func (ct RoleUser) Value() (driver.Value, error) {
    return string(ct), nil
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Role      RoleUser 	`gorm:"type:varchar(255);not null;default:Pelanggan"`
	Nama      string	`gorm:"type:varchar(255);not null;"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string	`gorm:"type:varchar(255);not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DataPelanggan 	*MasterPelanggan 		`gorm:"foreignKey:IDUser"`
	Tokens    		[]PersonalAccessToken	`gorm:"foreignKey:IDUser"`
}

func (User) TableName() string {
    return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}