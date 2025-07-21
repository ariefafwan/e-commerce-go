package repositories

import (
	"e-commerce-go/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
    GetAll() ([]models.User, error)
    GetByID(id string) (*models.User, error)
    Create(user *models.User) error
    Update(user *models.User) error
}

type userRepo struct {
    db *gorm.DB
}

func NewuserRepository(db *gorm.DB) UserRepository {
    return &userRepo{db}
}

func (r *userRepo) GetAll() ([]models.User, error) {
    var user []models.User
    err := r.db.Find(&user).Error
    return user, err
}

func (r *userRepo) GetByID(id string) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepo) Create(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepo) Update(user *models.User) error {
    return r.db.Save(user).Error
}
