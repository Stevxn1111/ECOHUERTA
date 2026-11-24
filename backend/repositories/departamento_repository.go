package repositories

import (
	"ecohuerta-backend/config"
	"ecohuerta-backend/models"
)

type DepartamentoRepository struct{}

func NewDepartamentoRepository() *DepartamentoRepository {
	return &DepartamentoRepository{}
}

func (r *DepartamentoRepository) GetAll() ([]models.Departamento, error) {
	rows, err := config.DB.Query("SELECT id_departamento, nombre, slug FROM departamentos ORDER BY nombre ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Departamento
	for rows.Next() {
		var d models.Departamento
		if err := rows.Scan(&d.IDDepartamento, &d.Nombre, &d.Slug); err != nil {
			return nil, err
		}
		lista = append(lista, d)
	}

	return lista, rows.Err()
}