package models

type DiaSemana int

const (
	Lunes DiaSemana = iota
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
	Domingo
)

type Semana struct {
	diaPreparacion DiaSemana
}
