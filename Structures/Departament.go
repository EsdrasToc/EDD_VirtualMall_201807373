package Structures

type Departament struct {
	Name  string `json:"Nombre"`
	Shops []Shop `json:"Tiendas"`
}

func (this Departament) GetShops() []Shop {
	return this.Shops
}
