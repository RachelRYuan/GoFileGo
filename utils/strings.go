package utils

import (
	"crypto/rand"
	"errors"
	"log"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenerateRandomString generates a random string of the specified length.
func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than zero")
	}

	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	for i := range b {
		b[i] = alphabet[b[i]%byte(len(alphabet))]
	}

	return string(b), nil
}

func main() {
	randomString, err := GenerateRandomString(10)
	if err != nil {
		log.Fatalf("Failed to generate random string: %v", err)
	}
	log.Println("Generated Random String:", randomString)
}
