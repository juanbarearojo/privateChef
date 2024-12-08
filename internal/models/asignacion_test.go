package models

import (
	"testing"
)

func setupTestEnvironment() (harina, azucar, levadura, huevos *Producto, inventario Inventario, recetas []Receta) {
	harina, _ = NewProducto("Harina", NoPerecedero, nil)
	azucar, _ = NewProducto("Azúcar", NoPerecedero, nil)
	fechaCaducidad := "15/12/2024"
	levadura, _ = NewProducto("Levadura", Perecedero, &fechaCaducidad)
	huevos, _ = NewProducto("Huevos", Perecedero, &fechaCaducidad)
	inventario = Inventario{
		ingredientes: map[Producto]uint64{
			*harina:   5,
			*azucar:   4,
			*levadura: 2,
			*huevos:   6,
		},
	}
	recetas = []Receta{
		{
			titulo: "Pan Casero",
			ingredientes: map[Producto]uint64{
				*harina:   2,
				*levadura: 1,
			},
		},
		{
			titulo: "Bizcocho",
			ingredientes: map[Producto]uint64{
				*harina: 3,
				*azucar: 2,
				*huevos: 2,
			},
		},
		{
			titulo: "Galletas",
			ingredientes: map[Producto]uint64{
				*harina: 1,
				*azucar: 1,
				*huevos: 1,
			},
		},
		{
			titulo: "Polvo Dulce",
			ingredientes: map[Producto]uint64{
				*harina: 2,
				*azucar: 3,
			},
		},
	}
	return
}

func TestRealizarAsignacionListaNoVacia(t *testing.T) {
	harina, azucar, levadura, _, inventario, recetas := setupTestEnvironment()
	recetasTest := []Receta{recetas[0], recetas[1]}
	inventario.ingredientes[*harina] = 2
	inventario.ingredientes[*azucar] = 1
	inventario.ingredientes[*levadura] = 1
	recetasAsignables := realizarAsignacion(recetasTest, inventario)

	if len(recetasAsignables) == 0 {
		t.Errorf("realizarAsignacion devolvió una lista vacía de recetas")
	}
}

func TestRealizarAsignacionCantidadCorrecta(t *testing.T) {
	harina, azucar, levadura, _, inventario, recetas := setupTestEnvironment()
	recetasTest := []Receta{recetas[0], recetas[1]}
	inventario.ingredientes[*harina] = 2
	inventario.ingredientes[*azucar] = 1
	inventario.ingredientes[*levadura] = 1
	recetasAsignables := realizarAsignacion(recetasTest, inventario)

	if len(recetasAsignables) != 1 {
		t.Errorf("Se esperaba 1 receta asignable, se obtuvieron %d", len(recetasAsignables))
	}
}

func TestAsignacionComparativa(t *testing.T) {
	_, _, _, _, inventarioInicial, recetas := setupTestEnvironment()

	asignacionHumana := []Receta{
		recetas[0],
		recetas[2],
	}

	inventarioHumano := inventarioInicial.Clone()
	inventarioHumano = inventarioHumano.aplicarAsignacion(asignacionHumana)
	desperdicioHumano := inventarioHumano.GetDesperdicio()
	asignacionOptima := realizarAsignacion(recetas, inventarioInicial)
	inventarioOptimo := inventarioInicial.Clone()
	inventarioOptimo = inventarioOptimo.aplicarAsignacion(asignacionOptima)
	desperdicioOptimo := inventarioOptimo.GetDesperdicio()

	if desperdicioHumano <= desperdicioOptimo {
		t.Errorf("El desperdicio de la asignación humana (%d) no es mayor que el desperdicio óptimo (%d)", desperdicioHumano, desperdicioOptimo)
	}
}
