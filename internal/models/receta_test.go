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

func testSumaPerecederosHelper(t *testing.T, recetaIndex int, sumaEsperada uint64) {
	env := setupTestEnvironment()
	receta := env.Recetas[recetaIndex]
	sumaObtenida := receta.SumaPerecederos()

	if sumaObtenida != sumaEsperada {
		t.Errorf("La suma de perecederos para la receta '%s' es incorrecta.\nEsperada: %d\nObtenida: %d",
			receta.titulo, sumaEsperada, sumaObtenida)
	}
}

func TestRecetaSumaPerecederos(t *testing.T) {
	recetaPan := 0
	sumaEsperadaPan := uint64(1)
	testSumaPerecederosHelper(t, recetaPan, sumaEsperadaPan)
}

func TestRecetaSumaPerecederosConMultiplesPerecederos(t *testing.T) {
	recetaBizcocho := 1
	sumaEsperadaBizcocho := uint64(2)
	testSumaPerecederosHelper(t, recetaBizcocho, sumaEsperadaBizcocho)
}

func TestRecetaSumaPerecederosSinPerecederos(t *testing.T) {
	recetaPolvoDulce := 3
	sumaEsperadaPolvo := uint64(0)
	testSumaPerecederosHelper(t, recetaPolvoDulce, sumaEsperadaPolvo)
}
