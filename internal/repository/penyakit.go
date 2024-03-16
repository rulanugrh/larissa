package repository

import (
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/util"
	"gorm.io/gorm"
)

type PenyakitInterface interface {
	Create(req domain.Penyakit) (*domain.Penyakit, error)
	ListAll() (*[]domain.Penyakit, error)
	FindID(id uint) (*domain.Penyakit, error)
	Delete(id uint) error
}

type penyakit struct {
	db *gorm.DB
}

func NewPenyakit(db *gorm.DB) PenyakitInterface {
	return &penyakit{
		db: db,
	}
}

func(p *penyakit) Create(req domain.Penyakit) (*domain.Penyakit, error) {
	err := p.db.Create(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	err = p.db.Model(&req.Obat).Association("Penyakit").Append(&req)
	if err != nil {
		return nil, util.Errors(err)
	}

	return &req, nil
}

func(p *penyakit) ListAll() (*[]domain.Penyakit, error) {
	var model []domain.Penyakit
	err := p.db.Find(model).Error

	if err != nil {
		return nil, util.Errors(err)
	}

	return &model, nil
}

func(p *penyakit) FindID(id uint) (*domain.Penyakit, error) {
	var model domain.Penyakit
	find := p.db.Where("id = ?", id).Find(&model)

	if find.RowsAffected == 0 {
		return nil, util.NotFound()
	}

	if find.Error != nil {
		return nil, util.Errors(find.Error)
	}

	return &model, nil
}

func(p *penyakit) Delete(id uint) error {
	var model domain.Penyakit
	find := p.db.Where("id = ?", id).Find(&model)

	if find.RowsAffected == 0 {
		return util.NotFound()
	}

	err := find.Delete(&model).Error
	if err != nil {
		return util.Errors(err)
	}

	return nil
}