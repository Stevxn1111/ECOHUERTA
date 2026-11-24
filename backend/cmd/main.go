package main

import (
	"ecohuerta-backend/config"
	"ecohuerta-backend/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	mux := routes.SetupRoutes()

	log.Println("Servidor corriendo en el puerto 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}