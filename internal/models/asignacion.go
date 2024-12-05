package models

func realizarAsignacion(recetas []Receta, inventario Inventario) []Receta {
	var mejorAsignacion []Receta
	desperdicioMinimo := inventario.GetDesperdicio()

	n := len(recetas)

	totalCombinaciones := 1 << n // 2^n combinaciones
	for i := 1; i < totalCombinaciones; i++ {
		var combinacion []Receta
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				combinacion = append(combinacion, recetas[j])
			}
		}
		inventarioTemporal := inventario.Clone()

		valida := true

		for _, receta := range combinacion {
			for producto, cantidadNecesaria := range receta.GetIngredientes() {
				cantidadDisponible, existe := inventarioTemporal.ingredientes[producto]
				if !existe || cantidadDisponible < cantidadNecesaria {
					valida = false
					break
				}
			}
			if !valida {
				break
			}

			for producto, cantidadNecesaria := range receta.GetIngredientes() {
				inventarioTemporal.ingredientes[producto] -= cantidadNecesaria
			}
		}

		if valida {
			desperdicioActual := inventarioTemporal.GetDesperdicio()
			if desperdicioActual < desperdicioMinimo {
				desperdicioMinimo = desperdicioActual
				mejorAsignacion = combinacion
			}
		}
	}

	return mejorAsignacion
}
