package Structures

import (
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	Cells []Cell `json:"Datos"`
}

func (this *Data) ReadJson(text []byte) []ScoreCategory {
	json.Unmarshal(text, &this)

	k := len(this.Cells) * len(this.Cells[0].Departaments)

	aux := make([]ScoreCategory, k*5)

	pos := 0
	k = 0

	for i := 0; i < len(this.Cells); i++ {
		for j := 0; j < len(this.Cells[i].Departaments); j++ {
			aux[pos].Departament = this.Cells[i].Departaments[j].Name
			aux[pos+1].Departament = this.Cells[i].Departaments[j].Name
			aux[pos+2].Departament = this.Cells[i].Departaments[j].Name
			aux[pos+3].Departament = this.Cells[i].Departaments[j].Name
			aux[pos+4].Departament = this.Cells[i].Departaments[j].Name

			aux[pos].Index = this.Cells[i].Index
			aux[pos+1].Index = this.Cells[i].Index
			aux[pos+2].Index = this.Cells[i].Index
			aux[pos+3].Index = this.Cells[i].Index
			aux[pos+4].Index = this.Cells[i].Index

			aux[pos].Score = 1
			aux[pos+1].Score = 2
			aux[pos+2].Score = 3
			aux[pos+3].Score = 4
			aux[pos+4].Score = 5

			pos += 5

			for l := 0; l < len(this.Cells[i].Departaments[j].Shops); l++ {

				if this.Cells[i].Departaments[j].Shops[l].Score < 6 && this.Cells[i].Departaments[j].Shops[l].Score > 0 {
					score := this.Cells[i].Departaments[j].Shops[l].Score
					aux[(k*5)+score-1].Add(this.Cells[i].Departaments[j].Shops[l])

				}
			}
			k++

			//aux[i].Order()
		}
	}

	return aux
}

func (this Data) GetCells() []Cell {
	return this.Cells
}

func (this *Data) ToJson(array []ScoreCategory) string {
	index := array[0].Index
	numIndex := 1
	for i := 0; i < len(array); i += 5 {
		if index != array[i].Index {
			index = array[i].Index
			numIndex++
		}
	}

	numDepartaments := (len(array) / 5) / numIndex
	this.Cells = make([]Cell, numIndex)

	for i := 0; i < len(this.Cells); i++ {
		this.Cells[i].Departaments = make([]Departament, numDepartaments)
	}

	pos := 0
	numShops := 0
	for i := 0; i < len(this.Cells); i++ {
		for j := 0; j < len(this.Cells[i].Departaments); j++ {
			this.Cells[i].Index = array[pos].Index
			this.Cells[i].Departaments[j].Name = array[pos].Departament
			numShops = array[pos].Lenght + array[pos+1].Lenght + array[pos+2].Lenght + array[pos+3].Lenght + array[pos+4].Lenght

			this.Cells[i].Departaments[j].Shops = make([]Shop, numShops)

			l := 0
			for k := 0; k < 5; k++ {
				aux := array[pos+k].first
				for aux != nil {
					this.Cells[i].Departaments[j].Shops[l] = *aux
					l++
					aux = aux.Next
				}
			}
			pos += 5
		}
	}

	json, _ := json.MarshalIndent(this, "", "\t")
	_ = ioutil.WriteFile("DatosGuardados.json", json, 0644)
	return string(json)
}

func (this Data) GetDepartamentShop(name string, score int) string {
	found := false
	departament := ""
	for i := 0; i < len(this.Cells); i++ {
		if string(name[0]) == this.Cells[i].Index {
			for j := 0; j < len(this.Cells[i].Departaments); j++ {
				for k := 0; k < len(this.Cells[i].Departaments[j].Shops); k++ {
					if this.Cells[i].Departaments[j].Shops[k].Name == name {
						departament = this.Cells[i].Departaments[j].Name
						found = true
						break
					}
				}

				if found {
					break
				}
			}
		}

		if found {
			break
		}
	}

	return departament
}
