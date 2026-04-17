package main

import (
	"log"
	"net/http"
	"user-api/internal/middleware"
	"github.com/gorilla/mux"
	"user-api/internal/handlers"
	"user-api/internal/services"
)

func main() {

	// Service
	service := services.NewUserService()

	// Handler
	handler := handlers.NewUserHandler(service)

	// Router
	r := mux.NewRouter()
	// Middleware
	r.Use(middleware.Logger)

	// Routes
	r.HandleFunc("/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/users", handler.AddUser).Methods("POST")
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}