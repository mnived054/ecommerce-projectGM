package handlers

import (
	"ecommerce-project/backend/models"
	"ecommerce-project/backend/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func AddAddressHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate user ID (this would come from a session or JWT token in a real application)
	userID := 1

	var address models.Address
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Add the address
	createdAddress, err := services.AddAddress(userID, address)
	if err != nil {
		http.Error(w, "Error adding address: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdAddress)
}

func GetAddressesHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate user ID (this would come from a session or JWT token in a real application)
	userID := 1

	addresses, err := services.GetAddresses(userID)
	if err != nil {
		http.Error(w, "Error retrieving addresses: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addresses)
}

func UpdateAddressHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate user ID (this would come from a session or JWT token in a real application)
	userID := 1

	// Get address ID from URL parameters
	addressID, err := strconv.Atoi(r.URL.Query().Get("address_id"))
	if err != nil {
		http.Error(w, "Invalid address ID", http.StatusBadRequest)
		return
	}

	var newAddress models.Address
	err = json.NewDecoder(r.Body).Decode(&newAddress)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update the address
	updatedAddress, err := services.UpdateAddress(userID, addressID, newAddress)
	if err != nil {
		http.Error(w, "Error updating address: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedAddress)
}

func DeleteAddressHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate user ID (this would come from a session or JWT token in a real application)
	userID := 1

	// Get address ID from URL parameters
	addressID, err := strconv.Atoi(r.URL.Query().Get("address_id"))
	if err != nil {
		http.Error(w, "Invalid address ID", http.StatusBadRequest)
		return
	}

	// Delete the address
	err = services.DeleteAddress(userID, addressID)
	if err != nil {
		http.Error(w, "Error deleting address: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
