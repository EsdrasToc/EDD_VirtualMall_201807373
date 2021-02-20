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

	r.HandleFunc("/cargartienda", UploadShops).Methods("POST")
	r.HandleFunc("/TiendaEspecifica", EspecificShop).Methods("POST")
	r.HandleFunc("/Eliminar", Delete).Methods("DELETE")
	r.HandleFunc("/id/{ID:[0-9]+}", SearchPosition).Methods("GET")
	r.HandleFunc("/getArreglo", Graph).Methods("GET")
	r.HandleFunc("/guardar", Save).Methods("GET")

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
