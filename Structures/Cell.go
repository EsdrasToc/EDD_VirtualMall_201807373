package Structures

type Cell struct {
	Index        string        `json:"Indice"`
	Departaments []Departament `json:"Departamentos"`
}

func (this Cell) GetDepartament() []Departament {
	return this.Departaments
}
