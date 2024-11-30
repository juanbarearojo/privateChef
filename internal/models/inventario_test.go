package models

import (
	"testing"
)

// Definimos un producto de ejemplo
func TestGetDesperdicioNoNegativo(t *testing.T) {
	// Crear un producto
	fechaCaducidadmanzana := "15/12/2024"
	producto := Producto{
		nombre:         "Manzana",
		tipo:           Perecedero,
		fechaCaducidad: &fechaCaducidadmanzana,
	}

	// Crear un inventario con el producto
	inventario := &Inventario{
		ingredientes: map[Producto]uint64{
			producto: 10, // Cantidad positiva
		},
	}

	// Llamar a la función GetDesperdicio
	desperdicio := inventario.GetDesperdicio()

	// Verificar que el valor no es negativo
	if desperdicio < 0 {
		t.Errorf("GetDesperdicio devolvió un valor negativo: %d", desperdicio)
	}
}
