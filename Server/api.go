package Server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/{ID:[0-9]+}", HolaMundoNumber)
	r.HandleFunc("/cargartienda", UploadShops).Methods("POST")
	r.HandleFunc("/TiendaEspecifica", EspecificShop).Methods("POST")
	r.HandleFunc("/Eliminar", Delete).Methods("DELETE")

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
