package models

import (
	"errors"
	"time"
)

type Ingrediente struct {
	productos      map[string]Producto
	cantidad       uint64
	fechaCaducidad *string
}

func NewIngrediente(productos map[string]Producto, producto Producto, cantidad uint64, fechaCaducidad *string) (*Ingrediente, error) {

	if fechaCaducidad != nil {
		fecha, err := time.Parse("02/01/2006", *fechaCaducidad)
		if err != nil {
			return nil, errors.New("la fecha de caducidad debe tener el formato DD/MM/YYYY")
		}

		anoActual := time.Now().Year()

		if fecha.Year() < anoActual {
			return nil, errors.New("la fecha de caducidad debe ser de este año o superior")
		}
	}

	if _, existe := productos[producto.nombre]; existe {
		return nil, errors.New("el producto ya existe")
	}

	productos[producto.nombre] = producto

	return &Ingrediente{
		productos:      productos,
		cantidad:       cantidad,
		fechaCaducidad: fechaCaducidad,
	}, nil
}
