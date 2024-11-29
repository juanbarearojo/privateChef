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
