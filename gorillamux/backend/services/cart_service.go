package services

import (
	"ecommerce-project/backend/models"
)

var Cart []models.Cart

func AddToCart(UserID int, product models.Cart) (*models.Cart, error) {

	product.ID = len(Cart) + 1
	product.UserID = UserID
	Cart = append(Cart, product)

	return &product, nil
}

func GetCart(userID int) ([]models.Cart, error) {
	var userCart []models.Cart

	for _, item := range Cart {
		if item.UserID == userID {
			userCart = append(userCart, item)
		}
	}

	return userCart, nil
}

func clearCart(userID int) {
	var filteredCart []models.Cart
	for _, item := range Cart {
		if item.UserID != userID {
			filteredCart = append(filteredCart, item)
		}
	}

	Cart = filteredCart
}
