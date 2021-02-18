package Structures

import (
	"encoding/json"
)

type Data struct {
	Cells []Cell `json:"Datos"`
}

func (this *Data) ReadJson(text []byte) [100]ScoreCategory {
	json.Unmarshal(text, &this)

	k := 0
	for i := 0; i < len(this.Cells); i++ {
		for j := 0; j < len(this.Cells[i].Departaments); j++ {
			k++
		}
	}

	var aux [100]ScoreCategory

	k = 0

	for i := 0; i < len(this.Cells); i++ {
		for j := 0; j < len(this.Cells[i].Departaments); j++ {

			for l := 0; l < len(this.Cells[i].Departaments[j].Shops); l++ {

				for m := 0; m < 5; m++ {
					aux[(k*5)+m].Departament = this.Cells[i].Departaments[j].Name
					aux[(k*5)+m].Index = this.Cells[i].Index
					aux[(k*5)+m].score = m + 1
				}

				if this.Cells[i].Departaments[j].Shops[l].Score < 6 && this.Cells[i].Departaments[j].Shops[l].Score > 0 {
					score := this.Cells[i].Departaments[j].Shops[l].Score
					/*aux[(k*5)+score-1].score = score
					aux[(k*5)+score-1].Departament = this.Cells[i].Departaments[j].Name
					aux[(k*5)+score-1].Index = this.Cells[i].Index*/
					aux[(k*5)+score-1].Add(this.Cells[i].Departaments[j].Shops[l])

				}
			}
			k++
		}
	}

	return aux
}

func (this Data) GetCells() []Cell {
	return this.Cells
}
