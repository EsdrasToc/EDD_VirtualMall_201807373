package Structures

type ScoreCategory struct {
	first       *Shop
	last        *Shop
	Index       string
	Departament string
	score       int
	lenght      int
}

func (this *ScoreCategory) Add(new Shop) {
	if this.first == nil {
		this.first = &new
	}

	if this.last == nil {
		this.last = &new
	} else {
		this.last.SetNext(new)
		new.SetPrevious(*this.last)
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
		if shop.Next != nil {
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
