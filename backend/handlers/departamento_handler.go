package handlers

import (
	"ecohuerta-backend/services"
	"encoding/json"
	"net/http"
)

type DepartamentoHandler struct {
	service *services.DepartamentoService
}

func NewDepartamentoHandler(service *services.DepartamentoService) *DepartamentoHandler {
	return &DepartamentoHandler{service: service}
}

func (h *DepartamentoHandler) GetDepartamentos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	departamentos, err := h.service.ListarDepartamentos()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error al obtener departamentos"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(departamentos)
}