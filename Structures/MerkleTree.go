package Structures

import (
	"fmt"
	"strconv"

	"github.com/fernet/fernet-go"
)

type MerkleTreeOrders struct {
	Root *MTreeONode
	N    *int
}

func (this *MerkleTreeOrders) Show() {
	this.Root.Show(0)
}

func (this *MerkleTreeOrders) AddOrder(order *Order) *MerkleTreeOrders {
	if this == nil {
		fmt.Println("Primero ingresó aqui")
		this = &MerkleTreeOrders{}
	}

	if this.Root == nil {
		fmt.Println("Y luego aqui")
		auxHash := strconv.FormatInt(order.User.Dpi, 10) + "," + order.Date
		hash, _ := fernet.EncryptAndSign([]byte(auxHash), &fernet.Key{})
		auxL := &MTreeONode{Leaf: true, Full: true, Hash: string(hash), Order: order}
		auxR := &MTreeONode{Leaf: true, Full: false, Hash: ""}
		this.Root = &MTreeONode{Leaf: false, Full: false, Left: auxL, Right: auxR}
		return this
	}

	if this.Root.Full {
		levels := this.Root.GetLevels()
		newTree := &MTreeONode{Leaf: false, Full: false, Hash: ""}
		newTree.CreateTree(levels)

		aux := this.Root
		this.Root = nil

		this.Root = &MTreeONode{Leaf: false, Full: false, Hash: "", Left: aux, Right: newTree}
	}

	hash, _ := fernet.EncryptAndSign([]byte(strconv.FormatInt(order.User.Dpi, 10)+","+order.Date), &fernet.Key{})
	this.Root = this.Root.AddNode(MTreeONode{Leaf: true, Full: true, Hash: string(hash), Order: order})

	return this
}

type MTreeONode struct {
	Order *Order
	Leaf  bool
	Full  bool
	Hash  string
	Left  *MTreeONode
	Right *MTreeONode
}

func (this *MTreeONode) AddNode(new MTreeONode) *MTreeONode {
	if this.Leaf {
		this = &new
		this.Full = true
		return this
	}

	if this.Left.Full {
		this.Right = this.Right.AddNode(new)
	} else {
		this.Left = this.Left.AddNode(new)
	}

	//encriptar izquierda y derecha
	info := this.Left.Hash + "," + this.Right.Hash
	hash, _ := fernet.EncryptAndSign([]byte(info), &fernet.Key{})
	this.Hash = string(hash)

	if this.Left.Full && this.Right.Full {
		this.Full = true
	}

	return this
}

func (this *MTreeONode) GetLevels() int {
	i := 0
	aux := this
	for aux != nil {
		i++
		aux = aux.Left
	}

	return i
}

func (this *MTreeONode) CreateTree(level int) {
	if level == 2 {
		this.Left = &MTreeONode{Leaf: true, Full: false, Hash: ""}
		this.Right = &MTreeONode{Leaf: true, Full: false, Hash: ""}
		return
	} else {
		this.Right = &MTreeONode{Leaf: false, Full: false, Hash: ""}
		this.Left = &MTreeONode{Leaf: false, Full: false, Hash: ""}
	}

	level--

	this.Left.CreateTree(level)
	this.Right.CreateTree(level)
}

func (this *MTreeONode) Show(h int) {
	fmt.Println("Level " + strconv.Itoa(h) + ": ")

	fmt.Print("[ ")
	fmt.Print(this.Hash)
	fmt.Print(" ")

	fmt.Println(" ]")

	if this.Left != nil {
		this.Left.Show(h + 1)
	}

	if this.Right != nil {
		this.Right.Show(h + 1)
	}

}
