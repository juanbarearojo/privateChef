package models

type Receta struct {
	titulo       string
	ingredientes map[Producto]uint64
}

func (r *Receta) GetIngredientes() map[Producto]uint64 {
	return r.ingredientes
}

func (r *Receta) SumaPerecederos() uint64 {
	var suma uint64 = 0
	for producto, cantidad := range r.ingredientes {
		if producto.tipo == Perecedero {
			suma += cantidad
		}
	}
	return suma
}
