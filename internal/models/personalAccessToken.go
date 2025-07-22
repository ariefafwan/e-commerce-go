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
			*ct = TipeToken(s) // Jika valid, tetapkan nilainya
			return nil
		default:
			// Jika tidak valid, kembalikan error
			return fmt.Errorf("nilai tipe token tidak valid: %s", s)
    }
}

func (ct TipeToken) Value() (driver.Value, error) {
    return string(ct), nil
}

type PersonalAccessToken struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	IDUser    uuid.UUID	`gorm:"type:char(36);not null;"`
	Nama      string	`gorm:"type:varchar(50);not null;"`
	Token     string	`gorm:"uniqueIndex"`
	Type      TipeToken	`gorm:"type:varchar(50);not null;"`
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