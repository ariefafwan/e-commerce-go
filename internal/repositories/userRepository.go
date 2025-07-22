package repositories

import (
	"e-commerce-go/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
    GetAll() ([]models.User, error)
    GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
    Create(user *models.User) error
    Update(user *models.User) error
    StoreAccessToken(token *models.PersonalAccessToken) error
    RevokeToken(token string, tokenType models.TipeToken) error
    RevokeAllUserTokens(IDUser string) error
    FindValidRefreshToken(token string) (*models.PersonalAccessToken, error)
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
    err := r.db.First(&user, "id = ?", id).Preload("DataPelanggan").Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Preload("DataPelanggan").Error
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

func (r *userRepo) StoreAccessToken(token *models.PersonalAccessToken) error {
	return r.db.Create(token).Error
}

func (r *userRepo) RevokeToken(token string, tokenType models.TipeToken) error {
	return r.db.Model(&models.PersonalAccessToken{}).
		Where("token = ? AND type = ? AND revoked_at IS NULL AND expired_at > ?", token, tokenType, time.Now()).
		Update("revoked_at", time.Now()).Error
}

func (r *userRepo) RevokeAllUserTokens(IDUser string) error {
	return r.db.Model(&models.PersonalAccessToken{}).
		Where("id_user = ? AND revoked_at IS NULL", IDUser).
		Update("revoked_at", time.Now()).Error
}

func (r *userRepo) FindValidRefreshToken(token string) (*models.PersonalAccessToken, error) {
	var pat models.PersonalAccessToken
	err := r.db.Where("token = ? AND type = ? AND revoked_at IS NULL AND expired_at > ?", token, models.RefreshToken, time.Now()).First(&pat).Error
	if err != nil {
		return nil, err
	}
	return &pat, nil
}
