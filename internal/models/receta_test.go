package models

import (
	"reflect"
	"testing"
)

func TestRecetaGetIngredientes(t *testing.T) {
	env := setupTestEnvironment()
	valorMinimoPanHarina := uint64(2)
	valorMinimoPanLevadura := uint64(1)
	ingredientesEsperados := map[Producto]uint64{
		*env.Harina:   valorMinimoPanHarina,
		*env.Levadura: valorMinimoPanLevadura,
	}
	ingredientesObtenidos := env.Recetas[RecetaPanCasero].GetIngredientes()
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
	sumaEsperadaPan := uint64(1)
	testSumaPerecederosHelper(t, RecetaPanCasero, sumaEsperadaPan)
}

func TestRecetaSumaPerecederosConMultiplesPerecederos(t *testing.T) {
	sumaEsperadaBizcocho := uint64(2)
	testSumaPerecederosHelper(t, RecetaBizcocho, sumaEsperadaBizcocho)
}

func TestRecetaSumaPerecederosSinPerecederos(t *testing.T) {
	sumaEsperadaPolvo := uint64(0)
	testSumaPerecederosHelper(t, RecetaPolvoDulce, sumaEsperadaPolvo)
}
