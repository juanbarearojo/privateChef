package models

import (
	"reflect"
	"testing"
)

func TestGetDesperdicioNoNegativo(t *testing.T) {
	_, _, _, _, inventario, _ := setupTestEnvironment()

	desperdicio := inventario.GetDesperdicio()
	if desperdicio < 0 {
		t.Errorf("GetDesperdicio devolvió un valor negativo: %d", desperdicio)
	}
}

func TestGetDesperdicioValorCorrecto(t *testing.T) {
	_, _, _, _, inventario, _ := setupTestEnvironment()

	// Perecederos: Levadura(2) + Huevos(6) = 8
	desperdicioEsperado := uint64(8)
	desperdicioActual := inventario.GetDesperdicio()

	if desperdicioActual != desperdicioEsperado {
		t.Errorf("GetDesperdicio devolvió %d, se esperaba %d", desperdicioActual, desperdicioEsperado)
	}
}

func TestInventarioCloneNoVacio(t *testing.T) {
	_, _, _, _, inventario, _ := setupTestEnvironment()

	inventarioClonado := inventario.Clone()

	if len(inventarioClonado.ingredientes) == 0 {
		t.Errorf("El inventario clonado está vacío")
	}
}

func TestInventarioCloneIgualAlOriginal(t *testing.T) {
	_, _, _, _, inventario, _ := setupTestEnvironment()

	inventarioClonado := inventario.Clone()

	if !reflect.DeepEqual(inventario.ingredientes, inventarioClonado.ingredientes) {
		t.Errorf("El inventario clonado no es igual al original.\nOriginal: %v\nClonado: %v", inventario.ingredientes, inventarioClonado.ingredientes)
	}
}

func TestAplicarAsignacionNoNull(t *testing.T) {
	_, _, _, _, inventario, recetas := setupTestEnvironment()

	// Probaremos con las dos primeras recetas: Pan Casero y Bizcocho
	recetasTest := []Receta{recetas[0], recetas[1]}
	nuevoInventario := inventario.aplicarAsignacion(recetasTest)

	if nuevoInventario == nil {
		t.Errorf("aplicarAsignacion devolvió un inventario nulo")
	}
}

func TestAplicarAsignacionInventarioCorrecto(t *testing.T) {
	harina, azucar, levadura, huevos, inventarioInicial, recetas := setupTestEnvironment()

	// Aplicamos "Pan Casero" y "Bizcocho"
	recetasTest := []Receta{recetas[0], recetas[1]}
	inventarioFinal := inventarioInicial.Clone().aplicarAsignacion(recetasTest)

	inventarioEsperado := Inventario{
		ingredientes: map[Producto]uint64{
			*harina:   0, // 5 - (2+3) = 0
			*azucar:   2, // 4 - 2 = 2
			*levadura: 1, // 2 - 1 = 1
			*huevos:   4, // 6 - 2 = 4
		},
	}

	if !reflect.DeepEqual(inventarioFinal.ingredientes, inventarioEsperado.ingredientes) {
		t.Errorf("El inventario resultante no coincide con el esperado.\nEsperado: %v\nObtenido: %v", inventarioEsperado.ingredientes, inventarioFinal.ingredientes)
	}
}
