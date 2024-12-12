package models

type TipoIngrediente string

const (
	Perecedero   TipoIngrediente = "perecedero"
	NoPerecedero TipoIngrediente = "no perecedero"
)

type Producto struct {
	nombre string
	tipo   TipoIngrediente
}
