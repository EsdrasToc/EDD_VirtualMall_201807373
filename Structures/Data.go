package Structures

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Cells []Cell `json:"Datos"`
}

func (this *Data) ReadJson(text []byte) []ScoreCategory {
	json.Unmarshal(text, &this)

	k := len(this.Cells) * len(this.Cells[0].Departaments)
	/*for i := 0; i < len(this.Cells); i++ {
		for j := 0; j < len(this.Cells[i].Departaments); j++ {
			k++
		}
	}*/

	//fmt.Println(k)
	//var aux [100]ScoreCategory
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

			aux[pos].score = 1
			aux[pos+1].score = 2
			aux[pos+2].score = 3
			aux[pos+3].score = 4
			aux[pos+4].score = 5

			pos += 5

			for l := 0; l < len(this.Cells[i].Departaments[j].Shops); l++ {

				if this.Cells[i].Departaments[j].Shops[l].Score < 6 && this.Cells[i].Departaments[j].Shops[l].Score > 0 {
					score := this.Cells[i].Departaments[j].Shops[l].Score
					aux[(k*5)+score-1].Add(this.Cells[i].Departaments[j].Shops[l])

				}
			}
			k++
		}
	}

	for i := 0; i < len(aux); i++ {
		fmt.Println(aux[i])
	}

	/*k = 0

	for i := 0; i < len(this.Cells); i++ {
		for j := 0; j < len(this.Cells[i].Departaments); j++ {

			for l := 0; l < len(this.Cells[i].Departaments[j].Shops); l++ {

				/*for m := 0; m < 5; m++ {
					aux[(k*5)+m].Departament = this.Cells[i].Departaments[j].Name
					aux[(k*5)+m].Index = this.Cells[i].Index
					aux[(k*5)+m].score = m + 1
				}

				if this.Cells[i].Departaments[j].Shops[l].Score < 6 && this.Cells[i].Departaments[j].Shops[l].Score > 0 {
					score := this.Cells[i].Departaments[j].Shops[l].Score
					/*aux[(k*5)+score-1].score = score
					aux[(k*5)+score-1].Departament = this.Cells[i].Departaments[j].Name
					aux[(k*5)+score-1].Index = this.Cells[i].Index
					aux[(k*5)+score-1].Add(this.Cells[i].Departaments[j].Shops[l])

				}
			}
			k++
		}
	}*/

	return aux
}

func (this Data) GetCells() []Cell {
	return this.Cells
}
