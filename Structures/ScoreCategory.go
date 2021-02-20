package Structures

import (
	"fmt"
	"strconv"
)

type ScoreCategory struct {
	first       *Shop
	last        *Shop
	Index       string
	Departament string
	Score       int
	Lenght      int
}

func (this *ScoreCategory) Add(new Shop) {
	last := this.last
	if this.first == nil {
		this.first = &new
	}

	if this.last == nil {
		this.last = &new
	} else {
		this.last.SetNext(new)
		new.SetPrevious(*last)
		this.last = &new
	}

	this.Lenght++
}

func (this *ScoreCategory) Search(name string, score int) (Shop, bool) {
	cicle := true
	shop := this.first

	if this.first == nil {
		return Shop{}, false
	}
	for cicle {
		if shop != nil {
			if shop.Score == score && shop.Name == name {
				return *shop, true
			} else {
				shop = shop.Next
			}
		} else {
			break
		}
	}
	return Shop{}, false
}

func (this *ScoreCategory) Delete(name string, score int) bool {
	shop := this.first

	if this.first == nil {
		return false
	}

	if shop.Score == score && shop.Name == name {
		this.first = this.first.Next
		this.Lenght--
		return true
	} else if this.last.Score == score && this.last.Name == name {
		fmt.Println(" o Entro aqui")
		this.last = this.last.Previous
		this.last.Next = nil
		this.Lenght--
		return true
	} else {
		for shop != nil {
			if shop.Score == score && shop.Name == name {
				shop.Previous.SetNext(*shop.Next)
				shop.Next.SetPrevious(*shop.Previous)
				shop = nil
				this.Lenght--
				return true
			}
			shop = shop.Next
		}
	}

	return false
}

func (this *ScoreCategory) Order() {
	var i, j *Shop
	i = this.first
	for i.Next != nil {
		j = i.Next
		for j != nil {
			if i.Name > j.Name {
				Swap(i, j)
			}
			j = j.Next
		}
		i = i.Next
	}
}

func Swap(a *Shop, b *Shop) {
	var aux *Shop

	aux = a

	a.Name = b.Name
	a.Description = b.Description
	a.Contact = b.Contact

	b.Name = aux.Name
	b.Description = aux.Description
	b.Contact = aux.Contact
}

func (this ScoreCategory) ToJson() string {
	aux := this.first
	content := ""

	for aux != nil {

		content = content + aux.ToJSON()

		aux = aux.Next
	}

	return ""
}

func (this ScoreCategory) ToGraph(i *int) (string, string) {
	aux := this.first
	content := ""
	edges := ""
	for aux != nil {
		*i++
		if aux == this.first {
			content = content + "node" + strconv.Itoa(*i) + "[label=\"" + aux.Name + "\"]\n"
			//edges = edges + "node" + strconv.Itoa(*i-1) + "->" + "node" + strconv.Itoa(*i-1) + "\n"
			fmt.Println("node" + strconv.Itoa(*i-1) + "->" + "node" + strconv.Itoa(*i-1) + "\n")
		} else {
			content = content + "node" + strconv.Itoa(*i) + "[label=\"" + aux.Name + "\"]\n"
			edges = edges + "node" + strconv.Itoa(*i) + "->" + "node" + strconv.Itoa(*i-1) + "\n"
			edges = edges + "node" + strconv.Itoa(*i-1) + "->" + "node" + strconv.Itoa(*i) + "\n"
		}
		aux = aux.Next
	}

	return content, edges
}
