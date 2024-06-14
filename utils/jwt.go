package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

// Sign creates a JWT token with the provided data, secret key environment variable name, and expiration duration.
func Sign(data map[string]interface{}, secretKeyEnvName string, expiresAt time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiresAt * time.Minute).Unix()

	jwtSecretKey := GodotEnv(secretKeyEnvName)

	claims := jwt.MapClaims{
		"exp":           expirationTime,
		"authorization": true,
	}

	for key, value := range data {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		logrus.Errorf("Error signing token: %v", err)
		return "", err
	}

	return signedToken, nil
}
