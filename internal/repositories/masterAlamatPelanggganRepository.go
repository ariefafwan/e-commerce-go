package repositories

import (
	"e-commerce-go/internal/models"
	"errors"

	"gorm.io/gorm"
)

type MasterAlamatPelangganRepository interface {
	GetAll() ([]models.MasterAlamatPelanggan, error)
	GetByID(id string) (*models.MasterAlamatPelanggan, error)
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

func (r *masterAlamatPelangganRepo) GetAll() ([]models.MasterAlamatPelanggan, error) {
	var alamat []models.MasterAlamatPelanggan
	err := r.db.Find(&alamat).Preload("DataPelanggan").Error
	return alamat, err
}

func (r *masterAlamatPelangganRepo) GetByID (id string) (*models.MasterAlamatPelanggan, error) {
	var alamat models.MasterAlamatPelanggan
	err := r.db.First(&alamat, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &alamat, nil
}

func (r *masterAlamatPelangganRepo) Create(alamat *models.MasterAlamatPelanggan) error {
	return r.db.Create(alamat).Error
}

func (r *masterAlamatPelangganRepo) Update(alamat *models.MasterAlamatPelanggan) error {
	return r.db.Save(alamat).Error
}

var ErrorAlamatPelanggan = errors.New("pelanggan harus memiliki setidaknya 1 alamat, silahkan update jika ingin merubah data ini")

func (r *masterAlamatPelangganRepo) Delete(id string) error {
	var data models.MasterAlamatPelanggan
	var count int64
	err := r.db.First(&data, "id = ?", id).Error
	if err != nil {
		return err
	}

	r.db.Model(&models.MasterAlamatPelanggan{}).Where("id_pelanggan = ?", data.IDPelanggan).Count(&count)
	if count < 2  {
		return ErrorAlamatPelanggan
	}

	return r.db.Delete(&models.MasterAlamatPelanggan{}, "id = ?", id).Error
}