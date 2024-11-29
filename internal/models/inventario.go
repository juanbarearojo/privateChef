package models

type Inventario struct {
	ingredientes map[Producto]uint64
}

func (i *Inventario) GetIngredientes() map[Producto]uint64 {
	return i.ingredientes
}
