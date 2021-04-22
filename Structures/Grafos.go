package Structures

import "encoding/json"

type Grafo struct {
	Titles []string
	Range  int
	L      [][]int
	C      []int
	D      []int
	Trange int
}

type ReadGrafo struct {
	Nodes []NodeGrafo `json:"Nodos"`
}

func (this *ReadGrafo) ToJSON(content []byte) *ReadGrafo {
	json.Unmarshal(content, &this)

	return this
}

type NodeGrafo struct {
	Name  string `json:"Nombre"`
	Links []Link `json:"Enlaces"`
}

type Link struct {
	Name     string `json:"Nombre"`
	Distance int    `json:"Distancia"`
}
