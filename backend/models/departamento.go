package models

type Departamento struct {
	IDDepartamento int    `json:"id_departamento"`
	Nombre         string `json:"nombre"`
	Slug           string `json:"slug"`
}