package repository

import (
	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/util"
)

type PenyakitInterface interface {
	Create(req domain.Penyakit) (*domain.Penyakit, error)
	ListAll() (*[]domain.Penyakit, error)
	FindID(id uint) (*domain.Penyakit, error)
	Delete(id uint) error
}

type penyakit struct {
	client *config.Postgres
}

func NewPenyakit(client *config.Postgres) PenyakitInterface {
	return &penyakit{
		client: client,
	}
}

func (p *penyakit) Create(req domain.Penyakit) (*domain.Penyakit, error) {
	err := p.client.DB.Create(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	err = p.client.DB.Preload("Obat").Find(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	err = p.client.DB.Model(&req.Obat).Association("Penyakit").Append(&req)
	if err != nil {
		return nil, util.Errors(err)
	}

	return &req, nil
}

func (p *penyakit) ListAll() (*[]domain.Penyakit, error) {
	var model []domain.Penyakit
	err := p.client.DB.Preload("Obat").Find(model).Error

	if err != nil {
		return nil, util.Errors(err)
	}

	return &model, nil
}

func (p *penyakit) FindID(id uint) (*domain.Penyakit, error) {
	var model domain.Penyakit
	find := p.client.DB.Preload("Obat").Where("id = ?", id).Find(&model)

	if find.RowsAffected == 0 {
		return nil, util.NotFound()
	}

	if find.Error != nil {
		return nil, util.Errors(find.Error)
	}

	return &model, nil
}

func (p *penyakit) Delete(id uint) error {
	var model domain.Penyakit
	find := p.client.DB.Where("id = ?", id).Find(&model)

	if find.RowsAffected == 0 {
		return util.NotFound()
	}

	err := find.Delete(&model).Error
	if err != nil {
		return util.Errors(err)
	}

	return nil
}
