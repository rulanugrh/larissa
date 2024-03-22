package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/middleware"
	"github.com/rulanugrh/larissa/internal/service"
	"github.com/rulanugrh/larissa/internal/util"
	"github.com/rulanugrh/larissa/pkg"
)

type KunjunganInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
}

type kunjungan struct {
	service service.KunjunganInterface
	gauge   *pkg.Data
}

func NewKunjungan(service service.KunjunganInterface, gauge *pkg.Data) KunjunganInterface {
	return &kunjungan{
		service: service,
		gauge:   gauge,
	}
}

func (k *kunjungan) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Kunjungan
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		k.gauge.KunjunganHistory.With(prometheus.Labels{"code": "500", "method": "POST", "type": "create"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	data, err := k.service.Create(req)
	if err != nil {
		k.gauge.KunjunganHistory.With(prometheus.Labels{"code": "400", "method": "POST", "type": "create"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	k.gauge.Kunjungan.Inc()
	k.gauge.KunjunganHistory.With(prometheus.Labels{"code": "200", "method": "POST", "type": "create"}).Observe(time.Since(time.Now()).Seconds())
	k.gauge.KunjunganUpgrade.With(prometheus.Labels{"type": "create"}).Inc()

	response := util.WriteJSON(util.Created("recorde data kunjungan", data))
	w.WriteHeader(201)
	w.Write(response)
	return
}

func (k *kunjungan) Find(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetUserID(r)
	if err != nil {
		k.gauge.KunjunganHistory.With(prometheus.Labels{"code": "401", "method": "GET", "type": "getByID"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.Unauthorized(err.Error()))
		w.WriteHeader(401)
		w.Write(response)
		return
	}

	data, err := k.service.Find(id)
	if err != nil {
		k.gauge.KunjunganHistory.With(prometheus.Labels{"code": "400", "method": "GET", "type": "getByID"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	k.gauge.KunjunganHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "getByID"}).Observe(time.Since(time.Now()).Seconds())
	k.gauge.KunjunganUpgrade.With(prometheus.Labels{"type": "getByID"}).Inc()

	response := util.WriteJSON(util.Success("find kunjungan by this user id", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}
