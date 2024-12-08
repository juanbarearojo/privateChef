package models

import (
	"reflect"
	"testing"
)

func TestRecetaGetIngredientes(t *testing.T) {
	harina, _, levadura, _, _, recetas := setupTestEnvironment()

	ingredientesEsperados := map[Producto]uint64{
		*harina:   2,
		*levadura: 1,
	}
	ingredientesObtenidos := recetas[0].GetIngredientes()
	if !reflect.DeepEqual(ingredientesObtenidos, ingredientesEsperados) {
		t.Errorf("Los ingredientes obtenidos no coinciden con los esperados.\nEsperados: %v\nObtenidos: %v", ingredientesEsperados, ingredientesObtenidos)
	}
}
func TestRecetaSumaPerecederos(t *testing.T) {
	_, _, _, _, _, recetas := setupTestEnvironment()

	sumaEsperada := uint64(1)
	sumaObtenida := recetas[0].SumaPerecederos()
	if sumaObtenida != sumaEsperada {
		t.Errorf("La suma de perecederos es incorrecta.\nEsperada: %d\nObtenida: %d", sumaEsperada, sumaObtenida)
	}
}

func TestRecetaSumaPerecederosConMultiplesPerecederos(t *testing.T) {
	_, _, _, _, _, recetas := setupTestEnvironment()

	sumaEsperada := uint64(2)
	sumaObtenida := recetas[1].SumaPerecederos()
	if sumaObtenida != sumaEsperada {
		t.Errorf("La suma de perecederos es incorrecta.\nEsperada: %d\nObtenida: %d", sumaEsperada, sumaObtenida)
	}
}

func TestRecetaSumaPerecederosSinPerecederos(t *testing.T) {
	_, _, _, _, _, recetas := setupTestEnvironment()
	sumaEsperada := uint64(0)
	sumaObtenida := recetas[3].SumaPerecederos()
	if sumaObtenida != sumaEsperada {
		t.Errorf("La suma de perecederos debe ser 0 cuando no hay perecederos.\nEsperada: %d\nObtenida: %d", sumaEsperada, sumaObtenida)
	}
}
