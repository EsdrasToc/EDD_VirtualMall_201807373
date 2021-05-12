package Structures

type Comments struct {
	Comments []*Comment
	InUse    int
	Size     int
	Percent  float32
}

func (this *Comments) ToJson() string {
	if this == nil {
		return ""
	}
	var text, text2 string
	var counter, counter2 int
	for i := 0; i < len(this.Comments); i++ {
		if this.Comments[i] != nil {
			counter++
		}
	}

	for i := 0; i < len(this.Comments); i++ {
		text2 = this.Comments[i].ToJson()

		if text2 != "" {
			counter2++
		}

		if this.Comments[i] == nil {
			continue
		}

		if counter2 == counter {
			//text = text + "{\n\"Comentarios\":[\n" + text2 + "\n]\n}\n"
			text = text + text2
		} else {
			//text = text + "{\n\"Comentarios\":[\n" + text2 + "\n]\n}\n,\n"
			text = text + text2
		}
	}

	return text
}

func (this *Comments) Init(i int) *Comments {
	//this.Comments = make([]*Comment, i)
	//this.Size = i

	return &Comments{Comments: make([]*Comment, i), Size: i}
}

func (this *Comments) AddComment(newComment *Comment) *Comments {
	if this == nil {
		this = this.Init(7)
	}

	i := this.HashFunction(newComment)

	this.Comments[i] = this.Comments[i].AddComment(newComment)

	if this.CalcPercent() >= 0.6 {
		var newTable *Comments

		newTable = newTable.Init(SearchNumber(this.Size))
		newTable.ReHash(this)
	}

	return this
}

func (this *Comments) ReHash(old *Comments) {
	salir := false
	for i := 0; i < len(old.Comments); i++ {
		salir = false
		for !salir {
			aux := old.Comments[i].GetFirst()

			if aux != nil {
				this.AddComment(aux)
			} else {
				salir = true
			}
		}

	}
}

func (this *Comments) CalcPercent() float32 {
	counter := 0
	for i := 0; i < len(this.Comments); i++ {
		if this.Comments[i] != nil {
			counter++
		}
	}

	var percent float32

	percent = float32(counter / this.Size)

	return percent
}

func (this *Comments) AddSubComment(newSComment *Comment, comment *Comment) {
	i := this.HashFunction(comment)

	aux := this.Comments[i].SearchComment(comment)

	aux.SubComment.AddComment(newSComment)

}

func (this *Comments) AddSubSubComment(newSSComment *Comment, SComment *Comment, comment *Comment) {
	i := this.HashFunction(comment)

	aux := this.Comments[i].SearchComment(comment)
	aux2 := aux.SubComment.SearchComment(SComment)

	aux2.SubComment.AddComment(newSSComment)

}

/*func (this *Comments) SearchComment(search *Comment) *Comment {

	id := this.HashFunction(search)

}*/

func (this *Comments) HashFunction(comment *Comment) int {
	var h int
	//h = int(this.Size * ((int(comment.User.Dpi) * 0.2525) % 1))
	h = int(comment.User.Dpi) % this.Size
	return h
}

type Comment struct {
	User       *Account `json:"Usuario"`
	Next       *Comment `json:"-"`
	SubComment *Comment `json:"-"`
	Content    string   `json:"Contenido"`
}

func (this *Comment) ToJson() string {

	if this == nil {
		return ""
	}

	text := ""

	aux := this

	for aux.Next != nil {
		text = text + "{\n\"Usuario\":" + this.User.AccountToJson() + ",\n\"SubComentarios\": [\n" + this.SubComment.ToJson() + "\n],\n\"Contenido\":\"" + this.Content + "\"\n},"
		aux = aux.Next
	}

	text = text + "{\n\"Usuario\":" + this.User.AccountToJson() + ",\n\"SubComentarios\": [\n" + this.SubComment.ToJson() + "\n],\n\"Contenido\":\"" + this.Content + "\"\n}"

	return text
}

func (this *Comment) GetFirst() *Comment {
	var temp, aux *Comment

	aux = this

	if aux == nil {
		return aux
	}

	temp = this.Next
	this = nil
	this = temp
	aux.Next = nil

	return aux

}

func (this *Comment) AddComment(comment *Comment) *Comment {
	aux := this

	if this == nil {
		this = comment
		return this
	} else {
		for aux != nil {
			if aux.Next == nil {
				aux.Next = comment
				return this
			}

			aux = aux.Next
		}
	}

	return this
}

func (this *Comment) SearchComment(comment *Comment) *Comment {
	aux := this

	for aux != nil {

		if aux.Content == comment.Content && aux.User.Dpi == aux.User.Dpi {
			break
		}

		aux = aux.Next
	}

	return aux
}

func SearchNumber(actual int) int {
	var counter int
	find := false

	for !find {
		actual++
		counter = 0
		for i := 1; i <= actual; i++ {
			if actual%i == 0 {
				counter++
			}
		}

		if counter == 2 {
			find = true
		}
	}

	return actual
}
