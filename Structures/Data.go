package Structures

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Cells []Cell `json:"Datos"`
}

func (this *Data) ReadJson(text []byte) {
	json.Unmarshal(text, &this)
	fmt.Print("\n\n\n\n")
}

func (this Data) GetCells() []Cell {
	return this.Cells
}
