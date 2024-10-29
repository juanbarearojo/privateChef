package models

type Receta struct {
	nombre       string
	tipoCocina   string
	ingredientes []Ingrediente
}

func NewReceta(nombre string, tipoCocina string, ingredientes []Ingrediente) Receta {
	return Receta{
		nombre:       nombre,
		tipoCocina:   tipoCocina,
		ingredientes: ingredientes,
	}
}

func (r Receta) GetNombre() string {
	return r.nombre
}

func (r Receta) GetTipoCocina() string {
	return r.tipoCocina
}

func (r Receta) GetIngredientes() []Ingrediente {
	return r.ingredientes
}

func (r *Receta) SetNombre(nombre string) {
	r.nombre = nombre
}

func (r *Receta) SetTipoCocina(tipoCocina string) {
	r.tipoCocina = tipoCocina
}

func (r *Receta) SetIngredientes(ingredientes []Ingrediente) {
	r.ingredientes = ingredientes
}
