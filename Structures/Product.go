package Structures

import (
	"fmt"
)

/*type AVLTreeNode struct { // Definir nodo
	key   int          `json:"Codigo"`
	high  int          `json:"-"`
	left  *AVLTreeNode `json:"-"`
	right *AVLTreeNode `json:"-"`
}*/

func (avl *Product) Insert(newNode Product) *Product { // Insertar valor
	value := newNode.Code
	if avl == nil { // Se inserta el nodo raíz
		fmt.Println("ENTRO AQUI XD")
		node := Product{
			Code:        value,
			Name:        newNode.Name,
			Description: newNode.Description,
			Price:       newNode.Price,
			Stock:       newNode.Stock,
			Image:       newNode.Image,
		}
		avl = &node
		avl.left = nil
		avl.right = nil
	} else if value < avl.Code { // Insertar en el subárbol izquierdo
		fmt.Println("NEL ACA XD")
		avl.left = avl.left.Insert(newNode)            // recorrer para encontrar la posición que se va a insertar
		if avl.left.height()-avl.right.height() == 2 { // Juzgar el equilibrio
			if value < avl.left.Code { // Izquierda Izquierda
				avl = avl.left_left()
			} else { // Acerca de
				avl = avl.left_right()
			}
		}
	} else if value > avl.Code { // Inserta el subárbol derecho
		fmt.Println("NEL ACA JAJAJA XD")
		avl.right = avl.right.Insert(newNode)
		if avl.right.height()-avl.left.height() == 2 {
			if value < avl.right.Code {
				avl = avl.right_left() // derecha izquierda
			} else {

				avl = avl.right_right() // derecha derecha
			}
		}
	} else {
		fmt.Println("the key", value, "has exists")
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
	fmt.Println("=== Producto: " + node.Name)
	fmt.Println("=== Stock: ")
	fmt.Println(node.Stock)
	inorden(node.right)
}

type Product struct {
	Name        string   `json:"Nombre"`
	Code        int      `json:"Codigo"`
	Description string   `json:"Descripcion"`
	Price       float64  `json:"Precio"`
	Stock       int      `json:"Cantidad"`
	Image       string   `json:"Imagen"`
	high        int      `json:"-"`
	left        *Product `json:"-"`
	right       *Product `json:"-"`
}