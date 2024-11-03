package models

import "time"

type Ingrediente struct {
	id              int
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

// Getters
func (i Ingrediente) GetId() int {
	return i.id
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

func (i Ingrediente) GetFechaCaducidad() *time.Time {
	return i.fechaCaducidad
}

// Setters
func (i *Ingrediente) SetCantidad(cantidad float32) {
	i.cantidad = cantidad
}

func (i *Ingrediente) SetFechaCaducidad(fechaCaducidad *time.Time) {
	if i.tipoIngrediente == Perecedero {
		i.fechaCaducidad = fechaCaducidad
	}
}
