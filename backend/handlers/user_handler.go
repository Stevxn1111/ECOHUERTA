package handlers

import (
	"ecohuerta-backend/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	service *services.UserService
}

type loginRequest struct {
	Usuario string `json:"usuario"`
	Clave   string `json:"clave"`
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Método no permitido"})
		return
	}

	var req struct {
		Nombre  string `json:"nombre"`
		Usuario string `json:"usuario"`
		Clave   string `json:"clave"`
		Rol     string `json:"rol"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Datos JSON inválidos"})
		return
	}

	if req.Nombre == "" || req.Usuario == "" || req.Clave == "" || req.Rol == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Todos los campos son obligatorios"})
		return
	}

	if err := h.service.RegisterUser(req.Nombre, req.Usuario, req.Clave, req.Rol); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "No se pudo registrar el usuario: " + err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"mensaje": "Usuario registrado exitosamente"})
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Método no permitido"})
		return
	}

	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Datos JSON inválidos"})
		return
	}

	if req.Usuario == "" || req.Clave == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Ingresa usuario y contraseña"})
		return
	}

	user, err := h.service.LoginUser(req.Usuario, req.Clave)
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"mensaje": "Inicio de sesión exitoso",
		"usuario": map[string]any{
			"id":      user.ID,
			"nombre":  user.Nombre,
			"usuario": user.Usuario,
			"rol":     user.Rol,
		},
	})
}

func (h *UserHandler) UsuariosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rol := r.URL.Query().Get("rol")
		users, err := h.service.ListUsers(rol)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, users)

	case http.MethodPost:
		var req struct {
			Nombre  string `json:"nombre"`
			Usuario string `json:"usuario"`
			Clave   string `json:"clave"`
			Rol     string `json:"rol"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Datos JSON inválidos"})
			return
		}

		if strings.TrimSpace(req.Nombre) == "" || strings.TrimSpace(req.Usuario) == "" || strings.TrimSpace(req.Clave) == "" || strings.TrimSpace(req.Rol) == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Nombre, usuario, contraseña y rol son obligatorios"})
			return
		}

		if err := h.service.RegisterUser(req.Nombre, req.Usuario, req.Clave, req.Rol); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusCreated, map[string]string{"mensaje": "Usuario creado exitosamente"})

	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Método no permitido"})
	}
}

func (h *UserHandler) UsuarioByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/usuarios/")
	idStr = strings.Trim(idStr, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "ID de usuario inválido"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, err := h.service.GetUserByID(id)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		if user == nil {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Usuario no encontrado"})
			return
		}
		writeJSON(w, http.StatusOK, user)

	case http.MethodPut, http.MethodPatch:
		var req struct {
			Nombre  string `json:"nombre"`
			Usuario string `json:"usuario"`
			Clave   string `json:"clave"`
			Rol     string `json:"rol"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Datos JSON inválidos"})
			return
		}

		if strings.TrimSpace(req.Nombre) == "" || strings.TrimSpace(req.Usuario) == "" || strings.TrimSpace(req.Rol) == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Nombre, usuario y rol son obligatorios"})
			return
		}

		actualizado, err := h.service.UpdateUser(id, req.Nombre, req.Usuario, req.Clave, req.Rol)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, actualizado)

	case http.MethodDelete:
		if err := h.service.DeleteUser(id); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"mensaje": "Usuario eliminado exitosamente"})

	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Método no permitido"})
	}
}