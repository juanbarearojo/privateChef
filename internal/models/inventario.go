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

func (i *Inventario) Clone() *Inventario {
	ingredientesClonados := make(map[Producto]uint64)
	for producto, cantidad := range i.ingredientes {
		ingredientesClonados[producto] = cantidad
	}

	return &Inventario{
		ingredientes: ingredientesClonados,
	}
}

func (i *Inventario) aplicarAsignacion(recetas []Receta) *Inventario {
	// Clonar el inventario para no modificar el original
	nuevoInventario := i.Clone()

	// Recorrer cada receta en la lista de recetas
	for _, receta := range recetas {
		// Obtener los ingredientes necesarios para la receta usando GetIngredientes()
		for producto, cantidadNecesaria := range receta.GetIngredientes() {
			// Restar la cantidad necesaria del inventario sin verificar si el producto existe
			nuevoInventario.ingredientes[producto] -= cantidadNecesaria
		}
	}

	return nuevoInventario
}
