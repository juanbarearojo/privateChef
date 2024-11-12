package models

type Inventario struct {
	ingredientes map[string]Ingrediente
}

func NewInventario() Inventario {
	return Inventario{
		ingredientes: make(map[string]Ingrediente),
	}
}
