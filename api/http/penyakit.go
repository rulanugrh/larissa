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

type PenyakitInterface interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	FindID(w http.ResponseWriter, r *http.Request)
}

type penyakit struct {
	service service.PenyakitInterface
	gauge *pkg.Data
}

func NewPenyakit(service service.PenyakitInterface, gauge *pkg.Data) PenyakitInterface {
	return &penyakit{
		service: service,
		gauge: gauge,
	}
}

func(p *penyakit) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := p.service.FindAll()
	if err != nil {
		p.gauge.PenyakitHistory.With(prometheus.Labels{"code": "400", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	p.gauge.PenyakitHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())
	p.gauge.Penyakit.Set(float64(len(*data)))

	response := util.WriteJSON(util.Success("data found", data))
	w.WriteHeader(200)
	w.Write(response)
}

func(p *penyakit) FindID(w http.ResponseWriter, r *http.Request) {
	id, err  := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/penyakit/find/"))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	data, err := p.service.FindID(uint(id))
	if err != nil {
		p.gauge.PenyakitHistory.With(prometheus.Labels{"code": "400", "method": "GET", "type": "getByID"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	p.gauge.PenyakitHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "getByID"}).Observe(time.Since(time.Now()).Seconds())

	response := util.WriteJSON(util.Success("data found", data))
	w.WriteHeader(200)
	w.Write(response)
}
