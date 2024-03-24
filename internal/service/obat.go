package service

import (
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
	orepo repository.ObatInterface
	log   pkg.ILogrust
}

func NewObat(orepo repository.ObatInterface, log pkg.ILogrust) ObatInterface {
	return &obat{
		orepo: orepo,
		log:   log,
	}
}

func (o *obat) FindID(id uint) (*web.Obat, error) {
	data, err := o.orepo.FindID(id)
	if err != nil {
		o.log.StartLogger("obat_service", "findID").Error(err.Error())
		return nil, util.Errors(err)
	}

	response := web.Obat{
		ID:           data.ID,
		Name:         data.Name,
		Composition:  data.Composition,
		Description:  data.Description,
		QtyAvailable: int(data.QtyAvailable),
		Price:        int(data.Price),
		QtyReserved:  int(data.QtyReserved),
		QtyOn:        int(data.QtyOn),
	}

	o.log.StartLogger("obat_service", "findID").Info("success get obat")
	return &response, nil
}

func (o *obat) FindAll() (*[]web.Obat, error) {
	data, err := o.orepo.FindAll()
	if err != nil {
		o.log.StartLogger("obat_service", "findAll").Error(err.Error())
		return nil, util.Errors(err)
	}

	var response []web.Obat
	for _, v := range *data {
		result := web.Obat{
			ID:           v.ID,
			Name:         v.Name,
			Composition:  v.Composition,
			Description:  v.Description,
			QtyAvailable: int(v.QtyAvailable),
			Price:        int(v.Price),
			QtyReserved:  int(v.QtyReserved),
			QtyOn:        int(v.QtyOn),
		}

		response = append(response, result)
	}

	o.log.StartLogger("obat_service", "findAll").Info("success get obat")
	return &response, nil
}
