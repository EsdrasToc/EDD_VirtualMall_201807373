package Structures

import (
	"fmt"
	"strconv"

	"crypto/sha256"
)

/*================================*/
/* ARBOLES DE MERKLE PARA ORDENES */
/*================================*/

type MerkleTreeOrders struct {
	Root *MTreeONode
	N    *int
}

func (this *MerkleTreeOrders) Show() {
	this.Root.Show(0)
}

func (this *MerkleTreeOrders) AddOrder(order *Order) *MerkleTreeOrders {
	if this == nil {
		//fmt.Println("Primero ingresó aqui")
		this = &MerkleTreeOrders{}
	}

	if this.Root == nil {
		//fmt.Println("Y luego aqui")
		h := sha256.New()
		auxHash := strconv.FormatInt(order.User.Dpi, 10) + "," + order.Date
		//hash, _ := fernet.EncryptAndSign([]byte(auxHash), &fernet.Key{})
		h.Write([]byte(auxHash))
		hash := h.Sum(nil)
		auxL := &MTreeONode{Leaf: true, Full: true, Hash: hash, Order: order}
		auxR := &MTreeONode{Leaf: true, Full: false}

		this.Root = &MTreeONode{Leaf: false, Full: false, Left: auxL, Right: auxR}
		info := string(this.Root.Left.Hash) + string(this.Root.Right.Hash)
		h.Write([]byte(info))
		//hash, _ := fernet.EncryptAndSign([]byte(info), &fernet.Key{})
		hash = h.Sum(nil)
		this.Root.Hash = hash
		return this
	}

	if this.Root.Full {
		levels := this.Root.GetLevels()
		newTree := &MTreeONode{Leaf: false, Full: false}
		newTree.CreateTree(levels)

		aux := this.Root
		this.Root = nil

		this.Root = &MTreeONode{Leaf: false, Full: false, Left: aux, Right: newTree}
	}

	h := sha256.New()
	h.Write([]byte(strconv.FormatInt(order.User.Dpi, 10) + "," + order.Date))
	//hash, _ := fernet.EncryptAndSign([]byte(strconv.FormatInt(order.User.Dpi, 10)+","+order.Date), &fernet.Key{})
	hash := h.Sum(nil)
	this.Root.AddNode(MTreeONode{Leaf: true, Full: true, Hash: hash, Order: order})

	return this
}

type MTreeONode struct {
	Order *Order
	Leaf  bool
	Full  bool
	Hash  []byte
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
	if this.Left != nil && this.Right != nil {
		info := string(this.Left.Hash) + string(this.Right.Hash)
		h := sha256.New()
		h.Write([]byte(info))
		//hash, _ := fernet.EncryptAndSign([]byte(info), &fernet.Key{})
		hash := h.Sum(nil)
		this.Hash = hash
	}

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
		this.Left = &MTreeONode{Leaf: true, Full: false}
		this.Right = &MTreeONode{Leaf: true, Full: false}
		return
	} else {
		this.Right = &MTreeONode{Leaf: false, Full: false}
		this.Left = &MTreeONode{Leaf: false, Full: false}
	}

	level--

	this.Left.CreateTree(level)
	this.Right.CreateTree(level)
}

func (this *MTreeONode) Show(h int) {
	fmt.Println("Level " + strconv.Itoa(h) + ": ")

	fmt.Print("[ ")
	fmt.Print(this.Hash)
	fmt.Println(len(this.Hash))
	fmt.Print(" ")

	fmt.Println(" ]")

	if this.Left != nil {
		this.Left.Show(h + 1)
	}

	if this.Right != nil {
		this.Right.Show(h + 1)
	}

}

/*================================*/
/* ARBOLES DE MERKLE PARA TIENDAS */
/*================================*/

type MerkleTreeShops struct {
	Root *MTreeSNode
	N    *int
}

func (this *MerkleTreeShops) Show() {
	this.Root.Show(0)
}

func (this *MerkleTreeShops) AddShop(shop *Shop) *MerkleTreeShops {
	if this == nil {
		//fmt.Println("Primero ingresó aqui")
		this = &MerkleTreeShops{}
	}

	if this.Root == nil {
		//fmt.Println("Y luego aqui")
		h := sha256.New()
		auxHash := strconv.Itoa(shop.Score) + "," + shop.Name + "," + shop.Description
		//hash, _ := fernet.EncryptAndSign([]byte(auxHash), &fernet.Key{})
		h.Write([]byte(auxHash))
		hash := h.Sum(nil)
		auxL := &MTreeSNode{Leaf: true, Full: true, Hash: hash, Shop: shop}
		auxR := &MTreeSNode{Leaf: true, Full: false}

		this.Root = &MTreeSNode{Leaf: false, Full: false, Left: auxL, Right: auxR}
		info := string(this.Root.Left.Hash) + string(this.Root.Right.Hash)
		h.Write([]byte(info))
		//hash, _ := fernet.EncryptAndSign([]byte(info), &fernet.Key{})
		hash = h.Sum(nil)
		this.Root.Hash = hash
		return this
	}

	if this.Root.Full {
		levels := this.Root.GetLevels()
		newTree := &MTreeSNode{Leaf: false, Full: false}
		newTree.CreateTree(levels)

		aux := this.Root
		this.Root = nil

		this.Root = &MTreeSNode{Leaf: false, Full: false, Left: aux, Right: newTree}
	}

	h := sha256.New()
	h.Write([]byte(strconv.Itoa(shop.Score) + "," + shop.Name + "," + shop.Description))
	//hash, _ := fernet.EncryptAndSign([]byte(strconv.FormatInt(order.User.Dpi, 10)+","+order.Date), &fernet.Key{})
	hash := h.Sum(nil)
	this.Root.AddNode(MTreeSNode{Leaf: true, Full: true, Hash: hash, Shop: shop})

	return this
}

type MTreeSNode struct {
	Shop  *Shop
	Leaf  bool
	Full  bool
	Hash  []byte
	Left  *MTreeSNode
	Right *MTreeSNode
}

func (this *MTreeSNode) AddNode(new MTreeSNode) *MTreeSNode {
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
	if this.Left != nil && this.Right != nil {
		info := string(this.Left.Hash) + string(this.Right.Hash)
		h := sha256.New()
		h.Write([]byte(info))
		//hash, _ := fernet.EncryptAndSign([]byte(info), &fernet.Key{})
		hash := h.Sum(nil)
		this.Hash = hash
	}

	if this.Left.Full && this.Right.Full {
		this.Full = true
	}

	return this
}

func (this *MTreeSNode) GetLevels() int {
	i := 0
	aux := this
	for aux != nil {
		i++
		aux = aux.Left
	}

	return i
}

func (this *MTreeSNode) CreateTree(level int) {
	if level == 2 {
		this.Left = &MTreeSNode{Leaf: true, Full: false}
		this.Right = &MTreeSNode{Leaf: true, Full: false}
		return
	} else {
		this.Right = &MTreeSNode{Leaf: false, Full: false}
		this.Left = &MTreeSNode{Leaf: false, Full: false}
	}

	level--

	this.Left.CreateTree(level)
	this.Right.CreateTree(level)
}

func (this *MTreeSNode) Show(h int) {
	fmt.Println("Level " + strconv.Itoa(h) + ": ")

	fmt.Print("[ ")
	fmt.Print(this.Hash)
	fmt.Println(len(this.Hash))
	fmt.Print(" ")

	fmt.Println(" ]")

	if this.Left != nil {
		this.Left.Show(h + 1)
	}

	if this.Right != nil {
		this.Right.Show(h + 1)
	}

}

/*==================================*/
/* ARBOLES DE MERKLE PARA PRODCUTOS */
/*==================================*/

type MerkleTreeProducts struct {
	Root *MTreePNode
	N    *int
}

func (this *MerkleTreeProducts) Show() {
	this.Root.Show(0)
}

func (this *MerkleTreeProducts) AddProduct(product *Product) *MerkleTreeProducts {
	if this == nil {
		//fmt.Println("Primero ingresó aqui")
		this = &MerkleTreeProducts{}
	}

	if this.Root == nil {
		//fmt.Println("Y luego aqui")
		h := sha256.New()
		auxHash := product.Name + "," + strconv.Itoa(product.Code) + "," + strconv.Itoa(product.Stock)
		//hash, _ := fernet.EncryptAndSign([]byte(auxHash), &fernet.Key{})
		h.Write([]byte(auxHash))
		hash := h.Sum(nil)
		auxL := &MTreePNode{Leaf: true, Full: true, Hash: hash, Product: product}
		auxR := &MTreePNode{Leaf: true, Full: false}

		this.Root = &MTreePNode{Leaf: false, Full: false, Left: auxL, Right: auxR}
		info := string(this.Root.Left.Hash) + string(this.Root.Right.Hash)
		h.Write([]byte(info))
		//hash, _ := fernet.EncryptAndSign([]byte(info), &fernet.Key{})
		hash = h.Sum(nil)
		this.Root.Hash = hash
		return this
	}

	if this.Root.Full {
		levels := this.Root.GetLevels()
		newTree := &MTreePNode{Leaf: false, Full: false}
		newTree.CreateTree(levels)

		aux := this.Root
		this.Root = nil

		this.Root = &MTreePNode{Leaf: false, Full: false, Left: aux, Right: newTree}
	}

	h := sha256.New()
	h.Write([]byte(product.Name + "," + strconv.Itoa(product.Code) + "," + strconv.Itoa(product.Stock)))
	//hash, _ := fernet.EncryptAndSign([]byte(strconv.FormatInt(order.User.Dpi, 10)+","+order.Date), &fernet.Key{})
	hash := h.Sum(nil)
	this.Root.AddNode(MTreePNode{Leaf: true, Full: true, Hash: hash, Product: product})

	return this
}

type MTreePNode struct {
	Product *Product
	Leaf    bool
	Full    bool
	Hash    []byte
	Left    *MTreePNode
	Right   *MTreePNode
}

func (this *MTreePNode) AddNode(new MTreePNode) *MTreePNode {
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
	if this.Left != nil && this.Right != nil {
		info := string(this.Left.Hash) + string(this.Right.Hash)
		h := sha256.New()
		h.Write([]byte(info))
		//hash, _ := fernet.EncryptAndSign([]byte(info), &fernet.Key{})
		hash := h.Sum(nil)
		this.Hash = hash
	}

	if this.Left.Full && this.Right.Full {
		this.Full = true
	}

	return this
}

func (this *MTreePNode) GetLevels() int {
	i := 0
	aux := this
	for aux != nil {
		i++
		aux = aux.Left
	}

	return i
}

func (this *MTreePNode) CreateTree(level int) {
	if level == 2 {
		this.Left = &MTreePNode{Leaf: true, Full: false}
		this.Right = &MTreePNode{Leaf: true, Full: false}
		return
	} else {
		this.Right = &MTreePNode{Leaf: false, Full: false}
		this.Left = &MTreePNode{Leaf: false, Full: false}
	}

	level--

	this.Left.CreateTree(level)
	this.Right.CreateTree(level)
}

func (this *MTreePNode) Show(h int) {
	fmt.Println("Level " + strconv.Itoa(h) + ": ")

	fmt.Print("[ ")
	fmt.Print(this.Hash)
	fmt.Println(len(this.Hash))
	fmt.Print(" ")

	fmt.Println(" ]")

	if this.Left != nil {
		this.Left.Show(h + 1)
	}

	if this.Right != nil {
		this.Right.Show(h + 1)
	}

}

/*=================================*/
/* ARBOLES DE MERKLE PARA USUARIOS */
/*=================================*/

type MerkleTreeUsers struct {
	Root *MTreeUNode
	N    *int
}

func (this *MerkleTreeUsers) Show() {
	this.Root.Show(0)
}

func (this *MerkleTreeUsers) AddUser(user *Account) *MerkleTreeUsers {
	if this == nil {
		//fmt.Println("Primero ingresó aqui")
		this = &MerkleTreeUsers{}
	}

	if this.Root == nil {
		//fmt.Println("Y luego aqui")
		h := sha256.New()
		auxHash := strconv.FormatInt(user.Dpi, 10) + "," + user.User + "," + user.Email + ","
		//auxHash := product.Name + "," + strconv.Itoa(product.Code) + "," + strconv.Itoa(product.Stock)
		//hash, _ := fernet.EncryptAndSign([]byte(auxHash), &fernet.Key{})
		h.Write([]byte(auxHash))
		hash := h.Sum(nil)
		auxL := &MTreeUNode{Leaf: true, Full: true, Hash: hash, User: user}
		auxR := &MTreeUNode{Leaf: true, Full: false}

		this.Root = &MTreeUNode{Leaf: false, Full: false, Left: auxL, Right: auxR}
		info := string(this.Root.Left.Hash) + string(this.Root.Right.Hash)
		h.Write([]byte(info))
		//hash, _ := fernet.EncryptAndSign([]byte(info), &fernet.Key{})
		hash = h.Sum(nil)
		this.Root.Hash = hash
		return this
	}

	if this.Root.Full {
		levels := this.Root.GetLevels()
		newTree := &MTreeUNode{Leaf: false, Full: false}
		newTree.CreateTree(levels)

		aux := this.Root
		this.Root = nil

		this.Root = &MTreeUNode{Leaf: false, Full: false, Left: aux, Right: newTree}
	}

	h := sha256.New()
	h.Write([]byte(strconv.FormatInt(user.Dpi, 10) + "," + user.User + "," + user.Email + ","))
	//hash, _ := fernet.EncryptAndSign([]byte(strconv.FormatInt(order.User.Dpi, 10)+","+order.Date), &fernet.Key{})
	hash := h.Sum(nil)
	this.Root.AddNode(MTreeUNode{Leaf: true, Full: true, Hash: hash, User: user})

	return this
}

type MTreeUNode struct {
	User  *Account
	Leaf  bool
	Full  bool
	Hash  []byte
	Left  *MTreeUNode
	Right *MTreeUNode
}

func (this *MTreeUNode) AddNode(new MTreeUNode) *MTreeUNode {
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
	if this.Left != nil && this.Right != nil {
		info := string(this.Left.Hash) + string(this.Right.Hash)
		h := sha256.New()
		h.Write([]byte(info))
		//hash, _ := fernet.EncryptAndSign([]byte(info), &fernet.Key{})
		hash := h.Sum(nil)
		this.Hash = hash
	}

	if this.Left.Full && this.Right.Full {
		this.Full = true
	}

	return this
}

func (this *MTreeUNode) GetLevels() int {
	i := 0
	aux := this
	for aux != nil {
		i++
		aux = aux.Left
	}

	return i
}

func (this *MTreeUNode) CreateTree(level int) {
	if level == 2 {
		this.Left = &MTreeUNode{Leaf: true, Full: false}
		this.Right = &MTreeUNode{Leaf: true, Full: false}
		return
	} else {
		this.Right = &MTreeUNode{Leaf: false, Full: false}
		this.Left = &MTreeUNode{Leaf: false, Full: false}
	}

	level--

	this.Left.CreateTree(level)
	this.Right.CreateTree(level)
}

func (this *MTreeUNode) Show(h int) {
	fmt.Println("Level " + strconv.Itoa(h) + ": ")

	fmt.Print("[ ")
	fmt.Print(this.Hash)
	fmt.Println(len(this.Hash))
	fmt.Print(" ")

	fmt.Println(" ]")

	if this.Left != nil {
		this.Left.Show(h + 1)
	}

	if this.Right != nil {
		this.Right.Show(h + 1)
	}

}
