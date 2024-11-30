package models

import (
	"testing"
)

func testRealizarAsignacionListaNoVacia(t *testing.T) {
	harina, err := NewProducto("Harina", NoPerecedero, nil)
	if err != nil {
		t.Fatalf("Error al crear Harina: %v", err)
	}
	azúcar, err := NewProducto("Azúcar", NoPerecedero, nil)
	if err != nil {
		t.Fatalf("Error al crear Azúcar: %v", err)
	}

	fechaCaducidad := "15/12/2024"
	levadura, err := NewProducto("Levadura", Perecedero, &fechaCaducidad)
	if err != nil {
		t.Fatalf("Error al crear Levadura: %v", err)
	}

	huevos, err := NewProducto("Huevos", Perecedero, &fechaCaducidad)
	if err != nil {
		t.Fatalf("Error al crear Huevos: %v", err)
	}

	inventario := Inventario{
		ingredientes: map[Producto]uint64{
			*harina:   2,
			*azúcar:   1,
			*levadura: 1,
		},
	}

	recetas := []Receta{
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
				*azúcar: 2,
				*huevos: 2,
			},
		},
	}

	recetasAsignables := realizarAsignacion(recetas, inventario)

	if len(recetasAsignables) == 0 {
		t.Errorf("realizarAsignacion devolvió una lista vacía de recetas")
	}
}

func testRealizarAsignacionCantidadCorrecta(t *testing.T) {
	harina, err := NewProducto("Harina", NoPerecedero, nil)
	if err != nil {
		t.Fatalf("Error al crear Harina: %v", err)
	}

	azúcar, err := NewProducto("Azúcar", NoPerecedero, nil)
	if err != nil {
		t.Fatalf("Error al crear Azúcar: %v", err)
	}

	fechaCaducidad := "15/12/2024"
	levadura, err := NewProducto("Levadura", Perecedero, &fechaCaducidad)
	if err != nil {
		t.Fatalf("Error al crear Levadura: %v", err)
	}

	huevos, err := NewProducto("Huevos", Perecedero, &fechaCaducidad)
	if err != nil {
		t.Fatalf("Error al crear Huevos: %v", err)
	}

	inventario := Inventario{
		ingredientes: map[Producto]uint64{
			*harina:   2,
			*azúcar:   1,
			*levadura: 1,
		},
	}

	recetas := []Receta{
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
				*azúcar: 2,
				*huevos: 2,
			},
		},
	}

	recetasAsignables := realizarAsignacion(recetas, inventario)

	if len(recetasAsignables) != 1 {
		t.Errorf("Se esperaba 1 receta asignable, se obtuvieron %d", len(recetasAsignables))
	}
}

func testAsignacionComparativa(t *testing.T) {
	// Crear productos
	harina, err := NewProducto("Harina", NoPerecedero, nil)
	if err != nil {
		t.Fatalf("Error al crear Harina: %v", err)
	}

	azucar, err := NewProducto("Azúcar", NoPerecedero, nil)
	if err != nil {
		t.Fatalf("Error al crear Azúcar: %v", err)
	}

	fechaCaducidad := "15/12/2024"
	levadura, err := NewProducto("Levadura", Perecedero, &fechaCaducidad)
	if err != nil {
		t.Fatalf("Error al crear Levadura: %v", err)
	}

	huevos, err := NewProducto("Huevos", Perecedero, &fechaCaducidad)
	if err != nil {
		t.Fatalf("Error al crear Huevos: %v", err)
	}

	// Crear inventario inicial
	inventarioInicial := Inventario{
		ingredientes: map[Producto]uint64{
			*harina:   5,
			*azucar:   4,
			*levadura: 2,
			*huevos:   6,
		},
	}

	// Definir recetas
	recetas := []Receta{
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
	}

	// Asignación "humana" (subóptima): elige recetas sin considerar el mínimo desperdicio
	asignacionHumana := []Receta{
		recetas[0], // "Pan Casero"
		recetas[2], // "Galletas"
	}

	// Aplicar asignación humana al inventario
	inventarioHumano := inventarioInicial.Clone()
	inventarioHumano = aplicarAsignacion(asignacionHumana, inventarioHumano)

	// Calcular desperdicio después de la asignación humana
	desperdicioHumano := inventarioHumano.GetDesperdicio()

	// Obtener la asignación óptima utilizando realizarAsignacion
	asignacionOptima := realizarAsignacion(recetas, inventarioInicial)

	// Aplicar asignación óptima al inventario
	inventarioOptimo := inventarioInicial.Clone()
	inventarioOptimo = aplicarAsignacion(asignacionOptima, inventarioOptimo)

	// Calcular desperdicio después de la asignación óptima
	desperdicioOptimo := inventarioOptimo.GetDesperdicio()

	// Verificar que el desperdicio humano es mayor que el óptimo
	if desperdicioHumano <= desperdicioOptimo {
		t.Errorf("El desperdicio de la asignación humana (%d) no es mayor que el desperdicio óptimo (%d)", desperdicioHumano, desperdicioOptimo)
	}
}
