package handlers

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"user-api/internal/models"
	"user-api/internal/services"
	"strconv"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetUsers()

	writeJSON(w, users, http.StatusOK)
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// request validation
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeError(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		writeError(w, "Name is required", http.StatusBadRequest)
		return
	}

	if user.Age <= 0 {
		writeError(w, "Age must be greater than 0", http.StatusBadRequest)
		return
	}

	// call service to add user
	h.service.AddUser(user)

	writeJSON(w, user, http.StatusCreated)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		writeError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		writeError(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, nil, http.StatusOK)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	// convert id to int 
	idInt, err1 := strconv.Atoi(id)

	if err1 != nil {
		writeError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User

	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		writeError(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// validation
	if updatedUser.Name == "" || updatedUser.Age <= 0 {
		writeError(w, "Invalid data", http.StatusBadRequest)
		return
	}

	ok := h.service.UpdateUser(idInt, updatedUser)

	if !ok {
		writeError(w, "User not found", http.StatusNotFound)
		return
	}

	writeJSON(w, updatedUser, http.StatusOK)
}

// Helper function to write error responses
func writeError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  nil,
		"error": message,
	})
}

// Helper function to write success responses
func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  data,
		"error": nil,
	})
}
