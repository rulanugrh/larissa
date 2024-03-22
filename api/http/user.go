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

type UserInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GotDoctor(w http.ResponseWriter, r *http.Request)
	GotNurse(w http.ResponseWriter, r *http.Request)
}

type user struct {
	service service.UserInterface
	gauge *pkg.Data
}

func NewUser(service service.UserInterface, gauge *pkg.Data) UserInterface {
	return &user{
		service: service,
		gauge: gauge,
	}
}

func(u *user) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.UserRegister
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.gauge.UserHistory.With(prometheus.Labels{"code": "500", "method": "POST", "type": "register"}).Observe(time.Since(time.Now()).Seconds())
		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	data, err := u.service.Register(req)
	if err != nil {
		u.gauge.UserHistory.With(prometheus.Labels{"code": "400", "method": "POST", "type": "register"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	u.gauge.User.Inc()
	u.gauge.UserHistory.With(prometheus.Labels{"code": "200", "method": "POST", "type": "register"}).Observe(time.Since(time.Now()).Seconds())
	u.gauge.UserUpgrade.With(prometheus.Labels{"type": "create"}).Inc()

	response := util.WriteJSON(util.Created("success register account", data))
	w.WriteHeader(201)
	w.Write(response)
	return
}

func(u *user) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.UserLogin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.gauge.UserHistory.With(prometheus.Labels{"code": "500", "method": "POST", "type": "login"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	data, err := u.service.Login(req)
	if err != nil {
		u.gauge.UserHistory.With(prometheus.Labels{"code": "400", "method": "POST", "type": "login"}).Observe(time.Since(time.Now()).Seconds())

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

	u.gauge.UserHistory.With(prometheus.Labels{"code": "200", "method": "POST", "type": "login"}).Observe(time.Since(time.Now()).Seconds())
	u.gauge.UserUpgrade.With(prometheus.Labels{"type": "login"}).Inc()

	w.WriteHeader(200)
	w.Write([]byte("sucess login"))
	return
}

func(u *user) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		u.gauge.UserHistory.With(prometheus.Labels{"code": "500", "method": "PUT", "type": "update"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.InternalServerError("cannot read request body"))
		w.WriteHeader(500)
		w.Write(response)
		return
	}

	id, err := middleware.GetUserID(r)
	if err != nil {
		u.gauge.UserHistory.With(prometheus.Labels{"code": "401", "method": "PUT", "type": "update"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.Unauthorized(err.Error()))
		w.WriteHeader(401)
		w.Write(response)
		return
	}

	err = u.service.Update(id, req)
	if err != nil {
		u.gauge.UserHistory.With(prometheus.Labels{"code": "400", "method": "PUT", "type": "update"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}


	u.gauge.UserHistory.With(prometheus.Labels{"code": "200", "method": "PUT", "type": "update"}).Observe(time.Since(time.Now()).Seconds())
	u.gauge.UserUpgrade.With(prometheus.Labels{"type": "update"}).Inc()

	w.WriteHeader(200)
	w.Write([]byte("success update user"))
	return
}

func(u *user) GotDoctor(w http.ResponseWriter, r *http.Request) {
	data, err := u.service.GotDoctor()
	if err != nil {
		u.gauge.UserHistory.With(prometheus.Labels{"code": "400", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	u.gauge.UserHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())

	response := util.WriteJSON(util.Created("success get doctor", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}

func(u *user) GotNurse(w http.ResponseWriter, r *http.Request) {
	data, err := u.service.GotNurse()
	if err != nil {
		u.gauge.UserHistory.With(prometheus.Labels{"code": "400", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())

		response := util.WriteJSON(util.BadRequest(err.Error()))
		w.WriteHeader(400)
		w.Write(response)
		return
	}

	u.gauge.UserHistory.With(prometheus.Labels{"code": "200", "method": "GET", "type": "get"}).Observe(time.Since(time.Now()).Seconds())

	response := util.WriteJSON(util.Created("success get nurse", data))
	w.WriteHeader(200)
	w.Write(response)
	return
}
