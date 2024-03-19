package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/service"
	"github.com/rulanugrh/larissa/internal/util"
	"github.com/rulanugrh/larissa/pkg"
)

type AdminInterface interface {
	Penyakit(w http.ResponseWriter, r *http.Request)
	Obat(w http.ResponseWriter, r *http.Request)
	UpdateObat(w http.ResponseWriter, r *http.Request)
	DeleteObat(w http.ResponseWriter, r *http.Request)
	DeletePenyakit(w http.ResponseWriter, r *http.Request)
	Reported(w http.ResponseWriter, r *http.Request)
	ListAllUser(w http.ResponseWriter, r *http.Request)
}

type admin struct {
	service service.AdminInterface
	gauge *pkg.Data
}

func NewAdmin(service service.AdminInterface, gauge *pkg.Data) AdminInterface {
	return &admin{
		service: service,
		gauge: gauge,
	}
}

func(a *admin) Penyakit(w http.ResponseWriter, r *http.Request) {
	var req domain.Penyakit
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		a.gauge.PenyakitHistory.With(prometheus.Labels{"code": "500", "method": "POST", "type": "create"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	data, err := a.service.CreatePenyakit(req)
	if err != nil {
		a.gauge.PenyakitHistory.With(prometheus.Labels{"code": "400", "method": "POST", "type": "create"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	a.gauge.Penyakit.Inc()
	a.gauge.PenyakitHistory.With(prometheus.Labels{"code": "200", "method": "POST", "type": "create"}).Observe(time.Since(time.Now()).Seconds())
	a.gauge.PenyakitUpgrade.With(prometheus.Labels{"type": "create"}).Inc()

	response := util.WriteJSON(util.Success("success add to database", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}

func(a *admin) Obat(w http.ResponseWriter, r *http.Request) {
	var req domain.Obat
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		a.gauge.ObatHistory.With(prometheus.Labels{"code": "500", "method": "POST", "type": "create"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	data, err := a.service.CreateObat(req)
	if err != nil {
		a.gauge.ObatHistory.With(prometheus.Labels{"code": "400", "method": "POST", "type": "create"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	a.gauge.Obat.Inc()
	a.gauge.ObatHistory.With(prometheus.Labels{"code": "200", "method": "POST", "type": "create"}).Observe(time.Since(time.Now()).Seconds())
	a.gauge.ObatUpgrade.With(prometheus.Labels{"type": "create"}).Inc()

	response := util.WriteJSON(util.Success("success add to database", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}

func(a *admin) UpdateObat(w http.ResponseWriter, r *http.Request) {
	var req domain.Obat
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		a.gauge.ObatHistory.With(prometheus.Labels{"code": "500", "method": "PUT", "type": "update"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/admin/update/obat/"))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	data, err := a.service.UpdateObat(uint(id), req)
	if err != nil {
		a.gauge.ObatHistory.With(prometheus.Labels{"code": "400", "method": "PUT", "type": "update"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	a.gauge.ObatHistory.With(prometheus.Labels{"code": "200", "method": "PUT", "type": "update"}).Observe(time.Since(time.Now()).Seconds())
	a.gauge.ObatUpgrade.With(prometheus.Labels{"type": "update"}).Inc()

	response := util.WriteJSON(util.Success("success update obat", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}

func(a *admin) DeleteObat(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/admin/delete/obat/"))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = a.service.DeleteObat(uint(id))
	if err != nil {
		a.gauge.ObatHistory.With(prometheus.Labels{"code": "400", "method": "DELETE", "type": "delete"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	a.gauge.ObatHistory.With(prometheus.Labels{"code": "204", "method": "DELETE", "type": "delete"}).Observe(time.Since(time.Now()).Seconds())
	a.gauge.ObatUpgrade.With(prometheus.Labels{"type": "delete"}).Inc()
	w.WriteHeader(204)
	w.Write([]byte("successfull delete data"))
	return
}

func(a *admin) DeletePenyakit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/admin/delete/penyakit/"))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = a.service.DeletePenyakit(uint(id))
	if err != nil {
		a.gauge.PenyakitHistory.With(prometheus.Labels{"code": "400", "method": "DELETE", "type": "delete"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	a.gauge.PenyakitHistory.With(prometheus.Labels{"code": "204", "method": "DELETE", "type": "delete"}).Observe(time.Since(time.Now()).Seconds())
	a.gauge.PenyakitUpgrade.With(prometheus.Labels{"type": "delete"}).Inc()
	w.WriteHeader(204)
	w.Write([]byte("successfull delete data"))
	return
}

func(a *admin) Reported(w http.ResponseWriter, r *http.Request) {
	data, err := a.service.Reported()
	if err != nil {
		a.gauge.KunjunganHistory.With(prometheus.Labels{"code": "400", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	a.gauge.KunjunganHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())
	a.gauge.Kunjungan.Set(float64(len(*data)))

	response := util.WriteJSON(util.Success("data reported found", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}

func(a *admin) ListAllUser(w http.ResponseWriter, r *http.Request) {
	data, err := a.service.ListAllUser()
	if err != nil {
		a.gauge.UserHistory.With(prometheus.Labels{"code": "400", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	a.gauge.UserHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())
	a.gauge.User.Set(float64(len(*data)))

	response := util.WriteJSON(util.Success("data user found", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}
