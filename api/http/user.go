package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/middleware"
	"github.com/rulanugrh/larissa/internal/service"
	"github.com/rulanugrh/larissa/internal/util"
)

type UserInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type user struct {
	service service.UserInterface
}

func NewUser(service service.UserInterface) UserInterface {
	return &user{
		service: service,
	}
}

func(u *user) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.UserRegister
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	data, err := u.service.Register(req)
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	response := util.WriteJSON(util.Created("success register account", data))
	w.WriteHeader(201)
	w.Write(response)
	return
}

func(u *user) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.UserLogin
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	data, err := u.service.Login(req)
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	err = util.SetCookie("token", data, w)
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("sucess login"))
	return
}

func(u *user) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	token := r.Header.Get("Authorization")
	claims, err := middleware.CheckToken(token)
	if err != nil {
		response := util.WriteJSON(util.Unauthorized(err.Error()))
		w.WriteHeader(401)
		w.Write(response)
		return
	}

	err = u.service.Update(claims.ID, req)
	if err != nil {
		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("success update user"))
	return
}
