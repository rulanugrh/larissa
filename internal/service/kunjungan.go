package service

import (
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/entity/web"
	"github.com/rulanugrh/larissa/internal/middleware"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/util"
	"github.com/rulanugrh/larissa/pkg"
)

type KunjunganInterface interface {
	Create(req domain.Kunjungan) (*web.Kunjungan, error)
	Find(userID uint) (*[]web.Kunjungan, error)
}

type kunjungan struct {
	krepo    repository.KunjunganInterface
	reported repository.ReportedInterface
	validate middleware.IValidation
	log pkg.ILogrust
}

func NewKunjungan(krepo repository.KunjunganInterface, reported repository.ReportedInterface, log pkg.ILogrust) KunjunganInterface {
	return &kunjungan{
		krepo:    krepo,
		reported: reported,
		validate: middleware.NewValidation(),
		log: log,
	}
}

func (k *kunjungan) Create(req domain.Kunjungan) (*web.Kunjungan, error) {
	err := k.validate.Validate(req)
	if err != nil {
		k.log.StartLogger("kunjungan_service", "create").Error("cannot validation field")
		return nil, k.validate.Error(err)
	}

	data, err := k.krepo.Create(req)
	if err != nil {
		k.log.StartLogger("kunjungan_service", "create").Error("cannot create into db")
		return nil, util.Errors(err)
	}

	err = k.reported.Create(data)
	if err != nil {
		k.log.StartLogger("kunjungan_service", "create").Error("cannot create into reported")
		return nil, util.Errors(err)
	}

	obats, err := k.krepo.GotPrice(req)
	if err != nil {
		k.log.StartLogger("kunjungan_service", "create").Error("cannot got price in db")
		return nil, util.Errors(err)
	}

	var penyakit []web.Penyakit
	for _, p := range data.Penyakit {
		var obat []web.Obat
		for _, o := range *obats {
			res := web.Obat{
				ID:           o.ID,
				Price:        int(o.Price),
				Composition:  o.Composition,
				Description:  o.Description,
				Name:         o.Name,
				QtyAvailable: int(o.QtyAvailable),
				QtyReserved:  int(o.QtyReserved),
				QtyOn:        int(o.QtyOn),
			}

			obat = append(obat, res)
		}

		py := web.Penyakit{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Obat:        obat,
		}

		penyakit = append(penyakit, py)
	}

	response := web.Kunjungan{
		FName:    data.User.FName,
		LName:    data.User.LName,
		Age:      data.User.Age,
		Address:  data.User.Address,
		Penyakit: penyakit,
	}

	k.log.StartLogger("kunjungan_service", "create").Info("success create kunjungan")
	return &response, nil
}

func (k *kunjungan) Find(userID uint) (*[]web.Kunjungan, error) {
	data, err := k.krepo.List(userID)
	if err != nil {
		k.log.StartLogger("kunjungan_service", "findByUserID").Error("cannot get list by this user id")
		return nil, util.Errors(err)
	}

	var response []web.Kunjungan
	var penyakit []web.Penyakit
	for _, dt := range *data {
		for _, p := range dt.Penyakit {
			var obat []web.Obat
			for _, o := range p.Obat {
				ob := web.Obat{
					ID:           o.ID,
					Name:         o.Name,
					Composition:  o.Composition,
					Description:  o.Description,
					Price:        int(o.Price),
					QtyAvailable: int(o.QtyAvailable),
					QtyReserved:  int(o.QtyReserved),
					QtyOn:        int(o.QtyOn),
				}

				obat = append(obat, ob)
			}

			py := web.Penyakit{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Obat:        obat,
			}

			penyakit = append(penyakit, py)
		}

		result := web.Kunjungan{
			FName:    dt.User.FName,
			LName:    dt.User.LName,
			Age:      dt.User.Age,
			Address:  dt.User.Address,
			Penyakit: penyakit,
		}

		response = append(response, result)
	}

	k.log.StartLogger("kunjungan_service", "findByUserID").Info("kunjungan found")
	return &response, nil
}
