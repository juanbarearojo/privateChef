package models

type Inventario struct {
	ingredientes []Ingrediente
}

func NewInventario(ingredientes []Ingrediente) Inventario {
	return Inventario{
		ingredientes: ingredientes,
	}
}
