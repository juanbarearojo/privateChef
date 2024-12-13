package models

import (
	"reflect"
	"testing"
)

func TestRecetaGetIngredientes(t *testing.T) {
	env := setupTestEnvironment()

	ingredientesEsperados := map[Producto]uint64{
		*env.Harina:   2,
		*env.Levadura: 1,
	}
	ingredientesObtenidos := env.Recetas[0].GetIngredientes()
	if !reflect.DeepEqual(ingredientesObtenidos, ingredientesEsperados) {
		t.Errorf("Los ingredientes obtenidos no coinciden con los esperados.\nEsperados: %v\nObtenidos: %v", ingredientesEsperados, ingredientesObtenidos)
	}
}

func TestRecetaSumaPerecederos(t *testing.T) {
	env := setupTestEnvironment()

	sumaEsperada := uint64(1)
	sumaObtenida := env.Recetas[0].SumaPerecederos()
	if sumaObtenida != sumaEsperada {
		t.Errorf("La suma de perecederos es incorrecta.\nEsperada: %d\nObtenida: %d", sumaEsperada, sumaObtenida)
	}
}

func TestRecetaSumaPerecederosConMultiplesPerecederos(t *testing.T) {
	env := setupTestEnvironment()

	sumaEsperada := uint64(2)
	sumaObtenida := env.Recetas[1].SumaPerecederos()
	if sumaObtenida != sumaEsperada {
		t.Errorf("La suma de perecederos es incorrecta.\nEsperada: %d\nObtenida: %d", sumaEsperada, sumaObtenida)
	}
}

func TestRecetaSumaPerecederosSinPerecederos(t *testing.T) {
	env := setupTestEnvironment()

	sumaEsperada := uint64(0)
	sumaObtenida := env.Recetas[3].SumaPerecederos()
	if sumaObtenida != sumaEsperada {
		t.Errorf("La suma de perecederos debe ser 0 cuando no hay perecederos.\nEsperada: %d\nObtenida: %d", sumaEsperada, sumaObtenida)
	}
}
