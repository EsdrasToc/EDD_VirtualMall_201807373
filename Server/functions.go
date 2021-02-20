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
		return
	}

	fmt.Fprint(w, "La tienda con el indice solicitado, no existe")
}

func Graph(w http.ResponseWriter, r *http.Request) {
	var counter, aux int
	var content, edges, auxContent, auxEdges, product string
	counter = 0
	for i := 0; i < len(vectorData); i++ {
		if i == 0 {
			content = content + "node" + strconv.Itoa(counter) + "[label=\"" + vectorData[i].Index + " " + vectorData[i].Departament + "\"]\n"
		} else {
			content = content + "node" + strconv.Itoa(counter) + "[label=\"" + vectorData[i].Index + " " + vectorData[i].Departament + " " + strconv.Itoa(vectorData[i].Score) + "\"]\n"
			if vectorData[i-1].Lenght != 0 {
				edges = edges + "node" + strconv.Itoa(aux) + "->node" + strconv.Itoa(counter) + "\n"
			}
		}
		edges = edges + "node" + strconv.Itoa(counter) + "->node" + strconv.Itoa(counter+1) + "\n"

		aux = counter
		auxContent, auxEdges = vectorData[i].ToGraph(&counter)

		content = content + auxContent
		edges = edges + auxEdges
		counter++
	}
	product = "Digraph G{\nrankdir=\"LR\"\n" + content + "\n\n" + edges + "}"

	fmt.Fprintln(w, product)
}

func Save(w http.ResponseWriter, r *http.Request) {
	var dataAux Structures.Data

	fmt.Fprintln(w, dataAux.ToJson(vectorData))
}
