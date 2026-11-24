package repositories

import (
	"database/sql"
	"ecohuerta-backend/config"
	"ecohuerta-backend/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *models.User) error {
	query := "INSERT INTO users (nombre, usuario, clave, rol) VALUES (?, ?, ?, ?)"
	result, err := config.DB.Exec(query, user.Nombre, user.Usuario, user.Clave, user.Rol)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

func (r *UserRepository) GetByUsuario(usuario string) (*models.User, error) {
	query := "SELECT id, nombre, usuario, clave, rol FROM users WHERE usuario = ?"
	var user models.User

	err := config.DB.QueryRow(query, usuario).Scan(
		&user.ID,
		&user.Nombre,
		&user.Usuario,
		&user.Clave,
		&user.Rol,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	query := "SELECT id, nombre, usuario, rol FROM users WHERE id = ?"
	var user models.User

	err := config.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Nombre,
		&user.Usuario,
		&user.Rol,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) List(rol string) ([]models.User, error) {
	var rows *sql.Rows
	var err error

	if rol != "" && rol != "todos" {
		query := "SELECT id, nombre, usuario, rol FROM users WHERE rol = ? ORDER BY id DESC"
		rows, err = config.DB.Query(query, rol)
	} else {
		query := "SELECT id, nombre, usuario, rol FROM users ORDER BY id DESC"
		rows, err = config.DB.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Nombre, &u.Usuario, &u.Rol); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(user *models.User) error {
	if user.Clave != "" {
		query := "UPDATE users SET nombre = ?, usuario = ?, clave = ?, rol = ? WHERE id = ?"
		_, err := config.DB.Exec(query, user.Nombre, user.Usuario, user.Clave, user.Rol, user.ID)
		return err
	}

	query := "UPDATE users SET nombre = ?, usuario = ?, rol = ? WHERE id = ?"
	_, err := config.DB.Exec(query, user.Nombre, user.Usuario, user.Rol, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := config.DB.Exec(query, id)
	return err
}