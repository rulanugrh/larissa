package repository

import (
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/util"
	"gorm.io/gorm"
)

type KunjunganInterface interface {
	Create(req domain.Kunjungan) (*domain.Kunjungan, error)
	List(id uint) (*[]domain.Kunjungan, error)
}

type kunjungan struct {
	db *gorm.DB
}

func NewKunjungan(db *gorm.DB) KunjunganInterface {
	return &kunjungan{
		db: db,
	}
}

func(k *kunjungan) Create(req domain.Kunjungan) (*domain.Kunjungan, error) {	
	err := k.db.Create(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	return &req, nil
}

func(k *kunjungan) List(id uint) (*[]domain.Kunjungan, error) {
	var model []domain.Kunjungan
	find := k.db.Where("id = ?").Find(&model)
	if find.RowsAffected == 0 {
		return nil, util.NotFound()
	}

	if find.Error != nil {
		return nil, util.Errors(find.Error)
	}

	return &model, nil
}