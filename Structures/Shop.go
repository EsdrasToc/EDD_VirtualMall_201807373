package Structures

import (
	"encoding/json"
)

type Shop struct {
	Name        string `json:"Nombre"`
	Description string `json:"Descripcion"`
	Contact     string `json:"Contacto"`
	Score       int    `json:"Calificacion"`
	Previous    *Shop  `json:"-"`
	Next        *Shop  `json:"-"`
	Node        `json:"-"`
}

func (this Shop) ToString() string {
	return "Nombre: " + this.Name + "\nDescripcion: " + this.Description + "\nContacto: " + this.Contact + "\nPunteo: " + string(this.Score)
}

func (this *Shop) SetNext(next Shop) {
	this.Next = &next
}

func (this *Shop) SetPrevious(previous Shop) {
	this.Previous = &previous
}

func (this Shop) ToJSON() string {

	file, _ := json.MarshalIndent(this, "", "\t")

	//_ = ioutil.WriteFile("BusquedaPosicionEspecifica.json", file, 0644)

	return string(file)
}
