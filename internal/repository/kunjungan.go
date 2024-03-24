package repository

import (
	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/util"
)

type KunjunganInterface interface {
	Create(req domain.Kunjungan) (*domain.Kunjungan, error)
	List(id uint) (*[]domain.Kunjungan, error)
	GotPrice(req domain.Kunjungan) (*[]domain.Obat, error)
}

type kunjungan struct {
	client *config.Postgres
}

func NewKunjungan(client *config.Postgres) KunjunganInterface {
	return &kunjungan{
		client: client,
	}
}

func (k *kunjungan) Create(req domain.Kunjungan) (*domain.Kunjungan, error) {

	err := k.client.DB.Create(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	err = k.client.DB.Preload("Penyakit").Preload("User").Preload("Category").Find(&req).Error
	if err != nil {
		return nil, util.Errors(err)
	}

	return &req, nil
}

func (k *kunjungan) GotPrice(req domain.Kunjungan) (*[]domain.Obat, error) {
	var penyakit []domain.Penyakit

	find := k.client.DB.Find(&penyakit, req.PenyakitID)
	if find.RowsAffected == 0 {
		return nil, util.Errors(find.Error)
	}

	var obat []domain.Obat
	for _, v := range penyakit {
		find := k.client.DB.Find(&obat, v.ObatID)
		if find.RowsAffected == 0 {
			return nil, util.Errors(find.Error)
		}

		var ob domain.Obat
		for _, o := range obat {
			ob.QtyAvailable = o.QtyAvailable - 1
			ob.QtyReserved += 1
			ob.QtyOn = o.QtyAvailable + o.QtyReserved

		}

		err := k.client.DB.Table("obats").Where("id IN ?", v.ObatID).Updates(domain.Obat{QtyAvailable: ob.QtyAvailable, QtyOn: ob.QtyOn, QtyReserved: ob.QtyReserved}).Error
		if err != nil {
			return nil, util.Errors(err)
		}

	}

	return &obat, nil
}

func (k *kunjungan) List(id uint) (*[]domain.Kunjungan, error) {
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
