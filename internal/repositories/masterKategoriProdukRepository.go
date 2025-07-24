package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"

	"github.com/gosimple/slug"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MasterKategoriProdukRepository interface {
	GetAll(q QueryParams) ([]dto.MasterKategoriProdukResponse, int64, error)
	GetByID(id string) (*dto.MasterKategoriProdukResponse, error)
	Create(kategori *models.MasterKategoriProduk) error
	Update(kategori *models.MasterKategoriProduk) error
	Delete(id string) error
}

type masterKategoriProdukRepo struct {
	db *gorm.DB
}

func NewMasterKategoriProdukRepository(db *gorm.DB) MasterKategoriProdukRepository {
	return &masterKategoriProdukRepo{db}
}

func (r *masterKategoriProdukRepo) GetAll(q QueryParams) ([]dto.MasterKategoriProdukResponse, int64, error) {
	var kategori []models.MasterKategoriProduk
	var total int64

	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := r.db.Model(&models.MasterKategoriProduk{}).Preload("DataParent")

	if q.Search != "" {
		query = query.Where("nama LIKE ?", "%"+q.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at " + q.Sort).
		Offset(offset).
		Limit(q.Limit).
		Find(&kategori).Error
	if err != nil {
		return nil, 0, err
	}

	var kategoriResponse []dto.MasterKategoriProdukResponse
	if err := copier.Copy(&kategoriResponse, &kategori); err != nil {
		return nil, 0, err
	}

	return kategoriResponse, total, nil
}

func (r *masterKategoriProdukRepo) GetByID(id string) (*dto.MasterKategoriProdukResponse, error) {
	var kategori models.MasterKategoriProduk
	err := r.db.Preload("DataParent").First(&kategori, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	var kategoriResponse dto.MasterKategoriProdukResponse
	if err := copier.Copy(&kategoriResponse, &kategori); err != nil {
		return nil, err
	}
	return &kategoriResponse, nil
}

func (r *masterKategoriProdukRepo) Create(kategori *models.MasterKategoriProduk) error {
	return r.db.Create(kategori).Error
}

func (r *masterKategoriProdukRepo) Update(kategori *models.MasterKategoriProduk) error {
	kategori.Slug = slug.Make(kategori.Nama)
	return r.db.Model(&models.MasterKategoriProduk{}).Where("id = ?", kategori.ID).Updates(kategori).Error
}

func (r *masterKategoriProdukRepo) Delete(id string) error {
	return r.db.Delete(&models.MasterKategoriProduk{}, "id = ?", id).Error
}



