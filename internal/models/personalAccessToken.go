package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"database/sql/driver"
	"fmt"
)

type TipeToken string

const (
    AccessToken    	TipeToken = "access-token"
    RefreshToken 	TipeToken = "refresh-token"
)

// scan dan value ini untuk type enum, dokumentasi : https://stackoverflow.com/questions/68637265/how-can-i-add-enum-in-gorm
func (ct *TipeToken) Scan(value interface{}) error {
    s, ok := value.(string)
    if !ok {
        b, ok := value.([]byte)
        if !ok {
            return fmt.Errorf("gagal scan tipe token: tipe token tidak dikenal")
        }
        s = string(b)
    }
    switch TipeToken(s) {
		case AccessToken, RefreshToken:
			*ct = TipeToken(s)
			return nil
		default:
			return fmt.Errorf("nilai tipe token tidak valid: %s", s)
    }
}

func (ct TipeToken) Value() (driver.Value, error) {
    return string(ct), nil
}

type PersonalAccessToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDUser    uuid.UUID	`gorm:"type:uuid;not null;"`
	Nama      string	`gorm:"type:varchar(255);not null;"`
	Token     string	`gorm:"uniqueIndex"`
	Type      TipeToken	`gorm:"type:varchar(255);not null;"`
	ExpiredAt time.Time
	RevokedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:IDUser"`
}

func (PersonalAccessToken) TableName() string {
    return "personal_access_tokens"
}

func (pct *PersonalAccessToken) BeforeCreate(tx *gorm.DB) (err error) {
	pct.ID = uuid.New()
	return
}