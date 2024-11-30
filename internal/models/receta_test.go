package models

import (
	"reflect"
	"testing"
)

func testRecetaGetIngredientes(t *testing.T) {
	// Crear productos de ejemplo
	harina := Producto{
		nombre:         "Harina",
		tipo:           NoPerecedero,
		fechaCaducidad: nil,
	}

	azucar := Producto{
		nombre:         "Azúcar",
		tipo:           NoPerecedero,
		fechaCaducidad: nil,
	}

	// Crear un mapa de ingredientes para la receta
	ingredientesEsperados := map[Producto]uint64{
		harina: 2,
		azucar: 1,
	}

	// Crear la receta con los ingredientes
	receta := Receta{
		titulo:       "Pastel de Harina",
		ingredientes: ingredientesEsperados,
	}

	// Obtener los ingredientes usando el método GetIngredientes
	ingredientesObtenidos := receta.GetIngredientes()

	// Verificar que los ingredientes obtenidos son iguales a los esperados
	if !reflect.DeepEqual(ingredientesObtenidos, ingredientesEsperados) {
		t.Errorf("Los ingredientes obtenidos no coinciden con los esperados.\nEsperados: %v\nObtenidos: %v", ingredientesEsperados, ingredientesObtenidos)
	}
}
