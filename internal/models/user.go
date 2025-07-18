package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Role      string    `gorm:"type:enum('Pelanggan','Admin')"`
	Nama      string
	Email     string    `gorm:"uniqueIndex"`
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Tokens    []PersonalAccessToken
}

func (User) TableName() string {
    return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}