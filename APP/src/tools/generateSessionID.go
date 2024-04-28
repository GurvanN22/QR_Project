package tools

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken() string {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}
	tokenString := base64.StdEncoding.EncodeToString(token)
	return tokenString
}
