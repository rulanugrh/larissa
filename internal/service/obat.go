package service

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rulanugrh/larissa/internal/entity/web"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/util"
	"github.com/rulanugrh/larissa/pkg"
)

type ObatInterface interface {
	FindID(id uint) (*web.Obat, error)
	FindAll() (*[]web.Obat, error)
}

type obat struct {
	orepo  repository.ObatInterface
	metric pkg.Data
}

func NewObat(orepo repository.ObatInterface, metric pkg.Data) ObatInterface {
	return &obat{
		orepo: orepo,
		metric: metric,
	}
}

func(o *obat) FindID(id uint) (*web.Obat, error)  {
	data, err := o.orepo.FindID(id)
	if err != nil {
		return nil, util.Errors(err)
	}

	response := web.Obat {
		ID: data.ID,
		Name: data.Name,
		Composition: data.Composition,
		Description: data.Description,
		QtyAvailable: int(data.QtyAvailable),
		Price: int(data.Price),
	}

	return &response, nil
}

func(o *obat) FindAll() (*[]web.Obat, error)  {
	data, err := o.orepo.FindAll()
	if err != nil {
		return nil, util.Errors(err)
	}

	var response []web.Obat
	for _, v := range *data {
		result := web.Obat {
			ID: v.ID,
			Name: v.Name,
			Composition: v.Composition,
			Description: v.Description,
			QtyAvailable: int(v.QtyAvailable),
			Price: int(v.Price),
		}

		response = append(response, result)
	}

	o.metric.ObatHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())
	o.metric.Obat.Set(float64(len(*data)))

	return &response, nil
}
