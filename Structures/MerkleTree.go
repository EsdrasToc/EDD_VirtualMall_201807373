package Structures

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	"crypto/sha256"
)

/*================================*/
/* ARBOLES DE MERKLE PARA ORDENES */
/*================================*/

type MerkleTreeOrders struct {
	Root *MTreeONode
	N    int
}

func (this *MerkleTreeOrders) GetHash() string {
	if this == nil {
		return ""
	}
	return this.Root.GetHash()
}

func (this *MerkleTreeShops) GetHash() string {
	if this == nil {
		return ""
	}
	return this.Root.GetHash()
}

func (this *MerkleTreeUsers) GetHash() string {
	if this == nil {
		return ""
	}
	return this.Root.GetHash()
}

func (this *MerkleTreeProducts) GetHash() string {
	if this == nil {
		return ""
	}
	return this.Root.GetHash()
}

func (this *MerkleTreeOrders) Show() {
	this.Root.Show(0)
}

func (this *MerkleTreeOrders) AddOrder(order *Order) *MerkleTreeOrders {
	if this == nil {
		this = &MerkleTreeOrders{}
	}

	if this.Root == nil {
		this.N = 3
		h := sha256.New()
		auxHash := strconv.FormatInt(order.User.Dpi, 10) + "," + order.Date
		h.Write([]byte(auxHash))
		hash := h.Sum(nil)
		auxL := &MTreeONode{Leaf: true, Full: true, Hash: hash, Order: order, id: 1}
		auxR := &MTreeONode{Leaf: true, Full: false, id: 2}

		this.Root = &MTreeONode{Leaf: false, Full: false, Left: auxL, Right: auxR, id: 0}
		info := this.Root.Left.GetHash() + this.Root.Right.GetHash()
		h.Write([]byte(info))
		hash = h.Sum(nil)
		this.Root.Hash = hash
		return this
	}

	if this.Root.Full {
		levels := this.Root.GetLevels()
		this.N++
		newTree := &MTreeONode{Leaf: false, Full: false, id: this.N}
		this.N += 3
		newTree.CreateTree(levels, this)

		aux := this.Root
		this.Root = nil

		this.N += 3
		this.Root = &MTreeONode{Leaf: false, Full: false, Left: aux, Right: newTree, id: this.N}
	}

	h := sha256.New()
	h.Write([]byte(strconv.FormatInt(order.User.Dpi, 10) + "," + order.Date))
	hash := h.Sum(nil)
	this.N += 3
	this.Root.AddNode(MTreeONode{Leaf: true, Full: true, Hash: hash, Order: order, id: this.N})
	this.N += 3

	return this
}

type MTreeONode struct {
	Order *Order
	Leaf  bool
	Full  bool
	Hash  []byte
	Left  *MTreeONode
	Right *MTreeONode
	id    int
}

func (this *MerkleTreeOrders) Graph() {
	var text string
	if this == nil {
		text = "digraph G{\n}"
	} else {
		this.N = 0
		nodes := this.Root.GraphNodes()
		this.N = 0
		lines := this.Root.GraphLines()
		//lines := ""

		text = "digraph G{\n" + nodes + "\n" + lines + "\n}"
	}

	data := []byte(text)
	err := ioutil.WriteFile("MerkleOrders.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "MerkleOrders.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("MerkleOrders.png", cmd, os.FileMode(mode))
}

func (this *MTreeONode) GraphNodes() string {
	if this == nil {
		return ""
	} else if this.Right == nil && this.Left == nil {
		if this.Order == nil {
			//return "node" + base64.URLEncoding.EncodeToString(this.Hash) + strconv.Itoa(*i) + "[label=<<tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>-1</td></tr>>]\n"
			return "node" + strconv.Itoa(this.id) + "[label=<<table><tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>-1</td></tr></table>>]\n"
		} else {
			//			return "node" + base64.URLEncoding.EncodeToString(this.Hash) + "[label=<<tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>" + this.Order.Date + "</td></tr><tr><td>" + this.Order.Shop + "</td></tr><tr><td>" + strconv.FormatInt(this.Order.User.Dpi, 10) + "</td></tr>>]\n"
			return "node" + strconv.Itoa(this.id) + "[label=<<table><tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>" + this.Order.Date + "</td></tr><tr><td>" + this.Order.Shop + "</td></tr><tr><td>" + strconv.FormatInt(this.Order.User.Dpi, 10) + "</td></tr></table>>]\n"
		}
	} else if this.Right == nil {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Left.GraphNodes()
	} else if this.Left == nil {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Right.GraphNodes()
	} else {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Right.GraphNodes() + this.Left.GraphNodes()
	}
}

func (this *MTreeONode) GraphLines() string {
	if this == nil {
		return ""
	} else if this.Right == nil && this.Left == nil {
		return ""
	} else if this.Right == nil {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Left.id) + this.Left.GraphLines()
	} else if this.Left == nil {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Right.id) + this.Right.GraphLines()
	} else {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Right.id) + "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Left.id) + this.Left.GraphLines() + this.Right.GraphLines()
	}
}

func (this *MTreeONode) GetHash() string {
	if this == nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(this.Hash)
}

func (this *MTreePNode) GetHash() string {
	if this == nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(this.Hash)
}

func (this *MTreeSNode) GetHash() string {
	if this == nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(this.Hash)
}

func (this *MTreeUNode) GetHash() string {
	if this == nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(this.Hash)
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
		info := this.Left.GetHash() + this.Right.GetHash()
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

func (this *MTreeONode) CreateTree(level int, merkle *MerkleTreeOrders) {
	if level == 2 {
		h := sha256.New()
		h.Write([]byte("-1"))
		merkle.N += 3
		this.Left = &MTreeONode{Leaf: true, Full: false, Hash: h.Sum(nil), id: merkle.N}
		merkle.N += 3
		this.Right = &MTreeONode{Leaf: true, Full: false, Hash: h.Sum(nil), id: merkle.N}
		merkle.N += 3
		return
	} else {
		merkle.N += 3
		this.Right = &MTreeONode{Leaf: false, Full: false, id: merkle.N}
		merkle.N += 3
		this.Left = &MTreeONode{Leaf: false, Full: false, id: merkle.N}
		merkle.N += 3
	}

	level--

	merkle.N += 3
	this.Left.CreateTree(level, merkle)
	merkle.N += 3
	this.Right.CreateTree(level, merkle)
	merkle.N += 3

	this.Hash = sha256.New().Sum([]byte(base64.URLEncoding.EncodeToString(this.Left.Hash) + base64.URLEncoding.EncodeToString(this.Right.Hash)))
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
func (this *MerkleTreeShops) Graph() {
	var text string
	if this == nil {
		text = "digraph G{\n}"
	} else {
		this.N = 0
		nodes := this.Root.GraphNodes()
		this.N = 0
		lines := this.Root.GraphLines()
		//lines := ""

		text = "digraph G{\n" + nodes + "\n" + lines + "\n}"
	}

	data := []byte(text)
	err := ioutil.WriteFile("MerkleShops.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "MerkleShops.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("MerkleShops.png", cmd, os.FileMode(mode))
}

func (this *MTreeSNode) GraphNodes() string {
	if this == nil {
		return ""
	} else if this.Right == nil && this.Left == nil {
		if this.Shop == nil {
			//return "node" + base64.URLEncoding.EncodeToString(this.Hash) + strconv.Itoa(*i) + "[label=<<tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>-1</td></tr>>]\n"
			return "node" + strconv.Itoa(this.id) + "[label=<<table><tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>-1</td></tr></table>>]\n"
		} else {
			//			return "node" + base64.URLEncoding.EncodeToString(this.Hash) + "[label=<<tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>" + this.Order.Date + "</td></tr><tr><td>" + this.Order.Shop + "</td></tr><tr><td>" + strconv.FormatInt(this.Order.User.Dpi, 10) + "</td></tr>>]\n"
			return "node" + strconv.Itoa(this.id) + "[label=<<table><tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>" + strconv.Itoa(this.Shop.Score) + "</td></tr><tr><td>" + this.Shop.Name + "</td></tr><tr><td>" + this.Shop.Description + "</td></tr></table>>]\n"
		}
	} else if this.Right == nil {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Left.GraphNodes()
	} else if this.Left == nil {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Right.GraphNodes()
	} else {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Right.GraphNodes() + this.Left.GraphNodes()
	}
}

func (this *MTreeSNode) GraphLines() string {
	if this == nil {
		return ""
	} else if this.Right == nil && this.Left == nil {
		return ""
	} else if this.Right == nil {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Left.id) + this.Left.GraphLines()
	} else if this.Left == nil {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Right.id) + this.Right.GraphLines()
	} else {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Right.id) + "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Left.id) + this.Left.GraphLines() + this.Right.GraphLines()
	}
}

type MerkleTreeShops struct {
	Root *MTreeSNode
	N    int
}

func (this *MerkleTreeShops) Show() {
	this.Root.Show(0)
}

func (this *MerkleTreeShops) AddShop(shop *Shop) *MerkleTreeShops {
	if this == nil {
		this = &MerkleTreeShops{}
	}

	if this.Root == nil {
		this.N = 3
		h := sha256.New()
		auxHash := strconv.Itoa(shop.Score) + "," + shop.Name + "," + shop.Description
		h.Write([]byte(auxHash))
		hash := h.Sum(nil)
		auxL := &MTreeSNode{Leaf: true, Full: true, Hash: hash, Shop: shop, id: 1}
		auxR := &MTreeSNode{Leaf: true, Full: false, id: 2}

		this.Root = &MTreeSNode{Leaf: false, Full: false, Left: auxL, Right: auxR, id: 0}
		info := this.Root.Left.GetHash() + this.Root.Right.GetHash()
		h.Write([]byte(info))
		hash = h.Sum(nil)
		this.Root.Hash = hash
		return this
	}

	if this.Root.Full {
		levels := this.Root.GetLevels()
		this.N++
		newTree := &MTreeSNode{Leaf: false, Full: false, id: this.N}
		this.N += 3
		newTree.CreateTree(levels, this)

		aux := this.Root
		this.Root = nil

		this.N += 3
		this.Root = &MTreeSNode{Leaf: false, Full: false, Left: aux, Right: newTree, id: this.N}
	}

	h := sha256.New()
	h.Write([]byte(strconv.Itoa(shop.Score) + "," + shop.Name + "," + shop.Description))
	hash := h.Sum(nil)
	this.N += 3
	this.Root.AddNode(MTreeSNode{Leaf: true, Full: true, Hash: hash, Shop: shop, id: this.N})
	this.N += 3

	return this
}

/*func (this *MerkleTreeShops) AddShop(shop *Shop) *MerkleTreeShops {
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
		info := this.Root.Left.GetHash() + this.Root.Right.GetHash()
		//info := string(this.Root.Left.Hash) + string(this.Root.Right.Hash)
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
}*/

type MTreeSNode struct {
	Shop  *Shop
	Leaf  bool
	Full  bool
	Hash  []byte
	Left  *MTreeSNode
	Right *MTreeSNode
	id    int
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
		//info := string(this.Left.Hash) + string(this.Right.Hash)
		info := this.Left.GetHash() + this.Right.GetHash()
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

func (this *MTreeSNode) CreateTree(level int, merkle *MerkleTreeShops) {
	if level == 2 {
		h := sha256.New()
		h.Write([]byte("-1"))
		merkle.N += 3
		this.Left = &MTreeSNode{Leaf: true, Full: false, Hash: h.Sum(nil), id: merkle.N}
		merkle.N += 3
		this.Right = &MTreeSNode{Leaf: true, Full: false, Hash: h.Sum(nil), id: merkle.N}
		merkle.N += 3
		return
	} else {
		merkle.N += 3
		this.Right = &MTreeSNode{Leaf: false, Full: false, id: merkle.N}
		merkle.N += 3
		this.Left = &MTreeSNode{Leaf: false, Full: false, id: merkle.N}
		merkle.N += 3
	}

	level--

	merkle.N += 3
	this.Left.CreateTree(level, merkle)
	merkle.N += 3
	this.Right.CreateTree(level, merkle)
	merkle.N += 3

	this.Hash = sha256.New().Sum([]byte(base64.URLEncoding.EncodeToString(this.Left.Hash) + base64.URLEncoding.EncodeToString(this.Right.Hash)))
}

/*func (this *MTreeSNode) CreateTree(level int) {
	if level == 2 {
		h := sha256.New()
		h.Write([]byte("-1"))
		this.Left = &MTreeSNode{Leaf: true, Full: false, Hash: h.Sum(nil)}
		this.Right = &MTreeSNode{Leaf: true, Full: false, Hash: h.Sum(nil)}
		return
	} else {
		this.Right = &MTreeSNode{Leaf: false, Full: false}
		this.Left = &MTreeSNode{Leaf: false, Full: false}
	}

	level--

	this.Left.CreateTree(level)
	this.Right.CreateTree(level)
}*/

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

func (this *MerkleTreeProducts) Graph() {
	var text string
	if this == nil {
		text = "digraph G{\n}"
	} else {
		this.N = 0
		nodes := this.Root.GraphNodes()
		this.N = 0
		lines := this.Root.GraphLines()
		//lines := ""

		text = "digraph G{\n" + nodes + "\n" + lines + "\n}"
	}

	data := []byte(text)
	err := ioutil.WriteFile("MerkleProducts.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "MerkleProducts.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("MerkleProducts.png", cmd, os.FileMode(mode))
}

func (this *MTreePNode) GraphNodes() string {
	if this == nil {
		return ""
	} else if this.Right == nil && this.Left == nil {
		if this.Product == nil {
			//return "node" + base64.URLEncoding.EncodeToString(this.Hash) + strconv.Itoa(*i) + "[label=<<tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>-1</td></tr>>]\n"
			return "node" + strconv.Itoa(this.id) + "[label=<<table><tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>-1</td></tr></table>>]\n"
		} else {
			//			return "node" + base64.URLEncoding.EncodeToString(this.Hash) + "[label=<<tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>" + this.Order.Date + "</td></tr><tr><td>" + this.Order.Shop + "</td></tr><tr><td>" + strconv.FormatInt(this.Order.User.Dpi, 10) + "</td></tr>>]\n"
			return "node" + strconv.Itoa(this.id) + "[label=<<table><tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>" + this.Product.Name + "</td></tr><tr><td>" + strconv.Itoa(this.Product.Code) + "</td></tr><tr><td>" + strconv.Itoa(this.Product.Stock) + "</td></tr></table>>]\n"
		}
	} else if this.Right == nil {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Left.GraphNodes()
	} else if this.Left == nil {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Right.GraphNodes()
	} else {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Right.GraphNodes() + this.Left.GraphNodes()
	}
}

func (this *MTreePNode) GraphLines() string {
	if this == nil {
		return ""
	} else if this.Right == nil && this.Left == nil {
		return ""
	} else if this.Right == nil {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Left.id) + this.Left.GraphLines()
	} else if this.Left == nil {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Right.id) + this.Right.GraphLines()
	} else {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Right.id) + "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Left.id) + this.Left.GraphLines() + this.Right.GraphLines()
	}
}

type MerkleTreeProducts struct {
	Root *MTreePNode
	N    int
}

func (this *MerkleTreeProducts) Show() {
	this.Root.Show(0)
}

func (this *MerkleTreeProducts) AddProduct(product *Product) *MerkleTreeProducts {
	if this == nil {
		this = &MerkleTreeProducts{}
	}

	if this.Root == nil {
		this.N = 3
		h := sha256.New()
		auxHash := product.Name + "," + strconv.Itoa(product.Code) + "," + strconv.Itoa(product.Stock)
		h.Write([]byte(auxHash))
		hash := h.Sum(nil)
		auxL := &MTreePNode{Leaf: true, Full: true, Hash: hash, Product: product, id: 1}
		auxR := &MTreePNode{Leaf: true, Full: false, id: 2}

		this.Root = &MTreePNode{Leaf: false, Full: false, Left: auxL, Right: auxR, id: 0}
		info := this.Root.Left.GetHash() + this.Root.Right.GetHash()
		h.Write([]byte(info))
		hash = h.Sum(nil)
		this.Root.Hash = hash
		return this
	}

	if this.Root.Full {
		levels := this.Root.GetLevels()
		this.N++
		newTree := &MTreePNode{Leaf: false, Full: false, id: this.N}
		this.N += 3
		newTree.CreateTree(levels, this)

		aux := this.Root
		this.Root = nil

		this.N += 3
		this.Root = &MTreePNode{Leaf: false, Full: false, Left: aux, Right: newTree, id: this.N}
	}

	h := sha256.New()
	h.Write([]byte(product.Name + "," + strconv.Itoa(product.Code) + "," + strconv.Itoa(product.Stock)))
	hash := h.Sum(nil)
	this.N += 3
	this.Root.AddNode(MTreePNode{Leaf: true, Full: true, Hash: hash, Product: product, id: this.N})
	this.N += 3

	return this
}

/*func (this *MerkleTreeProducts) AddProduct(product *Product) *MerkleTreeProducts {
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
		//info := string(this.Root.Left.Hash) + string(this.Root.Right.Hash)
		info := this.Root.Left.GetHash() + this.Root.Right.GetHash()
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
}*/

type MTreePNode struct {
	Product *Product
	Leaf    bool
	Full    bool
	Hash    []byte
	Left    *MTreePNode
	Right   *MTreePNode
	id      int
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
		//info := string(this.Left.Hash) + string(this.Right.Hash)
		info := this.Left.GetHash() + this.Right.GetHash()
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

func (this *MTreePNode) CreateTree(level int, merkle *MerkleTreeProducts) {
	if level == 2 {
		h := sha256.New()
		h.Write([]byte("-1"))
		merkle.N += 3
		this.Left = &MTreePNode{Leaf: true, Full: false, Hash: h.Sum(nil), id: merkle.N}
		merkle.N += 3
		this.Right = &MTreePNode{Leaf: true, Full: false, Hash: h.Sum(nil), id: merkle.N}
		merkle.N += 3
		return
	} else {
		merkle.N += 3
		this.Right = &MTreePNode{Leaf: false, Full: false, id: merkle.N}
		merkle.N += 3
		this.Left = &MTreePNode{Leaf: false, Full: false, id: merkle.N}
		merkle.N += 3
	}

	level--

	merkle.N += 3
	this.Left.CreateTree(level, merkle)
	merkle.N += 3
	this.Right.CreateTree(level, merkle)
	merkle.N += 3

	this.Hash = sha256.New().Sum([]byte(base64.URLEncoding.EncodeToString(this.Left.Hash) + base64.URLEncoding.EncodeToString(this.Right.Hash)))
}

/*func (this *MTreePNode) CreateTree(level int) {
	if level == 2 {
		h := sha256.New()
		h.Write([]byte("-1"))
		this.Left = &MTreePNode{Leaf: true, Full: false, Hash: h.Sum(nil)}
		this.Right = &MTreePNode{Leaf: true, Full: false, Hash: h.Sum(nil)}
		return
	} else {
		this.Right = &MTreePNode{Leaf: false, Full: false}
		this.Left = &MTreePNode{Leaf: false, Full: false}
	}

	level--

	this.Left.CreateTree(level)
	this.Right.CreateTree(level)
}*/

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

func (this *MerkleTreeUsers) Graph() {
	var text string
	if this == nil {
		text = "digraph G{\n}"
	} else {
		this.N = 0
		nodes := this.Root.GraphNodes()
		this.N = 0
		lines := this.Root.GraphLines()
		//lines := ""

		text = "digraph G{\n" + nodes + "\n" + lines + "\n}"
	}

	data := []byte(text)
	err := ioutil.WriteFile("MerkleUsers.dot", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "MerkleUsers.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("MerkleUsers.png", cmd, os.FileMode(mode))
}

func (this *MTreeUNode) GraphNodes() string {
	if this == nil {
		return ""
	} else if this.Right == nil && this.Left == nil {
		if this.User == nil {
			//return "node" + base64.URLEncoding.EncodeToString(this.Hash) + strconv.Itoa(*i) + "[label=<<tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>-1</td></tr>>]\n"
			return "node" + strconv.Itoa(this.id) + "[label=<<table><tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>-1</td></tr></table>>]\n"
		} else {
			//			return "node" + base64.URLEncoding.EncodeToString(this.Hash) + "[label=<<tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>" + this.Order.Date + "</td></tr><tr><td>" + this.Order.Shop + "</td></tr><tr><td>" + strconv.FormatInt(this.Order.User.Dpi, 10) + "</td></tr>>]\n"
			return "node" + strconv.Itoa(this.id) + "[label=<<table><tr><td>" + base64.URLEncoding.EncodeToString(this.Hash) + "</td></tr><tr><td>" + strconv.FormatInt(this.User.Dpi, 10) + "</td></tr><tr><td>" + this.User.User + "</td></tr><tr><td>" + this.User.Email + "</td></tr></table>>]\n"
		}
	} else if this.Right == nil {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Left.GraphNodes()
	} else if this.Left == nil {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Right.GraphNodes()
	} else {
		return "node" + strconv.Itoa(this.id) + "[label=\"" + base64.URLEncoding.EncodeToString(this.Hash) + "\"]\n" + this.Right.GraphNodes() + this.Left.GraphNodes()
	}
}

func (this *MTreeUNode) GraphLines() string {
	if this == nil {
		return ""
	} else if this.Right == nil && this.Left == nil {
		return ""
	} else if this.Right == nil {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Left.id) + this.Left.GraphLines()
	} else if this.Left == nil {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Right.id) + this.Right.GraphLines()
	} else {
		return "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Right.id) + "\nnode" + strconv.Itoa(this.id) + "->" + "node" + strconv.Itoa(this.Left.id) + this.Left.GraphLines() + this.Right.GraphLines()
	}
}

type MerkleTreeUsers struct {
	Root *MTreeUNode
	N    int
}

func (this *MerkleTreeUsers) Show() {
	this.Root.Show(0)
}

func (this *MerkleTreeUsers) AddUser(user *Account) *MerkleTreeUsers {
	if this == nil {
		this = &MerkleTreeUsers{}
	}

	if this.Root == nil {
		this.N = 3
		h := sha256.New()
		auxHash := strconv.FormatInt(user.Dpi, 10) + "," + user.User + "," + user.Email + ","
		h.Write([]byte(auxHash))
		hash := h.Sum(nil)
		auxL := &MTreeUNode{Leaf: true, Full: true, Hash: hash, User: user, id: 1}
		auxR := &MTreeUNode{Leaf: true, Full: false, id: 2}

		this.Root = &MTreeUNode{Leaf: false, Full: false, Left: auxL, Right: auxR, id: 0}
		info := this.Root.Left.GetHash() + this.Root.Right.GetHash()
		h.Write([]byte(info))
		hash = h.Sum(nil)
		this.Root.Hash = hash
		return this
	}

	if this.Root.Full {
		levels := this.Root.GetLevels()
		this.N++
		newTree := &MTreeUNode{Leaf: false, Full: false, id: this.N}
		this.N += 3
		newTree.CreateTree(levels, this)

		aux := this.Root
		this.Root = nil

		this.N += 3
		this.Root = &MTreeUNode{Leaf: false, Full: false, Left: aux, Right: newTree, id: this.N}
	}

	h := sha256.New()
	h.Write([]byte(strconv.FormatInt(user.Dpi, 10) + "," + user.User + "," + user.Email + ","))
	hash := h.Sum(nil)
	this.N += 3
	this.Root.AddNode(MTreeUNode{Leaf: true, Full: true, Hash: hash, User: user, id: this.N})
	this.N += 3

	return this
}

/*
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
		//info := string(this.Root.Left.Hash) + string(this.Root.Right.Hash)
		info := this.Root.Left.GetHash() + this.Root.Right.GetHash()
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
}*/

type MTreeUNode struct {
	User  *Account
	Leaf  bool
	Full  bool
	Hash  []byte
	Left  *MTreeUNode
	Right *MTreeUNode
	id    int
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
		//info := string(this.Left.Hash) + string(this.Right.Hash)
		info := this.Left.GetHash() + this.Right.GetHash()
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

func (this *MTreeUNode) CreateTree(level int, merkle *MerkleTreeUsers) {
	if level == 2 {
		h := sha256.New()
		h.Write([]byte("-1"))
		merkle.N += 3
		this.Left = &MTreeUNode{Leaf: true, Full: false, Hash: h.Sum(nil), id: merkle.N}
		merkle.N += 3
		this.Right = &MTreeUNode{Leaf: true, Full: false, Hash: h.Sum(nil), id: merkle.N}
		merkle.N += 3
		return
	} else {
		merkle.N += 3
		this.Right = &MTreeUNode{Leaf: false, Full: false, id: merkle.N}
		merkle.N += 3
		this.Left = &MTreeUNode{Leaf: false, Full: false, id: merkle.N}
		merkle.N += 3
	}

	level--

	merkle.N += 3
	this.Left.CreateTree(level, merkle)
	merkle.N += 3
	this.Right.CreateTree(level, merkle)
	merkle.N += 3

	this.Hash = sha256.New().Sum([]byte(base64.URLEncoding.EncodeToString(this.Left.Hash) + base64.URLEncoding.EncodeToString(this.Right.Hash)))
}

/*func (this *MTreeUNode) CreateTree(level int) {
	if level == 2 {
		h := sha256.New()
		h.Write([]byte("-1"))
		this.Left = &MTreeUNode{Leaf: true, Full: false, Hash: h.Sum(nil)}
		this.Right = &MTreeUNode{Leaf: true, Full: false, Hash: h.Sum(nil)}
		return
	} else {
		this.Right = &MTreeUNode{Leaf: false, Full: false}
		this.Left = &MTreeUNode{Leaf: false, Full: false}
	}

	level--

	this.Left.CreateTree(level)
	this.Right.CreateTree(level)
}*/

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
