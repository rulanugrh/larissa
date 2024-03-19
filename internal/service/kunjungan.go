package service

import (
	"github.com/prometheus/client_golang/prometheus"
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
	krepo repository.KunjunganInterface
	reported repository.ReportedInterface
	validate middleware.IValidation
	gauge pkg.Data
}

func NewKunjungan(krepo repository.KunjunganInterface, reported repository.ReportedInterface, gauge pkg.Data) KunjunganInterface {
	return &kunjungan{
		krepo: krepo,
		reported: reported,
		validate: middleware.NewValidation(),
		gauge: gauge,
	}
}

func(k *kunjungan) Create(req domain.Kunjungan) (*web.Kunjungan, error) {
	err := k.validate.Validate(req)
	if err != nil {
		return nil, k.validate.Error(err)
	}

	data, err := k.krepo.Create(req)
	if err != nil {
		return nil, util.Errors(err)
	}

	err = k.reported.Create(data)
	if err != nil {
		return nil, util.Errors(err)
	}

	var penyakit []web.Penyakit
	for _, p := range data.Penyakit {
		var obat []web.Obat
		for _, o := range p.Obat {
			ob := web.Obat {
				ID: o.ID,
				Name: o.Name,
				Composition: o.Composition,
				Description: o.Description,
				Price: int(o.Price),
				QtyAvailable: int(o.QtyAvailable),
			}

			obat = append(obat, ob)
		}

		py := web.Penyakit{
			ID: p.ID,
			Name: p.Name,
			Description: p.Description,
			Obat: obat,
		}

		penyakit = append(penyakit, py)
	}

	response := web.Kunjungan{
		FName: data.User.FName,
		LName: data.User.LName,
		Age: data.User.Age,
		Address: data.User.Address,
		Penyakit: penyakit,
	}


	k.gauge.Kunjungan.Inc()
	k.gauge.KunjunganUpgrade.With(prometheus.Labels{"type": "create"}).Inc()
	return &response, nil
}

func(k *kunjungan) Find(userID uint) (*[]web.Kunjungan, error) {
	data, err := k.krepo.List(userID)
	if err != nil {
		return nil, util.Errors(err)
	}

	var response []web.Kunjungan
	var penyakit []web.Penyakit
	for _, dt := range *data {
		for _, p := range dt.Penyakit {
			var obat []web.Obat
			for _, o := range p.Obat {
				ob := web.Obat {
					ID: o.ID,
					Name: o.Name,
					Composition: o.Composition,
					Description: o.Description,
					Price: int(o.Price),
					QtyAvailable: int(o.QtyAvailable),
				}

				obat = append(obat, ob)
			}

			py := web.Penyakit{
				ID: p.ID,
				Name: p.Name,
				Description: p.Description,
				Obat: obat,
			}

			penyakit = append(penyakit, py)
		}

		result := web.Kunjungan {
			FName: dt.User.FName,
			LName: dt.User.LName,
			Age: dt.User.Age,
			Address: dt.User.Address,
			Penyakit: penyakit,
		}

		response = append(response, result)
	}

	return &response, nil
}
