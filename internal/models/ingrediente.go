package models

import "time"

type Ingrediente struct {
	id              int
	nombre          string
	tipoIngrediente TipoIngrediente
	cantidad        float32
	fechaCaducidad  *time.Time
}

func NewIngrediente(id int, nombre string, tipoIngrediente TipoIngrediente, cantidad float32, fechaCaducidad *time.Time) Ingrediente {
	ingrediente := Ingrediente{
		id:              id,
		nombre:          nombre,
		tipoIngrediente: tipoIngrediente,
		cantidad:        cantidad,
	}

	if tipoIngrediente == Perecedero {
		ingrediente.fechaCaducidad = fechaCaducidad
	}

	return ingrediente
}
