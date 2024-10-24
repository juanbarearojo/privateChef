package models

type Receta struct {
	Nombre        string
	TipoCocina    string
	Ingredientes  []Ingrediente
	Instrucciones string
}
