package models

type Producto struct {
	nombre          string
	tipoIngrediente TipoIngrediente
}

func NewProducto(nombre string, tipoIngrediente TipoIngrediente) Producto {
	return Producto{
		nombre:          nombre,
		tipoIngrediente: tipoIngrediente,
	}
}
