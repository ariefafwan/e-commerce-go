package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MasterPelangganRepository interface {
	GetAll(q QueryParams) ([]dto.MasterPelangganResponse, int64, error)
	GetByID(id string) (*dto.MasterPelangganResponse, error)
	Create(pelanggan *models.MasterPelanggan) error
	Update(pelanggan *models.MasterPelanggan) error
	Delete(id string) error
}

type masterPelangganRepo struct {
	db *gorm.DB
}

func NewMasterPelangganRepository(db *gorm.DB) MasterPelangganRepository {
	return &masterPelangganRepo{db}
}

func (r *masterPelangganRepo) GetAll(q QueryParams) ([]dto.MasterPelangganResponse, int64, error) {
	var pelanggan []models.MasterPelanggan
	var total int64

	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := r.db.Model(&models.MasterPelanggan{}).Preload("DataUser").Preload("DataAlamat")

	if q.Search != "" {
		query = query.Where("nama_lengkap LIKE ?", "%"+q.Search+"%").
					Or("nama_panggilan LIKE ?", "%"+q.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at " + q.Sort).
		Offset(offset).
		Limit(q.Limit).
		Find(&pelanggan).Error
	if err != nil {
		return nil, 0, err
	}

	var pelangganResponse []dto.MasterPelangganResponse
	if err := copier.Copy(&pelangganResponse, &pelanggan); err != nil {
		return nil, 0, err
	}
	return pelangganResponse, total, err
}

func (r *masterPelangganRepo) GetByID(id string) (*dto.MasterPelangganResponse, error) {
	var pelanggan models.MasterPelanggan
	err := r.db.Preload("DataUser").Preload("DataAlamat").First(&pelanggan, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	var pelangganResponse dto.MasterPelangganResponse
	if err := copier.Copy(&pelangganResponse, &pelanggan); err != nil {
		return nil, err
	}
	return &pelangganResponse, nil
}

func (r *masterPelangganRepo) Create(pelanggan *models.MasterPelanggan) error {
	return r.db.Create(pelanggan).Error
}

func (r *masterPelangganRepo) Update(pelanggan *models.MasterPelanggan) error {
	return r.db.Model(&models.MasterPelanggan{}).Where("id = ?", pelanggan.ID).Updates(pelanggan).Error
}

func (r *masterPelangganRepo) Delete(id string) error {
	return r.db.Delete(&models.MasterPelanggan{}, "id = ?", id).Error
}