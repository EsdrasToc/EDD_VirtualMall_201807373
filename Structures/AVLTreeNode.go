package Structures

import "fmt"

type AVLTreeNode struct { // Definir nodo
	key   int          `json:"Codigo"`
	high  int          `json:"-"`
	left  *AVLTreeNode `json:"-"`
	right *AVLTreeNode `json:"-"`
}

func (avl *AVLTreeNode) Insert(value int) *AVLTreeNode { // Insertar valor
	if avl == nil { // Se inserta el nodo raíz
		node := &AVLTreeNode{
			key: value,
		}
		avl = node
		avl.left = nil
		avl.right = nil
	} else if value < avl.key { // Insertar en el subárbol izquierdo
		avl.left = avl.left.Insert(value)              // recorrer para encontrar la posición que se va a insertar
		if avl.left.height()-avl.right.height() == 2 { // Juzgar el equilibrio
			if value < avl.left.key { // Izquierda Izquierda
				avl = avl.left_left()
			} else { // Acerca de
				avl = avl.left_right()
			}
		}
	} else if value > avl.key { // Inserta el subárbol derecho
		avl.right = avl.right.Insert(value)
		if avl.right.height()-avl.left.height() == 2 {
			if value < avl.right.key {
				avl = avl.right_left() // derecha izquierda
			} else {

				avl = avl.right_right() // derecha derecha
			}
		}
	} else {
		fmt.Println("the key", value, "has exists")
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

func (avl *AVLTreeNode) left_left() *AVLTreeNode {

	k := avl.left      // Determine la nueva raíz
	avl.left = k.right // rotar
	k.right = avl
	avl.high = max1(avl.left.height(), avl.right.height()) + 1
	k.high = max1(k.left.height(), avl.high) + 1
	return k
}

func (avl *AVLTreeNode) left_right() *AVLTreeNode {
	avl.left = avl.left.left_left() // Primero gira a la izquierda el subárbol izquierdo
	return avl.right_right()        // Rota a la derecha este nodo
}

func (avl *AVLTreeNode) right_left() *AVLTreeNode {
	avl.right = avl.right.right_right()
	return avl.left_left()
}

func (avl *AVLTreeNode) right_right() *AVLTreeNode {

	k := avl.right
	avl.right = k.left
	k.left = avl

	avl.high = max1(avl.left.height(), avl.right.height()) + 1
	k.high = max1(avl.high, k.right.height()) + 1

	return k
}

func (avl *AVLTreeNode) height() int {
	if avl == nil {
		return 0
	} else {
		return avl.high
	}
}

type Product struct {
	name        string  `json:"Nombre"`
	description string  `json:"Descripcion"`
	price       float32 `json:"Precio"`
	stock       int     `json:"Cantidad"`
	image       string  `json:"Imagen"`
	AVLTreeNode
}
