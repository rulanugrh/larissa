package service

import (
	"github.com/rulanugrh/larissa/internal/entity/web"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/util"
	"github.com/rulanugrh/larissa/pkg"
)

type PenyakitInterface interface {
	FindID(id uint) (*web.Penyakit, error)
	FindAll() (*[]web.Penyakit, error)
}

type penyakit struct {
	repository repository.PenyakitInterface
	gauge pkg.Data
}

func NewPenyakit(repository repository.PenyakitInterface, gauge pkg.Data) PenyakitInterface {
	return &penyakit{
		repository: repository,
		gauge: gauge,
	}
}

func (p *penyakit) FindID(id uint) (*web.Penyakit, error) {
	data, err := p.repository.FindID(id)
	if err != nil {
		return nil, util.Errors(err)
	}

	var obat []web.Obat
	for _, v := range data.Obat {
		result := web.Obat{
			ID:           v.ID,
			Price:        int(v.Price),
			Name:         v.Name,
			Composition:  v.Composition,
			QtyAvailable: int(v.QtyAvailable),
			Description:  v.Description,
		}

		obat = append(obat, result)
	}

	response := web.Penyakit{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Obat:        obat,
	}

	return &response, nil
}

func (p *penyakit) FindAll() (*[]web.Penyakit, error) {
	data, err := p.repository.ListAll()
	if err != nil {
		return nil, util.Errors(err)
	}

	var response []web.Penyakit
	var obat []web.Obat
	for _, v := range *data {
		for _, o := range v.Obat {
			result := web.Obat{
				ID:           o.ID,
				Name:         o.Name,
				Composition:  o.Composition,
				Description:  o.Description,
				QtyAvailable: int(o.QtyAvailable),
				Price:        int(o.Price),
			}

			obat = append(obat, result)
		}

		result := web.Penyakit{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Obat:        obat,
		}

		response = append(response, result)
	}

	p.gauge.Penyakit.Set(float64(len(*data)))
	return &response, nil
}
