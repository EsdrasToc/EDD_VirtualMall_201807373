package Server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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

	//w.Header().Set("Access-Control-Allow-Origin", "*")

	//log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))

	r.HandleFunc("/cargartienda", UploadShops).Methods("POST")
	r.HandleFunc("/TiendaEspecifica", EspecificShop).Methods("POST")
	r.HandleFunc("/Eliminar", Delete).Methods("DELETE")
	r.HandleFunc("/id/{ID:[0-9]+}", SearchPosition).Methods("GET")
	r.HandleFunc("/getArreglo", Graph).Methods("GET")
	r.HandleFunc("/guardar", Save).Methods("GET")
	r.HandleFunc("/AddInventory", AddInventory).Methods("POST")
	r.HandleFunc("/getshops", getShops).Methods("GET")
	r.HandleFunc("/getProducts/{Name:.+}/{Score:[1-9]}", getProducts).Methods("GET")
	r.HandleFunc("/putPurchase", putPurchase).Methods("PUT")
	r.HandleFunc("/postOrders", addOrders).Methods("POST")
	r.HandleFunc("/getYears", getYears).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
