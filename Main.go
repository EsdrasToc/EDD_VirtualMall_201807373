package main

import (
	"io/ioutil"

	"./Structures"

	//	"encoding/json"
	//"./Utilities"

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

	/*for i := 0; i < len(datos.GetCells()); i++ {
		fmt.Println(datos.GetCells()[i].GetDepartament())
		fmt.Print("\n\n\n\n")
	}*/
	/*for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}*/

	text, err = ioutil.ReadFile("./Buscar.json")

	if err != nil {
		fmt.Println("ocurrio un error")
	}
	search.ReadJson(text)
	search.EspecificSearchEngine(array)
}
