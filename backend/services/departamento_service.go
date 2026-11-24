package services

import (
	"ecohuerta-backend/models"
	"ecohuerta-backend/repositories"
)

type DepartamentoService struct {
	repo *repositories.DepartamentoRepository
}

func NewDepartamentoService(repo *repositories.DepartamentoRepository) *DepartamentoService {
	return &DepartamentoService{repo: repo}
}

func (s *DepartamentoService) ListarDepartamentos() ([]models.Departamento, error) {
	return s.repo.GetAll()
}