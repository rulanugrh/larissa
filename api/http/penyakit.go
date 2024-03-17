package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/rulanugrh/larissa/internal/service"
	"github.com/rulanugrh/larissa/internal/util"
)

type PenyakitInterface interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	FindID(w http.ResponseWriter, r *http.Request)
}

type penyakit struct {
	service service.PenyakitInterface
}

func NewPenyakit(service service.PenyakitInterface) PenyakitInterface {
	return &penyakit{
		service: service,
	}
}

func(p *penyakit) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := p.service.FindAll()
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

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
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response := util.WriteJSON(util.Success("data found", data))
	w.WriteHeader(200)
	w.Write(response)
}
