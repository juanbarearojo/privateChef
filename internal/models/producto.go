package models

import (
	"errors"
	"time"
)

type TipoIngrediente string

const (
	Perecedero   TipoIngrediente = "perecedero"
	NoPerecedero TipoIngrediente = "no perecedero"
)

type Producto struct {
	nombre         string
	tipo           TipoIngrediente
	fechaCaducidad *string
}

func NewProducto(nombre string, tipo TipoIngrediente, fechaCaducidad *string) (*Producto, error) {
	if tipo == Perecedero && fechaCaducidad != nil {
		fecha, err := time.Parse("02/01/2006", *fechaCaducidad)
		if err != nil {
			return nil, errors.New("la fecha de caducidad debe tener el formato DD/MM/YYYY")
		}

		anoActual := time.Now().Year()
		if fecha.Year() < anoActual {
			return nil, errors.New("la fecha de caducidad debe ser de este año o superior")
		}
	} else if tipo == Perecedero && fechaCaducidad == nil {
		return nil, errors.New("los productos perecederos deben tener una fecha de caducidad")
	}

	return &Producto{
		nombre:         nombre,
		tipo:           tipo,
		fechaCaducidad: fechaCaducidad,
	}, nil
}
