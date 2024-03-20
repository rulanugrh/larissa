package repository

import (
	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/util"
)

type KunjunganInterface interface {
	Create(req domain.Kunjungan) (*domain.Kunjungan, error)
	List(id uint) (*[]domain.Kunjungan, error)
}

type kunjungan struct {
	client *config.Postgres
}

func NewKunjungan(client *config.Postgres) KunjunganInterface {
	return &kunjungan{
		client: client,
	}
}

func(k *kunjungan) Create(req domain.Kunjungan) (*domain.Kunjungan, error) {
	err := k.client.DB.Create(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	err = k.client.DB.Preload("Penyakit").Preload("User").Find(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	return &req, nil
}

func(k *kunjungan) List(id uint) (*[]domain.Kunjungan, error) {
	var model []domain.Kunjungan
	find := k.client.DB.Where("id = ?").Find(&model)
	if find.RowsAffected == 0 {
		return nil, util.NotFound()
	}

	if find.Error != nil {
		return nil, util.Errors(find.Error)
	}

	err := k.client.DB.Preload("Penyakit").Preload("User").Find(&model).Error
	if err != nil {
		return nil, util.Errors(err)
	}


	return &model, nil
}
