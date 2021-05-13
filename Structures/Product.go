package Structures

import (
	"fmt"
	"strconv"
)

func (avl *Product) Insert(newNode Product) *Product { // Insertar valor
	value := newNode.Code
	if avl == nil { // Se inserta el nodo raíz
		node := Product{
			Code:           value,
			Name:           newNode.Name,
			Description:    newNode.Description,
			Price:          newNode.Price,
			Stock:          newNode.Stock,
			Image:          newNode.Image,
			Almacenamiento: newNode.Almacenamiento,
		}
		avl = &node
		avl.left = nil
		avl.right = nil
	} else if value < avl.Code { // Insertar en el subárbol izquierdo
		avl.left = avl.left.Insert(newNode)            // recorrer para encontrar la posición que se va a insertar
		if avl.left.height()-avl.right.height() == 2 { // Juzgar el equilibrio
			if value < avl.left.Code { // Izquierda Izquierda
				avl = avl.left_left()
			} else { // Acerca de
				avl = avl.left_right()
			}
		}
	} else if value > avl.Code { // Inserta el subárbol derecho
		avl.right = avl.right.Insert(newNode)
		if avl.right.height()-avl.left.height() == 2 {
			if value < avl.right.Code {
				avl = avl.right_left() // derecha izquierda
			} else {

				avl = avl.right_right() // derecha derecha
			}
		}
	} else {
		avl.Stock += newNode.Stock
	}
	avl.high = max1(avl.left.height(), avl.right.height()) + 1 // Actualizar altura

	return avl
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (avl *Product) left_left() *Product {

	k := avl.left      // Determine la nueva raíz
	avl.left = k.right // rotar
	k.right = avl
	avl.high = max1(avl.left.height(), avl.right.height()) + 1
	k.high = max1(k.left.height(), avl.high) + 1
	return k
}

func (avl *Product) left_right() *Product {
	avl.left = avl.left.left_left() // Primero gira a la izquierda el subárbol izquierdo
	return avl.right_right()        // Rota a la derecha este nodo
}

func (avl *Product) right_left() *Product {
	avl.right = avl.right.right_right()
	return avl.left_left()
}

func (avl *Product) right_right() *Product {

	k := avl.right
	avl.right = k.left
	k.left = avl

	avl.high = max1(avl.left.height(), avl.right.height()) + 1
	k.high = max1(avl.high, k.right.height()) + 1

	return k
}

func (avl *Product) height() int {
	if avl == nil {
		return 0
	} else {
		return avl.high
	}
}

func inorden(node *Product) {
	if node == nil {
		return
	}

	inorden(node.left)
	//fmt.Println("=== Producto: " + node.Name)
	//fmt.Println("=== Stock: ")
	//fmt.Println(node.Stock)
	inorden(node.right)
}

type Product struct {
	Name           string    `json:"Nombre"`
	Code           int       `json:"Codigo"`
	Description    string    `json:"Descripcion"`
	Price          float64   `json:"Precio"`
	Stock          int       `json:"Cantidad"`
	Image          string    `json:"Imagen"`
	Almacenamiento string    `json:"Almacenamiento"`
	Comments       *Comments `json:"-"`
	high           int       `json:"-"`
	left           *Product  `json:"-"`
	right          *Product  `json:"-"`
}

func (this *Product) SearchProduct(product *Product) *Product {
	if this == nil {
		return nil
	}

	if this.Code == product.Code && this.Name == product.Name {
		fmt.Println("SE ENCONTRO EL PRODUCTO")
		return this
	}

	left := this.left.SearchProduct(product)
	right := this.right.SearchProduct(product)
	if left != nil {
		return left
	} else {
		return right
	}
}

func (this *Product) ToJson() string {
	//file, _ := json.MarshalIndent(this, "", "\t")

	var text string

	text = "{\n\"Nombre\":\"" + this.Name + "\",\n\"Codigo\":" + strconv.Itoa(this.Code) + ",\n\"Descripcion\":\"" + this.Description + "\",\n\"Precio\":" + strconv.FormatFloat(this.Price, 'E', -1, 64) + ",\n\"Cantidad\":" + strconv.Itoa(this.Stock) + ",\n\"Imagen\":\"" + this.Image + "\",\n\"Almacenamiento\":\"" + this.Almacenamiento + "\",\n\"Comentarios\":[" + this.Comments.ToJson() + "]\n}"
	fmt.Println(text)
	return text
}

func (this *Product) GetProducts() string {
	if this == nil {
		return ""
	} else if this.right == nil && this.left == nil {
		return this.ToJson()
	} else if this.right == nil {
		return this.ToJson() + ",\n" + this.left.GetProducts()
	} else if this.left == nil {
		return this.ToJson() + ",\n" + this.right.GetProducts()
	} else {
		return this.ToJson() + ",\n" + this.right.GetProducts() + ",\n" + this.left.GetProducts()
	}
}

func (this *Product) ToString() string {

	if this == nil {
		return ""
	}

	var text string

	text = this.Name + strconv.Itoa(this.Code) + this.Description + strconv.Itoa(this.Stock) + this.Image + this.Almacenamiento

	return text
}
