package handlers

import (
	"ecohuerta-backend/models"
	"ecohuerta-backend/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type PlantaHandler struct {
	service *services.PlantaService
}

func NewPlantaHandler(service *services.PlantaService) *PlantaHandler {
	return &PlantaHandler{service: service}
}

func esRolDocenteOAdmin(r *http.Request) bool {
	rol := strings.ToLower(strings.TrimSpace(r.Header.Get("X-User-Role")))
	return rol == "docente" || rol == "admin"
}

func (h *PlantaHandler) GetTiposPlanta(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	tipos, err := h.service.ListarTiposPlanta()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al obtener tipos de planta"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipos)
}

func (h *PlantaHandler) GetPlantasPorDepartamento(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/departamentos/")
	path = strings.TrimSuffix(path, "/plantas")
	slug := strings.Trim(path, "/")

	if slug == "" || strings.Contains(slug, "/") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Ruta de departamento inválida"})
		return
	}

	var idTipo int
	if tipoParam := r.URL.Query().Get("tipo"); tipoParam != "" {
		parsed, err := strconv.Atoi(tipoParam)
		if err == nil && parsed > 0 {
			idTipo = parsed
		}
	}

	plantas, err := h.service.ListarPlantasPorDepartamento(slug, idTipo)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al obtener las plantas del departamento"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plantas)
}

// CreatePlanta maneja POST /api/plantas
func (h *PlantaHandler) CreatePlanta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Método no permitido"})
		return
	}

	if !esRolDocenteOAdmin(r) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "Solo docentes o administradores pueden crear plantas"})
		return
	}

	var nueva models.Planta
	if err := json.NewDecoder(r.Body).Decode(&nueva); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos JSON inválidos"})
		return
	}

	creada, err := h.service.CrearPlanta(&nueva)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creada)
}

// PlantaByIDHandler maneja GET, PUT y DELETE en /api/plantas/{id}
func (h *PlantaHandler) PlantaByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/plantas/")
	idStr = strings.Trim(idStr, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID de planta inválido"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		planta, err := h.service.ObtenerPlantaPorID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Planta no encontrada"})
			return
		}
		json.NewEncoder(w).Encode(planta)

	case http.MethodPut:
		if !esRolDocenteOAdmin(r) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"error": "Solo docentes o administradores pueden editar plantas"})
			return
		}
		var editada models.Planta
		if err := json.NewDecoder(r.Body).Decode(&editada); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Datos JSON inválidos"})
			return
		}
		editada.IDPlanta = id
		actualizada, err := h.service.ActualizarPlanta(&editada)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(actualizada)

	case http.MethodDelete:
		if !esRolDocenteOAdmin(r) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"error": "Solo docentes o administradores pueden eliminar plantas"})
			return
		}
		if err := h.service.EliminarPlanta(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "No se pudo eliminar la planta"})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Planta eliminada correctamente"})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Método no permitido"})
	}
}