package models

import "time"

type Ingrediente struct {
	producto       Producto
	cantidad       float32
	fechaCaducidad *time.Time
}

func NewIngrediente(producto Producto, cantidad float32, fechaCaducidad *time.Time) Ingrediente {
	return Ingrediente{
		producto:       producto,
		cantidad:       cantidad,
		fechaCaducidad: fechaCaducidad,
	}
}
