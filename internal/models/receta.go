package models

type Receta struct {
	nombre        string
	tipoCocina    string
	ingredientes  []Ingrediente
	instrucciones string
}
