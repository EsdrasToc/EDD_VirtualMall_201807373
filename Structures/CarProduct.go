package Structures

type CarProduct struct {
	Shop_    Shop      `json:"Tienda"`
	Date     string    `json:"Fecha"`
	Products []Product `json:"Producto"`
}
