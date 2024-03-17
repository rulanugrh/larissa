package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/middleware"
	"github.com/rulanugrh/larissa/internal/service"
	"github.com/rulanugrh/larissa/internal/util"
)

type KunjunganInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
}

type kunjungan struct {
	service service.KunjunganInterface
}

func NewKunjungan(service service.KunjunganInterface) KunjunganInterface {
	return &kunjungan{
		service: service,
	}
}

func(k *kunjungan) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Kunjungan
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	data, err := k.service.Create(req)
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response := util.WriteJSON(util.Created("recorde data kunjungan", data))
	w.WriteHeader(201)
	w.Write(response)
	return
}

func(k *kunjungan) Find(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	check, err := middleware.CheckToken(token)
	if err != nil {
		response := util.WriteJSON(util.Unauthorized(err.Error()))
		w.WriteHeader(401)
		w.Write(response)
		return
	}

	data, err := k.service.Find(check.ID)
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response := util.WriteJSON(util.Success("find kunjungan by this user id", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}
