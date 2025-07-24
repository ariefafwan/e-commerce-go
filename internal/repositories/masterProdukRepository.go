package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MasterProdukRepository interface {
	GetAll(q QueryParams) ([]dto.MasterProdukResponse, int64, error)
	GetAllByKategori(slug string, q QueryParams) ([]dto.MasterProdukResponse, int64,error)
	GetByID(id string) (dto.MasterProdukResponse, error)
	GetBySlug(slug string) (dto.MasterProdukResponse, error)
	UpdateStatus(id string, status string) error
	Create(produk *models.MasterProduk, kategoris []string) error
	Update(produk *models.MasterProduk, kategoris []string) error
	Delete(id string) error
}

type masterProdukRepo struct {
	db *gorm.DB
}

func NewMasterProdukRepository(db *gorm.DB) MasterProdukRepository {
	return &masterProdukRepo{db}
}

func (m *masterProdukRepo) GetAll(q QueryParams) ([]dto.MasterProdukResponse, int64, error) {
	var produk []models.MasterProduk
	var total int64

	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := m.db.Model(&models.MasterProduk{}).Preload("DataKategori").Preload("DataGaleri").Preload("DataVariant")

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
		Find(&produk).Error
	if err != nil {
		return nil, 0, err
	}

	var produkResponse []dto.MasterProdukResponse
	if err := copier.Copy(&produkResponse, &produk); err != nil {
		return nil, 0, err
	}
	return produkResponse, total, err
}

func (m *masterProdukRepo) GetAllByKategori(slug string, q QueryParams) ([]dto.MasterProdukResponse, int64,error) {
	var produk []models.MasterProduk
	var kategori models.MasterKategoriProduk
	err := m.db.Model(&models.MasterKategoriProduk{}).Where("slug = ?", slug).First(&kategori).Error
	if err != nil {
		return nil, 0, err
	}

	var kategoriIDs []uuid.UUID
	kategoriIDs = append(kategoriIDs, kategori.ID)

	var childIDs []uuid.UUID
	if err := m.db.Model(&models.MasterKategoriProduk{}).
		Where("id_parent = ?", kategori.ID).
		Pluck("id", &childIDs).Error; err != nil {
		return nil, 0, err
	}

	kategoriIDs = append(kategoriIDs, childIDs...)

	var total int64
	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := m.db.Model(&models.MasterProduk{}).Preload("DataKategori").
		Preload("DataGaleri").
		Preload("DataVariant").
		Where("id IN (?)",
			m.db.Table("master_produk_kategori_produk").
				Select("id_produk").
				Where("id_kategori IN ?", kategoriIDs),
		)

	if q.Search != "" {
		query = query.Where("nama LIKE ?", "%"+q.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = query.
		Order("created_at " + q.Sort).
		Offset(offset).
		Limit(q.Limit).
		Find(&produk).Error
	if err != nil {
		return nil, 0, err
	}

	var produkResponse []dto.MasterProdukResponse
	err = copier.Copy(&produkResponse, &produk)
	if err != nil {
		return nil, 0, err
	}
	return produkResponse, total, err
}

func (m *masterProdukRepo) GetByID(id string) (dto.MasterProdukResponse, error) {
	var produk models.MasterProduk
	err := m.db.Preload("DataKategori").Preload("DataGaleri").Preload("DataVariant").First(&produk, "id = ?", id).Error
	if err != nil {
		return dto.MasterProdukResponse{}, err
	}

	var produkResponse dto.MasterProdukResponse
	if err := copier.Copy(&produkResponse, &produk); err != nil {
		return dto.MasterProdukResponse{}, err
	}
	return produkResponse, nil
}

func (m *masterProdukRepo) GetBySlug(slug string) (dto.MasterProdukResponse, error) {
	var produk models.MasterProduk
	err := m.db.Preload("DataKategori").Preload("DataGaleri").Preload("DataVariant").First(&produk, "slug = ?", slug).Error
	if err != nil {
		return dto.MasterProdukResponse{}, err
	}

	var produkResponse dto.MasterProdukResponse
	if err := copier.Copy(&produkResponse, &produk); err != nil {
		return dto.MasterProdukResponse{}, err
	}
	return produkResponse, nil
}

func (m *masterProdukRepo) Create(produk *models.MasterProduk, kategoris []string) error {
	var kategori []models.MasterKategoriProduk
	err := m.db.Where("id IN ?", kategoris).Find(&kategori).Error
	if err != nil {
		return err
	}

	if len(kategori) != len(kategoris) {
		return errors.New("kategori tidak ditemukan")
	}

	produk.DataKategori = kategori	
	return m.db.Create(&produk).Error
}

func (m *masterProdukRepo) Update(produk *models.MasterProduk, kategoris []string) error {
	var data models.MasterProduk
	err := m.db.First(&data, "id = ?", produk.ID).Error
	if err != nil {
		return err
	}
	if data.Nama != produk.Nama {
		baseSlug := slug.Make(produk.Nama)
		slug := baseSlug
		counter := 1
		for {
			var count int64
			err := m.db.Model(&models.MasterProduk{}).Where("slug = ?", data.Slug).Count(&count).Error
			if err != nil {
				return err
			}
			if count == 0 {
				break
			}
			slug = fmt.Sprintf("%s-%d", baseSlug, counter)
			counter++
		}
		produk.Slug = slug
	}
	var kategori []models.MasterKategoriProduk
	err = m.db.Where("id IN ?", kategoris).Find(&kategori).Error
	if err != nil {
		return err
	}
	produk.DataKategori = kategori
	return m.db.Model(&models.MasterProduk{}).Where("id = ?", produk.ID).Updates(&produk).Error
}


func (m *masterProdukRepo) UpdateStatus(id string, status string) error {
	return m.db.Model(&models.MasterProduk{}).Where("id = ?", id).Update("status", status).Error
}

func (m *masterProdukRepo) Delete(id string) error {
	return m.db.Delete(&models.MasterProduk{}, "id = ?", id).Error
}