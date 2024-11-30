package models

type Inventario struct {
	ingredientes map[Producto]uint64
}

func (i *Inventario) GetIngredientes() map[Producto]uint64 {
	return i.ingredientes
}

func (i *Inventario) GetDesperdicio() uint64 {
	var desperdicio uint64 = 0
	for producto, cantidad := range i.ingredientes {
		if producto.tipo == Perecedero {
			desperdicio += cantidad
		}
	}
	return desperdicio
}
