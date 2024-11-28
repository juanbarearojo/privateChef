package models

import (
	"errors"
)

type Ingrediente struct {
	productos map[string]Producto
	cantidad  uint64
}

func NewIngrediente(productos map[string]Producto, producto Producto, cantidad uint64) (*Ingrediente, error) {

	if _, existe := productos[producto.nombre]; existe {
		return nil, errors.New("el producto ya existe")
	}

	productos[producto.nombre] = producto

	return &Ingrediente{
		productos: productos,
		cantidad:  cantidad,
	}, nil
}
