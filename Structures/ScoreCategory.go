package Structures

import "fmt"

type ScoreCategory struct {
	first       *Shop
	last        *Shop
	Index       string
	Departament string
	score       int
	lenght      int
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

	this.lenght++
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
		//fmt.Println("Entró aqui")
		return false
	}

	if shop.Score == score && shop.Name == name {
		fmt.Println(this.first)
		this.first = this.first.Next
		fmt.Println(this.first)
		this.lenght--
		return true
	} else if this.last.Score == score && this.last.Name == name {
		fmt.Println(" o Entro aqui")
		this.last = this.last.Previous
		this.last.Next = nil
		this.lenght--
		return true
	} else {
		for shop != nil {
			if shop.Score == score && shop.Name == name {
				fmt.Println(shop.Next)
				fmt.Println(shop.Previous)
				shop.Previous.SetNext(*shop.Next)
				shop.Next.SetPrevious(*shop.Previous)
				shop = nil
				//fmt.Println("Entró aqui tambien")
				this.lenght--
				return true
			}
			shop = shop.Next
		}
	}

	return false
}
