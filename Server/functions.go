package Server

import (
	"fmt"
	"strconv"

	"io/ioutil"

	"net/http"

	"../Structures"

	"github.com/gorilla/mux"
)

var data Structures.Data
var vectorData []Structures.ScoreCategory
var finder Structures.Search

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

func SearchPosition(w http.ResponseWriter, r *http.Request) {
	x, err := mux.Vars(r)["ID"]
	id, _ := strconv.Atoi(x)

	if err && id <= len(vectorData) {
		fmt.Fprintln(w, vectorData[id-1].ToJson())
	}

	fmt.Fprint(w, "La tienda con el indice solicitado, no existe")
}

func ShowData(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(vectorData); i++ {
		fmt.Fprintln(w, vectorData[i].Departament)
	}
}
