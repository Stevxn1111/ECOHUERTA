package models

type User struct {
	ID      int    `json:"id"`
	Nombre  string `json:"nombre"`
	Usuario string `json:"usuario"`
	Clave   string `json:"clave,omitempty"`
	Rol     string `json:"rol"`
}