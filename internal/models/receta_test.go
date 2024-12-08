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
