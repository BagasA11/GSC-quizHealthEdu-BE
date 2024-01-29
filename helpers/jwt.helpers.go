package helpers

import (
	"BagasA11/go-POS/configs"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(ID uint, username string) (string, error) {
	claims := &configs.JWTClaims{
		ID:       ID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Hour)),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(configs.JWT_KEY)
	if err != nil {
		log.Panic()
		return "", err
	}
	return accessToken, err
}

func ValidateToken(inputToken string) (*configs.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(inputToken, &configs.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*configs.JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func UpdateToken(refreshToken string) (string, error) {
	claims, err := ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}
	newToken, err := GenerateAccessToken(claims.ID, claims.Username)
	if err != nil {
		return "", err
	}
	return newToken, err
}
