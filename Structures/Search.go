package Structures

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Search struct {
	Departament string `json:"Departamento"`
	Name        string `json:"Nombre"`
	Score       int    `json:"Calificacion"`
}

func (this *Search) ReadJson(text []byte) {
	json.Unmarshal(text, &this)
}

func (this Search) EspecificSearchEngine(array []ScoreCategory) string {

	for i := 0; i < len(array); i++ {
		if this.Departament == array[i].Departament {
			object, null := array[i].Search(this.Name, this.Score)

			if null {
				return object.ToJSON()
			}
		}
	}

	return "No se encontró ninguna tienda con dichos parametros"
}

func (this Search) Delete(array []ScoreCategory, w http.ResponseWriter) *[]ScoreCategory {

	for i := 0; i < len(array); i++ {
		if this.Departament == array[i].Departament {
			deleted := array[i].Delete(this.Name, this.Score)
			if deleted {
				fmt.Fprintf(w, "Eliminado correctamente")
				return &array
			}
		}
	}

	fmt.Fprintf(w, "No se encontró ninguna tienda con dichos parametros")
	return &array
}
