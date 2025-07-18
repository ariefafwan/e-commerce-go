package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonalAccessToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	IDUser    uuid.UUID
	Nama      string
	Token     string
	Type      string    `gorm:"type:enum('access-token','refresh-token')"`
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