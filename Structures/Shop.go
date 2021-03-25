package Structures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Shop struct {
	Name        string   `json:"Nombre"`
	Description string   `json:"Descripcion"`
	Contact     string   `json:"Contacto"`
	Score       int      `json:"Calificacion"`
	Logo        string   `json:"Logo"`
	Inventory   *Product `json:"-"`
	Previous    *Shop    `json:"-"`
	Next        *Shop    `json:"-"`
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

	_ = ioutil.WriteFile("BusquedaPosicionEspecifica.json", file, 0644)

	return string(file)
}

func (this Shop) ToJSONRequest() string {

	file, _ := json.MarshalIndent(this, "", "\t")

	return string(file)
}

func (this Shop) GetProducts() string {
	return this.Inventory.GetProducts()
}

func (this *Shop) AddProducts(products []Product) {
	//aux := Product{}
	for i := 0; i < len(products); i++ {
		/*if this.Inventory == nil {
			fmt.Println("ALJKSDFLKJASHDF")
			fmt.Println("Producto: ")
			fmt.Println(products[i])
			this.Inventory = &products[i]
		} else {*/
		fmt.Println("Producto: ")
		fmt.Println(products[i])
		this.Inventory = this.Inventory.Insert(products[i])

	}
}
