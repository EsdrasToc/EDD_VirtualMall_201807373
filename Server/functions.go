package Server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"io/ioutil"

	"net/http"

	"../Structures"

	"github.com/gorilla/mux"

	"strings"
)

var data Structures.Data
var vectorData []Structures.ScoreCategory
var finder Structures.Search
var yearOrder *Structures.Year

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

	data := []byte(product)
	err := ioutil.WriteFile("VectorDeListas.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "VectorDeListas.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("VectorDeListas.png", cmd, os.FileMode(mode))
}

func Save(w http.ResponseWriter, r *http.Request) {
	var dataAux Structures.Data

	fmt.Fprintln(w, dataAux.ToJson(vectorData))
}

func AddInventory(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var inventory Structures.Inventory

	vectorData = inventory.ReadJson(body, vectorData)

	for i := 0; i < len(vectorData); i++ {
		vectorData[i].PrintInv()
	}

	fmt.Fprintln(w, "Se han añadido los productos correctamente")
}

func getShops(w http.ResponseWriter, r *http.Request) {
	var shops string
	for i := 0; i < len(vectorData); i++ {
		shops = shops + vectorData[i].ToJson()
	}

	fmt.Println(len(shops))

	shops = "[\n" + strings.TrimSuffix(shops, ",\n") + "]"

	fmt.Fprintln(w, shops)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	name, err := mux.Vars(r)["Name"]
	x, err := mux.Vars(r)["Score"]
	score, _ := strconv.Atoi(x)

	var shop *Structures.Shop
	var find bool

	products := "[]"
	if err {
		for i := 0; i < len(vectorData); i++ {
			shop, find = vectorData[i].Search(name, score)
			if find {
				products = "[" + shop.GetProducts() + "]"
				break
			}
		}
	}

	fmt.Fprintln(w, products)
}

func putPurchase(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(body))
	var find bool
	var shop *Structures.Shop
	var car []Structures.CarProduct

	json.Unmarshal(body, &car)

	fmt.Println(car)

	for i := 0; i < len(car); i++ {
		for j := 0; j < len(vectorData); j++ {
			shop, find = vectorData[j].Search(car[i].Shop_.Name, car[i].Shop_.Score)
			if find {
				for k := 0; k < len(car[i].Products); k++ {
					car[i].Products[k].Stock *= -1
				}
				shop.AddProducts(car[i].Products)
			}
		}
	}

}

func addOrders(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var orders Structures.Orders

	json.Unmarshal(body, &orders)

	fmt.Println(orders)

	for i := 0; i < len(orders.Orders); i++ {
		years := strings.Split(orders.Orders[i].Date, "-")
		//primero buscamos si el año ya existe, de lo contrario lo agregamos
		year, err := strconv.Atoi(years[2])
		fmt.Println(year)
		if err == nil {
			anio := yearOrder.SearchYear(year)
			if anio.Value != year {
				yearOrder = yearOrder.Insert(Structures.Year{Value: year})
				anio = yearOrder.SearchYear(year)
			} else {
				fmt.Println("El año ya existe")
			}

			//Ahora que ya existe, ingresamos el mes. Si y solo si aun no existe
			month, err := strconv.Atoi(years[1])
			if err == nil {
				mes := anio.SearchMonth(month)

				if mes.Value != month {
					var newMonth Structures.Calendar
					newMonth.Init(month)
					anio.AddMonth(&newMonth)
					mes = anio.SearchMonth(month)
				} else {
					fmt.Println("El mes ya existe")
				}

				//Ahora añadimos la orden al mes
				day, err := strconv.Atoi(years[0])
				if err == nil {
					mes.AddOrder(orders.Orders[i].Departament, day, &orders.Orders[i])
				}
			}
		}
	}

	Structures.YearInorden(yearOrder)
	fmt.Fprintln(w, "Hola mundo")
}
