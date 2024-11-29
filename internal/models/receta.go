package models

type Receta struct {
	titulo       string
	ingredientes map[Producto]uint64
}

func (r *Receta) GetIngredientes() map[Producto]uint64 {
	return r.ingredientes
}
