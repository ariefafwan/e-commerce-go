package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MasterAlamatPelangganRepository interface {
	GetAll(q QueryParams) ([]dto.MasterAlamatPelangganResponse, int64,error)
	GetAllByPelanggan(id string, q QueryParams) ([]dto.MasterAlamatPelangganResponse, int64, error)
	GetByID(id string) (*dto.MasterAlamatPelangganResponse, error)
	SetAlamatUtama(id string) error
	Create(alamat *models.MasterAlamatPelanggan) error
	Update(alamat *models.MasterAlamatPelanggan) error
	Delete(id string) error
}

type masterAlamatPelangganRepo struct {
	db *gorm.DB
}

func NewMasterAlamatPelangganRepository(db *gorm.DB) MasterAlamatPelangganRepository {
	return &masterAlamatPelangganRepo{db}
}

func (r *masterAlamatPelangganRepo) GetAll(q QueryParams) ([]dto.MasterAlamatPelangganResponse, int64,error) {
	var alamat []models.MasterAlamatPelanggan
	var total int64

	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := r.db.Model(&models.MasterAlamatPelanggan{}).Preload("DataKecamatan.DataKota.DataProvinsi").Preload("DataPelanggan")

	if q.Search != "" {
		query = query.Where("label LIKE ?", "%"+q.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at " + q.Sort).
		Offset(offset).
		Limit(q.Limit).
		Find(&alamat).Error
	if err != nil {
		return nil, 0, err
	}
	
	var alamatResponse []dto.MasterAlamatPelangganResponse
	if err := copier.Copy(&alamatResponse, &alamat); err != nil {
		return nil, 0, err
	}
	return alamatResponse, total, err
}

func (r *masterAlamatPelangganRepo) GetAllByPelanggan(id string, q QueryParams) ([]dto.MasterAlamatPelangganResponse, int64, error) {
	var alamat []models.MasterAlamatPelanggan
	var total int64

	offset := (q.Page - 1) * q.Limit

	if q.Sort != "asc" && q.Sort != "desc" {
		q.Sort = "asc"
	}

	query := r.db.Model(&models.MasterAlamatPelanggan{})

	if q.Search != "" {
		query = query.Where("id_pelanggan = ? AND label COLLATE SQL_Latin1_General_CP1_CI_AS LIKE ?", id, "%"+q.Search+"%").Preload("DataKecamatan.DataKota.DataProvinsi").Preload("DataPelanggan")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at " + q.Sort).
		Offset(offset).
		Limit(q.Limit).
		Find(&alamat).Error
	if err != nil {
		return nil, 0, err
	}

	var alamatResponse []dto.MasterAlamatPelangganResponse
	if err := copier.Copy(&alamatResponse, &alamat); err != nil {
		return nil, 0, err
	}
	return alamatResponse, total, err
}

func (r *masterAlamatPelangganRepo) GetByID(id string) (*dto.MasterAlamatPelangganResponse, error) {
	var alamat models.MasterAlamatPelanggan
	err := r.db.First(&alamat, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	var alamatResponse dto.MasterAlamatPelangganResponse
	if err := copier.Copy(&alamatResponse, &alamat); err != nil {
		return nil, err
	}

	return &alamatResponse, nil
}

func (r *masterAlamatPelangganRepo) Create(alamat *models.MasterAlamatPelanggan) error {
	var count int64
	err := r.db.
			Model(&models.MasterAlamatPelanggan{}).
			Where("id_pelanggan = ?", alamat.IDPelanggan).
			Count(&count).Error
	if err != nil {
		return errors.New("failed to get data")
	}

	if count < 1 {
		alamat.IsDefault = true
	} else {
		alamat.IsDefault = false
	}

	return r.db.Create(alamat).Error
}

func (r *masterAlamatPelangganRepo) Update(alamat *models.MasterAlamatPelanggan) error {
	return r.db.Model(&models.MasterAlamatPelanggan{}).Where("id = ?", alamat.ID).Updates(alamat).Error
}

var ErrorAlamatPelanggan = errors.New("pelanggan harus memiliki setidaknya 1 alamat, silahkan update jika ingin merubah data ini")

func (r *masterAlamatPelangganRepo) Delete(id string) error {
	var data models.MasterAlamatPelanggan
	var count int64

	if err := r.db.First(&data, "id = ?", id).Error; err != nil {
		return err
	}
	
	if err := r.db.
		Model(&models.MasterAlamatPelanggan{}).
		Where("id_pelanggan = ?", data.IDPelanggan).
		Count(&count).Error; err != nil {
		return err
	}

	if count < 2 {
		return ErrorAlamatPelanggan
	} else if data.IsDefault {
		var dataTerakhir models.MasterAlamatPelanggan
		if err := r.db.
			Last(&dataTerakhir, "id != ? AND id_pelanggan = ?", data.ID, data.IDPelanggan).Error; err != nil {
			return errors.New("failed to update default alamat 1")
		}

		if err := r.db.
			Model(&models.MasterAlamatPelanggan{}).
			Where("id = ?", dataTerakhir.ID).
			Update("is_default", true).Error; err != nil {
			return errors.New("failed to update default alamat")
		}
	}

	return r.db.Delete(&data, "id = ?", id).Error
}

func (r *masterAlamatPelangganRepo) SetAlamatUtama(id string) error {
	data, err := r.GetByID(id)
	if err != nil {
		return err
	}

	var count int64
	r.db.First(&models.MasterAlamatPelanggan{}, "id != ? AND id_pelanggan = ? AND is_default = ?", id, data.IDPelanggan, true).
		Count(&count)
		
	if count > 0 {
		if err := r.db.
			Model(&models.MasterAlamatPelanggan{}).
			Where("id != ? AND id_pelanggan = ? AND is_default = ?", id, data.IDPelanggan, true).
			Update("is_default", false).Error; err != nil {
			return errors.New("failed to update default alamat")
		}
	}

	return r.db.
		Model(&models.MasterAlamatPelanggan{}).
		Where("id = ?", id).
		Update("is_default", true).Error
}