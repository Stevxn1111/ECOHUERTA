package models

type Planta struct {
	IDPlanta        int    `json:"id_planta"`
	Nombre          string `json:"nombre"`
	Descripcion     string `json:"descripcion"`
	Caracteristicas string `json:"caracteristicas"`
	Cuidados        string `json:"cuidados"`
	Beneficios      string `json:"beneficios"`
	IDTipo          int    `json:"id_tipo"`
	TipoNombre      string `json:"tipo_nombre"`
	IDDepartamento  int    `json:"id_departamento"`
	DeptoNombre     string `json:"departamento_nombre"`
	DeptoSlug       string `json:"departamento_slug"`
	Imagen          string `json:"imagen"`
}