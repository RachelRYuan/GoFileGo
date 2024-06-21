package utils

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
func VerifyTokenHeader(ctx *gin.Context, secretPublicKeyEnvName string) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	if tokenHeader == "" {
		return nil, errors.New("Authorization header is missing")
	}

	parts := strings.Fields(tokenHeader)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return nil, errors.New("Authorization header format must be Bearer {token}")
	}

	accessToken := parts[1]
	jwtSecretKey := GodotEnv(secretPublicKeyEnvName)
	if jwtSecretKey == "" {
		return nil, errors.New("Secret key is missing or not set")
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

// VerifyToken verifies the JWT token provided as a string.
func VerifyToken(accessToken, secretPublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey := GodotEnv(secretPublicKeyEnvName)
	if jwtSecretKey == "" {
		return nil, errors.New("Secret key is missing or not set")
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

// DecodeToken decodes the JWT token into an AccessToken struct.
func DecodeToken(token *jwt.Token) (AccessToken, error) {
	var accessToken AccessToken
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return accessToken, errors.New("Invalid token claims")
	}

	jsonString, err := json.Marshal(claims)
	if err != nil {
		return accessToken, err
	}

	err = json.Unmarshal(jsonString, &accessToken)
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}
