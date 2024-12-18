package models

import "sort"

func realizarAsignacion(recetas []Receta, inventario Inventario) []Receta {
	sort.Slice(recetas, func(i, j int) bool {
		return recetas[i].SumaPerecederos() > recetas[j].SumaPerecederos()
	})

	var asignacion []Receta
	inventarioTemporal := inventario.Clone()

	for _, receta := range recetas {
		if sePuedePreparar(receta, inventarioTemporal) {
			for producto, cantidadNecesaria := range receta.GetIngredientes() {
				inventarioTemporal.ingredientes[producto] -= cantidadNecesaria
			}
			asignacion = append(asignacion, receta)
		}
	}

	return asignacion
}

func sePuedePreparar(r Receta, inv *Inventario) bool {
	for producto, cantidadNecesaria := range r.GetIngredientes() {
		cantidadDisponible, existe := inv.ingredientes[producto]
		if !existe || cantidadDisponible < cantidadNecesaria {
			return false
		}
	}
	return true
}
