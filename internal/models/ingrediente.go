package models

import "time"

type Ingrediente struct {
	id              int
	nombre          string
	tipoIngrediente TipoIngrediente
	cantidad        float32
	fechaCaducidad  *time.Time
}
