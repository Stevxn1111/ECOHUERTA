package services

import (
	"ecohuerta-backend/models"
	"ecohuerta-backend/repositories"
	"errors"
	"strings"
)

type PlantaService struct {
	repo *repositories.PlantaRepository
}

func NewPlantaService(repo *repositories.PlantaRepository) *PlantaService {
	return &PlantaService{repo: repo}
}

func (s *PlantaService) ListarTiposPlanta() ([]models.TipoPlanta, error) {
	return s.repo.GetTiposPlanta()
}

func (s *PlantaService) ListarPlantasPorDepartamento(slug string, idTipo int) ([]models.Planta, error) {
	limpio := strings.TrimSpace(strings.ToLower(slug))
	if limpio == "" {
		return nil, errors.New("el slug del departamento es obligatorio")
	}
	return s.repo.GetByDepartamentoSlug(limpio, idTipo)
}

func (s *PlantaService) ObtenerPlantaPorID(id int) (*models.Planta, error) {
	if id <= 0 {
		return nil, errors.New("id de planta inválido")
	}
	return s.repo.GetByID(id)
}

func (s *PlantaService) CrearPlanta(p *models.Planta) (*models.Planta, error) {
	if err := validarPlanta(p); err != nil {
		return nil, err
	}
	id, err := s.repo.Create(p)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(id)
}

func (s *PlantaService) ActualizarPlanta(p *models.Planta) (*models.Planta, error) {
	if p.IDPlanta <= 0 {
		return nil, errors.New("id de planta inválido")
	}
	if err := validarPlanta(p); err != nil {
		return nil, err
	}
	if err := s.repo.Update(p); err != nil {
		return nil, err
	}
	return s.repo.GetByID(p.IDPlanta)
}

func (s *PlantaService) EliminarPlanta(id int) error {
	if id <= 0 {
		return errors.New("id de planta inválido")
	}
	return s.repo.Delete(id)
}

func validarPlanta(p *models.Planta) error {
	p.Nombre = strings.TrimSpace(p.Nombre)
	p.Descripcion = strings.TrimSpace(p.Descripcion)
	p.Caracteristicas = strings.TrimSpace(p.Caracteristicas)
	p.Cuidados = strings.TrimSpace(p.Cuidados)
	p.Beneficios = strings.TrimSpace(p.Beneficios)
	p.Imagen = strings.TrimSpace(p.Imagen)

	if p.Nombre == "" {
		return errors.New("el nombre de la planta es obligatorio")
	}
	if p.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}
	if p.IDTipo <= 0 {
		return errors.New("debes seleccionar un tipo de planta")
	}
	if p.IDDepartamento <= 0 {
		return errors.New("debes seleccionar un departamento")
	}
	return nil
}