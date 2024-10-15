package models

type TipoIngrediente int

const (
	Perecedero TipoIngrediente = iota
	NoPerecedero
	UsoComun
)

type Ingrediente struct {
	nombre          string
	tipoIngrediente TipoIngrediente
	cantidad        float64
}
