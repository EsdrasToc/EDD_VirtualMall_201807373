package Structures

import (
	"encoding/json"
)

type Inventory struct {
	Inventories []IProducts `json:"Inventarios"`
}

type IProducts struct {
	Shop        string    `json:"Tienda"`
	Departament string    `json:"Departamento"`
	Score       int       `json:"Calificacion"`
	Products    []Product `json:"Productos"`
}

func (this *Inventory) ReadJson(text []byte, data []ScoreCategory) []ScoreCategory {
	json.Unmarshal(text, &this)
	shop := &Shop{}
	real := false
	for i := 0; i < len(this.Inventories); i++ {
		for j := 0; j < len(data); j++ {
			if data[j].Departament == this.Inventories[i].Departament && data[j].Score == this.Inventories[i].Score {
				shop, real = data[j].Search(this.Inventories[i].Shop, this.Inventories[i].Score)
				if real {
					shop.AddProducts(this.Inventories[i].Products)
				}
			}
		}

	}

	return data
}
