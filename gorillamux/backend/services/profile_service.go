package services

import (
	"ecommerce-project/backend/models"
	"fmt"
)

var profiles = []models.Profile{
	{ID: 1, Username: "Nived", Email: "mnived054@gmail.com", Address: "Hyderabad-502319"},
}

func GetProfile(userID int) (*models.Profile, error) {
	for _, profile := range profiles {
		if profile.ID == userID {
			return &profile, nil
		}
	}
	return nil, fmt.Errorf("profile not found")
}

func UpdateProfile(userID int, profile models.Profile) (*models.Profile, error) {
	for i, p := range profiles {
		if p.ID == userID {
			profiles[i].Username = profile.Username
			profiles[i].Email = profile.Email
			profiles[i].Address = profile.Address
			return &profile, nil
		}
	}
	return nil, fmt.Errorf("profile not found")
}
