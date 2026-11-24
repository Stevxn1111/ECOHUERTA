package services

import (
	"ecohuerta-backend/models"
	"ecohuerta-backend/repositories"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func isValidRol(rol string) bool {
	switch rol {
	case "estudiante", "docente", "admin":
		return true
	default:
		return false
	}
}

func (s *UserService) RegisterUser(nombre, usuario, clave, rol string) error {
	nombre = strings.TrimSpace(nombre)
	usuario = strings.TrimSpace(usuario)
	rol = strings.TrimSpace(strings.ToLower(rol))

	if !isValidRol(rol) {
		return errors.New("rol no válido")
	}

	existente, err := s.repo.GetByUsuario(usuario)
	if err != nil {
		return err
	}
	if existente != nil {
		return errors.New("el nombre de usuario ya está en uso")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(clave), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Nombre:  nombre,
		Usuario: usuario,
		Clave:   string(hashedPassword),
		Rol:     rol,
	}

	return s.repo.Create(user)
}

func (s *UserService) LoginUser(usuario, clave string) (*models.User, error) {
	user, err := s.repo.GetByUsuario(strings.TrimSpace(usuario))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuario o contraseña incorrectos")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Clave), []byte(clave))
	if err != nil {
		return nil, errors.New("usuario o contraseña incorrectos")
	}

	return user, nil
}

func (s *UserService) ListUsers(rol string) ([]models.User, error) {
	rol = strings.TrimSpace(strings.ToLower(rol))
	if rol != "" && rol != "todos" && !isValidRol(rol) {
		return nil, errors.New("filtro de rol no válido")
	}
	return s.repo.List(rol)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUser(id int, nombre, usuario, clave, rol string) (*models.User, error) {
	nombre = strings.TrimSpace(nombre)
	usuario = strings.TrimSpace(usuario)
	rol = strings.TrimSpace(strings.ToLower(rol))

	if !isValidRol(rol) {
		return nil, errors.New("rol no válido")
	}

	actual, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if actual == nil {
		return nil, errors.New("usuario no encontrado")
	}

	otro, err := s.repo.GetByUsuario(usuario)
	if err != nil {
		return nil, err
	}
	if otro != nil && otro.ID != id {
		return nil, errors.New("el nombre de usuario ya está en uso")
	}

	user := &models.User{
		ID:      id,
		Nombre:  nombre,
		Usuario: usuario,
		Rol:     rol,
	}

	if strings.TrimSpace(clave) != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(clave), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Clave = string(hashedPassword)
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return s.repo.GetByID(id)
}

func (s *UserService) DeleteUser(id int) error {
	actual, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if actual == nil {
		return errors.New("usuario no encontrado")
	}
	return s.repo.Delete(id)
}