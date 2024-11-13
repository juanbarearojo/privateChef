package models

type TipoIngrediente string

const (
	Perecedero TipoIngrediente = "perecedero"
)

type Producto struct {
	nombre string
	tipo   *TipoIngrediente
}

func NewProducto(nombre string, tipo *TipoIngrediente) Producto {
	return Producto{
		nombre: nombre,
		tipo:   tipo,
	}
}
