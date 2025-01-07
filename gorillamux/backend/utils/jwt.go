package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key") // You can store this in an environment variable for security

// Claims struct represents the custom claims for the JWT token (you can add more fields as needed)
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token for the given user ID
func GenerateJWT(userID int) (string, error) {
	// Set the token expiration time (e.g., 1 hour)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims, which includes the user ID and expiration time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "ecommerce-project", // Issuer can be your app name or domain
		},
	}

	// Create a new JWT token with the specified claims and signing method (HS256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

// ValidateJWT validates a JWT token and returns the claims (userID) if valid
func ValidateJWT(tokenString string) (*Claims, error) {
	// Parse the token and validate it
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used for signing the token
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// If the token is valid, return the claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ExtractUserIDFromJWT extracts the user ID from a valid JWT token
func ExtractUserIDFromJWT(tokenString string) (int, error) {
	// Validate the JWT token and extract the claims
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return 0, err
	}

	// Return the user ID from the claims
	return claims.UserID, nil
}
