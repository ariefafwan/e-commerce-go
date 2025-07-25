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
	GetAllByKategori(slug string, q QueryParams) ([]dto.MasterProdukResponse, int64, error)
	GetProdukNonAktif(q QueryParams) ([]dto.MasterProdukResponse, int64, error)
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

	query := m.db.Model(&models.MasterProduk{}).
					Preload("DataKategori").
					Preload("DataGaleri").
					Preload("DataVariant").
					Where("status = ?", "Aktif").
					Where("EXISTS (?)",
							m.db.Table("master_produk_variant").
								Select("1").
								Where("master_produk_variant.id_produk = master_produk.id AND master_produk_variant.stok > 0"),
						)

	if q.Search != "" {
		query = query.Where("nama ILIKE ?", "%"+q.Search+"%")
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
		Where("status = ?", "Aktif").
		Where("id IN (?)",
			m.db.Table("master_produk_kategori_produk").
				Select("id_produk").
				Where("id_kategori IN ?", kategoriIDs),
		).
		Where("EXISTS (?)",
			m.db.Table("master_produk_variant").
				Select("1").
				Where("master_produk_variant.id_produk = master_produk.id AND master_produk_variant.stok > 0"),
		)

	if q.Search != "" {
		query = query.Where("nama ILIKE ?", "%"+q.Search+"%")
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

func (m *masterProdukRepo) GetProdukNonAktif(q QueryParams) ([]dto.MasterProdukResponse, int64, error) {
	var produk []models.MasterProduk
	var total int64

	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := m.db.Model(&models.MasterProduk{}).Preload("DataKategori").
		Preload("DataGaleri").
		Preload("DataVariant").
		Where("status = ?", "Tidak Aktif")

	if q.Search != "" {
		query = query.Where("nama ILIKE ?", "%"+q.Search+"%")
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
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var kategori []models.MasterKategoriProduk
	err := tx.Where("id IN ?", kategoris).Find(&kategori).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(kategori) != len(kategoris) {
		tx.Rollback()
		return errors.New("kategori tidak ditemukan")
	}

	if err := tx.Create(produk).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(produk).Association("DataKategori").Append(&kategori); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (m *masterProdukRepo) Update(produk *models.MasterProduk, kategoris []string) error {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var data models.MasterProduk
	err := tx.First(&data, "id = ?", produk.ID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if data.Nama != produk.Nama {
		baseSlug := slug.Make(produk.Nama)
		slug := baseSlug
		counter := 1
		for {
			var count int64
			err := tx.Model(&models.MasterProduk{}).Where("slug = ?", slug).Count(&count).Error
			if err != nil {
				tx.Rollback()
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
	err = tx.Where("id IN ?", kategoris).Find(&kategori).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(kategori) != len(kategoris) {
		tx.Rollback()
		return errors.New("kategori tidak ditemukan")
	}

	err = tx.Model(&models.MasterProduk{}).Where("id = ?", produk.ID).Updates(produk).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&produk).Association("DataKategori").Replace(kategori); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}


func (m *masterProdukRepo) UpdateStatus(id string, status string) error {
	return m.db.Model(&models.MasterProduk{}).Where("id = ?", id).Update("status", status).Error
}

func (m *masterProdukRepo) Delete(id string) error {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var produk models.MasterProduk
	if err := tx.Preload("DataKategori").First(&produk, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&produk).Association("DataKategori").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&produk).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}