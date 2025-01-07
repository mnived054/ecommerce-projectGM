package handlers

import (
	"ecommerce-project/backend/models"
	"ecommerce-project/backend/services"
	"encoding/json"
	"net/http"
)

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID := 1

	profile, err := services.GetProfile(userID)
	if err != nil {
		http.Error(w, "Error retrieving profile:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID := 1

	var profile models.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updateProfile, err := services.UpdateProfile(userID, profile)
	if err != nil {
		http.Error(w, "Error updating profile:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateProfile)
}
