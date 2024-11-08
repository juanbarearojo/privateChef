package models

import "time"

type Ingrediente struct {
	nombre          string
	tipoIngrediente TipoIngrediente
	cantidad        float32
	fechaCaducidad  *time.Time
}

func NewIngrediente(nombre string, tipoIngrediente TipoIngrediente, cantidad float32, fechaCaducidad *time.Time) Ingrediente {
	ingrediente := Ingrediente{
		nombre:          nombre,
		tipoIngrediente: tipoIngrediente,
		cantidad:        cantidad,
	}

	if tipoIngrediente == Perecedero {
		ingrediente.fechaCaducidad = fechaCaducidad
	}

	return ingrediente
}
