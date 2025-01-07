package services

import (
	"ecommerce-project/backend/models"
	"errors"
)

var addresses []models.Address

func AddAddress(userID int, address models.Address) (*models.Address, error) {
	address.ID = len(addresses) + 1
	address.UserID = userID
	addresses = append(addresses, address)

	return &address, nil
}

func GetAddresses(userID int) ([]models.Address, error) {
	var userAddresses []models.Address
	for _, address := range addresses {
		if address.UserID == userID {
			userAddresses = append(userAddresses, address)
		}
	}
	if len(userAddresses) == 0 {
		return nil, errors.New("no address found for the user")
	}
	return userAddresses, nil
}

func UpdateAddress(userID, addressID int, newAddress models.Address) (*models.Address, error) {
	for i, address := range addresses {
		if address.ID == addressID && address.UserID == userID {
			addresses[i] = newAddress
			addresses[i].ID = addressID
			addresses[i].UserID = userID
			return &addresses[i], nil
		}
	}
	return nil, errors.New("address not found")
}

func DeleteAddress(userID, addressID int) error {
	for i, address := range addresses {
		if address.ID == addressID && address.UserID == userID {
			addresses = append(addresses[:i], addresses[i+1:]...)
			return nil
		}
	}
	return errors.New("address not found")
}
