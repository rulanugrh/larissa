package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/rulanugrh/larissa/api/http"
	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/service"
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
}

func (api *API) KunjunganRoute(r *mux.Router) {
	app := r.PathPrefix("/api/kunjungan/").Subrouter()
	app.HandleFunc("/create", api.kunjungan.Create).Methods("POST")
	app.HandleFunc("/find", api.kunjungan.Find).Methods("GET")
}

func (api *API) ObatRoute(r *mux.Router) {
	app := r.PathPrefix("/api/obat/").Subrouter()
	app.HandleFunc("/get", api.obat.FindAll).Methods("GET")
	app.HandleFunc("/find/{id}", api.obat.FindID).Methods("GET")
}

func (api *API) PenyakitRoute(r *mux.Router) {
	app := r.PathPrefix("/api/penyakit/").Subrouter()
	app.HandleFunc("/get", api.penyakit.FindAll).Methods("GET")
	app.HandleFunc("/find/{id}", api.penyakit.FindID).Methods("GET")
}

func (api *API) AdminRoute(r *mux.Router) {
	app := r.PathPrefix("/api/admin/").Subrouter()
	app.HandleFunc("/create/pbat", api.admin.Obat).Methods("POST")
	app.HandleFunc("/create/penyakit", api.admin.Penyakit).Methods("POST")
	app.HandleFunc("/update/obat/{id}", api.admin.UpdateObat).Methods("PUT")
	app.HandleFunc("/delete/obat/{id}", api.admin.DeleteObat).Methods("DELETE")
	app.HandleFunc("/delete/penyakit/{id}", api.admin.DeletePenyakit).Methods("DELETE")
	app.HandleFunc("/reported", api.admin.Reported).Methods("GET")
}

func main() {
	conf := config.GetConfig()
	postgres := config.InitializePostgres(conf)
	postgres.NewConnection()

	mongo := config.InitializeMongo(conf)

	userRepository := repository.NewUser(postgres)
	obatRepository := repository.NewObat(postgres)
	kunjunganRepository := repository.NewKunjungan(postgres)
	penyakitRepository := repository.NewPenyakit(postgres)
	reportedRepository := repository.NewReported(mongo.Conn, conf)

	userService := service.NewUser(userRepository)
	obatService := service.NewObat(obatRepository)
	kunjunganServices := service.NewKunjungan(kunjunganRepository, reportedRepository)
	penyakitService := service.NewPenyakit(penyakitRepository)
	adminService := service.NewAdmin(obatRepository, penyakitRepository, reportedRepository)

	api := API{
		obat:      handler.NewObat(obatService),
		user:      handler.NewUser(userService),
		kunjungan: handler.NewKunjungan(kunjunganServices),
		penyakit:  handler.NewPenyakit(penyakitService),
		admin:     handler.NewAdmin(adminService),
	}

	routes := mux.NewRouter()

	// add routes
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

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("cannot running http service because: %s", err.Error())
	}

	log.Printf("running at: %s", dsn)

}
