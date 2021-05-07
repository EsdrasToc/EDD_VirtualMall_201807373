package Structures

import "container/list"

type Comments struct {
	Comments []list.List
	InUse    int
	Size     int
	Percent  float32
}

func (this *Comments) Add(newComment *Comment) {
	if this.InUse == 0 {
		this.Comments = make([]list.List, 7)
	}

	i := this.HashFunction(newComment)

	this.Comments[i].PushBack(newComment)
}

func (this *Comments) HashFunction(comment *Comment) int {
	var h int
	//h = this.Size * ((int(comment.User.Dpi) * 0.2525) % 1)

	return h
}

type Comment struct {
	User       *Account
	SubComment list.List
	Content    string
}
