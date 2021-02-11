package main

import (
	"io/ioutil"

	"./Structures"

	//	"encoding/json"

	"fmt"
)

func main() {
	var datos Structures.Data
	text, err := ioutil.ReadFile("./Categorias.json")

	if err != nil {
		fmt.Println("ocurrio un error")
	}

	datos.ReadJson(text)

	for i := 0; i < len(datos.GetCells()); i++ {
		fmt.Println(datos.GetCells()[i].GetDepartament())
		fmt.Print("\n\n\n\n")
	}
}
