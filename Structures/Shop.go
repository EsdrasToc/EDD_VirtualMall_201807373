package Structures

type Shop struct {
	Name        string  `json:"Nombre"`
	Description string  `json:"Descripcion"`
	Contact     string  `json:"Contacto"`
	Score       float32 `json:"Calificacion"`
}

func (this Shop) GetName() string {
	return this.Name
}
