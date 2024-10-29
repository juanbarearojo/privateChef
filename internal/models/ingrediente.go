package models

type Ingrediente struct {
	nombre          string
	tipoIngrediente TipoIngrediente
	cantidad        float32
}

func NewIngrediente(nombre string, tipoIngrediente TipoIngrediente, cantidad float32) Ingrediente {
	return Ingrediente{
		nombre:          nombre,
		tipoIngrediente: tipoIngrediente,
		cantidad:        cantidad,
	}
}

func (i Ingrediente) GetNombre() string {
	return i.nombre
}

func (i Ingrediente) GetTipoIngrediente() TipoIngrediente {
	return i.tipoIngrediente
}

func (i Ingrediente) GetCantidad() float32 {
	return i.cantidad
}

func (i *Ingrediente) SetNombre(nombre string) {
	i.nombre = nombre
}

func (i *Ingrediente) SetTipoIngrediente(tipoIngrediente TipoIngrediente) {
	i.tipoIngrediente = tipoIngrediente
}

func (i *Ingrediente) SetCantidad(cantidad float32) {
	i.cantidad = cantidad
}
