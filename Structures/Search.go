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

func (this Search) Delete(array *[100]ScoreCategory) {

	for i := 0; i < len(array); i++ {
		if this.Departament == array[i].Departament {
			deleted := array[i].Delete(this.Name, this.Score)
			fmt.Println(deleted)
			fmt.Println(array[i].first)
			if deleted {
				fmt.Println("Eliminado correctamente")
				return
			}
		}
	}

	fmt.Println("=========== No se encontró ninguna tienda con dichos parametros")
}
