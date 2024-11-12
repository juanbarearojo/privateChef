package models

type Producto struct {
	nombre string
}

func NewProducto(nombre string) Producto {
	return Producto{
		nombre: nombre,
	}
}
