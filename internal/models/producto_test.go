package models

import (
	"testing"
)

func TestNewProductoPerecederoInvalidDateFormat(t *testing.T) {
	nombre := "Yogur"
	tipo := Perecedero
	fechaCaducidad := "2023-12-15" // Formato incorrecto

	_, err := NewProducto(nombre, tipo, &fechaCaducidad)
	if err == nil {
		t.Fatalf("Se esperaba error por formato de fecha inválido, pero se obtuvo nil")
	}
	expectedError := "la fecha de caducidad debe tener el formato DD/MM/YYYY"
	if err.Error() != expectedError {
		t.Errorf("Se esperaba el error '%s', pero se obtuvo '%s'", expectedError, err.Error())
	}
}

func TestNewProductoPerecederoPastDate(t *testing.T) {
	nombre := "Pan"
	tipo := Perecedero
	fechaCaducidad := "15/12/2021" // Año pasado

	_, err := NewProducto(nombre, tipo, &fechaCaducidad)
	if err == nil {
		t.Fatalf("Se esperaba error por fecha de caducidad en el pasado, pero se obtuvo nil")
	}
	expectedError := "la fecha de caducidad debe ser de este año o superior"
	if err.Error() != expectedError {
		t.Errorf("Se esperaba el error '%s', pero se obtuvo '%s'", expectedError, err.Error())
	}
}

func TestNewProductoPerecederoNilFechaCaducidad(t *testing.T) {
	nombre := "Queso"
	tipo := Perecedero
	var fechaCaducidad *string = nil

	prod, err := NewProducto(nombre, tipo, fechaCaducidad)
	if err != nil {
		t.Fatalf("No se esperaba error cuando fechaCaducidad es nil, pero se obtuvo: %v", err)
	}
	if prod.fechaCaducidad != nil {
		t.Errorf("Se esperaba que fechaCaducidad fuera nil, pero se obtuvo %v", *prod.fechaCaducidad)
	}
}
