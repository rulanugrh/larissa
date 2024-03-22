package repository

import (
	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/util"
)

type ObatInterface interface {
	Create(req domain.Obat) (*domain.Obat, error)
	FindID(id uint) (*domain.Obat, error)
	FindAll() (*[]domain.Obat, error)
	Update(id uint, req domain.Obat) (*domain.Obat, error)
	Delete(id uint) error
}

type obat struct {
	client *config.Postgres
}

func NewObat(client *config.Postgres) ObatInterface {
	return &obat{
		client: client,
	}
}

func (o *obat) Create(req domain.Obat) (*domain.Obat, error) {
	find := o.client.DB.Where("name = ?", req.Name).Find(&req)
	if find.RowsAffected != 0 {
		return nil, util.DataHasBeenUsed()
	}

	err := find.Create(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	err = o.client.DB.Preload("Penyakit").Find(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	return &req, nil
}

func (o *obat) FindID(id uint) (*domain.Obat, error) {
	var model domain.Obat
	err := o.client.DB.Where("id = ?", id).Preload("Penyakit").Find(&model).Error

	if err != nil {
		return nil, util.Errors(err)
	}

	return &model, nil
}

func (o *obat) FindAll() (*[]domain.Obat, error) {
	var finds []domain.Obat

	err := o.client.DB.Preload("Penyakit").Find(&finds).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	return &finds, nil
}

func (o *obat) Update(id uint, req domain.Obat) (*domain.Obat, error) {
	var model domain.Obat
	find := o.client.DB.Where("id = ?", id).Find(&model)
	if find.RowsAffected == 0 {
		return nil, util.NotFound()
	}

	err := find.Model(&req).Updates(&model).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	return &model, nil
}

func (o *obat) Delete(id uint) error {
	var model domain.Obat
	find := o.client.DB.Where("id  = ?", id).Find(&model)
	if find.RowsAffected == 0 {
		return util.NotFound()
	}

	err := find.Delete(&model).Error
	if err != nil {
		return util.Errors(err)
	}

	return nil
}
