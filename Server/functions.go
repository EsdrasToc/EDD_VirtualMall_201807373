package Server

import (
	"fmt"

	"io/ioutil"

	"net/http"

	"../Structures"

	"github.com/gorilla/mux"
)

var data Structures.Data
var vectorData []Structures.ScoreCategory
var finder Structures.Search

func HolaMundoNumber(w http.ResponseWriter, r *http.Request) {
	id, err := mux.Vars(r)["ID"]

	if !err {
		fmt.Println("Ocurrio un error, chale")
		return
	}

	fmt.Fprintf(w, "Hola mundo, este es mi primer server con Go: "+string(id))
}

func UploadShops(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	vectorData = data.ReadJson(body)
	fmt.Fprintf(w, "Informacion guardada correctamente")
}

func EspecificShop(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	finder.ReadJson(body)
	fmt.Fprintf(w, finder.EspecificSearchEngine(vectorData))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	finder.ReadJson(body)
	vectorData = *finder.Delete(vectorData, w)
}
