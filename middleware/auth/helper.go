package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/napakornsk/go-rest/config"
)

func (h *AuthService) validateToken(tokenString *string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(*tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return config.PbcKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func (h *AuthService) isTokenExpire(claims jwt.MapClaims) bool {
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return true
	}
	return false
}
