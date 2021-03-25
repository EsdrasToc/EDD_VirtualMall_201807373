package Structures

type CarProduct struct {
	Shop_    Shop      `json:"Tienda"`
	Products []Product `json:"Producto"`
}
