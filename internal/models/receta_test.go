package models

import (
	"testing"
)

func TestRealizarAsignacionListaNoVacia(t *testing.T) {
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

func TestRealizarAsignacionCantidadCorrecta(t *testing.T) {
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
