package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/service"
	"github.com/rulanugrh/larissa/internal/util"
)

type AdminInterface interface {
	Penyakit(w http.ResponseWriter, r *http.Request)
	Obat(w http.ResponseWriter, r *http.Request)
	UpdateObat(w http.ResponseWriter, r *http.Request)
	DeleteObat(w http.ResponseWriter, r *http.Request)
	DeletePenyakit(w http.ResponseWriter, r *http.Request)
	Reported(w http.ResponseWriter, r *http.Request)
}

type admin struct {
	service service.AdminInterface
}

func NewAdmin(service service.AdminInterface) AdminInterface {
	return &admin{
		service: service,
	}
}

func(a *admin) Penyakit(w http.ResponseWriter, r *http.Request) {
	var req domain.Penyakit
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	data, err := a.service.CreatePenyakit(req)
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response := util.WriteJSON(util.Success("success add to database", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}

func(a *admin) Obat(w http.ResponseWriter, r *http.Request) {
	var req domain.Obat
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	data, err := a.service.CreateObat(req)
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response := util.WriteJSON(util.Success("success add to database", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}

func(a *admin) UpdateObat(w http.ResponseWriter, r *http.Request) {
	var req domain.Obat
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
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
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

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
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	w.WriteHeader(200)
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
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("successfull delete data"))
	return
}

func(a *admin) Reported(w http.ResponseWriter, r *http.Request) {
	data, err := a.service.Reported()
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response := util.WriteJSON(util.Success("data reported found", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}
