package models

type Inventario struct {
	ingredientes map[string]Ingrediente
}

func (i *Inventario) GetIngredientes() map[string]Ingrediente {
	return i.ingredientes
}
