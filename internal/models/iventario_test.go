package models

import (
	"testing"
)

// Definimos un producto de ejemplo
func TestGetDesperdicio(t *testing.T) {
	// Crear un producto
	producto := Producto{
		nombre: "Manzana",
		tipo:   Perecedero,
		// Suponiendo que no necesitamos fechaCaducidad para este test
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
