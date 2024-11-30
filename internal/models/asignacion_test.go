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