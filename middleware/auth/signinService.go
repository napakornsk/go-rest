package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/napakornsk/go-rest/config"
	"github.com/napakornsk/go-rest/orm/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) SigninService(model *model.User) (*string, error) {
	data, err := s.FindUser(model.Username)
	if err != nil {
		log.Printf("Error while searching user: %v\n", err)
		return nil, err
	}

	if data.ID == 0 {
		return nil, errors.New("Invalid credential")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(model.Password)); err != nil {
		return nil, errors.New("Invalid credential")
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": data.Username,
		"id":       data.ID,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(config.PrvKey)
	if err != nil {
		return nil, errors.New("Failed to generate token: " + err.Error())
	}

	return &tokenString, nil
}
