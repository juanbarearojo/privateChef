package models

type Inventario struct {
	ingredientes []Ingrediente
}

func NewInventario(ingredientes []Ingrediente) Inventario {
	return Inventario{
		ingredientes: ingredientes,
	}
}

func (i Inventario) GetIngredientes() []Ingrediente {
	return i.ingredientes
}

func (i *Inventario) SetIngredientes(ingredientes []Ingrediente) {
	i.ingredientes = ingredientes
}
