package service

import (
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/entity/web"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/util"
)

type KunjunganInterface interface {
	Create(req domain.Kunjungan) (*web.Kunjungan, error)
	Find(userID uint) (*[]web.Kunjungan, error)
}

type kunjungan struct {
	krepo repository.KunjunganInterface
	reported repository.ReportedInterface
}

func NewKunjungan(krepo repository.KunjunganInterface, reported repository.ReportedInterface) KunjunganInterface {
	return &kunjungan{
		krepo: krepo,
		reported: reported,
	}
}

func(k *kunjungan) Create(req domain.Kunjungan) (*web.Kunjungan, error) {
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