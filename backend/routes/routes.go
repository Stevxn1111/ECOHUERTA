package routes

import (
	"ecohuerta-backend/handlers"
	"ecohuerta-backend/repositories"
	"ecohuerta-backend/services"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	deptoRepo := repositories.NewDepartamentoRepository()
	deptoService := services.NewDepartamentoService(deptoRepo)
	deptoHandler := handlers.NewDepartamentoHandler(deptoService)

	plantaRepo := repositories.NewPlantaRepository()
	plantaService := services.NewPlantaService(plantaRepo)
	plantaHandler := handlers.NewPlantaHandler(plantaService)

	// Endpoints de autenticación y gestión de usuarios
	mux.HandleFunc("/api/registro", userHandler.RegisterHandler)
	mux.HandleFunc("/api/login", userHandler.LoginHandler)
	mux.HandleFunc("/api/usuarios", userHandler.UsuariosHandler)
	mux.HandleFunc("/api/usuarios/", userHandler.UsuarioByIDHandler)

	// Endpoints de departamentos, tipos de planta y plantas
	mux.HandleFunc("/api/departamentos", deptoHandler.GetDepartamentos)
	mux.HandleFunc("/api/departamentos/", plantaHandler.GetPlantasPorDepartamento)
	mux.HandleFunc("/api/tipos-planta", plantaHandler.GetTiposPlanta)
	mux.HandleFunc("/api/plantas", plantaHandler.CreatePlanta)
	mux.HandleFunc("/api/plantas/", plantaHandler.PlantaByIDHandler)

	// Servir archivos estáticos del frontend (corriendo desde ./backend)
	fileServer := http.FileServer(http.Dir("../front"))
	mux.Handle("/", fileServer)

	return mux
}