package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/rulanugrh/larissa/internal/service"
	"github.com/rulanugrh/larissa/internal/util"
)

type ObatInterface interface {
	FindID(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
}

type obat struct {
	service service.ObatInterface
}

func NewObat(service service.ObatInterface) ObatInterface {
	return &obat{
		service: service,
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
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response := util.WriteJSON(util.Success("find obat by this id", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}

func(o *obat) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := o.service.FindAll()
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response := util.WriteJSON(util.Success("sucessfull, obat found", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}