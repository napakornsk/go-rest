package service

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/napakornsk/go-rest/config"
	"github.com/napakornsk/go-rest/orm/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *PortfolioSrv) SignupUser(model *model.User) (*uint, error) {
	// hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), 10)
	if err != nil {
		log.Printf("Error while hashing password: %v\n", err)
		return nil, err
	}

	model.Password = string(hashPassword)

	id, err := s.srv.Create(model)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return id, nil
}

func (s *PortfolioSrv) SigninUser(model *model.User) (*string, error) {
	data, err := s.srv.GetUserByUsername(&model.Username)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}

	if data.ID == 0 {
		return nil, errors.New("Invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(model.Password)); err != nil {
		return nil, errors.New("Invalid username or password")
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
