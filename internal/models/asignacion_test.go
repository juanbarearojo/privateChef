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

func TestSePuedePrepararPanCasero(t *testing.T) {
	harina, _, levadura, _, inventario, recetas := setupTestEnvironment()
	receta := recetas[0] // "Pan Casero"
	inventario.ingredientes[*harina] = 2
	inventario.ingredientes[*levadura] = 1

	puede := sePuedePreparar(receta, &inventario)
	if !puede {
		t.Errorf("Se esperaba que 'Pan Casero' se pudiera preparar, pero no")
	}
}

func TestNoSePuedePrepararPanCasero(t *testing.T) {
	harina, _, levadura, _, inventario, recetas := setupTestEnvironment()
	receta := recetas[0] // "Pan Casero"
	inventario.ingredientes[*harina] = 2
	inventario.ingredientes[*levadura] = 0

	puede := sePuedePreparar(receta, &inventario)
	if puede {
		t.Errorf("Se esperaba que 'Pan Casero' no se pudiera preparar, pero sí")
	}
}

func TestRealizarAsignacionListaNoVacia(t *testing.T) {
	_, _, _, _, inventario, recetas := setupTestEnvironment()
	recetasAsignables := realizarAsignacion(recetas, inventario)
	if len(recetasAsignables) == 0 {
		t.Errorf("realizarAsignacion devolvió una lista vacía de recetas")
	}
}

func TestRealizarAsignacionInventarioVacio(t *testing.T) {
	_, _, _, _, inventario, recetas := setupTestEnvironment()
	inventario.ingredientes = make(map[Producto]uint64)
	recetasAsignables := realizarAsignacion(recetas, inventario)
	if len(recetasAsignables) != 0 {
		t.Errorf("Se esperaba una lista vacía de recetas asignables cuando el inventario está vacío, pero se obtuvo %d recetas", len(recetasAsignables))
	}
}

func TestRealizarAsignacionCantidadCorrecta(t *testing.T) {
	harina, azucar, levadura, _, inventario, recetas := setupTestEnvironment()

	valorMinimoHacerPanHarina := uint64(2)
	valorMinimoHacerPanAzucar := uint64(1)
	valorMinimoHacerPanLevadura := uint64(1)
	inventario.ingredientes[*harina] = valorMinimoHacerPanHarina
	inventario.ingredientes[*azucar] = valorMinimoHacerPanAzucar
	inventario.ingredientes[*levadura] = valorMinimoHacerPanLevadura
	recetasAsignables := realizarAsignacion(recetas, inventario)

	numeroRecetasAsignablesEsperadas := 1
	if len(recetasAsignables) != numeroRecetasAsignablesEsperadas {
		t.Errorf("Se esperaba 1 receta asignable, se obtuvieron %d", len(recetasAsignables))
	}
}

func TestRealizarAsignacionUnicaPosibilidad(t *testing.T) {
	harina, azucar, levadura, _, inventario, recetas := setupTestEnvironment()
	inventario.ingredientes[*harina] = 2
	inventario.ingredientes[*azucar] = 1
	inventario.ingredientes[*levadura] = 1
	recetasAsignables := realizarAsignacion(recetas, inventario)

	recetaAsignada := recetasAsignables[0]
	if recetaAsignada.titulo != "Pan Casero" {
		t.Errorf("Se esperaba que la receta asignada fuera 'Pan Casero', pero fue '%s'", recetaAsignada.titulo)
	}
}

func ejecutarComparativaAsignacion(t *testing.T, indicesRecetas []int) {
	_, _, _, _, inventarioInicial, recetas := setupTestEnvironment()

	asignacionHumana := make([]Receta, len(indicesRecetas))
	for i, idx := range indicesRecetas {
		asignacionHumana[i] = recetas[idx]
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

func TestRealizarAsignacionComparativa(t *testing.T) {
	indicesRecetas := []int{0, 2}
	ejecutarComparativaAsignacion(t, indicesRecetas)
}

func TestRealizarAsignacionComparativa2(t *testing.T) {
	indicesRecetas := []int{1, 3}
	ejecutarComparativaAsignacion(t, indicesRecetas)
}

func TestRealizarAsignacionComparativa3(t *testing.T) {
	indicesRecetas := []int{2, 3}
	ejecutarComparativaAsignacion(t, indicesRecetas)
}

func TestRealizarAsignacionDesperdicioCero(t *testing.T) {
	harina, azucar, levadura, huevos, inventario, recetas := setupTestEnvironment()
	inventario.ingredientes[*harina] = 3
	inventario.ingredientes[*azucar] = 1
	inventario.ingredientes[*levadura] = 1
	inventario.ingredientes[*huevos] = 1
	recetasAsignables := realizarAsignacion(recetas, inventario)

	desperdicio := inventario.Clone().aplicarAsignacion(recetasAsignables).GetDesperdicio()
	if desperdicio != 0 {
		t.Errorf("Se esperaba desperdicio cero, pero se obtuvo %d", desperdicio)
	}
}
