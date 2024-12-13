package models

import (
	"reflect"
	"testing"
)

func TestGetDesperdicioNoNegativo(t *testing.T) {
	env := setupTestEnvironment()

	desperdicio := env.Inventario.GetDesperdicio()
	if desperdicio < 0 {
		t.Errorf("GetDesperdicio devolvió un valor negativo: %d", desperdicio)
	}
}

func TestGetDesperdicioWithEmptyInventory(t *testing.T) {
	env := setupTestEnvironment()

	env.Inventario.ingredientes = make(map[Producto]uint64)

	desperdicio := env.Inventario.GetDesperdicio()
	if desperdicio != 0 {
		t.Errorf("GetDesperdicio devolvió %d, se esperaba 0 para inventario vacío", desperdicio)
	}
}

func TestGetDesperdicioValorCorrecto(t *testing.T) {
	env := setupTestEnvironment()

	desperdicioEsperado := uint64(8)
	desperdicioActual := env.Inventario.GetDesperdicio()
	if desperdicioActual != desperdicioEsperado {
		t.Errorf("GetDesperdicio devolvió %d, se esperaba %d", desperdicioActual, desperdicioEsperado)
	}
}

func TestInventarioCloneNoVacio(t *testing.T) {
	env := setupTestEnvironment()

	inventarioClonado := env.Inventario.Clone()
	if len(inventarioClonado.ingredientes) == 0 {
		t.Errorf("El inventario clonado está vacío")
	}
}

func TestInventarioCloneIgualAlOriginal(t *testing.T) {
	env := setupTestEnvironment()

	inventarioClonado := env.Inventario.Clone()
	if !reflect.DeepEqual(env.Inventario.ingredientes, inventarioClonado.ingredientes) {
		t.Errorf("El inventario clonado no es igual al original.\nOriginal: %v\nClonado: %v", env.Inventario.ingredientes, inventarioClonado.ingredientes)
	}
}

func TestAplicarAsignacionNoNull(t *testing.T) {
	env := setupTestEnvironment()

	recetasTest := []Receta{env.Recetas[0], env.Recetas[1]}
	nuevoInventario := env.Inventario.aplicarAsignacion(recetasTest)
	if nuevoInventario == nil {
		t.Errorf("aplicarAsignacion devolvió un inventario nulo")
	}
}

func TestAplicarAsignacionInventarioCorrecto(t *testing.T) {
	env := setupTestEnvironment()

	recetasTest := []Receta{env.Recetas[0], env.Recetas[1]}
	inventarioFinal := env.Inventario.Clone().aplicarAsignacion(recetasTest)
	inventarioEsperado := Inventario{
		ingredientes: map[Producto]uint64{
			*env.Harina:   0,
			*env.Azucar:   2,
			*env.Levadura: 1,
			*env.Huevos:   4,
		},
	}

	if !reflect.DeepEqual(inventarioFinal.ingredientes, inventarioEsperado.ingredientes) {
		t.Errorf("El inventario resultante no coincide con el esperado.\nEsperado: %v\nObtenido: %v", inventarioEsperado.ingredientes, inventarioFinal.ingredientes)
	}
}
