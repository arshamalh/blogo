package tools

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(subject string, lifespan time.Duration, secret string) (string, error) {
	payload := jwt.StandardClaims{
		Subject:   subject,
		ExpiresAt: time.Now().Add(lifespan).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(secret))
}

// Genarate a random unique hashed string from given string
func RandomUniqueHash(data string) string {
	data = strconv.Itoa(int(time.Now().Unix())) + data + strconv.Itoa(rand.Intn(999))
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

func ExtractTokenData(token string, secret string) (*jwt.StandardClaims, error) {
	jwt_token, err := jwt.ParseWithClaims(
		token, &jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	} else if !jwt_token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	payload := jwt_token.Claims.(*jwt.StandardClaims)
	return payload, nil
}
