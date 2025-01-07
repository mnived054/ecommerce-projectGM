package services

import (
	"ecommerce-project/backend/models"
	"fmt"
	"time"
)

// Simulating an in-memory order list (for simplicity)
var orders []models.Order

// CreateOrder processes an order for the user from their cart
func CreateOrder(userID int, cartItems []models.Cart) (*models.Order, error) {
	var totalPrice float64
	var productNames []string

	// Calculate the total price of the cart
	for _, item := range cartItems {
		if item.UserID == userID {
			totalPrice += float64(item.Quantity) * item.Price
			productNames = append(productNames, item.Product)
		}
	}

	// If the cart is empty, return an error
	if len(productNames) == 0 {
		return nil, fmt.Errorf("no products in cart")
	}

	// Create the order
	order := models.Order{
		ID:         len(orders) + 1,
		UserID:     userID,
		TotalPrice: totalPrice,
		Products:   fmt.Sprintf("%v", productNames),
		Status:     "pending", // The order is created as "pending"
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	// Save the order
	orders = append(orders, order)

	// Clear the cart after placing the order
	clearCart(userID)

	return &order, nil
}

// GetOrders retrieves all orders for a user
func GetOrders(userID int) ([]models.Order, error) {
	var userOrders []models.Order
	for _, order := range orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders, nil
}
