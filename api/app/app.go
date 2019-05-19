package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nirmoy/vmctl/api/app/handler"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/servers", a.GetAllServer)
	a.Post("/servers", a.CreateServer)
	a.Get("/servers/{uuid}", a.GetServerByUUID)
	a.Get("/servers/{uuid}/status", a.GetServerStatusByUUID)
	a.Delete("/servers/{uuid}", a.DeleteServerByUUID)
	a.Get("/check/{name}", a.CheckServer)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) GetAllServer(w http.ResponseWriter, r *http.Request) {
	handler.GetAllServer(w, r)
}

func (a *App) GetServerByUUID(w http.ResponseWriter, r *http.Request) {
	handler.GetServerByUUID(w, r)
}

func (a *App) GetServerStatusByUUID(w http.ResponseWriter, r *http.Request) {
	handler.GetServerStatusByUUID(w, r)
}

func (a *App) CreateServer(w http.ResponseWriter, r *http.Request) {
	handler.CreateServer(w, r)
}

func (a *App) DeleteServerByUUID(w http.ResponseWriter, r *http.Request) {
	handler.DeleteServerByUUID(w, r)
}

func (a *App) CheckServer(w http.ResponseWriter, r *http.Request) {
	handler.CheckServer(w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
