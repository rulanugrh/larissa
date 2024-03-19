package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rulanugrh/larissa/internal/service"
	"github.com/rulanugrh/larissa/internal/util"
	"github.com/rulanugrh/larissa/pkg"
)

type ObatInterface interface {
	FindID(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
}

type obat struct {
	service service.ObatInterface
	gauge *pkg.Data
}

func NewObat(service service.ObatInterface, gauge *pkg.Data) ObatInterface {
	return &obat{
		service: service,
		gauge: gauge,
	}
}

func(o *obat) FindID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/obat/find/"))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	data, err := o.service.FindID(uint(id))
	if err != nil {
		o.gauge.ObatHistory.With(prometheus.Labels{"code": "400", "method": "GET", "type": "getByID"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	o.gauge.ObatHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "getByID"}).Observe(time.Since(time.Now()).Seconds())

	response := util.WriteJSON(util.Success("find obat by this id", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}

func(o *obat) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := o.service.FindAll()
	if err != nil {
		o.gauge.ObatHistory.With(prometheus.Labels{"code": "400", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	o.gauge.ObatHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())
	o.gauge.Obat.Set(float64(len(*data)))

	response := util.WriteJSON(util.Success("sucessfull, obat found", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}
