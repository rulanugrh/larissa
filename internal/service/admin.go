package service

import (
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/entity/web"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/util"
)

type AdminInterface interface {
	CreatePenyakit(req domain.Penyakit) (*web.PenyakitCreated, error)
	CreateObat(req domain.Obat) (*web.ObatCreated, error)
	UpdateObat(id uint, req domain.Obat) (*web.ObatUpdated, error)
	DeleteObat(id uint) error
	DeletePenyakit(id uint) error
	Reported() (*[]web.Reported, error)
}

type admin struct {
	obat repository.ObatInterface
	penyakit repository.PenyakitInterface
	reported repository.ReportedInterface
}

func NewAdmin(obat repository.ObatInterface, penyakit repository.PenyakitInterface, reported repository.ReportedInterface) AdminInterface {
	return &admin{
		obat: obat,
		penyakit: penyakit,
		reported: reported,
	}
}

func(a *admin) CreatePenyakit(req domain.Penyakit) (*web.PenyakitCreated, error) {
	data, err := a.penyakit.Create(req)
	if err != nil {
		return nil, util.Errors(err)
	}

	response := web.PenyakitCreated{
		ID: data.ID,
		Name: data.Name,
		Description: data.Description,
	}

	return &response, nil
}

func(a *admin) CreateObat(req domain.Obat) (*web.ObatCreated, error) {
	data, err := a.obat.Create(req)
	if err != nil {
		return nil, util.Errors(err)
	}

	response := web.ObatCreated {
		ID: data.ID,
		Price: data.Price,
		QtyAvailable: data.QtyAvailable,
		Description: data.Description,
		Composition: data.Composition,
		Name: data.Name,
	}

	return &response, nil
}

func(a *admin) UpdateObat(id uint, req domain.Obat) (*web.ObatUpdated, error) {
	data, err := a.obat.Update(id, req)
	if err != nil {
		return nil, util.Errors(err)
	}

	response := web.ObatUpdated {
		ID: data.ID,
		Name: data.Name,
		Description: data.Description,
		Price: data.Price,
		QtyAvailable: data.QtyAvailable,
	}

	return &response, nil
}

func(a *admin) DeleteObat(id uint) error {
	err := a.obat.Delete(id)
	if err != nil {
		return util.Errors(err)
	}

	return nil
}

func(a *admin) DeletePenyakit(id uint) error {
	err := a.penyakit.Delete(id)
	if err != nil {
		return util.Errors(err)
	}

	return nil
}

func(a *admin) Reported() (*[]web.Reported, error) {
	data, err := a.reported.List()
	if err != nil {
		return nil, util.Errors(err)
	}

	var response []web.Reported
	for _, v := range *data {
		result := web.Reported {
			ID: v.ID,
			Pengunjung: v.Pengunjung,
			Age: v.Age,
			Address: v.Address,
		}

		response = append(response, result)
	}

	return &response, nil
}