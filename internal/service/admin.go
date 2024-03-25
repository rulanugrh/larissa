package service

import (
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/entity/web"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/util"
	"github.com/rulanugrh/larissa/pkg"
)

type AdminInterface interface {
	CreatePenyakit(req domain.Penyakit) (*web.PenyakitCreated, error)
	CreateObat(req domain.Obat) (*web.ObatCreated, error)
	UpdateObat(id uint, req domain.Obat) (*web.ObatUpdated, error)
	DeleteObat(id uint) error
	DeletePenyakit(id uint) error
	Reported() (*[]web.Reported, error)
	ListAllUser() (*[]web.User, error)
}

type admin struct {
	obat     repository.ObatInterface
	penyakit repository.PenyakitInterface
	reported repository.ReportedInterface
	user     repository.UserInterface
	log pkg.ILogrust
}

func NewAdmin(obat repository.ObatInterface, penyakit repository.PenyakitInterface, reported repository.ReportedInterface, user repository.UserInterface, log pkg.ILogrust) AdminInterface {
	return &admin{
		obat:     obat,
		penyakit: penyakit,
		reported: reported,
		user:     user,
		log: log,
	}
}

func (a *admin) CreatePenyakit(req domain.Penyakit) (*web.PenyakitCreated, error) {
	data, err := a.penyakit.Create(req)
	if err != nil {
		a.log.StartLogger("penyakit_service", "create").Error(err.Error())
		return nil, util.Errors(err)
	}

	response := web.PenyakitCreated{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
	}
	a.log.StartLogger("penyakit_service", "create").Info("success operation")
	return &response, nil
}

func (a *admin) CreateObat(req domain.Obat) (*web.ObatCreated, error) {
	data, err := a.obat.Create(req)
	if err != nil {
		a.log.StartLogger("obat_service", "create").Error(err.Error())
		return nil, util.Errors(err)
	}

	response := web.ObatCreated{
		ID:           data.ID,
		Price:        data.Price,
		QtyAvailable: data.QtyAvailable,
		Description:  data.Description,
		Composition:  data.Composition,
		Name:         data.Name,
	}

	a.log.StartLogger("obat_service", "create").Info("success operation")
	return &response, nil
}

func (a *admin) UpdateObat(id uint, req domain.Obat) (*web.ObatUpdated, error) {
	data, err := a.obat.Update(id, req)
	if err != nil {
		a.log.StartLogger("obat_service", "update").Error(err.Error())
		return nil, util.Errors(err)
	}

	response := web.ObatUpdated{
		ID:           data.ID,
		Name:         data.Name,
		Description:  data.Description,
		Price:        data.Price,
		QtyAvailable: data.QtyAvailable,
	}

	a.log.StartLogger("obat_service", "update").Info("success update")
	return &response, nil
}

func (a *admin) DeleteObat(id uint) error {
	err := a.obat.Delete(id)
	if err != nil {
		a.log.StartLogger("obat_service", "delete").Error(err.Error())
		return util.Errors(err)
	}

	a.log.StartLogger("obat_service", "delete").Info("success delete")
	return nil
}

func (a *admin) DeletePenyakit(id uint) error {
	err := a.penyakit.Delete(id)
	if err != nil {
		a.log.StartLogger("penyakit_service", "delete").Error(err.Error())
		return util.Errors(err)
	}

	a.log.StartLogger("penyakit_service", "delete").Info("success delete")
	return nil
}

func (a *admin) Reported() (*[]web.Reported, error) {
	data, err := a.reported.List()
	if err != nil {
		a.log.StartLogger("reported_service", "findAll").Error(err.Error())
		return nil, util.Errors(err)
	}

	var response []web.Reported
	for _, v := range *data {
		result := web.Reported{
			ID:         v.ID,
			Pengunjung: v.Pengunjung,
			Age:        v.Age,
			Address:    v.Address,
			Category:   v.Category,
		}

		response = append(response, result)
	}

	a.log.StartLogger("reported_service", "findAll").Info("success get all data")
	return &response, nil
}

func (a *admin) ListAllUser() (*[]web.User, error) {
	data, err := a.user.ListAll()
	if err != nil {
		a.log.StartLogger("user_service", "findAll").Error(err.Error())
		return nil, util.Errors(err)
	}

	var response []web.User
	for _, v := range *data {
		result := web.User{
			ID:      v.ID,
			FName:   v.FName,
			LName:   v.LName,
			Email:   v.Email,
			Age:     v.Age,
			Address: v.Address,
			TTL:     v.TTL,
		}

		response = append(response, result)
	}

	a.log.StartLogger("user_service", "findAll").Info("success get all user")
	return &response, nil
}
