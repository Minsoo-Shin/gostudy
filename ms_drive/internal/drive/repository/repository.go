package repository

import (
	"github.com/Minsoo-Shin/ms_drive/entity"
	"gorm.io/gorm"
)

type DriverRepository interface {
	Create(entity *entity.File) (entity.File, error)
	Get(id int64) (entity.File, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) DriverRepository {
	repository := repository{
		db: db,
	}
	return &repository
}

func (r repository) Create(e *entity.File) (entity.File, error) {
	var ret entity.File

	if err := r.db.Create(e).Scan(&ret).Error; err != nil {
		return entity.File{}, nil
	}

	return ret, nil
}

func (r repository) Get(id int64) (entity.File, error) {
	var ret entity.File

	if err := r.db.
		Where(entity.File{ID: id}).
		First(&ret).Error; err != nil {
		return ret, err
	}
	return ret, nil
}
