package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	handler "github.com/rulanugrh/larissa/api/http"
	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/middleware"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/service"
	"github.com/rulanugrh/larissa/pkg"
)

type API struct {
	user      handler.UserInterface
	obat      handler.ObatInterface
	admin     handler.AdminInterface
	kunjungan handler.KunjunganInterface
	penyakit  handler.PenyakitInterface
}

func (api *API) UserRoute(r *mux.Router) {
	app := r.PathPrefix("/api/user/").Subrouter()
	app.HandleFunc("/register", api.user.Register).Methods("POST")
	app.HandleFunc("/login", api.user.Login).Methods("POST")
	app.HandleFunc("/update", api.user.Update).Methods("PUT")
	app.HandleFunc("/doctor", api.user.GotDoctor).Methods("GET")
	app.HandleFunc("/nurse", api.user.GotNurse).Methods("GET")
}

func (api *API) KunjunganRoute(r *mux.Router) {
	app := r.PathPrefix("/api/kunjungan/").Subrouter()
	app.Use(middleware.GeneralVerify)
	app.HandleFunc("/create", api.kunjungan.Create).Methods("POST")
	app.HandleFunc("/find", api.kunjungan.Find).Methods("GET")
}

func (api *API) ObatRoute(r *mux.Router) {
	app := r.PathPrefix("/api/obat/").Subrouter()
	app.Use(middleware.GeneralVerify)
	app.HandleFunc("/get", api.obat.FindAll).Methods("GET")
	app.HandleFunc("/find/{id}", api.obat.FindID).Methods("GET")
}

func (api *API) PenyakitRoute(r *mux.Router) {
	app := r.PathPrefix("/api/penyakit/").Subrouter()
	app.Use(middleware.GeneralVerify)
	app.HandleFunc("/get", api.penyakit.FindAll).Methods("GET")
	app.HandleFunc("/find/{id}", api.penyakit.FindID).Methods("GET")
}

func (api *API) AdminRoute(r *mux.Router) {
	app := r.PathPrefix("/api/admin/").Subrouter()
	app.Use(middleware.AdminVerify)
	app.HandleFunc("/create/pbat", api.admin.Obat).Methods("POST")
	app.HandleFunc("/create/penyakit", api.admin.Penyakit).Methods("POST")
	app.HandleFunc("/update/obat/{id}", api.admin.UpdateObat).Methods("PUT")
	app.HandleFunc("/delete/obat/{id}", api.admin.DeleteObat).Methods("DELETE")
	app.HandleFunc("/delete/penyakit/{id}", api.admin.DeletePenyakit).Methods("DELETE")
	app.HandleFunc("/reported", api.admin.Reported).Methods("GET")
	app.HandleFunc("/user", api.admin.ListAllUser).Methods("GET")

}

func main() {
	command := os.Args[1]
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	m := pkg.NewMetric(reg)
	gauge := m.SetGauge()
	log := pkg.NewLogger()

	conf := config.GetConfig()
	postgres := config.InitializePostgres(conf)
	postgres.NewConnection()
	mongo := config.InitializeMongo(conf)

	userRepository := repository.NewUser(postgres)
	obatRepository := repository.NewObat(postgres)
	kunjunganRepository := repository.NewKunjungan(postgres)
	penyakitRepository := repository.NewPenyakit(postgres)
	reportedRepository := repository.NewReported(mongo.Conn, conf)

	userService := service.NewUser(userRepository, log)
	obatService := service.NewObat(obatRepository, log)
	kunjunganServices := service.NewKunjungan(kunjunganRepository, reportedRepository, log)
	penyakitService := service.NewPenyakit(penyakitRepository, log)
	adminService := service.NewAdmin(obatRepository, penyakitRepository, reportedRepository, userRepository)

	api := API{
		obat:      handler.NewObat(obatService, gauge),
		user:      handler.NewUser(userService, gauge),
		kunjungan: handler.NewKunjungan(kunjunganServices, gauge),
		penyakit:  handler.NewPenyakit(penyakitService, gauge),
		admin:     handler.NewAdmin(adminService, gauge),
	}

	routes := mux.NewRouter()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	// add routes
	routes.Handle("/metric", promHandler).Methods("GET")
	api.UserRoute(routes)
	api.AdminRoute(routes)
	api.ObatRoute(routes)
	api.PenyakitRoute(routes)
	api.KunjunganRoute(routes)

	dsn := fmt.Sprintf("%s:%s", conf.App.Host, conf.App.Port)
	server := http.Server{
		Addr:    dsn,
		Handler: routes,
	}

	if command == "migrate" {
		postgres.Migration()

	} else if command == "seed" {
		err := postgres.Seeder()
		if err != nil {
			log.Printf("error seeder to db: %s", err.Error())
		}

	} else if command == "serve" {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("cannot running http service because: %s", err.Error())
		}

		log.Printf("running at: %s", dsn)

	} else {
		log.Println("error command")
	}

}
