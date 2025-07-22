package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("RAHASIA-KU")

type JWTClaims struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"` // "access-token" atau "refresh-token"
	jwt.RegisteredClaims
}

func GenerateJWT(userID, tokenType string, duration time.Duration) (string, time.Time) {
	exp := time.Now().Add(duration)
	claims := JWTClaims{
		UserID: userID,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString(jwtKey)
	return signed, exp
}

func VerifyJWT(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*JWTClaims)
	return claims, nil
}
