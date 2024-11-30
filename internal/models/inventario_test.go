package models

import (
	"reflect"
	"testing"
)

func TestGetDesperdicioNoNegativo(t *testing.T) {
	fechaCaducidadmanzana := "15/12/2024"
	producto := Producto{
		nombre:         "Manzana",
		tipo:           Perecedero,
		fechaCaducidad: &fechaCaducidadmanzana,
	}

	inventario := &Inventario{
		ingredientes: map[Producto]uint64{
			producto: 10, // Cantidad positiva
		},
	}

	desperdicio := inventario.GetDesperdicio()

	if desperdicio < 0 {
		t.Errorf("GetDesperdicio devolvió un valor negativo: %d", desperdicio)
	}
}

func TestAplicarAsignacionNoNull(t *testing.T) {
	// Crear productos
	fechaCaducidad := "15/12/2024"
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

	levadura := Producto{
		nombre:         "Levadura",
		tipo:           Perecedero,
		fechaCaducidad: &fechaCaducidad,
	}

	huevos := Producto{
		nombre:         "Huevos",
		tipo:           Perecedero,
		fechaCaducidad: &fechaCaducidad,
	}

	// Crear inventario inicial
	inventario := Inventario{
		ingredientes: map[Producto]uint64{
			harina:   5,
			azucar:   3,
			levadura: 2,
			huevos:   6,
		},
	}

	// Crear lista de recetas
	recetas := []Receta{
		{
			titulo: "Pan Casero",
			ingredientes: map[Producto]uint64{
				harina:   2,
				levadura: 1,
			},
		},
		{
			titulo: "Bizcocho",
			ingredientes: map[Producto]uint64{
				harina: 3,
				azucar: 2,
				huevos: 2,
			},
		},
	}

	// Llamar a la función aplicarAsignacion
	nuevoInventario := inventario.aplicarAsignacion(recetas)

	// Verificar que el nuevo inventario no es nulo
	if nuevoInventario == nil {
		t.Errorf("aplicarAsignacion devolvió un inventario nulo")
	}
}

func TestAplicarAsignacionInventarioCorrecto(t *testing.T) {
	fechaCaducidad := "15/12/2024"
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

	levadura := Producto{
		nombre:         "Levadura",
		tipo:           Perecedero,
		fechaCaducidad: &fechaCaducidad,
	}

	huevos := Producto{
		nombre:         "Huevos",
		tipo:           Perecedero,
		fechaCaducidad: &fechaCaducidad,
	}

	inventarioInicial := Inventario{
		ingredientes: map[Producto]uint64{
			harina:   5,
			azucar:   3,
			levadura: 2,
			huevos:   6,
		},
	}

	recetas := []Receta{
		{
			titulo: "Pan Casero",
			ingredientes: map[Producto]uint64{
				harina:   2,
				levadura: 1,
			},
		},
		{
			titulo: "Bizcocho",
			ingredientes: map[Producto]uint64{
				harina: 3,
				azucar: 2,
				huevos: 2,
			},
		},
	}

	inventarioEsperado := Inventario{
		ingredientes: map[Producto]uint64{
			harina:   0, // 5 - (2 + 3) = 0
			azucar:   1, // 3 - 2 = 1
			levadura: 1, // 2 - 1 = 1
			huevos:   4, // 6 - 2 = 4
		},
	}

	nuevoInventario := inventarioInicial.Clone()
	nuevoInventario = aplicarAsignacion(recetas, nuevoInventario)

	if !reflect.DeepEqual(nuevoInventario.ingredientes, inventarioEsperado.ingredientes) {
		t.Errorf("El inventario resultante no coincide con el esperado.\nEsperado: %v\nObtenido: %v", inventarioEsperado.ingredientes, nuevoInventario.ingredientes)
	}
}
