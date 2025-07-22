package repositories

import (
	"e-commerce-go/internal/dto"
	"e-commerce-go/internal/models"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MasterAlamatPelangganRepository interface {
	GetAll() ([]dto.MasterAlamatPelangganResponse, error)
	GetAllByPelanggan(id string) ([]dto.MasterAlamatPelangganResponse, error)
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

func (r *masterAlamatPelangganRepo) GetAll() ([]dto.MasterAlamatPelangganResponse, error) {
	var alamat []models.MasterAlamatPelanggan
	err := r.db.Find(&alamat).Preload("DataPelanggan").Error
	if err != nil {
		return []dto.MasterAlamatPelangganResponse{}, err
	}
	
	var alamatResponse []dto.MasterAlamatPelangganResponse
	if err := copier.Copy(&alamatResponse, &alamat); err != nil {
		return []dto.MasterAlamatPelangganResponse{}, err
	}
	return alamatResponse, err
}

func (r *masterAlamatPelangganRepo) GetAllByPelanggan(id string) ([]dto.MasterAlamatPelangganResponse, error) {
	var alamat []models.MasterAlamatPelanggan
	err := r.db.Find(&alamat, "id_pelanggan = ?", id).Error
	if err != nil {
		return []dto.MasterAlamatPelangganResponse{}, err
	}

	var alamatResponse []dto.MasterAlamatPelangganResponse
	if err := copier.Copy(&alamatResponse, &alamat); err != nil {
		return []dto.MasterAlamatPelangganResponse{}, err
	}
	return alamatResponse, err
}

func (r *masterAlamatPelangganRepo) GetByID (id string) (*dto.MasterAlamatPelangganResponse, error) {
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
	return r.db.Save(alamat).Error
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
	}

	return r.db.Delete(&data, "id = ?", id).Error
}

func (r *masterAlamatPelangganRepo) SetAlamatUtama(id string) error {
	if err := r.db.
		Model(&models.MasterAlamatPelanggan{}).
		Where("id != ? AND is_default = ?", id, true).
		Update("is_default", false).Error; err != nil {
		return err
	}

	return r.db.
		Model(&models.MasterAlamatPelanggan{}).
		Where("id = ?", id).
		Update("is_default", true).Error
}