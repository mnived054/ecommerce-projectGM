package handlers

import (
	"ecommerce-project/backend/models"
	"ecommerce-project/backend/services"
	"encoding/json"
	"net/http"
)

// AddToCartHandler adds a product to the user's cart
func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate user ID (in a real app, this would come from a JWT token or session)
	userID := 1

	var product models.Cart
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Add the product to the cart
	newProduct, err := services.AddToCart(userID, product)
	if err != nil {
		http.Error(w, "Error adding to cart: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

// GetCartHandler retrieves all items in the user's cart
func GetCartHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate user ID (in a real app, this would come from a JWT token or session)
	userID := 1

	cartItems, err := services.GetCart(userID)
	if err != nil {
		http.Error(w, "Error retrieving cart items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cartItems)
}
