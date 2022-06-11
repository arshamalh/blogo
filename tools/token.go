package tools

import (
	"fmt"
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

func ExtractTokenData(token string, secret string) (*jwt.StandardClaims, error) {
	jwt_token, err := jwt.ParseWithClaims(
		token, &jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	payload := jwt_token.Claims.(*jwt.StandardClaims)
	if err != nil {
		// This error usually means that the token is expired
		return payload, err
	} else if !jwt_token.Valid {
		// This error usually means that the token is invalid because it was not signed with the correct secret
		// We have access to payloads here, but we don't return it.
		return nil, fmt.Errorf("invalid token")
	}
	return payload, nil
}
