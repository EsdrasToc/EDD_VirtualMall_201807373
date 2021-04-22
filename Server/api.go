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
	r.HandleFunc("/getGraphYears", graphYears).Methods("GET")
	r.HandleFunc("/getGraphMonths/{Anio:[0-9]+}", graphMonths).Methods("GET")
	r.HandleFunc("/getMonth/{Anio:[0-9]+}/{Mes:[0-9]+}", graphMonths).Methods("Get")
	r.HandleFunc("/AddUsers", addAccounts).Methods("POST")
	r.HandleFunc("/Authenticate", authenticate).Methods("POST")
	r.HandleFunc("/putOrder", putOrder).Methods("PUT")
	r.HandleFunc("/Month", GetMonth).Methods("PUT")
	r.HandleFunc("/GraphAccounts", GraphAccounts).Methods("GET")
	r.HandleFunc("/NewUser", addUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
