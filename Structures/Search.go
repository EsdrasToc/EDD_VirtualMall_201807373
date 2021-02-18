package Structures

import (
	"encoding/json"
	"fmt"
	//"Structures"
)

type Search struct {
	Departament string `json:"Departamento"`
	Name        string `json:"Nombre"`
	Score       int    `json:"Calificacion"`
}

func (this *Search) ReadJson(text []byte) {
	json.Unmarshal(text, &this)
}

func (this Search) EspecificSearchEngine(array [100]ScoreCategory) {

	for i := 0; i < len(array); i++ {
		if this.Departament == array[i].Departament {
			object, null := array[i].Search(this.Name, this.Score)

			if null {
				object.ToJSON()
				return
			}
		}
	}

	fmt.Println("No se encontró ninguna tienda con dichos parametros")
}
