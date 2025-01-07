package handlers

import (
	"ecommerce-project/backend/services"
	"encoding/json"
	"net/http"
)

// CreateOrderHandler creates an order based on the user's cart
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate user ID (in a real app, this would come from a JWT token or session)
	userID := 1

	// Retrieve the user's cart (in a real app, this would be a database call)
	cartItems, err := services.GetCart(userID)
	if err != nil {
		http.Error(w, "Error retrieving cart items", http.StatusInternalServerError)
		return
	}

	// Create the order
	order, err := services.CreateOrder(userID, cartItems)
	if err != nil {
		http.Error(w, "Error creating order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// GetOrdersHandler retrieves all orders for the user
func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate user ID (in a real app, this would come from a JWT token or session)
	userID := 1

	orders, err := services.GetOrders(userID)
	if err != nil {
		http.Error(w, "Error retrieving orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
