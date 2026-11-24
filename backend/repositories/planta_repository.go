package repositories

import (
	"database/sql"
	"ecohuerta-backend/config"
	"ecohuerta-backend/models"
)

type PlantaRepository struct{}

func NewPlantaRepository() *PlantaRepository {
	return &PlantaRepository{}
}

func (r *PlantaRepository) GetTiposPlanta() ([]models.TipoPlanta, error) {
	rows, err := config.DB.Query("SELECT id_tipo, nombre FROM tipos_planta ORDER BY nombre ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tipos []models.TipoPlanta
	for rows.Next() {
		var t models.TipoPlanta
		if err := rows.Scan(&t.IDTipo, &t.Nombre); err != nil {
			return nil, err
		}
		tipos = append(tipos, t)
	}
	return tipos, rows.Err()
}

func (r *PlantaRepository) GetByDepartamentoSlug(slug string, idTipo int) ([]models.Planta, error) {
	query := `
		SELECT
			p.id_planta,
			p.nombre,
			p.descripcion,
			p.caracteristicas,
			p.cuidados,
			p.beneficios,
			p.id_tipo,
			COALESCE(tp.nombre, '') AS tipo_nombre,
			p.id_departamento,
			d.nombre AS departamento_nombre,
			d.slug AS departamento_slug,
			COALESCE(p.imagen, '') AS imagen
		FROM plantas p
		INNER JOIN departamentos d ON p.id_departamento = d.id_departamento
		LEFT JOIN tipos_planta tp ON p.id_tipo = tp.id_tipo
		WHERE d.slug = ?`

	args := []interface{}{slug}
	if idTipo > 0 {
		query += " AND p.id_tipo = ?"
		args = append(args, idTipo)
	}
	query += " ORDER BY p.nombre ASC"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plantas := make([]models.Planta, 0)
	for rows.Next() {
		var p models.Planta
		var img sql.NullString
		if err := rows.Scan(
			&p.IDPlanta,
			&p.Nombre,
			&p.Descripcion,
			&p.Caracteristicas,
			&p.Cuidados,
			&p.Beneficios,
			&p.IDTipo,
			&p.TipoNombre,
			&p.IDDepartamento,
			&p.DeptoNombre,
			&p.DeptoSlug,
			&img,
		); err != nil {
			return nil, err
		}
		if img.Valid {
			p.Imagen = img.String
		}
		plantas = append(plantas, p)
	}

	return plantas, rows.Err()
}

func (r *PlantaRepository) GetByID(id int) (*models.Planta, error) {
	query := `
		SELECT
			p.id_planta,
			p.nombre,
			p.descripcion,
			p.caracteristicas,
			p.cuidados,
			p.beneficios,
			p.id_tipo,
			COALESCE(tp.nombre, '') AS tipo_nombre,
			p.id_departamento,
			d.nombre AS departamento_nombre,
			d.slug AS departamento_slug,
			COALESCE(p.imagen, '') AS imagen
		FROM plantas p
		INNER JOIN departamentos d ON p.id_departamento = d.id_departamento
		LEFT JOIN tipos_planta tp ON p.id_tipo = tp.id_tipo
		WHERE p.id_planta = ?`

	var p models.Planta
	var img sql.NullString
	err := config.DB.QueryRow(query, id).Scan(
		&p.IDPlanta,
		&p.Nombre,
		&p.Descripcion,
		&p.Caracteristicas,
		&p.Cuidados,
		&p.Beneficios,
		&p.IDTipo,
		&p.TipoNombre,
		&p.IDDepartamento,
		&p.DeptoNombre,
		&p.DeptoSlug,
		&img,
	)
	if err != nil {
		return nil, err
	}
	if img.Valid {
		p.Imagen = img.String
	}
	return &p, nil
}

func (r *PlantaRepository) Create(p *models.Planta) (int, error) {
	res, err := config.DB.Exec(`
		INSERT INTO plantas (nombre, descripcion, caracteristicas, cuidados, beneficios, id_tipo, id_departamento, imagen)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		p.Nombre, p.Descripcion, p.Caracteristicas, p.Cuidados, p.Beneficios, p.IDTipo, p.IDDepartamento, p.Imagen,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func (r *PlantaRepository) Update(p *models.Planta) error {
	_, err := config.DB.Exec(`
		UPDATE plantas
		SET nombre = ?, descripcion = ?, caracteristicas = ?, cuidados = ?, beneficios = ?, id_tipo = ?, id_departamento = ?, imagen = ?
		WHERE id_planta = ?`,
		p.Nombre, p.Descripcion, p.Caracteristicas, p.Cuidados, p.Beneficios, p.IDTipo, p.IDDepartamento, p.Imagen, p.IDPlanta,
	)
	return err
}

func (r *PlantaRepository) Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM plantas WHERE id_planta = ?", id)
	return err
}