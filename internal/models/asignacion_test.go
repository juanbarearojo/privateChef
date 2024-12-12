package models

import (
	"testing"
)

type TestEnv struct {
	Harina     *Producto
	Azucar     *Producto
	Levadura   *Producto
	Huevos     *Producto
	Inventario Inventario
	Recetas    []Receta
}

func setupTestEnvironment() *TestEnv {
	env := &TestEnv{}

	env.Harina = &Producto{nombre: "Harina", tipo: NoPerecedero}
	env.Azucar = &Producto{nombre: "Azúcar", tipo: NoPerecedero}
	env.Levadura = &Producto{nombre: "Levadura", tipo: Perecedero}
	env.Huevos = &Producto{nombre: "Huevos", tipo: Perecedero}

	env.Inventario = Inventario{
		ingredientes: map[Producto]uint64{
			*env.Harina:   5,
			*env.Azucar:   4,
			*env.Levadura: 2,
			*env.Huevos:   6,
		},
	}

	env.Recetas = []Receta{
		{
			titulo: "Pan Casero",
			ingredientes: map[Producto]uint64{
				*env.Harina:   2,
				*env.Levadura: 1,
			},
		},
		{
			titulo: "Bizcocho",
			ingredientes: map[Producto]uint64{
				*env.Harina: 3,
				*env.Azucar: 2,
				*env.Huevos: 2,
			},
		},
		{
			titulo: "Galletas",
			ingredientes: map[Producto]uint64{
				*env.Harina: 1,
				*env.Azucar: 1,
				*env.Huevos: 1,
			},
		},
		{
			titulo: "Polvo Dulce",
			ingredientes: map[Producto]uint64{
				*env.Harina: 2,
				*env.Azucar: 3,
			},
		},
	}

	return env
}

func TestSePuedePrepararPanCasero(t *testing.T) {
	env := setupTestEnvironment()
	receta := env.Recetas[0] // "Pan Casero"
	env.Inventario.ingredientes[*env.Harina] = 2
	env.Inventario.ingredientes[*env.Levadura] = 1

	puede := sePuedePreparar(receta, &env.Inventario)
	if !puede {
		t.Errorf("Se esperaba que 'Pan Casero' se pudiera preparar, pero no")
	}
}

func TestNoSePuedePrepararPanCasero(t *testing.T) {
	env := setupTestEnvironment()
	receta := env.Recetas[0] // "Pan Casero"
	env.Inventario.ingredientes[*env.Harina] = 2
	env.Inventario.ingredientes[*env.Levadura] = 0

	puede := sePuedePreparar(receta, &env.Inventario)
	if puede {
		t.Errorf("Se esperaba que 'Pan Casero' no se pudiera preparar, pero sí")
	}
}

func TestRealizarAsignacionListaNoVacia(t *testing.T) {
	env := setupTestEnvironment()
	recetasAsignables := realizarAsignacion(env.Recetas, env.Inventario)
	if len(recetasAsignables) == 0 {
		t.Errorf("realizarAsignacion devolvió una lista vacía de recetas")
	}
}

func TestRealizarAsignacionInventarioVacio(t *testing.T) {
	env := setupTestEnvironment()
	env.Inventario.ingredientes = make(map[Producto]uint64)
	recetasAsignables := realizarAsignacion(env.Recetas, env.Inventario)
	if len(recetasAsignables) != 0 {
		t.Errorf("Se esperaba una lista vacía de recetas asignables cuando el inventario está vacío, pero se obtuvo %d recetas", len(recetasAsignables))
	}
}

func TestRealizarAsignacionCantidadCorrecta(t *testing.T) {
	env := setupTestEnvironment()

	valorMinimoHacerPanHarina := uint64(2)
	valorMinimoHacerPanAzucar := uint64(1)
	valorMinimoHacerPanLevadura := uint64(1)
	env.Inventario.ingredientes[*env.Harina] = valorMinimoHacerPanHarina
	env.Inventario.ingredientes[*env.Azucar] = valorMinimoHacerPanAzucar
	env.Inventario.ingredientes[*env.Levadura] = valorMinimoHacerPanLevadura
	recetasAsignables := realizarAsignacion(env.Recetas, env.Inventario)

	numeroRecetasAsignablesEsperadas := 1
	if len(recetasAsignables) != numeroRecetasAsignablesEsperadas {
		t.Errorf("Se esperaba 1 receta asignable, se obtuvieron %d", len(recetasAsignables))
	}
}

func TestRealizarAsignacionUnicaPosibilidad(t *testing.T) {
	env := setupTestEnvironment()
	env.Inventario.ingredientes[*env.Harina] = 2
	env.Inventario.ingredientes[*env.Azucar] = 1
	env.Inventario.ingredientes[*env.Levadura] = 1
	recetasAsignables := realizarAsignacion(env.Recetas, env.Inventario)

	if len(recetasAsignables) == 0 {
		t.Fatalf("Se esperaba al menos una receta asignable, pero no se asignó ninguna")
	}

	recetaAsignada := recetasAsignables[0]
	if recetaAsignada.titulo != "Pan Casero" {
		t.Errorf("Se esperaba que la receta asignada fuera 'Pan Casero', pero fue '%s'", recetaAsignada.titulo)
	}
}

func ejecutarComparativaAsignacion(t *testing.T, indicesRecetas []int) {
	env := setupTestEnvironment()

	asignacionHumana := make([]Receta, len(indicesRecetas))
	for i, idx := range indicesRecetas {
		asignacionHumana[i] = env.Recetas[idx]
	}

	inventarioHumano := env.Inventario.Clone()
	inventarioHumano = inventarioHumano.aplicarAsignacion(asignacionHumana)
	desperdicioHumano := inventarioHumano.GetDesperdicio()
	asignacionOptima := realizarAsignacion(env.Recetas, env.Inventario)
	inventarioOptimo := env.Inventario.Clone()
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
	env := setupTestEnvironment()
	env.Inventario.ingredientes[*env.Harina] = 3
	env.Inventario.ingredientes[*env.Azucar] = 1
	env.Inventario.ingredientes[*env.Levadura] = 1
	env.Inventario.ingredientes[*env.Huevos] = 1
	recetasAsignables := realizarAsignacion(env.Recetas, env.Inventario)

	desperdicio := env.Inventario.Clone().aplicarAsignacion(recetasAsignables).GetDesperdicio()
	if desperdicio != 0 {
		t.Errorf("Se esperaba desperdicio cero, pero se obtuvo %d", desperdicio)
	}
}
