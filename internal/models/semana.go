package models

type DiaSemana string

const (
	Lunes     DiaSemana = "Lunes"
	Martes    DiaSemana = "Martes"
	Miercoles DiaSemana = "Miercoles"
	Jueves    DiaSemana = "Jueves"
	Viernes   DiaSemana = "Viernes"
	Sabado    DiaSemana = "Sabado"
	Domingo   DiaSemana = "Domingo"
)

type Semana struct {
	DiaPreparacion DiaSemana
}
