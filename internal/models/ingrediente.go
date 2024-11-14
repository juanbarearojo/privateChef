package models

import (
	"errors"
	"time"
)

type Ingrediente struct {
	producto       Producto
	cantidad       float32
	fechaCaducidad string
}

func NewIngrediente(producto Producto, cantidad float32, fechaCaducidad string, inventario *Inventario) (*Ingrediente, error) {
	if cantidad <= 0 {
		return nil, errors.New("la cantidad debe ser mayor que 0")
	}

	fecha, err := time.Parse("02/01/2006", fechaCaducidad)
	if err != nil {
		return nil, errors.New("la fecha de caducidad debe tener el formato DD/MM/YYYY")
	}

	if fecha.Before(time.Now()) {
		return nil, errors.New("la fecha de caducidad debe ser superior al año actual")
	}

	ingredientes := inventario.GetIngredientes()
	if _, existe := ingredientes[producto.nombre]; existe {
		return nil, errors.New("el ingrediente ya está en el inventario")
	}

	return &Ingrediente{
		producto:       producto,
		cantidad:       cantidad,
		fechaCaducidad: fechaCaducidad,
	}, nil
}
