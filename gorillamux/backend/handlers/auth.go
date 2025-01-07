package handlers

import (
	"ecommerce-project/backend/models"
	"ecommerce-project/backend/services"
	"encoding/json"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.RegisterUser(user)
	if err != nil {
		http.Error(w, "Error signing up", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	token, err := services.LoginUser(credentials)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Authorization", "Bearer"+token)
	w.WriteHeader(http.StatusOK)

}
