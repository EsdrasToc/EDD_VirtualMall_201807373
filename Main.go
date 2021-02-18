package main

import (
	"io/ioutil"

	"./Structures"

	"fmt"
)

func main() {
	var datos Structures.Data
	var search Structures.Search

	text, err := ioutil.ReadFile("./Categorias.json")

	if err != nil {
		fmt.Println("ocurrio un error")
	}

	array := datos.ReadJson(text)

	text, err = ioutil.ReadFile("./Buscar.json")

	if err != nil {
		fmt.Println("ocurrio un error")
	}
	search.ReadJson(text)
	search.Delete(&array)
	search.EspecificSearchEngine(array)
}
